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

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	bmcv1alpha1 "github.com/tinkerbell/rufio/api/v1alpha1"
)

// Index key for Job Owner Name
const jobOwnerKey = ".metadata.controller"

// JobReconciler reconciles a Job object
type JobReconciler struct {
	client client.Client
	logger logr.Logger
}

// NewJobReconciler returns a new JobReconciler
func NewJobReconciler(client client.Client, logger logr.Logger) *JobReconciler {
	return &JobReconciler{
		client: client,
		logger: logger,
	}
}

//+kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=jobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=jobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=jobs/finalizers,verbs=update

// Reconcile runs a Job.
// Creates the individual Tasks on the cluster.
// Watches for Task and creates next Job Task based on conditions.
func (r *JobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.logger.WithValues("Job", req.NamespacedName)
	logger.Info("Reconciling Job")

	// Fetch the bmcJob object
	bmcJob := &bmcv1alpha1.Job{}
	err := r.client.Get(ctx, req.NamespacedName, bmcJob)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		logger.Error(err, "Failed to get Job")
		return ctrl.Result{}, err
	}

	// Deletion is a noop.
	if !bmcJob.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	// Job is Completed or Failed is noop.
	if bmcJob.HasCondition(bmcv1alpha1.JobCompleted, bmcv1alpha1.ConditionTrue) ||
		bmcJob.HasCondition(bmcv1alpha1.JobFailed, bmcv1alpha1.ConditionTrue) {
		return ctrl.Result{}, nil
	}

	// Create a patch from the initial Job object
	// Patch is used to update Status after reconciliation
	bmcJobPatch := client.MergeFrom(bmcJob.DeepCopy())

	return r.reconcile(ctx, bmcJob, bmcJobPatch, logger)
}

func (r *JobReconciler) reconcile(ctx context.Context, bmj *bmcv1alpha1.Job, bmjPatch client.Patch, logger logr.Logger) (ctrl.Result, error) {
	// Check if Job is not currently Running
	// Initialize the StartTime for the Job
	// Set the Job to Running condition True
	if !bmj.HasCondition(bmcv1alpha1.JobRunning, bmcv1alpha1.ConditionTrue) {
		now := metav1.Now()
		bmj.Status.StartTime = &now
		bmj.SetCondition(bmcv1alpha1.JobRunning, bmcv1alpha1.ConditionTrue)
	}

	// Get Machine object for the Job
	// Requeue if error
	machine := &bmcv1alpha1.Machine{}
	err := r.getMachine(ctx, bmj.Spec.MachineRef, machine)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("get Job %s/%s MachineRef: %v", bmj.Namespace, bmj.Name, err)
	}

	// List all Task owned by Job
	bmcTasks := &bmcv1alpha1.TaskList{}
	err = r.client.List(ctx, bmcTasks, client.MatchingFields{jobOwnerKey: bmj.Name})
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to list owned Tasks for Job %s/%s", bmj.Namespace, bmj.Name)
	}

	completedTasksCount := 0
	// Iterate Task Items.
	// Count the number of completed tasks.
	// Set the Job condition Failed True if Task has failed.
	// If the Task has neither Completed or Failed is noop.
	for _, task := range bmcTasks.Items {
		if task.HasCondition(bmcv1alpha1.TaskCompleted, bmcv1alpha1.ConditionTrue) {
			completedTasksCount += 1
			continue
		}

		if task.HasCondition(bmcv1alpha1.TaskFailed, bmcv1alpha1.ConditionTrue) {
			err := fmt.Errorf("Task %s/%s failed", task.Namespace, task.Name)
			bmj.SetCondition(bmcv1alpha1.JobFailed, bmcv1alpha1.ConditionTrue, bmcv1alpha1.WithJobConditionMessage(err.Error()))
			patchErr := r.patchStatus(ctx, bmj, bmjPatch)
			if patchErr != nil {
				return ctrl.Result{}, utilerrors.NewAggregate([]error{patchErr, err})
			}

			return ctrl.Result{}, err
		}

		return ctrl.Result{}, nil
	}

	// Check if all Job tasks have Completed
	// Set the Task CompletionTime
	// Set Task Condition Completed True
	if completedTasksCount == len(bmj.Spec.Tasks) {
		bmj.SetCondition(bmcv1alpha1.JobCompleted, bmcv1alpha1.ConditionTrue)
		now := metav1.Now()
		bmj.Status.CompletionTime = &now
		err = r.patchStatus(ctx, bmj, bmjPatch)
		return ctrl.Result{}, err
	}

	// Create the first Task for the Job
	if err := r.createTaskWithOwner(ctx, *bmj, completedTasksCount, machine.Spec.Connection); err != nil {
		// Set the Job condition Failed True
		bmj.SetCondition(bmcv1alpha1.JobFailed, bmcv1alpha1.ConditionTrue, bmcv1alpha1.WithJobConditionMessage(err.Error()))
		patchErr := r.patchStatus(ctx, bmj, bmjPatch)
		if patchErr != nil {
			return ctrl.Result{}, utilerrors.NewAggregate([]error{patchErr, err})
		}

		return ctrl.Result{}, err
	}

	// Patch the status at the end of reconcile loop
	err = r.patchStatus(ctx, bmj, bmjPatch)
	return ctrl.Result{}, err
}

