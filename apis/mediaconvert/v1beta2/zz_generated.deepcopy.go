//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queue) DeepCopyInto(out *Queue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queue.
func (in *Queue) DeepCopy() *Queue {
	if in == nil {
		return nil
	}
	out := new(Queue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Queue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueInitParameters) DeepCopyInto(out *QueueInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PricingPlan != nil {
		in, out := &in.PricingPlan, &out.PricingPlan
		*out = new(string)
		**out = **in
	}
	if in.ReservationPlanSettings != nil {
		in, out := &in.ReservationPlanSettings, &out.ReservationPlanSettings
		*out = new(ReservationPlanSettingsInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueInitParameters.
func (in *QueueInitParameters) DeepCopy() *QueueInitParameters {
	if in == nil {
		return nil
	}
	out := new(QueueInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueList) DeepCopyInto(out *QueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Queue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueList.
func (in *QueueList) DeepCopy() *QueueList {
	if in == nil {
		return nil
	}
	out := new(QueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueObservation) DeepCopyInto(out *QueueObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PricingPlan != nil {
		in, out := &in.PricingPlan, &out.PricingPlan
		*out = new(string)
		**out = **in
	}
	if in.ReservationPlanSettings != nil {
		in, out := &in.ReservationPlanSettings, &out.ReservationPlanSettings
		*out = new(ReservationPlanSettingsObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueObservation.
func (in *QueueObservation) DeepCopy() *QueueObservation {
	if in == nil {
		return nil
	}
	out := new(QueueObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueParameters) DeepCopyInto(out *QueueParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.PricingPlan != nil {
		in, out := &in.PricingPlan, &out.PricingPlan
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ReservationPlanSettings != nil {
		in, out := &in.ReservationPlanSettings, &out.ReservationPlanSettings
		*out = new(ReservationPlanSettingsParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueParameters.
func (in *QueueParameters) DeepCopy() *QueueParameters {
	if in == nil {
		return nil
	}
	out := new(QueueParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSpec) DeepCopyInto(out *QueueSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSpec.
func (in *QueueSpec) DeepCopy() *QueueSpec {
	if in == nil {
		return nil
	}
	out := new(QueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStatus) DeepCopyInto(out *QueueStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStatus.
func (in *QueueStatus) DeepCopy() *QueueStatus {
	if in == nil {
		return nil
	}
	out := new(QueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservationPlanSettingsInitParameters) DeepCopyInto(out *ReservationPlanSettingsInitParameters) {
	*out = *in
	if in.Commitment != nil {
		in, out := &in.Commitment, &out.Commitment
		*out = new(string)
		**out = **in
	}
	if in.RenewalType != nil {
		in, out := &in.RenewalType, &out.RenewalType
		*out = new(string)
		**out = **in
	}
	if in.ReservedSlots != nil {
		in, out := &in.ReservedSlots, &out.ReservedSlots
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservationPlanSettingsInitParameters.
func (in *ReservationPlanSettingsInitParameters) DeepCopy() *ReservationPlanSettingsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReservationPlanSettingsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservationPlanSettingsObservation) DeepCopyInto(out *ReservationPlanSettingsObservation) {
	*out = *in
	if in.Commitment != nil {
		in, out := &in.Commitment, &out.Commitment
		*out = new(string)
		**out = **in
	}
	if in.RenewalType != nil {
		in, out := &in.RenewalType, &out.RenewalType
		*out = new(string)
		**out = **in
	}
	if in.ReservedSlots != nil {
		in, out := &in.ReservedSlots, &out.ReservedSlots
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservationPlanSettingsObservation.
func (in *ReservationPlanSettingsObservation) DeepCopy() *ReservationPlanSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(ReservationPlanSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservationPlanSettingsParameters) DeepCopyInto(out *ReservationPlanSettingsParameters) {
	*out = *in
	if in.Commitment != nil {
		in, out := &in.Commitment, &out.Commitment
		*out = new(string)
		**out = **in
	}
	if in.RenewalType != nil {
		in, out := &in.RenewalType, &out.RenewalType
		*out = new(string)
		**out = **in
	}
	if in.ReservedSlots != nil {
		in, out := &in.ReservedSlots, &out.ReservedSlots
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservationPlanSettingsParameters.
func (in *ReservationPlanSettingsParameters) DeepCopy() *ReservationPlanSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(ReservationPlanSettingsParameters)
	in.DeepCopyInto(out)
	return out
}
