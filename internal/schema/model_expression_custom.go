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
	"sort"
	"strings"
)

// GetSliceIds identifies Expression elements for list reconciliation
// (mergeSlice/GetElementById, see internal/resource/model_utils.go).
//
// The auto-generated version (formerly in model_expression.go, removed -
// see its DO-NOT-EDIT header) matched only by Name. That's correct for a
// top-level Expression element (always named), but ExpressionResourceModel
// is reused verbatim for the embedded `sub_expression` list, and SMC never
// names an embedded sub-expression (it isn't a standalone element). Name-
// only matching therefore always failed for every sub-expression entry, so
// mergeSlice always skipped them (see its "dest element not found" path),
// leaving every fpro-tagged field (admin_domain/etag/id/key/link/locked/
// read_only/system/system_key/trashed) permanently unknown after apply -
// reported upstream as "Provider returned invalid result object after
// apply ... unknown value for sub_expression[N].*" for any expression using
// sub_expression at all.
//
// Entries are identified by a recursive Operator+NeRef+SubExpression
// signature, never by Name: the plan side may carry a name SMC drops from
// the API response. ne_ref and nested sub_expressions are member sets, so
// both are sorted to tolerate SMC reordering them; nested content is folded
// in so siblings that differ only in their children don't share an id.
// Returns nil (no usable id) when Operator, any NeRef or any nested entry
// is null/unknown.
func (r *ExpressionResourceModel) GetSliceIds(ctx context.Context) []string {
	sig, ok := r.signature()
	if !ok {
		return nil
	}
	return []string{sig}
}

func (r *ExpressionResourceModel) signature() (string, bool) {
	if r.Operator.IsUnknown() || r.Operator.IsNull() {
		return "", false
	}

	var neRefs []string
	if r.NeRef != nil {
		for _, ref := range *r.NeRef {
			if ref.IsUnknown() || ref.IsNull() {
				return "", false
			}
			neRefs = append(neRefs, ref.ValueString())
		}
	}
	sort.Strings(neRefs)

	var subs []string
	if r.SubExpression != nil {
		for i := range *r.SubExpression {
			sub, ok := (*r.SubExpression)[i].signature()
			if !ok {
				return "", false
			}
			subs = append(subs, sub)
		}
	}
	sort.Strings(subs)

	return r.Operator.ValueString() + "|" + strings.Join(neRefs, ",") + "|(" + strings.Join(subs, ";") + ")", true
}
