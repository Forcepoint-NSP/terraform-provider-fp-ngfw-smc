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

package smcvalidator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// fqdnRegex: dot-separated labels; each label starts/ends with alphanumeric, may contain hyphens.
var fqdnRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

func FQDN() validator.String { return fqdnValidator{} }

type fqdnValidator struct{}

func (v fqdnValidator) Description(_ context.Context) string {
	return "value must be a valid fully qualified domain name"
}
func (v fqdnValidator) MarkdownDescription(_ context.Context) string {
	return "value must be a valid fully qualified domain name"
}
func (v fqdnValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	if len(val) > 253 || !fqdnRegex.MatchString(val) {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid FQDN",
			"Expected a fully qualified domain name (e.g. host.example.com), got: "+val)
	}
}
