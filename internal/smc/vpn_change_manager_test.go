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

package smc

import (
	"testing"
)

func TestIsVpnUrl(t *testing.T) {
	url := "http://localhost:8082/7.4/elements/vpn/268438672/gateway_tree_nodes/central"
	prefix := isVpnUrl(url)
	if prefix != "http://localhost:8082/7.4/elements/vpn/268438672" {
		t.Errorf("expected vpn url prefix, got: %s", prefix)
	}

	prefix = isVpnUrl("http://localhost:8082/7.4/elements/single_fw/268571118")
	if prefix != "" {
		t.Errorf("expected empty string for non-vpn url, got: %s", prefix)
	}
}
