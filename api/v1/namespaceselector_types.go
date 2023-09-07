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

// NamespaceSelectorSpec defines the desired state of NamespaceSelector
type NamespaceSelectorSpec struct {
	// Namespace is the name of the namespace
	Namespace string `json:"namespace"`

	// RequiredLabel is the label that must be present in this namespace
	RequiredLabel string `json:"requiredLabel"`
}

// NamespaceSelectorStatus defines the observed state of NamespaceSelector
type NamespaceSelectorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NamespaceSelector is the Schema for the namespaceselectors API
type NamespaceSelector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceSelectorSpec   `json:"spec,omitempty"`
	Status NamespaceSelectorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NamespaceSelectorList contains a list of NamespaceSelector
type NamespaceSelectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceSelector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NamespaceSelector{}, &NamespaceSelectorList{})
}
