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

// ZeebeClientSpec defines the desired state of ZeebeClient
type ZeebeClientSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ClientName string `json:"clientName"`
	ClusterId  string `json:"clusterId"`
	// +kubebuilder:validation:Optional
	ClientId string `json:"clientId"`
	// +kubebuilder:validation:Optional
	SecretName string `json:"secretName"`
	// +kubebuilder:validation:Optional
	ConfigMapName string `json:"configMapName"`
}

// ZeebeClientStatus defines the observed state of ZeebeClient
type ZeebeClientStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// ZeebeClient is the Schema for the zeebeclients API
type ZeebeClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZeebeClientSpec   `json:"spec,omitempty"`
	Status ZeebeClientStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZeebeClientList contains a list of ZeebeClient
type ZeebeClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZeebeClient `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZeebeClient{}, &ZeebeClientList{})
}
