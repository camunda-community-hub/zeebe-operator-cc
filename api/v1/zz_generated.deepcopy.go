// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZeebeClient) DeepCopyInto(out *ZeebeClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZeebeClient.
func (in *ZeebeClient) DeepCopy() *ZeebeClient {
	if in == nil {
		return nil
	}
	out := new(ZeebeClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZeebeClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZeebeClientList) DeepCopyInto(out *ZeebeClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ZeebeClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZeebeClientList.
func (in *ZeebeClientList) DeepCopy() *ZeebeClientList {
	if in == nil {
		return nil
	}
	out := new(ZeebeClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZeebeClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZeebeClientSpec) DeepCopyInto(out *ZeebeClientSpec) {
	*out = *in
	out.ZeebeClientDetails = in.ZeebeClientDetails
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZeebeClientSpec.
func (in *ZeebeClientSpec) DeepCopy() *ZeebeClientSpec {
	if in == nil {
		return nil
	}
	out := new(ZeebeClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZeebeClientStatus) DeepCopyInto(out *ZeebeClientStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZeebeClientStatus.
func (in *ZeebeClientStatus) DeepCopy() *ZeebeClientStatus {
	if in == nil {
		return nil
	}
	out := new(ZeebeClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZeebeCluster) DeepCopyInto(out *ZeebeCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZeebeCluster.
func (in *ZeebeCluster) DeepCopy() *ZeebeCluster {
	if in == nil {
		return nil
	}
	out := new(ZeebeCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZeebeCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZeebeClusterList) DeepCopyInto(out *ZeebeClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ZeebeCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZeebeClusterList.
func (in *ZeebeClusterList) DeepCopy() *ZeebeClusterList {
	if in == nil {
		return nil
	}
	out := new(ZeebeClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZeebeClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZeebeClusterSpec) DeepCopyInto(out *ZeebeClusterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZeebeClusterSpec.
func (in *ZeebeClusterSpec) DeepCopy() *ZeebeClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ZeebeClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZeebeClusterStatus) DeepCopyInto(out *ZeebeClusterStatus) {
	*out = *in
	out.ClusterStatus = in.ClusterStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZeebeClusterStatus.
func (in *ZeebeClusterStatus) DeepCopy() *ZeebeClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ZeebeClusterStatus)
	in.DeepCopyInto(out)
	return out
}
