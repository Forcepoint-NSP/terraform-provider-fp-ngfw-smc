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

// this test should be in package resource with model_drift.go, but we
// need structs from the schema to test and we would have a circular
// dependency (see model_utils_test.go).

package provider

import (
	"context"
	"os"
	"testing"

	"github.com/terraform-providers/terraform-provider-smc/internal/apijson"
	"github.com/terraform-providers/terraform-provider-smc/internal/resource"
	"github.com/terraform-providers/terraform-provider-smc/internal/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
)

// state has address+comment set by the user; the API reports a drifted
// address (should be overlaid) and a comment the user never set
// (should NOT be imported).
func TestMergeResourceDrift_ScalarOverlay(t *testing.T) {
	setupTest()

	var state_json = `{
		"address": "192.168.1.2",
		"name": "AExampleHost"
	}`

	var api_json = `{
		"address": "10.0.0.9",
		"comment": "injected by SMC",
		"link": [
			{
				"href": "http://localhost:8082/7.3/elements/host/5253",
				"rel": "self",
				"type": "host"
			}
		],
		"name": "AExampleHost"
	}`

	var state schema.HostResourceModel
	var api schema.HostResourceModel
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	if err := apijson.UnmarshalRoot([]byte(state_json), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(api_json), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.Address.ValueString() != "10.0.0.9" {
		t.Errorf("expected address drift to be overlaid, got %q", state.Address.ValueString())
	}
	if !state.Comment.IsNull() {
		t.Errorf("expected comment to stay null (never set by user), got %q", state.Comment.ValueString())
	}
	if state.Lk.IsNull() || state.Lk.IsUnknown() {
		t.Errorf("expected Lk to be refreshed by drift merge, but it is null or unknown")
	}
}

// Reproduces a reported false-positive: SMC masks the password field
// in its GET response ("*****") instead of echoing back the real
// value; that placeholder must never be treated as real drift.
func TestMergeResourceDrift_RedactedPasswordNotDrift(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.InternalUserResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "tf_user", "password": "SuperSecretPassword"}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.InternalUserResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "tf_user", "password": "*****"}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.Password.ValueString() != "SuperSecretPassword" {
		t.Errorf("expected the masked API value to be ignored, got %q", state.Password.ValueString())
	}
}

// Reproduces a reported false-positive: SMC returns unique_id as
// "{name}@#@#{value}" even though it's configured and stored as the
// plain value - the name prefix is SMC's internal encoding, not part
// of the attribute's real content.
func TestMergeResourceDrift_CompositeUniqueIdPrefixStripped(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.InternalUserResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "tf_user", "unique_id": "cn=tf_user,dc=stonegate,domain=InternalDomain"}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.InternalUserResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "tf_user", "unique_id": "tf_user@#@#cn=tf_user,dc=stonegate,domain=InternalDomain"}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.UniqueId.ValueString() != "cn=tf_user,dc=stonegate,domain=InternalDomain" {
		t.Errorf("expected the name@#@# prefix to be stripped, got %q", state.UniqueId.ValueString())
	}
}

// A genuine change to unique_id (not just the name-prefix quirk) must
// still be detected as real drift.
func TestMergeResourceDrift_CompositeUniqueIdRealChangeDetected(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.InternalUserResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "tf_user", "unique_id": "cn=tf_user,dc=stonegate,domain=InternalDomain"}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.InternalUserResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "tf_user", "unique_id": "tf_user@#@#cn=someone_else,dc=stonegate,domain=InternalDomain"}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.UniqueId.ValueString() != "cn=someone_else,dc=stonegate,domain=InternalDomain" {
		t.Errorf("expected a genuine unique_id change to be detected, got %q", state.UniqueId.ValueString())
	}
}

