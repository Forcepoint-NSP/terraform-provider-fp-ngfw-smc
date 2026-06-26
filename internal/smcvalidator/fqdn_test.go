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

package smcvalidator_test

import (
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-smc/internal/smcvalidator"
)

func TestFQDN_valid(t *testing.T) {
	for _, val := range []string{
		"example.com",
		"sub.example.com",
		"my-host.example.org",
		"a.b.c.d.example.com",
	} {
		resp := validateString(smcvalidator.FQDN(), val)
		if resp.Diagnostics.HasError() {
			t.Errorf("FQDN(%q): expected no error, got: %s", val, resp.Diagnostics)
		}
	}
}

func TestFQDN_invalid(t *testing.T) {
	for _, val := range []string{
		"not-a-fqdn",
		"192.168.1.1",
		"-bad.example.com",
		"bad-.example.com",
		strings.Repeat("a", 254) + ".com",
	} {
		resp := validateString(smcvalidator.FQDN(), val)
		if !resp.Diagnostics.HasError() {
			t.Errorf("FQDN(%q): expected error but got none", val)
		}
	}
}
