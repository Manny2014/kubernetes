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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ObjectPrinterSpec defines the desired state of ObjectPrinter
type ObjectPrinterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +comment Message to be printed by the resource
	// +kubebuilder:validation:MinLength=0
	Message string `json:"message"`

	// +comment Cron schedule
	// +kubebuilder:validation:MinValue=1
	PrintCount *int64 `json:"printCount"`
}

// ObjectPrinterStatus defines the observed state of ObjectPrinter
type ObjectPrinterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	PrintCount int64 `json:"printCount"`

	// +optional
	FailureCount int64 `json:"failureCount"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ObjectPrinter is the Schema for the objectprinters API
type ObjectPrinter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectPrinterSpec   `json:"spec,omitempty"`
	Status ObjectPrinterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectPrinterList contains a list of ObjectPrinter
type ObjectPrinterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectPrinter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ObjectPrinter{}, &ObjectPrinterList{})
}