// Reproduces a reported false-positive: state has a pointer-to-slice
// attribute ("ip") set by Terraform, but a particular API response
// omits the key entirely (decodes as a nil pointer, not an empty
// list). This must not clear state - and unlike a scalar attribute,
// a pointer field cleared to nil could never recover on a later read.
func TestMergeResourceDrift_SrcNilPointerDoesNotClobberState(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.IpAddressListResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"ip": ["104.209.194.165", "104.209.194.52", "104.209.233.92"]}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.IpAddressListResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}
	if api.Ip != nil {
		t.Fatalf("test setup error: expected api.Ip to be a nil pointer, got %+v", api.Ip)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.Ip == nil {
		t.Fatal("expected ip to survive an API response that omitted it, got nil")
	}
	got := *state.Ip
	if len(got) != 3 || got[0].ValueString() != "104.209.194.165" {
		t.Errorf("expected ip to survive unchanged, got %+v", got)
	}
}

// An explicit empty list from the API (as opposed to the key being
// absent) is a genuine, real change and must still be detected.
func TestMergeResourceDrift_SrcExplicitEmptyListDetected(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.IpAddressListResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"ip": ["104.209.194.165"]}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.IpAddressListResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"ip": []}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}
	if api.Ip == nil {
		t.Fatalf("test setup error: expected api.Ip to be a non-nil (empty) pointer, got nil")
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.Ip == nil || len(*state.Ip) != 0 {
		t.Errorf("expected an explicit empty list to be detected as real drift, got %+v", state.Ip)
	}
}

// Reproduces a reported false-positive: SMC returns the same group
// membership as the plain scalar list "element", just in a different
// order, which must not be treated as a reorder/add/remove.
func TestMergeResourceDrift_ScalarSliceReorderNotDrift(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.RpcServiceGroupResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{
		"name": "tf_rpc_service_group",
		"element": [
			"http://mysmc:8082/7.5/elements/rpc_service/280",
			"http://mysmc:8082/7.5/elements/rpc_service/884"
		]
	}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.RpcServiceGroupResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{
		"name": "tf_rpc_service_group",
		"element": [
			"http://mysmc:8082/7.5/elements/rpc_service/884",
			"http://mysmc:8082/7.5/elements/rpc_service/280"
		]
	}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	got := *state.Element
	if len(got) != 2 || got[0].ValueString() != "http://mysmc:8082/7.5/elements/rpc_service/280" ||
		got[1].ValueString() != "http://mysmc:8082/7.5/elements/rpc_service/884" {
		t.Errorf("expected state's original order to be preserved (same members, reordered), got %+v", got)
	}
}

// A genuine membership change (not just reordering) must still be
// detected as real drift.
func TestMergeResourceDrift_ScalarSliceRealChangeDetected(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.RpcServiceGroupResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "tf_rpc_service_group", "element": ["http://mysmc:8082/7.5/elements/rpc_service/280"]}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.RpcServiceGroupResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "tf_rpc_service_group", "element": ["http://mysmc:8082/7.5/elements/rpc_service/884"]}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	got := *state.Element
	if len(got) != 1 || got[0].ValueString() != "http://mysmc:8082/7.5/elements/rpc_service/884" {
		t.Errorf("expected a genuine membership change to be detected, got %+v", got)
	}
}

// Reproduces a reported false-positive: state has a value set by
// Terraform, but the API response for this particular refresh omits
// the attribute entirely (decodes as null) - e.g. an implicitly
// created sub-resource whose GET response only conditionally includes
// some boolean fields. An absent value from the API is not strong
// enough evidence of a deliberate removal to clobber a known value.
func TestMergeResourceDrift_SrcNullDoesNotClobberState(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.HostResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"address": "192.168.1.2", "comment": "kept via terraform"}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.HostResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"address": "192.168.1.2"}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}
	if !api.Comment.IsNull() {
		t.Fatalf("test setup error: expected api.Comment to be null, got %q", api.Comment.ValueString())
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.Comment.ValueString() != "kept via terraform" {
		t.Errorf("expected comment to survive an API response that omitted it, got %q (null=%v)",
			state.Comment.ValueString(), state.Comment.IsNull())
	}
}

