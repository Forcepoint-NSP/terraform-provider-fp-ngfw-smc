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

package schema

import (
	"github.com/terraform-providers/terraform-provider-smc/internal/apijson"
)

func (r *SingleFirewallResource) ModelToJson(data *SingleFirewallResourceModel) ([]byte, error) {
	// the API expects an empty array instead of null for
	// WebAuthentication.EnabledInterface (error 400, parameter1 is null)
	if data.WebAuthentication != nil && data.WebAuthentication.EnabledInterface == nil {
		data.WebAuthentication.EnabledInterface = &[]EnabledInterfaceEntryResourceModel{}
	}
	return apijson.MarshalRoot(*data)
}
