// +build !ignore_autogenerated

/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigClaim) DeepCopyInto(out *ConfigClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigClaim.
func (in *ConfigClaim) DeepCopy() *ConfigClaim {
	if in == nil {
		return nil
	}
	out := new(ConfigClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigClaimList) DeepCopyInto(out *ConfigClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigClaimList.
func (in *ConfigClaimList) DeepCopy() *ConfigClaimList {
	if in == nil {
		return nil
	}
	out := new(ConfigClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigClaimStatus) DeepCopyInto(out *ConfigClaimStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigClaimStatus.
func (in *ConfigClaimStatus) DeepCopy() *ConfigClaimStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigReference) DeepCopyInto(out *ConfigReference) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigReference.
func (in *ConfigReference) DeepCopy() *ConfigReference {
	if in == nil {
		return nil
	}
	out := new(ConfigReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigReference) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigReferenceList) DeepCopyInto(out *ConfigReferenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigReferenceList.
func (in *ConfigReferenceList) DeepCopy() *ConfigReferenceList {
	if in == nil {
		return nil
	}
	out := new(ConfigReferenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigReferenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigReferenceStatus) DeepCopyInto(out *ConfigReferenceStatus) {
	*out = *in
	if in.Refs != nil {
		in, out := &in.Refs, &out.Refs
		*out = make([]*Reference, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Reference)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigReferenceStatus.
func (in *ConfigReferenceStatus) DeepCopy() *ConfigReferenceStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigReferenceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlledBy) DeepCopyInto(out *ControlledBy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlledBy.
func (in *ControlledBy) DeepCopy() *ControlledBy {
	if in == nil {
		return nil
	}
	out := new(ControlledBy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraInfo) DeepCopyInto(out *ExtraInfo) {
	*out = *in
	in.ReleaseExtraInfo.DeepCopyInto(&out.ReleaseExtraInfo)
	out.IngressExtraInfo = in.IngressExtraInfo
	out.WorkloadExtraInfo = in.WorkloadExtraInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraInfo.
func (in *ExtraInfo) DeepCopy() *ExtraInfo {
	if in == nil {
		return nil
	}
	out := new(ExtraInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressExtraInfo) DeepCopyInto(out *IngressExtraInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressExtraInfo.
func (in *IngressExtraInfo) DeepCopy() *IngressExtraInfo {
	if in == nil {
		return nil
	}
	out := new(IngressExtraInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reference) DeepCopyInto(out *Reference) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]*Reference, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Reference)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.ExtraInfo.DeepCopyInto(&out.ExtraInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reference.
func (in *Reference) DeepCopy() *Reference {
	if in == nil {
		return nil
	}
	out := new(Reference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseExtraInfo) DeepCopyInto(out *ReleaseExtraInfo) {
	*out = *in
	if in.ControlledBy != nil {
		in, out := &in.ControlledBy, &out.ControlledBy
		*out = new(ControlledBy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseExtraInfo.
func (in *ReleaseExtraInfo) DeepCopy() *ReleaseExtraInfo {
	if in == nil {
		return nil
	}
	out := new(ReleaseExtraInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadExtraInfo) DeepCopyInto(out *WorkloadExtraInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadExtraInfo.
func (in *WorkloadExtraInfo) DeepCopy() *WorkloadExtraInfo {
	if in == nil {
		return nil
	}
	out := new(WorkloadExtraInfo)
	in.DeepCopyInto(out)
	return out
}
