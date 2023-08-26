/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GrafanaOrganizationSpec defines the desired state of GrafanaOrganization
type GrafanaOrganizationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Required
	// Wheter it is or it isn't Grafana main organization
	IsMainOrg bool `json:"ismainorg"`
	// // +kubebuilder:validation:Required
	// // Grafana Organization tenant
	// Tenant string `json:"tenant"`
	// // +listType=set
	// // List of groups that will have admin role in the Grafana Organization
	// AdminGroups []Group `json:"adminGroups,omitempty"`
	// // +listType=set
	// // List of groups that will have editor role in the Grafana Organization
	// EditorGroups []Group `json:"editorGroups,omitempty"`
	// // +listType=set
	// // List of groups that will have viewer role in the Grafana Organization
	// ViewerGroups []Group `json:"viewerGroups,omitempty"`
}

// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=256
type Group string

// GrafanaOrganizationStatus defines the observed state of GrafanaOrganization
type GrafanaOrganizationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// General status of the resource
	// +kubebuilder:validation:Enum=Pending;Ready;Error
	Phase string `json:"phase,omitempty"`
	// Represents the observations of a GrafanaOrganization current state.
	// Known .status.conditions.type are: "Available", "Progressing", and "Degraded"
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
	// List of errors in last reconciliation
	Errors []string `json:"errors,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GrafanaOrganization is the Schema for the grafanaorganizations API
type GrafanaOrganization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GrafanaOrganizationSpec   `json:"spec,omitempty"`
	Status GrafanaOrganizationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GrafanaOrganizationList contains a list of GrafanaOrganization
type GrafanaOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrafanaOrganization `json:"items"`
}

func (in *GrafanaOrganizationList) Find(name string) *GrafanaOrganization {
	for _, organization := range in.Items {
		if organization.Name == name {
			return &organization
		}
	}
	return nil
}

func init() {
	SchemeBuilder.Register(&GrafanaOrganization{}, &GrafanaOrganizationList{})
}
