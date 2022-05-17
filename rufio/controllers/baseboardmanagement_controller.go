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
	"strings"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	bmcv1alpha1 "github.com/tinkerbell/rufio/api/v1alpha1"
)

// BMCClient represents a baseboard management controller client.
// It defines a set of methods to connect and interact with a BMC.
type BMCClient interface {
	// Close ends the connection with the bmc.
	Close(ctx context.Context) error
	// GetPowerState fetches the current power status of the bmc.
	GetPowerState(ctx context.Context) (string, error)
	// SetPowerState power controls the bmc to the input power state.
	SetPowerState(ctx context.Context, state string) (bool, error)
	// SetBootDevice sets the boot device on the bmc.
	// Currently this sets the first boot device.
	// setPersistent, if true will set the boot device permanently. If false, sets one time boot.
	// efiBoot, if true passes efiboot options while setting boot device.
	SetBootDevice(ctx context.Context, bootDevice string, setPersistent, efiBoot bool) (bool, error)
}

// BMCClientFactoryFunc defines a func that returns a BMCClient
type BMCClientFactoryFunc func(ctx context.Context, hostIP, port, username, password string) (BMCClient, error)

// BaseboardManagementReconciler reconciles a BaseboardManagement object
type BaseboardManagementReconciler struct {
	client           client.Client
	bmcClientFactory BMCClientFactoryFunc
	logger           logr.Logger
}

// NewBaseboardManagementReconciler returns a new BaseboardManagementReconciler
func NewBaseboardManagementReconciler(client client.Client, bmcClientFactory BMCClientFactoryFunc, logger logr.Logger) *BaseboardManagementReconciler {
	return &BaseboardManagementReconciler{
		client:           client,
		bmcClientFactory: bmcClientFactory,
		logger:           logger,
	}
}

// baseboardManagementFieldReconciler defines a function to reconcile BaseboardManagement spec field
type baseboardManagementFieldReconciler func(context.Context, *bmcv1alpha1.BaseboardManagement, BMCClient) error

//+kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=baseboardmanagements,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=baseboardmanagements/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bmc.tinkerbell.org,resources=baseboardmanagements/finalizers,verbs=update

// Reconcile ensures the state of a BaseboardManagement.
// Gets the BaseboardManagement object and uses the SecretReference to initialize a BMC Client.
// Ensures the BMC power is set to the desired state.
// Updates the status and conditions accordingly.
func (r *BaseboardManagementReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.logger.WithValues("BaseboardManagement", req.NamespacedName)
	logger.Info("Reconciling BaseboardManagement")

	// Fetch the BaseboardManagement object
	baseboardManagement := &bmcv1alpha1.BaseboardManagement{}
	err := r.client.Get(ctx, req.NamespacedName, baseboardManagement)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		logger.Error(err, "Failed to get BaseboardManagement")
		return ctrl.Result{}, err
	}

	// Create a patch from the initial BaseboardManagement object
	// Patch is used to update Status after reconciliation
	baseboardManagementPatch := client.MergeFrom(baseboardManagement.DeepCopy())

	// Deletion is a noop.
	if !baseboardManagement.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	return r.reconcile(ctx, baseboardManagement, baseboardManagementPatch, logger)
}

func (r *BaseboardManagementReconciler) reconcile(ctx context.Context, bm *bmcv1alpha1.BaseboardManagement, bmPatch client.Patch, logger logr.Logger) (ctrl.Result, error) {
	// Fetching username, password from SecretReference
	// Requeue if error fetching secret
	username, password, err := r.resolveAuthSecretRef(ctx, bm.Spec.Connection.AuthSecretRef)
	if err != nil {
		return ctrl.Result{Requeue: true}, fmt.Errorf("resolving BaseboardManagement %s/%s SecretReference: %v", bm.Namespace, bm.Name, err)
	}

	// TODO (pokearu): Remove port hardcoding
	// Initializing BMC Client
	bmcClient, err := r.bmcClientFactory(ctx, bm.Spec.Connection.Host, "623", username, password)
	if err != nil {
		logger.Error(err, "BMC connection failed", "host", bm.Spec.Connection.Host)
		result, setConditionErr := r.setCondition(ctx, bm, bmPatch, bmcv1alpha1.ConnectionError, err.Error())
		if setConditionErr != nil {
			return result, utilerrors.NewAggregate([]error{fmt.Errorf("failed to set conditions: %v", setConditionErr), err})
		}
		return result, err
	}

	// Close BMC connection after reconcilation
	defer func() {
		err = bmcClient.Close(ctx)
		if err != nil {
			logger.Error(err, "BMC close connection failed", "host", bm.Spec.Connection.Host)
		}
	}()

	// fieldReconcilers defines BaseboardManagement spec field reconciler functions
	fieldReconcilers := []baseboardManagementFieldReconciler{
		r.reconcilePower,
	}
	for _, reconiler := range fieldReconcilers {
		if err := reconiler(ctx, bm, bmcClient); err != nil {
			logger.Error(err, "Failed to reconcile BaseboardManagement", "host", bm.Spec.Connection.Host)
		}
	}

	// Patch the status after each reconciliation
	return r.reconcileStatus(ctx, bm, bmPatch)
}

