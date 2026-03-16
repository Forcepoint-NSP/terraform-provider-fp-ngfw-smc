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
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (r *SingleNodeInterfaceResourceModel) GetSliceIds(ctx context.Context) []string {
	if !r.Dynamic.IsNull() && !r.Dynamic.IsUnknown() && r.Dynamic.ValueBool() {
		if !r.DynamicIndex.IsNull() && !r.DynamicIndex.IsUnknown() {
			return []string{fmt.Sprintf("%d", r.DynamicIndex.ValueInt64())}
		}

		if !r.DynamicIpv6Index.IsNull() && !r.DynamicIpv6Index.IsUnknown() {
			return []string{fmt.Sprintf("%d", r.DynamicIpv6Index.ValueInt64())}
		}
		tflog.Error(ctx, "Dynamic is set but neither DynamicIndex nor DynamicIpv6Index is set")
		return nil
	}
	if !r.Address.IsNull() && !r.Address.IsUnknown() {
		return []string{r.Address.ValueString()}
	}

	return nil
}