// Reproduces a reported false-positive: SMC returns an IPv4 CIDR with
// a leading zero in an octet ("192.168.100.00/24"), which is the same
// subnet as Terraform's canonical "192.168.100.0/24" but a different
// string - this must not be treated as drift.
func TestMergeResourceDrift_IPv4LeadingZeroNotDrift(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.SingleNodeInterfaceResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"network_value": "192.168.100.0/24"}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.SingleNodeInterfaceResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"network_value": "192.168.100.00/24"}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.NetworkValue.ValueString() != "192.168.100.0/24" {
		t.Errorf("expected network_value to stay in its canonical form (same subnet, no real drift), got %q",
			state.NetworkValue.ValueString())
	}
}

// A genuinely different subnet must still be detected as real drift.
func TestMergeResourceDrift_IPv4RealChangeDetected(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.SingleNodeInterfaceResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"network_value": "192.168.100.0/24"}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}

	var api schema.SingleNodeInterfaceResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"network_value": "192.168.200.0/24"}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.NetworkValue.ValueString() != "192.168.200.0/24" {
		t.Errorf("expected a genuinely different subnet to be detected as drift, got %q", state.NetworkValue.ValueString())
	}
}

// from_ref has no json tag (it's Terraform-side bookkeeping for
// sub-resource creation, not an SMC API field), so a fresh API decode
// always leaves it null. Drift merge must not overlay it onto state,
// or a sub-resource loses its from_ref on every refresh.
func TestMergeResourceDrift_FromRefPreserved(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state schema.Ipv4AccessRuleResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "allow_ssh"}`), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}
	state.FromRef = types.StringValue("http://mysmc:8082/7.5/elements/fw_policy/1000273/fw_ipv4_access_rule")

	var api schema.Ipv4AccessRuleResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"name": "allow_ssh"}`), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.FromRef.IsNull() || state.FromRef.ValueString() == "" {
		t.Errorf("expected from_ref to be preserved, got %q (null=%v)", state.FromRef.ValueString(), state.FromRef.IsNull())
	}
}

// state has a Secondary list set; a set list is replaced wholesale from
// the API since there is no per-element identity to reconcile against.
// An unset Secondary list must not be populated from the API.
func TestMergeResourceDrift_ScalarSliceReplace(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state_json = `{"address": "192.168.1.2", "name": "AExampleHost", "secondary": ["212.20.1.1", "123.6.5.19"]}`
	var api_json = `{"address": "192.168.1.2", "name": "AExampleHost", "secondary": ["9.9.9.9"]}`

	var state schema.HostResourceModel
	var api schema.HostResourceModel
	if err := apijson.UnmarshalRoot([]byte(state_json), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(api_json), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	if state.Secondary == nil || len(*state.Secondary) != 1 || (*state.Secondary)[0].ValueString() != "9.9.9.9" {
		t.Errorf("expected secondary to be replaced with API value, got %+v", state.Secondary)
	}

	var stateUnset schema.HostResourceModel
	if err := apijson.UnmarshalRoot([]byte(`{"address": "192.168.1.2", "name": "AExampleHost"}`), &stateUnset); err != nil {
		t.Fatalf("Failed to unmarshal stateUnset: %v", err)
	}
	if err := resource.MergeResourceDrift(ctx, &api, &stateUnset); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}
	if stateUnset.Secondary != nil {
		t.Errorf("expected secondary to stay unset, got %+v", stateUnset.Secondary)
	}
}

