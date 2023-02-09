// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package scaling_policy

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.PolicyName, b.ko.Spec.PolicyName) {
		delta.Add("Spec.PolicyName", a.ko.Spec.PolicyName, b.ko.Spec.PolicyName)
	} else if a.ko.Spec.PolicyName != nil && b.ko.Spec.PolicyName != nil {
		if *a.ko.Spec.PolicyName != *b.ko.Spec.PolicyName {
			delta.Add("Spec.PolicyName", a.ko.Spec.PolicyName, b.ko.Spec.PolicyName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PolicyType, b.ko.Spec.PolicyType) {
		delta.Add("Spec.PolicyType", a.ko.Spec.PolicyType, b.ko.Spec.PolicyType)
	} else if a.ko.Spec.PolicyType != nil && b.ko.Spec.PolicyType != nil {
		if *a.ko.Spec.PolicyType != *b.ko.Spec.PolicyType {
			delta.Add("Spec.PolicyType", a.ko.Spec.PolicyType, b.ko.Spec.PolicyType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ResourceID, b.ko.Spec.ResourceID) {
		delta.Add("Spec.ResourceID", a.ko.Spec.ResourceID, b.ko.Spec.ResourceID)
	} else if a.ko.Spec.ResourceID != nil && b.ko.Spec.ResourceID != nil {
		if *a.ko.Spec.ResourceID != *b.ko.Spec.ResourceID {
			delta.Add("Spec.ResourceID", a.ko.Spec.ResourceID, b.ko.Spec.ResourceID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ScalableDimension, b.ko.Spec.ScalableDimension) {
		delta.Add("Spec.ScalableDimension", a.ko.Spec.ScalableDimension, b.ko.Spec.ScalableDimension)
	} else if a.ko.Spec.ScalableDimension != nil && b.ko.Spec.ScalableDimension != nil {
		if *a.ko.Spec.ScalableDimension != *b.ko.Spec.ScalableDimension {
			delta.Add("Spec.ScalableDimension", a.ko.Spec.ScalableDimension, b.ko.Spec.ScalableDimension)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ServiceNamespace, b.ko.Spec.ServiceNamespace) {
		delta.Add("Spec.ServiceNamespace", a.ko.Spec.ServiceNamespace, b.ko.Spec.ServiceNamespace)
	} else if a.ko.Spec.ServiceNamespace != nil && b.ko.Spec.ServiceNamespace != nil {
		if *a.ko.Spec.ServiceNamespace != *b.ko.Spec.ServiceNamespace {
			delta.Add("Spec.ServiceNamespace", a.ko.Spec.ServiceNamespace, b.ko.Spec.ServiceNamespace)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.StepScalingPolicyConfiguration, b.ko.Spec.StepScalingPolicyConfiguration) {
		delta.Add("Spec.StepScalingPolicyConfiguration", a.ko.Spec.StepScalingPolicyConfiguration, b.ko.Spec.StepScalingPolicyConfiguration)
	} else if a.ko.Spec.StepScalingPolicyConfiguration != nil && b.ko.Spec.StepScalingPolicyConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType, b.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType) {
			delta.Add("Spec.StepScalingPolicyConfiguration.AdjustmentType", a.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType, b.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType)
		} else if a.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType != nil && b.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType != nil {
			if *a.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType != *b.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType {
				delta.Add("Spec.StepScalingPolicyConfiguration.AdjustmentType", a.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType, b.ko.Spec.StepScalingPolicyConfiguration.AdjustmentType)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.StepScalingPolicyConfiguration.Cooldown, b.ko.Spec.StepScalingPolicyConfiguration.Cooldown) {
			delta.Add("Spec.StepScalingPolicyConfiguration.Cooldown", a.ko.Spec.StepScalingPolicyConfiguration.Cooldown, b.ko.Spec.StepScalingPolicyConfiguration.Cooldown)
		} else if a.ko.Spec.StepScalingPolicyConfiguration.Cooldown != nil && b.ko.Spec.StepScalingPolicyConfiguration.Cooldown != nil {
			if *a.ko.Spec.StepScalingPolicyConfiguration.Cooldown != *b.ko.Spec.StepScalingPolicyConfiguration.Cooldown {
				delta.Add("Spec.StepScalingPolicyConfiguration.Cooldown", a.ko.Spec.StepScalingPolicyConfiguration.Cooldown, b.ko.Spec.StepScalingPolicyConfiguration.Cooldown)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType, b.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType) {
			delta.Add("Spec.StepScalingPolicyConfiguration.MetricAggregationType", a.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType, b.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType)
		} else if a.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType != nil && b.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType != nil {
			if *a.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType != *b.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType {
				delta.Add("Spec.StepScalingPolicyConfiguration.MetricAggregationType", a.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType, b.ko.Spec.StepScalingPolicyConfiguration.MetricAggregationType)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude, b.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude) {
			delta.Add("Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude", a.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude, b.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude)
		} else if a.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude != nil && b.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude != nil {
			if *a.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude != *b.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude {
				delta.Add("Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude", a.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude, b.ko.Spec.StepScalingPolicyConfiguration.MinAdjustmentMagnitude)
			}
		}
		if !reflect.DeepEqual(a.ko.Spec.StepScalingPolicyConfiguration.StepAdjustments, b.ko.Spec.StepScalingPolicyConfiguration.StepAdjustments) {
			delta.Add("Spec.StepScalingPolicyConfiguration.StepAdjustments", a.ko.Spec.StepScalingPolicyConfiguration.StepAdjustments, b.ko.Spec.StepScalingPolicyConfiguration.StepAdjustments)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration, b.ko.Spec.TargetTrackingScalingPolicyConfiguration) {
		delta.Add("Spec.TargetTrackingScalingPolicyConfiguration", a.ko.Spec.TargetTrackingScalingPolicyConfiguration, b.ko.Spec.TargetTrackingScalingPolicyConfiguration)
	} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification) {
			delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification)
		} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification != nil {
			if !reflect.DeepEqual(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions) {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Dimensions)
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName) {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName)
			} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName != nil {
				if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName {
					delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.MetricName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace) {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace)
			} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace != nil {
				if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace {
					delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Namespace)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic) {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic)
			} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic != nil {
				if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic {
					delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Statistic)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit) {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit)
			} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit != nil {
				if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit {
					delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.CustomizedMetricSpecification.Unit)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn) {
			delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn)
		} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn != nil {
			if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.DisableScaleIn)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification) {
			delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification)
		} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType) {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType)
			} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType != nil {
				if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType {
					delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.PredefinedMetricType)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel) {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel)
			} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel != nil {
				if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel {
					delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.PredefinedMetricSpecification.ResourceLabel)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown) {
			delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown)
		} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown != nil {
			if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleInCooldown)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown) {
			delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown)
		} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown != nil {
			if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.ScaleOutCooldown)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue) {
			delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.TargetValue", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue)
		} else if a.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue != nil && b.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue != nil {
			if *a.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue != *b.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue {
				delta.Add("Spec.TargetTrackingScalingPolicyConfiguration.TargetValue", a.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue, b.ko.Spec.TargetTrackingScalingPolicyConfiguration.TargetValue)
			}
		}
	}

	return delta
}
