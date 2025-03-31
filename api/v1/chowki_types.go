/*
Copyright 2025.

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

// ChowkiSpec defines the desired state of Chowki.
type ChowkiSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Chowki. Edit chowki_types.go to remove/update
	Name        string `json:"name,omitempty"`
	Kind        string `json:"kind,omitempty"`
	ResoureType string `json:"resoureType,omitempty"`
	ResoureName string `json:"resourceName,omitempty"`
}

// ChowkiStatus defines the observed state of Chowki.
type ChowkiStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Chowki is the Schema for the chowkis API.
type Chowki struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChowkiSpec   `json:"spec,omitempty"`
	Status ChowkiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ChowkiList contains a list of Chowki.
type ChowkiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Chowki `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Chowki{}, &ChowkiList{})
}