// reconcilePower ensures the BaseboardManagement Power is in the desired state.
func (r *BaseboardManagementReconciler) reconcilePower(ctx context.Context, bm *bmcv1alpha1.BaseboardManagement, bmcClient BMCClient) error {
	powerStatus, err := bmcClient.GetPowerState(ctx)
	if err != nil {
		return fmt.Errorf("failed to get power state: %v", err)
	}

	// If BaseboardManagement has desired power state then return
	if bm.Spec.Power == bmcv1alpha1.PowerState(strings.ToLower(powerStatus)) {
		// Update status to represent current power state
		bm.Status.Power = bm.Spec.Power
		return nil
	}

	// Setting baseboard management to desired power state
	_, err = bmcClient.SetPowerState(ctx, string(bm.Spec.Power))
	if err != nil {
		return fmt.Errorf("failed to set power state: %v", err)
	}

	// Update status to represent current power state
	bm.Status.Power = bm.Spec.Power

	return nil
}

// setCondition updates the status.Condition if the condition type is present.
// Appends if new condition is found.
// Patches the BaseboardManagement status.
func (r *BaseboardManagementReconciler) setCondition(ctx context.Context, bm *bmcv1alpha1.BaseboardManagement, bmPatch client.Patch, cType bmcv1alpha1.BaseboardManagementConditionType, message string) (ctrl.Result, error) {
	currentConditions := bm.Status.Conditions
	for i := range currentConditions {
		// If condition exists, update the message if different
		if currentConditions[i].Type == cType {
			if currentConditions[i].Message != message {
				bm.Status.Conditions[i].Message = message
				return r.patchStatus(ctx, bm, bmPatch)
			}
			return ctrl.Result{}, nil
		}
	}

	// Append new condition to Conditions
	condition := bmcv1alpha1.BaseboardManagementCondition{
		Type:    cType,
		Message: message,
	}
	bm.Status.Conditions = append(bm.Status.Conditions, condition)

	return r.patchStatus(ctx, bm, bmPatch)
}

// reconcileStatus updates the Power and Conditions and patches BaseboardManagement status.
func (r *BaseboardManagementReconciler) reconcileStatus(ctx context.Context, bm *bmcv1alpha1.BaseboardManagement, bmPatch client.Patch) (ctrl.Result, error) {
	// TODO: (pokearu) modify conditions to model current state.
	// Add condition Status to represent if object has a condition
	// insted of clearing the conditions.
	bm.Status.Conditions = []bmcv1alpha1.BaseboardManagementCondition{}

	return r.patchStatus(ctx, bm, bmPatch)
}

// patchStatus patches the specifies patch on the BaseboardManagement.
func (r *BaseboardManagementReconciler) patchStatus(ctx context.Context, bm *bmcv1alpha1.BaseboardManagement, patch client.Patch) (ctrl.Result, error) {
	err := r.client.Status().Patch(ctx, bm, patch)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to patch BaseboardManagement %s/%s status: %v", bm.Namespace, bm.Name, err)
	}

	return ctrl.Result{}, nil
}

// resolveAuthSecretRef Gets the Secret from the SecretReference.
// Returns the username and password encoded in the Secret.
func (r *BaseboardManagementReconciler) resolveAuthSecretRef(ctx context.Context, secretRef corev1.SecretReference) (string, string, error) {
	secret := &corev1.Secret{}
	key := types.NamespacedName{Namespace: secretRef.Namespace, Name: secretRef.Name}

	if err := r.client.Get(ctx, key, secret); err != nil {
		if apierrors.IsNotFound(err) {
			return "", "", fmt.Errorf("secret %s not found: %v", key, err)
		}

		return "", "", fmt.Errorf("failed to retrieve secret %s : %v", secretRef, err)
	}

	username, ok := secret.Data["username"]
	if !ok {
		return "", "", fmt.Errorf("'username' required in BaseboardManagement secret")
	}

	password, ok := secret.Data["password"]
	if !ok {
		return "", "", fmt.Errorf("'password' required in BaseboardManagement secret")
	}

	return string(username), string(password), nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BaseboardManagementReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bmcv1alpha1.BaseboardManagement{}).
		Complete(r)
}
