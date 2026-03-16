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
)

func (r *RoutingNodeResourceModel) GetSliceIds(ctx context.Context) []string {
	var ids []string
	if (!r.Href.IsNull()) && (!r.Href.IsUnknown()) {
		ids = append(ids, r.Href.ValueString())
	}
	if (!r.Name.IsNull()) && (!r.Name.IsUnknown()) {
		ids = append(ids, r.Name.ValueString())
	}
	if len(ids) != 0 {
		return ids
	}
	return nil
}
