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

// ConfigMapPrinterSpec defines the desired state of ConfigMapPrinter
type ConfigMapPrinterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +comment Configmap Data
	Information map[string]string `json:"information,omitempty" protobuf:"bytes,2,rep,name=data"`

	// +comment Whether the parent should hold the source of truth, or let the children evolved independently
	Managed bool `json:"managed"`

}

// ConfigMapPrinterStatus defines the observed state of ConfigMapPrinter
type ConfigMapPrinterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	Revision int64 `json:"revision"`
}

// +kubebuilder:object:root=true

// ConfigMapPrinter is the Schema for the configmapprinters API
type ConfigMapPrinter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigMapPrinterSpec   `json:"spec,omitempty"`
	Status ConfigMapPrinterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigMapPrinterList contains a list of ConfigMapPrinter
type ConfigMapPrinterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigMapPrinter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigMapPrinter{}, &ConfigMapPrinterList{})
}