// getMachine Gets the Machine from MachineRef
func (r *JobReconciler) getMachine(ctx context.Context, reference corev1.ObjectReference, machine *bmcv1alpha1.Machine) error {
	key := types.NamespacedName{Namespace: reference.Namespace, Name: reference.Name}
	err := r.client.Get(ctx, key, machine)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("Machine %s not found: %v", key, err)
		}
		return fmt.Errorf("failed to get Machine %s: %v", key, err)
	}

	return nil
}

// createTaskWithOwner creates a Task object with an OwnerReference set to the Job
func (r *JobReconciler) createTaskWithOwner(ctx context.Context, bmj bmcv1alpha1.Job, taskIndex int, conn bmcv1alpha1.Connection) error {
	isController := true
	bmcTask := &bmcv1alpha1.Task{
		ObjectMeta: metav1.ObjectMeta{
			Name:      bmcv1alpha1.FormatTaskName(bmj, taskIndex),
			Namespace: bmj.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: bmj.APIVersion,
					Kind:       bmj.Kind,
					Name:       bmj.Name,
					UID:        bmj.ObjectMeta.UID,
					Controller: &isController,
				},
			},
		},
		Spec: bmcv1alpha1.TaskSpec{
			Task:       bmj.Spec.Tasks[taskIndex],
			Connection: conn,
		},
	}

	err := r.client.Create(ctx, bmcTask)
	if err != nil {
		return fmt.Errorf("failed to create Task %s/%s: %v", bmcTask.Namespace, bmcTask.Name, err)
	}

	return nil
}

// patchStatus patches the specified patch on the Job.
func (r *JobReconciler) patchStatus(ctx context.Context, bmj *bmcv1alpha1.Job, patch client.Patch) error {
	err := r.client.Status().Patch(ctx, bmj, patch)
	if err != nil {
		return fmt.Errorf("failed to patch Job %s/%s status: %v", bmj.Namespace, bmj.Name, err)
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *JobReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(
		ctx,
		&bmcv1alpha1.Task{},
		jobOwnerKey,
		bmcTaskOwnerIndexFunc,
	); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&bmcv1alpha1.Job{}).
		Watches(
			&source.Kind{Type: &bmcv1alpha1.Task{}},
			&handler.EnqueueRequestForOwner{
				OwnerType:    &bmcv1alpha1.Job{},
				IsController: true,
			}).
		Complete(r)
}

// bmcTaskOwnerIndexFunc is Indexer func which returns the owner name for obj.
func bmcTaskOwnerIndexFunc(obj client.Object) []string {
	task, ok := obj.(*bmcv1alpha1.Task)
	if !ok {
		return nil
	}

	owner := metav1.GetControllerOf(task)
	if owner == nil {
		return nil
	}

	// Check if owner is Job
	if owner.Kind != "Job" || owner.APIVersion != bmcv1alpha1.GroupVersion.String() {
		return nil
	}

	return []string{owner.Name}
}