// state has nodes 1 and 2; the API now reports node 2 (with a changed
// comment, drift) and a new node 3, and no longer reports node 1. The
// merged result should keep node 2 (state-anchored merge), add node 3
// (fully from the API), and drop node 1.
func TestMergeResourceDrift_ListReconciliation(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state_json = `{
		"name": "mycluster",
		"nodes": [
			{"firewall_node": {"name": "node1", "nodeid": 1, "comment": "user comment 1"}},
			{"firewall_node": {"name": "node2", "nodeid": 2, "comment": "user comment 2"}}
		]
	}`

	var api_json = `{
		"name": "mycluster",
		"nodes": [
			{"firewall_node": {"name": "node2-renamed", "nodeid": 2, "comment": "smc comment 2"}},
			{"firewall_node": {"name": "node3", "nodeid": 3, "comment": "smc comment 3"}}
		]
	}`

	var state schema.FirewallClusterResourceModel
	var api schema.FirewallClusterResourceModel
	if err := apijson.UnmarshalRoot([]byte(state_json), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(api_json), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	nodes := *state.Nodes
	if len(nodes) != 2 {
		t.Fatalf("expected 2 nodes after reconciliation, got %d", len(nodes))
	}

	node2 := nodes[0].FirewallNode
	if node2.Nodeid.ValueInt64() != 2 {
		t.Fatalf("expected first node to still be nodeid 2 (existing state order preserved), got %d", node2.Nodeid.ValueInt64())
	}
	if node2.Comment.ValueString() != "smc comment 2" {
		t.Errorf("expected node 2 comment drift to be overlaid, got %q", node2.Comment.ValueString())
	}

	node3 := nodes[1].FirewallNode
	if node3.Nodeid.ValueInt64() != 3 {
		t.Fatalf("expected second node to be nodeid 3 (added), got %d", node3.Nodeid.ValueInt64())
	}
	if node3.Name.ValueString() != "node3" || node3.Comment.ValueString() != "smc comment 3" {
		t.Errorf("expected added node 3 to be fully populated from API, got name=%q comment=%q",
			node3.Name.ValueString(), node3.Comment.ValueString())
	}

	for _, n := range nodes {
		if n.FirewallNode.Nodeid.ValueInt64() == 1 {
			t.Errorf("expected node 1 to be pruned (no longer reported by API)")
		}
	}
}

// Reproduces a reported false-diff: state has a single existing node,
// and the API now also reports a new node, listed BEFORE the existing
// one in the API's own order (as SMC can do when a nested list gets an
// out-of-band addition). Nested lists back ListNestedBlock, which
// diffs config against the refreshed state positionally, so the
// existing node must stay at its original position (matching config)
// and the new node must land after it - not wherever the API happens
// to list it - or the existing node looks modified and the new one
// looks like an unrelated removal.
func TestMergeResourceDrift_ListReconciliation_PreservesExistingOrder(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var state_json = `{
		"name": "mycluster",
		"nodes": [
			{"firewall_node": {"name": "node0", "nodeid": 0}}
		]
	}`

	// API lists the newly added node first, then the pre-existing one.
	var api_json = `{
		"name": "mycluster",
		"nodes": [
			{"firewall_node": {"name": "node1", "nodeid": 1}},
			{"firewall_node": {"name": "node0", "nodeid": 0}}
		]
	}`

	var state schema.FirewallClusterResourceModel
	var api schema.FirewallClusterResourceModel
	if err := apijson.UnmarshalRoot([]byte(state_json), &state); err != nil {
		t.Fatalf("Failed to unmarshal state: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(api_json), &api); err != nil {
		t.Fatalf("Failed to unmarshal api: %v", err)
	}

	if err := resource.MergeResourceDrift(ctx, &api, &state); err != nil {
		t.Fatalf("MergeResourceDrift failed: %v", err)
	}

	nodes := *state.Nodes
	if len(nodes) != 2 {
		t.Fatalf("expected 2 nodes after reconciliation, got %d", len(nodes))
	}

	if nodes[0].FirewallNode.Nodeid.ValueInt64() != 0 {
		t.Errorf("expected existing node 0 to stay at position 0 (matching config), got nodeid %d",
			nodes[0].FirewallNode.Nodeid.ValueInt64())
	}
	if nodes[1].FirewallNode.Nodeid.ValueInt64() != 1 {
		t.Errorf("expected new node 1 to be appended at position 1, got nodeid %d",
			nodes[1].FirewallNode.Nodeid.ValueInt64())
	}
}
