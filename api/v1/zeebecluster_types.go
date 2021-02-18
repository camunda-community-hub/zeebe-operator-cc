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
	cc "github.com/salaboy/camunda-cloud-go-client/pkg/cc/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ZeebeClusterSpec defines the desired state of ZeebeCluster
type ZeebeClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

}

// ZeebeClusterStatus defines the observed state of ZeebeCluster
type ZeebeClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ClusterId     string           `json:"clusterId"`
	ClusterStatus cc.ClusterStatus `json:"clusterStatus"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// ZeebeCluster is the Schema for the zeebeclusters API
type ZeebeCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZeebeClusterSpec   `json:"spec,omitempty"`
	Status ZeebeClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZeebeClusterList contains a list of ZeebeCluster
type ZeebeClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZeebeCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZeebeCluster{}, &ZeebeClusterList{})
}
