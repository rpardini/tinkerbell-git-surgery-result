package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WorkflowSpec struct {
	// HardwareRef is a reference to a Hardware resource this workflow will execute on.
	// If no namespace is specified the Workflow's namespace is assumed.
	HardwareRef corev1.LocalObjectReference `json:"hardwareRef,omitempty"`

	// TemplateRef is a reference to a Template resource used to render workflow actions.
	// If no namespace is specified the Workflow's namespace is assumed.
	TemplateRef corev1.LocalObjectReference `json:"templateRef,omitempty"`

	// TemplateParams are a list of key-value pairs that are injected into templates at render
	// time. TemplateParams are exposed to templates using a top level .Params key.
	//
	// For example, TemplateParams = {"foo": "bar"}, the foo key can be accessed via .Params.foo.
	// +optional
	TemplateParams map[string]string `json:"templateParams,omitempty"`

	// TimeoutSeconds defines the time the workflow has to complete. The timer begins when the first
	// action is requested. When set to 0, no timeout is applied.
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	TimeoutSeconds int64 `json:"timeout,omitempty"`
}

type WorkflowStatus struct {
	// Actions is a list of action states.
	Actions []ActionStatus `json:"actions"`

	// StartedAt is the time the first action was requested. Nil indicates the Workflow has not
	// started.
	StartedAt *metav1.Time `json:"startedAt,omitempty"`

	// LastTransition is the observed time when State transitioned last.
	LastTransition *metav1.Time `json:"lastTransitioned,omitempty"`

	// State describes the current state of the workflow. For the workflow to enter the
	// WorkflowStateSucceeded state all actions must be in ActionStateSucceeded. The Workflow will
	// enter a WorkflowStateFailed if 1 or more Actions fails.
	State WorkflowState `json:"state,omitempty"`

	// Conditions details a set of observations about the Workflow.
	Conditions Conditions `json:"conditions"`
}

// ActionStatus describes status information about an action.
type ActionStatus struct {
	// Rendered is the rendered action.
	Rendered Action `json:"rendered,omitempty"`

	// ID uniquely identifies the action status.
	ID string `json:"id"`

	// StartedAt is the time the action was started as reported by the client. Nil indicates the
	// Action has not started.
	StartedAt *metav1.Time `json:"startedAt,omitempty"`

	// LastTransition is the observed time when State transitioned last.
	LastTransition *metav1.Time `json:"lastTransitioned,omitempty"`

	// State describes the current state of the action.
	State ActionState `json:"state,omitempty"`

	// FailureReason is a short CamelCase word or phrase describing why the Action entered
	// ActionStateFailed.
	FailureReason string `json:"failureReason,omitempty"`

	// FailureMessage is a free-form user friendly message describing why the Action entered the
	// ActionStateFailed state. Typically, this is an elaboration on the Reason.
	FailureMessage string `json:"failureMessage,omitempty"`
}

// State describes the point in time state of a Workflow.
type WorkflowState string

const (
	// WorkflowStatePending indicates the workflow is in a pending state.
	WorkflowStatePending WorkflowState = "Pending"

	// WorkflowStateRunning indicates the first Action has been requested and the Workflow is in
	// progress.
	WorkflowStateRunning WorkflowState = "Running"

	// WorkflowStateSucceeded indicates all Workflow actions have successfully completed.
	WorkflowStateSucceeded WorkflowState = "Succeeded"

	// WorkflowStateFailed indicates an Action entered a failure state.
	WorkflowStateFailed WorkflowState = "Failed"
)

// ActionState describes a point in time state of an Action.
type ActionState string

const (
	// ActionStatePending indicates an Action is awaiting execution.
	ActionStatePending ActionState = "Pending"

	// ActionStateRunning indicates an Action has begun execution.
	ActionStateRunning ActionState = "Running"

	// ActionStateSucceeded indicates an Action completed execution successfully.
	ActionStateSucceeded ActionState = "Succeeded"

	// ActionStatFailed indicates an Action failed to execute. Users may inspect the associated
	// Workflow resource to gain deeper insights into why the action failed.
	ActionStateFailed ActionState = "Failed"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=tinkerbell,shortName=wf
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state",description="State of the workflow such as Pending,Running etc"
// +kubebuilder:printcolumn:name="Hardware",type="string",JSONPath=".spec.hardwareRef",description="Hardware object that runs the workflow"
// +kubebuilder:printcolumn:name="Template",type="string",JSONPath=".spec.templateRef",description="Template to run on the associated Hardware"
// +kubebuilder:unservedversion

// Workflow describes a set of actions to be run on a specific Hardware. Workflows execute
// once and should be considered ephemeral.
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkflowSpec   `json:"spec,omitempty"`
	Status WorkflowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workflow `json:"items,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Workflow{}, &WorkflowList{})
}
