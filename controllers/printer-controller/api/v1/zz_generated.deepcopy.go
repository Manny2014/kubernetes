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
func (in *ConfigMapPrinter) DeepCopyInto(out *ConfigMapPrinter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapPrinter.
func (in *ConfigMapPrinter) DeepCopy() *ConfigMapPrinter {
	if in == nil {
		return nil
	}
	out := new(ConfigMapPrinter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigMapPrinter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapPrinterList) DeepCopyInto(out *ConfigMapPrinterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigMapPrinter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapPrinterList.
func (in *ConfigMapPrinterList) DeepCopy() *ConfigMapPrinterList {
	if in == nil {
		return nil
	}
	out := new(ConfigMapPrinterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigMapPrinterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapPrinterSpec) DeepCopyInto(out *ConfigMapPrinterSpec) {
	*out = *in
	if in.Information != nil {
		in, out := &in.Information, &out.Information
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapPrinterSpec.
func (in *ConfigMapPrinterSpec) DeepCopy() *ConfigMapPrinterSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigMapPrinterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapPrinterStatus) DeepCopyInto(out *ConfigMapPrinterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapPrinterStatus.
func (in *ConfigMapPrinterStatus) DeepCopy() *ConfigMapPrinterStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigMapPrinterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectPrinter) DeepCopyInto(out *ObjectPrinter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectPrinter.
func (in *ObjectPrinter) DeepCopy() *ObjectPrinter {
	if in == nil {
		return nil
	}
	out := new(ObjectPrinter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectPrinter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectPrinterList) DeepCopyInto(out *ObjectPrinterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectPrinter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectPrinterList.
func (in *ObjectPrinterList) DeepCopy() *ObjectPrinterList {
	if in == nil {
		return nil
	}
	out := new(ObjectPrinterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectPrinterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectPrinterSpec) DeepCopyInto(out *ObjectPrinterSpec) {
	*out = *in
	if in.PrintCount != nil {
		in, out := &in.PrintCount, &out.PrintCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectPrinterSpec.
func (in *ObjectPrinterSpec) DeepCopy() *ObjectPrinterSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectPrinterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectPrinterStatus) DeepCopyInto(out *ObjectPrinterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectPrinterStatus.
func (in *ObjectPrinterStatus) DeepCopy() *ObjectPrinterStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectPrinterStatus)
	in.DeepCopyInto(out)
	return out
}
