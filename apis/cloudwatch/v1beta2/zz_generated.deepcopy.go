//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionsSuppressorInitParameters) DeepCopyInto(out *ActionsSuppressorInitParameters) {
	*out = *in
	if in.Alarm != nil {
		in, out := &in.Alarm, &out.Alarm
		*out = new(string)
		**out = **in
	}
	if in.ExtensionPeriod != nil {
		in, out := &in.ExtensionPeriod, &out.ExtensionPeriod
		*out = new(float64)
		**out = **in
	}
	if in.WaitPeriod != nil {
		in, out := &in.WaitPeriod, &out.WaitPeriod
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionsSuppressorInitParameters.
func (in *ActionsSuppressorInitParameters) DeepCopy() *ActionsSuppressorInitParameters {
	if in == nil {
		return nil
	}
	out := new(ActionsSuppressorInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionsSuppressorObservation) DeepCopyInto(out *ActionsSuppressorObservation) {
	*out = *in
	if in.Alarm != nil {
		in, out := &in.Alarm, &out.Alarm
		*out = new(string)
		**out = **in
	}
	if in.ExtensionPeriod != nil {
		in, out := &in.ExtensionPeriod, &out.ExtensionPeriod
		*out = new(float64)
		**out = **in
	}
	if in.WaitPeriod != nil {
		in, out := &in.WaitPeriod, &out.WaitPeriod
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionsSuppressorObservation.
func (in *ActionsSuppressorObservation) DeepCopy() *ActionsSuppressorObservation {
	if in == nil {
		return nil
	}
	out := new(ActionsSuppressorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionsSuppressorParameters) DeepCopyInto(out *ActionsSuppressorParameters) {
	*out = *in
	if in.Alarm != nil {
		in, out := &in.Alarm, &out.Alarm
		*out = new(string)
		**out = **in
	}
	if in.ExtensionPeriod != nil {
		in, out := &in.ExtensionPeriod, &out.ExtensionPeriod
		*out = new(float64)
		**out = **in
	}
	if in.WaitPeriod != nil {
		in, out := &in.WaitPeriod, &out.WaitPeriod
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionsSuppressorParameters.
func (in *ActionsSuppressorParameters) DeepCopy() *ActionsSuppressorParameters {
	if in == nil {
		return nil
	}
	out := new(ActionsSuppressorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeAlarm) DeepCopyInto(out *CompositeAlarm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeAlarm.
func (in *CompositeAlarm) DeepCopy() *CompositeAlarm {
	if in == nil {
		return nil
	}
	out := new(CompositeAlarm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompositeAlarm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeAlarmInitParameters) DeepCopyInto(out *CompositeAlarmInitParameters) {
	*out = *in
	if in.ActionsEnabled != nil {
		in, out := &in.ActionsEnabled, &out.ActionsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ActionsSuppressor != nil {
		in, out := &in.ActionsSuppressor, &out.ActionsSuppressor
		*out = new(ActionsSuppressorInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.AlarmActions != nil {
		in, out := &in.AlarmActions, &out.AlarmActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AlarmActionsRefs != nil {
		in, out := &in.AlarmActionsRefs, &out.AlarmActionsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AlarmActionsSelector != nil {
		in, out := &in.AlarmActionsSelector, &out.AlarmActionsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.AlarmDescription != nil {
		in, out := &in.AlarmDescription, &out.AlarmDescription
		*out = new(string)
		**out = **in
	}
	if in.AlarmRule != nil {
		in, out := &in.AlarmRule, &out.AlarmRule
		*out = new(string)
		**out = **in
	}
	if in.InsufficientDataActions != nil {
		in, out := &in.InsufficientDataActions, &out.InsufficientDataActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OkActions != nil {
		in, out := &in.OkActions, &out.OkActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OkActionsRefs != nil {
		in, out := &in.OkActionsRefs, &out.OkActionsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OkActionsSelector != nil {
		in, out := &in.OkActionsSelector, &out.OkActionsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeAlarmInitParameters.
func (in *CompositeAlarmInitParameters) DeepCopy() *CompositeAlarmInitParameters {
	if in == nil {
		return nil
	}
	out := new(CompositeAlarmInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeAlarmList) DeepCopyInto(out *CompositeAlarmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CompositeAlarm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeAlarmList.
func (in *CompositeAlarmList) DeepCopy() *CompositeAlarmList {
	if in == nil {
		return nil
	}
	out := new(CompositeAlarmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompositeAlarmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeAlarmObservation) DeepCopyInto(out *CompositeAlarmObservation) {
	*out = *in
	if in.ActionsEnabled != nil {
		in, out := &in.ActionsEnabled, &out.ActionsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ActionsSuppressor != nil {
		in, out := &in.ActionsSuppressor, &out.ActionsSuppressor
		*out = new(ActionsSuppressorObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.AlarmActions != nil {
		in, out := &in.AlarmActions, &out.AlarmActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AlarmDescription != nil {
		in, out := &in.AlarmDescription, &out.AlarmDescription
		*out = new(string)
		**out = **in
	}
	if in.AlarmRule != nil {
		in, out := &in.AlarmRule, &out.AlarmRule
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InsufficientDataActions != nil {
		in, out := &in.InsufficientDataActions, &out.InsufficientDataActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OkActions != nil {
		in, out := &in.OkActions, &out.OkActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeAlarmObservation.
func (in *CompositeAlarmObservation) DeepCopy() *CompositeAlarmObservation {
	if in == nil {
		return nil
	}
	out := new(CompositeAlarmObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeAlarmParameters) DeepCopyInto(out *CompositeAlarmParameters) {
	*out = *in
	if in.ActionsEnabled != nil {
		in, out := &in.ActionsEnabled, &out.ActionsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ActionsSuppressor != nil {
		in, out := &in.ActionsSuppressor, &out.ActionsSuppressor
		*out = new(ActionsSuppressorParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.AlarmActions != nil {
		in, out := &in.AlarmActions, &out.AlarmActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AlarmActionsRefs != nil {
		in, out := &in.AlarmActionsRefs, &out.AlarmActionsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AlarmActionsSelector != nil {
		in, out := &in.AlarmActionsSelector, &out.AlarmActionsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.AlarmDescription != nil {
		in, out := &in.AlarmDescription, &out.AlarmDescription
		*out = new(string)
		**out = **in
	}
	if in.AlarmRule != nil {
		in, out := &in.AlarmRule, &out.AlarmRule
		*out = new(string)
		**out = **in
	}
	if in.InsufficientDataActions != nil {
		in, out := &in.InsufficientDataActions, &out.InsufficientDataActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OkActions != nil {
		in, out := &in.OkActions, &out.OkActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OkActionsRefs != nil {
		in, out := &in.OkActionsRefs, &out.OkActionsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OkActionsSelector != nil {
		in, out := &in.OkActionsSelector, &out.OkActionsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeAlarmParameters.
func (in *CompositeAlarmParameters) DeepCopy() *CompositeAlarmParameters {
	if in == nil {
		return nil
	}
	out := new(CompositeAlarmParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeAlarmSpec) DeepCopyInto(out *CompositeAlarmSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeAlarmSpec.
func (in *CompositeAlarmSpec) DeepCopy() *CompositeAlarmSpec {
	if in == nil {
		return nil
	}
	out := new(CompositeAlarmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeAlarmStatus) DeepCopyInto(out *CompositeAlarmStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeAlarmStatus.
func (in *CompositeAlarmStatus) DeepCopy() *CompositeAlarmStatus {
	if in == nil {
		return nil
	}
	out := new(CompositeAlarmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricAlarm) DeepCopyInto(out *MetricAlarm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricAlarm.
func (in *MetricAlarm) DeepCopy() *MetricAlarm {
	if in == nil {
		return nil
	}
	out := new(MetricAlarm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricAlarm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricAlarmInitParameters) DeepCopyInto(out *MetricAlarmInitParameters) {
	*out = *in
	if in.ActionsEnabled != nil {
		in, out := &in.ActionsEnabled, &out.ActionsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AlarmActions != nil {
		in, out := &in.AlarmActions, &out.AlarmActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AlarmDescription != nil {
		in, out := &in.AlarmDescription, &out.AlarmDescription
		*out = new(string)
		**out = **in
	}
	if in.ComparisonOperator != nil {
		in, out := &in.ComparisonOperator, &out.ComparisonOperator
		*out = new(string)
		**out = **in
	}
	if in.DatapointsToAlarm != nil {
		in, out := &in.DatapointsToAlarm, &out.DatapointsToAlarm
		*out = new(float64)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
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
	if in.EvaluateLowSampleCountPercentiles != nil {
		in, out := &in.EvaluateLowSampleCountPercentiles, &out.EvaluateLowSampleCountPercentiles
		*out = new(string)
		**out = **in
	}
	if in.EvaluationPeriods != nil {
		in, out := &in.EvaluationPeriods, &out.EvaluationPeriods
		*out = new(float64)
		**out = **in
	}
	if in.ExtendedStatistic != nil {
		in, out := &in.ExtendedStatistic, &out.ExtendedStatistic
		*out = new(string)
		**out = **in
	}
	if in.InsufficientDataActions != nil {
		in, out := &in.InsufficientDataActions, &out.InsufficientDataActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.MetricQuery != nil {
		in, out := &in.MetricQuery, &out.MetricQuery
		*out = make([]MetricQueryInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.OkActions != nil {
		in, out := &in.OkActions, &out.OkActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
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
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdMetricID != nil {
		in, out := &in.ThresholdMetricID, &out.ThresholdMetricID
		*out = new(string)
		**out = **in
	}
	if in.TreatMissingData != nil {
		in, out := &in.TreatMissingData, &out.TreatMissingData
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricAlarmInitParameters.
func (in *MetricAlarmInitParameters) DeepCopy() *MetricAlarmInitParameters {
	if in == nil {
		return nil
	}
	out := new(MetricAlarmInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricAlarmList) DeepCopyInto(out *MetricAlarmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricAlarm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricAlarmList.
func (in *MetricAlarmList) DeepCopy() *MetricAlarmList {
	if in == nil {
		return nil
	}
	out := new(MetricAlarmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricAlarmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricAlarmObservation) DeepCopyInto(out *MetricAlarmObservation) {
	*out = *in
	if in.ActionsEnabled != nil {
		in, out := &in.ActionsEnabled, &out.ActionsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AlarmActions != nil {
		in, out := &in.AlarmActions, &out.AlarmActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AlarmDescription != nil {
		in, out := &in.AlarmDescription, &out.AlarmDescription
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ComparisonOperator != nil {
		in, out := &in.ComparisonOperator, &out.ComparisonOperator
		*out = new(string)
		**out = **in
	}
	if in.DatapointsToAlarm != nil {
		in, out := &in.DatapointsToAlarm, &out.DatapointsToAlarm
		*out = new(float64)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
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
	if in.EvaluateLowSampleCountPercentiles != nil {
		in, out := &in.EvaluateLowSampleCountPercentiles, &out.EvaluateLowSampleCountPercentiles
		*out = new(string)
		**out = **in
	}
	if in.EvaluationPeriods != nil {
		in, out := &in.EvaluationPeriods, &out.EvaluationPeriods
		*out = new(float64)
		**out = **in
	}
	if in.ExtendedStatistic != nil {
		in, out := &in.ExtendedStatistic, &out.ExtendedStatistic
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InsufficientDataActions != nil {
		in, out := &in.InsufficientDataActions, &out.InsufficientDataActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.MetricQuery != nil {
		in, out := &in.MetricQuery, &out.MetricQuery
		*out = make([]MetricQueryObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.OkActions != nil {
		in, out := &in.OkActions, &out.OkActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
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
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdMetricID != nil {
		in, out := &in.ThresholdMetricID, &out.ThresholdMetricID
		*out = new(string)
		**out = **in
	}
	if in.TreatMissingData != nil {
		in, out := &in.TreatMissingData, &out.TreatMissingData
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricAlarmObservation.
func (in *MetricAlarmObservation) DeepCopy() *MetricAlarmObservation {
	if in == nil {
		return nil
	}
	out := new(MetricAlarmObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricAlarmParameters) DeepCopyInto(out *MetricAlarmParameters) {
	*out = *in
	if in.ActionsEnabled != nil {
		in, out := &in.ActionsEnabled, &out.ActionsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AlarmActions != nil {
		in, out := &in.AlarmActions, &out.AlarmActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AlarmDescription != nil {
		in, out := &in.AlarmDescription, &out.AlarmDescription
		*out = new(string)
		**out = **in
	}
	if in.ComparisonOperator != nil {
		in, out := &in.ComparisonOperator, &out.ComparisonOperator
		*out = new(string)
		**out = **in
	}
	if in.DatapointsToAlarm != nil {
		in, out := &in.DatapointsToAlarm, &out.DatapointsToAlarm
		*out = new(float64)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
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
	if in.EvaluateLowSampleCountPercentiles != nil {
		in, out := &in.EvaluateLowSampleCountPercentiles, &out.EvaluateLowSampleCountPercentiles
		*out = new(string)
		**out = **in
	}
	if in.EvaluationPeriods != nil {
		in, out := &in.EvaluationPeriods, &out.EvaluationPeriods
		*out = new(float64)
		**out = **in
	}
	if in.ExtendedStatistic != nil {
		in, out := &in.ExtendedStatistic, &out.ExtendedStatistic
		*out = new(string)
		**out = **in
	}
	if in.InsufficientDataActions != nil {
		in, out := &in.InsufficientDataActions, &out.InsufficientDataActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.MetricQuery != nil {
		in, out := &in.MetricQuery, &out.MetricQuery
		*out = make([]MetricQueryParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.OkActions != nil {
		in, out := &in.OkActions, &out.OkActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
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
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdMetricID != nil {
		in, out := &in.ThresholdMetricID, &out.ThresholdMetricID
		*out = new(string)
		**out = **in
	}
	if in.TreatMissingData != nil {
		in, out := &in.TreatMissingData, &out.TreatMissingData
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricAlarmParameters.
func (in *MetricAlarmParameters) DeepCopy() *MetricAlarmParameters {
	if in == nil {
		return nil
	}
	out := new(MetricAlarmParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricAlarmSpec) DeepCopyInto(out *MetricAlarmSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricAlarmSpec.
func (in *MetricAlarmSpec) DeepCopy() *MetricAlarmSpec {
	if in == nil {
		return nil
	}
	out := new(MetricAlarmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricAlarmStatus) DeepCopyInto(out *MetricAlarmStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricAlarmStatus.
func (in *MetricAlarmStatus) DeepCopy() *MetricAlarmStatus {
	if in == nil {
		return nil
	}
	out := new(MetricAlarmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricInitParameters) DeepCopyInto(out *MetricInitParameters) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
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
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.Stat != nil {
		in, out := &in.Stat, &out.Stat
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricInitParameters.
func (in *MetricInitParameters) DeepCopy() *MetricInitParameters {
	if in == nil {
		return nil
	}
	out := new(MetricInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricObservation) DeepCopyInto(out *MetricObservation) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
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
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.Stat != nil {
		in, out := &in.Stat, &out.Stat
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricObservation.
func (in *MetricObservation) DeepCopy() *MetricObservation {
	if in == nil {
		return nil
	}
	out := new(MetricObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricParameters) DeepCopyInto(out *MetricParameters) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
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
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.Stat != nil {
		in, out := &in.Stat, &out.Stat
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricParameters.
func (in *MetricParameters) DeepCopy() *MetricParameters {
	if in == nil {
		return nil
	}
	out := new(MetricParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQueryInitParameters) DeepCopyInto(out *MetricQueryInitParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(MetricInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.ReturnData != nil {
		in, out := &in.ReturnData, &out.ReturnData
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQueryInitParameters.
func (in *MetricQueryInitParameters) DeepCopy() *MetricQueryInitParameters {
	if in == nil {
		return nil
	}
	out := new(MetricQueryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQueryObservation) DeepCopyInto(out *MetricQueryObservation) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(MetricObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.ReturnData != nil {
		in, out := &in.ReturnData, &out.ReturnData
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQueryObservation.
func (in *MetricQueryObservation) DeepCopy() *MetricQueryObservation {
	if in == nil {
		return nil
	}
	out := new(MetricQueryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricQueryParameters) DeepCopyInto(out *MetricQueryParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(MetricParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(float64)
		**out = **in
	}
	if in.ReturnData != nil {
		in, out := &in.ReturnData, &out.ReturnData
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricQueryParameters.
func (in *MetricQueryParameters) DeepCopy() *MetricQueryParameters {
	if in == nil {
		return nil
	}
	out := new(MetricQueryParameters)
	in.DeepCopyInto(out)
	return out
}
