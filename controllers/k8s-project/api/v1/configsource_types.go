/*

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
	sources "k8s-project/pkg/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	MaxFailures = 2
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConfigSourceSpec defines the desired state of ConfigSource
type ConfigSourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +comment Source information
	SourceType sources.SourceType `json:"sourceType"`

	// +comment Source Configuration
	SourceConfig sources.SourceConfig `json:"sourceConfig"`
}

// ConfigSourceStatus defines the observed state of ConfigSource
type ConfigSourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +comment Object sync status
	// +optional
	Synced bool `json:"synced"`

	// +comment Number of retrys attempted
	// +optional
	FailureCount int64 `json:"failureCount"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ConfigSource is the Schema for the configsources API
type ConfigSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSourceSpec   `json:"spec,omitempty"`
	Status ConfigSourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigSourceList contains a list of ConfigSource
type ConfigSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigSource{}, &ConfigSourceList{})
}
