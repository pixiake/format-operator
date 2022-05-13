/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FormatSpec defines the desired state of Format
type FormatSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Format. Edit format_types.go to remove/update
	Foo []string `json:"foo,omitempty"`
	//+kubebuilder:pruning:PreserveUnknownFields
	Test runtime.RawExtension `json:"test,omitempty"`
	//+kubebuilder:pruning:PreserveUnknownFields
	Test2 runtime.RawExtension `json:"test2,omitempty"`
}

// FormatStatus defines the observed state of Format
type FormatStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	//+kubebuilder:pruning:PreserveUnknownFields
	Results []runtime.RawExtension `json:"results,omitempty"`
}

//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Format is the Schema for the formats API
type Format struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FormatSpec   `json:"spec,omitempty"`
	Status FormatStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FormatList contains a list of Format
type FormatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Format `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Format{}, &FormatList{})
}
