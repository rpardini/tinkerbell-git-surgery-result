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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PowerState represents power state the BaseboardManagement.
type PowerState string

// BootDevice represents boot device of the BaseboardManagement.
type BootDevice string

// BaseboardManagementConditionType represents the condition of the BaseboardManagement.
type BaseboardManagementConditionType string

const (
	On  PowerState = "on"
	Off PowerState = "off"
)

const (
	PXE   BootDevice = "pxe"
	Disk  BootDevice = "disk"
	BIOS  BootDevice = "bios"
	CDROM BootDevice = "cdrom"
	Safe  BootDevice = "safe"
)

const (
	// ConnectionError represents failure to connect to the BaseboardManagement.
	ConnectionError BaseboardManagementConditionType = "ConnectionError"
)

// BaseboardManagementSpec defines the desired state of BaseboardManagement
type BaseboardManagementSpec struct {

	// Connection represents the BaseboardManagement connectivity information.
	Connection Connection `json:"connection"`

	// Power is the desired power state of the BaseboardManagement.
	// +kubebuilder:validation:Enum=On;Off
	Power PowerState `json:"power"`
}

type Connection struct {
	// Host is the host IP address or hostname of the BaseboardManagement.
	// +kubebuilder:validation:MinLength=1
	Host string `json:"host"`

	// AuthSecretRef is the SecretReference that contains authentication information of the BaseboardManagement.
	// The Secret must contain username and password keys.
	AuthSecretRef corev1.SecretReference `json:"authSecretRef"`

	// InsecureTLS specifies trusted TLS connections.
	InsecureTLS bool `json:"insecureTLS"`
}

// BaseboardManagementStatus defines the observed state of BaseboardManagement
type BaseboardManagementStatus struct {
	// Power is the current power state of the BaseboardManagement.
	// +kubebuilder:validation:Enum=On;Off
	// +optional
	Power PowerState `json:"powerState,omitempty"`

	// Conditions represents the latest available observations of an object's current state.
	// +optional
	Conditions []BaseboardManagementCondition `json:"conditions,omitempty"`
}

type BaseboardManagementCondition struct {
	// Type of the BaseboardManagement condition.
	Type BaseboardManagementConditionType `json:"type"`

	// Message represents human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// BaseboardManagementRef defines the reference information to a BaseboardManagement resource.
type BaseboardManagementRef struct {
	// Name is unique within a namespace to reference a BaseboardManagement resource.
	Name string `json:"name"`

	// Namespace defines the space within which the BaseboardManagement name must be unique.
	Namespace string `json:"namespace"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=baseboardmanagements,scope=Namespaced,categories=tinkerbell,singular=baseboardmanagement,shortName=bm

// BaseboardManagement is the Schema for the baseboardmanagements API
type BaseboardManagement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaseboardManagementSpec   `json:"spec,omitempty"`
	Status BaseboardManagementStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BaseboardManagementList contains a list of BaseboardManagement
type BaseboardManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaseboardManagement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaseboardManagement{}, &BaseboardManagementList{})
}
