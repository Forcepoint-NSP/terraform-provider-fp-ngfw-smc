// Copyright 2026 Forcepoint LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// PlanModifiers for a readOnly attribute carry a UseStateForUnknown
// modifier only when drift detection is enabled for the resource: that
// modifier promises the prior state's value is still accurate, and only
// Terraform Read (drift detection) keeps that promise true (see
// abstract_resource.go's Read, gated by the same "drift_detection"
// config flag). Without it, a readOnly attribute the SMC API can
// legitimately recompute on Update (e.g. a normalized href, SMC-64735)
// would carry a stale carried-forward value into the plan and mismatch
// what Update actually returns - "Provider produced inconsistent
// result after apply". When drift detection is disabled, these return
// nil, so the attribute is planned Unknown on every update instead
// (the framework's default for a Computed attribute with no plan
// modifier), exactly matching provider behavior before drift detection
// existed.

func StringPlanModifiers(driftDetectionEnabled bool) []planmodifier.String {
	if !driftDetectionEnabled {
		return nil
	}
	return []planmodifier.String{stringplanmodifier.UseStateForUnknown()}
}

func BoolPlanModifiers(driftDetectionEnabled bool) []planmodifier.Bool {
	if !driftDetectionEnabled {
		return nil
	}
	return []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()}
}

func Int64PlanModifiers(driftDetectionEnabled bool) []planmodifier.Int64 {
	if !driftDetectionEnabled {
		return nil
	}
	return []planmodifier.Int64{int64planmodifier.UseStateForUnknown()}
}

func Float64PlanModifiers(driftDetectionEnabled bool) []planmodifier.Float64 {
	if !driftDetectionEnabled {
		return nil
	}
	return []planmodifier.Float64{float64planmodifier.UseStateForUnknown()}
}

func MapPlanModifiers(driftDetectionEnabled bool) []planmodifier.Map {
	if !driftDetectionEnabled {
		return nil
	}
	return []planmodifier.Map{mapplanmodifier.UseStateForUnknown()}
}

func ListPlanModifiers(driftDetectionEnabled bool) []planmodifier.List {
	if !driftDetectionEnabled {
		return nil
	}
	return []planmodifier.List{listplanmodifier.UseStateForUnknown()}
}
