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

// this test should be in package resource with model_update_payload.go,
// but we need structs from the schema to test and we would have a
// circular dependency (see model_utils_test.go).

package provider

import (
	"context"
	"os"
	"testing"

	"github.com/terraform-providers/terraform-provider-smc/internal/apijson"
	"github.com/terraform-providers/terraform-provider-smc/internal/resource"
	"github.com/terraform-providers/terraform-provider-smc/internal/schema"

	"github.com/hashicorp/terraform-plugin-log/tflogtest"
)

// Reproduces the reported bug: a comment set only via the SMC GUI (not
// in the plan) must survive an update that only changes another
// attribute, instead of being sent as null and cleared server-side.
// Attributes the plan does set must still win, and nested lists remain
// plan-authoritative for membership/structure.
func TestMergeResourceUpdatePayload_PreservesUnsetScalar(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	// the plan: user is only changing "name", never configured "comment".
	var plan_json = `{"address": "192.168.1.2", "name": "renamed-by-terraform", "secondary": ["1.1.1.1"]}`
	// base: freshly read from SMC, has a comment set via the GUI.
	var base_json = `{"address": "192.168.1.2", "name": "AExampleHost", "comment": "set via SMC GUI", "secondary": ["9.9.9.9"]}`

	var plan schema.HostResourceModel
	var base schema.HostResourceModel
	if err := apijson.UnmarshalRoot([]byte(plan_json), &plan); err != nil {
		t.Fatalf("Failed to unmarshal plan: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(base_json), &base); err != nil {
		t.Fatalf("Failed to unmarshal base: %v", err)
	}

	if err := resource.MergeResourceUpdatePayload(ctx, &base /*src*/, &plan /*dest*/); err != nil {
		t.Fatalf("MergeResourceUpdatePayload failed: %v", err)
	}

	if plan.Comment.ValueString() != "set via SMC GUI" {
		t.Errorf("expected comment (unset in plan) to be backfilled from base, got %q", plan.Comment.ValueString())
	}
	if plan.Name.ValueString() != "renamed-by-terraform" {
		t.Errorf("expected name (set in plan) to win, got %q", plan.Name.ValueString())
	}
	if plan.Secondary == nil || len(*plan.Secondary) != 1 || (*plan.Secondary)[0].ValueString() != "1.1.1.1" {
		t.Errorf("expected secondary list to stay plan-authoritative, got %+v", plan.Secondary)
	}
}

// Regression test: computed fields nested inside lists-of-objects
// (e.g. each cluster node's own "key", or the deeper routing_node ->
// routing_node case reported live) must still be backfilled from the
// fresh read at every nesting level, not just at the top level.
func TestMergeResourceUpdatePayload_BackfillsNestedComputedFields(t *testing.T) {
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	// plan: the user's config for the node, with no "key" (computed,
	// never set by a plan).
	var plan_json = `{
		"name": "mycluster",
		"nodes": [
			{"firewall_node": {"name": "node1", "nodeid": 1}}
		]
	}`

	// base: freshly read from SMC, has the real key for that node.
	var base_json = `{
		"name": "mycluster",
		"nodes": [
			{"firewall_node": {"name": "node1", "nodeid": 1, "key": 9999}}
		]
	}`

	var plan schema.FirewallClusterResourceModel
	var base schema.FirewallClusterResourceModel
	if err := apijson.UnmarshalRoot([]byte(plan_json), &plan); err != nil {
		t.Fatalf("Failed to unmarshal plan: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(base_json), &base); err != nil {
		t.Fatalf("Failed to unmarshal base: %v", err)
	}

	if err := resource.MergeResourceUpdatePayload(ctx, &base /*src*/, &plan /*dest*/); err != nil {
		t.Fatalf("MergeResourceUpdatePayload failed: %v", err)
	}

	node := (*plan.Nodes)[0].FirewallNode
	if node.Key.ValueInt64() != 9999 {
		t.Errorf("expected nested node's key to be backfilled from base, got %d", node.Key.ValueInt64())
	}
	if node.Name.ValueString() != "node1" {
		t.Errorf("expected plan's own nested field to be unaffected, got %q", node.Name.ValueString())
	}
}
