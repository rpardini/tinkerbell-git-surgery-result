/*
Copyright 2022 Tinkerbell.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	bmcv1alpha1 "github.com/tinkerbell/rufio/api/v1alpha1"
)

const powerActionRequeueAfter = 3 * time.Second

// TaskReconciler reconciles a Task object
type TaskReconciler struct {
	client           client.Client
	bmcClientFactory BMCClientFactoryFunc
}

// NewTaskReconciler returns a new TaskReconciler
func NewTaskReconciler(client client.Client, bmcClientFactory BMCClientFactoryFunc) *TaskReconciler {
	return &TaskReconciler{
		client:           client,
		bmcClientFactory: bmcClientFactory,
	}
}

//+kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=tasks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=tasks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=tasks/finalizers,verbs=update

// Reconcile runs a Task.
// Establishes a connection to the BMC.
// Runs the specified action in the Task.
func (r *TaskReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrl.LoggerFrom(ctx)
	logger.Info("Reconciling Task")

	// Fetch the Task object
	task := &bmcv1alpha1.Task{}
	if err := r.client.Get(ctx, req.NamespacedName, task); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		logger.Error(err, "Failed to get Task")
		return ctrl.Result{}, err
	}

	// Deletion is a noop.
	if !task.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	// Task is Completed or Failed is noop.
	if task.HasCondition(bmcv1alpha1.TaskFailed, bmcv1alpha1.ConditionTrue) ||
		task.HasCondition(bmcv1alpha1.TaskCompleted, bmcv1alpha1.ConditionTrue) {
		return ctrl.Result{}, nil
	}

	// Create a patch from the initial Task object
	// Patch is used to update Status after reconciliation
	taskPatch := client.MergeFrom(task.DeepCopy())

	return r.reconcile(ctx, task, taskPatch, logger)
}

func (r *TaskReconciler) reconcile(ctx context.Context, task *bmcv1alpha1.Task, taskPatch client.Patch, logger logr.Logger) (ctrl.Result, error) {
	// Fetching username, password from SecretReference in Connection.
	// Requeue if error fetching secret
	username, password, err := resolveAuthSecretRef(ctx, r.client, task.Spec.Connection.AuthSecretRef)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("resolving connection secret for task %s/%s: %v", task.Namespace, task.Name, err)
	}

	// Initializing BMC Client
	bmcClient, err := r.bmcClientFactory(ctx, task.Spec.Connection.Host, strconv.Itoa(task.Spec.Connection.Port), username, password)
	if err != nil {
		logger.Error(err, "BMC connection failed", "host", task.Spec.Connection.Host)
		task.SetCondition(bmcv1alpha1.TaskFailed, bmcv1alpha1.ConditionTrue, bmcv1alpha1.WithTaskConditionMessage(fmt.Sprintf("Failed to connect to BMC: %v", err)))
		patchErr := r.patchStatus(ctx, task, taskPatch)
		if patchErr != nil {
			return ctrl.Result{}, utilerrors.NewAggregate([]error{patchErr, err})
		}

		return ctrl.Result{}, err
	}

	defer func() {
		// Close BMC connection after reconcilation
		if err := bmcClient.Close(ctx); err != nil {
			logger.Error(err, "BMC close connection failed", "host", task.Spec.Connection.Host)
		}
	}()

	// Task has StartTime, we check the status.
	// Requeue if actions did not complete.
	if !task.Status.StartTime.IsZero() {
		jobRunningTime := time.Since(task.Status.StartTime.Time)
		// TODO(pokearu): add timeout for tasks on API spec
		if jobRunningTime >= 3*time.Minute {
			timeOutErr := fmt.Errorf("bmc task timeout: %d", jobRunningTime)
			// Set Task Condition Failed True
			task.SetCondition(bmcv1alpha1.TaskFailed, bmcv1alpha1.ConditionTrue, bmcv1alpha1.WithTaskConditionMessage(timeOutErr.Error()))
			patchErr := r.patchStatus(ctx, task, taskPatch)
			if patchErr != nil {
				return ctrl.Result{}, utilerrors.NewAggregate([]error{patchErr, timeOutErr})
			}

			return ctrl.Result{}, timeOutErr
		}

		result, err := r.checkTaskStatus(ctx, task.Spec.Task, bmcClient)
		if err != nil {
			return result, fmt.Errorf("bmc task status check: %s", err)
		}

		if !result.IsZero() {
			return result, nil
		}

		// Set the Task CompletionTime
		now := metav1.Now()
		task.Status.CompletionTime = &now
		// Set Task Condition Completed True
		task.SetCondition(bmcv1alpha1.TaskCompleted, bmcv1alpha1.ConditionTrue)
		if err := r.patchStatus(ctx, task, taskPatch); err != nil {
			return result, err
		}

		return result, nil
	}

	// Set the Task StartTime
	now := metav1.Now()
	task.Status.StartTime = &now
	// run the specified Task in Task
	if err := r.runTask(ctx, task.Spec.Task, bmcClient); err != nil {
		// Set Task Condition Failed True
		task.SetCondition(bmcv1alpha1.TaskFailed, bmcv1alpha1.ConditionTrue, bmcv1alpha1.WithTaskConditionMessage(err.Error()))
		patchErr := r.patchStatus(ctx, task, taskPatch)
		if patchErr != nil {
			return ctrl.Result{}, utilerrors.NewAggregate([]error{patchErr, err})
		}

		return ctrl.Result{}, err
	}

	if err := r.patchStatus(ctx, task, taskPatch); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// runTask executes the defined Task in a Task
func (r *TaskReconciler) runTask(ctx context.Context, task bmcv1alpha1.Action, bmcClient BMCClient) error {
	if task.PowerAction != nil {
		_, err := bmcClient.SetPowerState(ctx, string(*task.PowerAction))
		if err != nil {
			return fmt.Errorf("failed to perform PowerAction: %v", err)
		}
	}

	if task.OneTimeBootDeviceAction != nil {
		// OneTimeBootDeviceAction currently sets the first boot device from Devices.
		// setPersistent is false.
		_, err := bmcClient.SetBootDevice(ctx, string(task.OneTimeBootDeviceAction.Devices[0]), false, task.OneTimeBootDeviceAction.EFIBoot)
		if err != nil {
			return fmt.Errorf("failed to perform OneTimeBootDeviceAction: %v", err)
		}
	}

	return nil
}

// checkTaskStatus checks if Task action completed.
// This is currently limited only to a few PowerAction types.
func (r *TaskReconciler) checkTaskStatus(ctx context.Context, task bmcv1alpha1.Action, bmcClient BMCClient) (ctrl.Result, error) {
	// TODO(pokearu): Extend to all actions.
	if task.PowerAction != nil {
		rawState, err := bmcClient.GetPowerState(ctx)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to get power state: %v", err)
		}

		state, err := convertRawBMCPowerState(rawState)
		if err != nil {
			return ctrl.Result{}, err
		}

		switch *task.PowerAction {
		case bmcv1alpha1.PowerOn:
			if bmcv1alpha1.On != state {
				return ctrl.Result{RequeueAfter: powerActionRequeueAfter}, nil
			}
		case bmcv1alpha1.PowerHardOff, bmcv1alpha1.PowerSoftOff:
			if bmcv1alpha1.Off != state {
				return ctrl.Result{RequeueAfter: powerActionRequeueAfter}, nil
			}
		}
	}

	// Other Task action types do not support checking status. So noop.
	return ctrl.Result{}, nil
}

// patchStatus patches the specified patch on the Task.
func (r *TaskReconciler) patchStatus(ctx context.Context, task *bmcv1alpha1.Task, patch client.Patch) error {
	err := r.client.Status().Patch(ctx, task, patch)
	if err != nil {
		return fmt.Errorf("failed to patch Task %s/%s status: %v", task.Namespace, task.Name, err)
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TaskReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bmcv1alpha1.Task{}).
		Complete(r)
}
