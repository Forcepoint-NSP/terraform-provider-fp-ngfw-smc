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
	"testing"

	"github.com/terraform-providers/terraform-provider-smc/internal/smcvalidator"
)

func TestIPv4Network_valid(t *testing.T) {
	for _, val := range []string{
		"192.168.0.0/24", "10.0.0.0/8", "0.0.0.0/0",
		// leading zeros in octets must be accepted (normalised before parsing)
		"192.168.100.00/24", "010.0.0.0/8", "192.168.001.0/24",
	} {
		resp := validateString(smcvalidator.IPv4Network(), val)
		if resp.Diagnostics.HasError() {
			t.Errorf("IPv4Network(%q): expected no error, got: %s", val, resp.Diagnostics)
		}
	}
}

func TestIPv4Network_invalid(t *testing.T) {
	for _, val := range []string{"192.168.1.1", "2001:db8::/32", "not-a-cidr", "192.168.0.0/33"} {
		resp := validateString(smcvalidator.IPv4Network(), val)
		if !resp.Diagnostics.HasError() {
			t.Errorf("IPv4Network(%q): expected error but got none", val)
		}
	}
}

func TestIPv6Network_valid(t *testing.T) {
	for _, val := range []string{"2001:db8::/32", "fe80::/10", "::/0"} {
		resp := validateString(smcvalidator.IPv6Network(), val)
		if resp.Diagnostics.HasError() {
			t.Errorf("IPv6Network(%q): expected no error, got: %s", val, resp.Diagnostics)
		}
	}
}

func TestIPv6Network_invalid(t *testing.T) {
	for _, val := range []string{"192.168.0.0/24", "not-a-cidr", "2001:db8::1"} {
		resp := validateString(smcvalidator.IPv6Network(), val)
		if !resp.Diagnostics.HasError() {
			t.Errorf("IPv6Network(%q): expected error but got none", val)
		}
	}
}
