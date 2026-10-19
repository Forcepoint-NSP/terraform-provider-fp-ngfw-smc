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

package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/terraform-providers/terraform-provider-smc/internal/apijson"
	"github.com/terraform-providers/terraform-provider-smc/internal/resource"
	"github.com/terraform-providers/terraform-provider-smc/internal/schema"
)

// TestMergeExpressionSubExpressionReorderedNeRef covers SMC echoing a
// sub_expression's ne_ref members back in a different order than submitted
// (a member set, not an ordered list): the sorted Operator+NeRef signature in
// ExpressionResourceModel.GetSliceIds still matches.
func TestMergeExpressionSubExpressionReorderedNeRef(t *testing.T) {
	setupTest()

	configJSON := `
{
    "name": "not accor not guest not sw-admin",
    "operator": "intersection",
    "sub_expression": [
        {
            "ne_ref": ["http://localhost:8082/7.4/elements/group/100", "http://localhost:8082/7.4/elements/group/101"],
            "operator": "intersection"
        },
        {
            "ne_ref": ["http://localhost:8082/7.4/elements/network/200"],
            "operator": "exclusion"
        }
    ]
}`

	// SMC echoes the first sub_expression's ne_ref members back reordered -
	// same set, different order, so the Operator+NeRef signature no longer
	// matches byte-for-byte.
	createdJSON := `
{
    "name": "not accor not guest not sw-admin",
    "operator": "intersection",
    "key": 500,
    "sub_expression": [
        {
            "ne_ref": ["http://localhost:8082/7.4/elements/group/101", "http://localhost:8082/7.4/elements/group/100"],
            "operator": "intersection",
            "key": 501
        },
        {
            "ne_ref": ["http://localhost:8082/7.4/elements/network/200"],
            "operator": "exclusion",
            "key": 502
        }
    ]
}`

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var configData schema.ExpressionResourceModel
	var createdData schema.ExpressionResourceModel

	if err := apijson.UnmarshalRoot([]byte(configJSON), &configData); err != nil {
		t.Fatalf("Failed to unmarshal configData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(createdJSON), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}

	if err := resource.MergeResourceModels(ctx, &createdData, &configData); err != nil {
		t.Fatalf("MergeResourceModels failed: %v", err)
	}

	apiData, _ := apijson.MarshalRoot(configData)
	fmt.Printf("After merge, configData: %s\n", apiData)

	sub0 := (*configData.SubExpression)[0]
	if sub0.Key.IsUnknown() || sub0.Key.IsNull() {
		t.Errorf("sub_expression[0].key is unknown/null after merge (expected 501)")
	} else if sub0.Key.ValueInt64() != 501 {
		t.Errorf("sub_expression[0].key: expected 501, got %d", sub0.Key.ValueInt64())
	}

	sub1 := (*configData.SubExpression)[1]
	if sub1.Key.IsUnknown() || sub1.Key.IsNull() {
		t.Errorf("sub_expression[1].key is unknown/null after merge (expected 502)")
	} else if sub1.Key.ValueInt64() != 502 {
		t.Errorf("sub_expression[1].key: expected 502, got %d", sub1.Key.ValueInt64())
	}
}

// TestMergeExpressionSubExpression is the regression test for SMC-67247's
// "Provider returned invalid result object after apply ... unknown value
// for sub_expression[0].read_only" (and every other fpro-tagged metadata
// field, for every entry, on any smc_expression using sub_expression).
//
// Root cause: ExpressionResourceModel.GetSliceIds() (see
// model_expression_custom.go) used to match only by Name. That's correct
// for a top-level Expression element (always named), but the same struct
// is reused verbatim for the embedded sub_expression list, whose entries
// SMC never names (they aren't standalone elements). Name-only matching
// therefore always failed for every sub-expression, so mergeSlice always
// skipped them (see its "dest element not found" path in model_utils.go),
// leaving their computed fields permanently unknown. Fixed by falling back
// to an Operator+NeRef signature when Name is absent, and (belt-and-braces,
// see TestMergeExpressionSubExpressionPositionalFallback) by mergeSlice's
// own positional fallback for whatever that signature still can't match.
const (
	subExprGroup100 = "http://localhost:8082/7.4/elements/group/100"
	subExprGroup101 = "http://localhost:8082/7.4/elements/group/101"
	subExprNet200   = "http://localhost:8082/7.4/elements/network/200"
)

func mergeExpressionJSON(t *testing.T, configJSON, createdJSON string) schema.ExpressionResourceModel {
	t.Helper()
	setupTest()
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var configData, createdData schema.ExpressionResourceModel
	if err := apijson.UnmarshalRoot([]byte(configJSON), &configData); err != nil {
		t.Fatalf("Failed to unmarshal configData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(createdJSON), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}
	if err := resource.MergeResourceModels(ctx, &createdData, &configData); err != nil {
		t.Fatalf("MergeResourceModels failed: %v", err)
	}
	return configData
}

// TestMergeExpressionNamedSubExpression covers the documented usage: the
// config names each sub_expression, SMC's response carries no name.
func TestMergeExpressionNamedSubExpression(t *testing.T) {
	configJSON := fmt.Sprintf(`{
    "name": "e", "operator": "intersection",
    "sub_expression": [
        {"name": "first", "ne_ref": ["%s", "%s"], "operator": "intersection"},
        {"name": "second", "ne_ref": ["%s"], "operator": "exclusion"}
    ]}`, subExprGroup100, subExprGroup101, subExprNet200)
	createdJSON := fmt.Sprintf(`{
    "name": "e", "operator": "intersection", "key": 500,
    "sub_expression": [
        {"ne_ref": ["%s"], "operator": "exclusion", "key": 502},
        {"ne_ref": ["%s", "%s"], "operator": "intersection", "key": 501}
    ]}`, subExprNet200, subExprGroup100, subExprGroup101)

	got := mergeExpressionJSON(t, configJSON, createdJSON)

	// API order differs from config order: keys must follow the signature.
	for i, want := range []int64{501, 502} {
		key := (*got.SubExpression)[i].Key
		if key.IsUnknown() || key.IsNull() || key.ValueInt64() != want {
			t.Errorf("sub_expression[%d].key: expected %d, got %v", i, want, key)
		}
	}
}

// TestMergeExpressionSubExpressionIdMismatch: a sub_expression whose
// signature has no counterpart in the API response must stay unknown, not
// be paired positionally with an unrelated entry.
func TestMergeExpressionSubExpressionIdMismatch(t *testing.T) {
	configJSON := fmt.Sprintf(`{
    "name": "e", "operator": "intersection",
    "sub_expression": [{"ne_ref": ["%s"], "operator": "intersection"}]}`, subExprGroup100)
	createdJSON := fmt.Sprintf(`{
    "name": "e", "operator": "intersection", "key": 500,
    "sub_expression": [{"ne_ref": ["%s"], "operator": "exclusion", "key": 502}]}`, subExprNet200)

	got := mergeExpressionJSON(t, configJSON, createdJSON)

	if key := (*got.SubExpression)[0].Key; !key.IsUnknown() && !key.IsNull() {
		t.Errorf("sub_expression[0].key must not be set from a mismatched entry, got %v", key)
	}
}

// TestMergeExpressionNestedOnlySiblings: siblings with no ne_ref that differ
// only by their own sub_expression must each get their own keys.
func TestMergeExpressionNestedOnlySiblings(t *testing.T) {
	configJSON := fmt.Sprintf(`{
    "name": "e", "operator": "intersection",
    "sub_expression": [
        {"operator": "union", "sub_expression": [{"operator": "union", "ne_ref": ["%s"]}]},
        {"operator": "union", "sub_expression": [{"operator": "union", "ne_ref": ["%s"]}]}
    ]}`, subExprGroup100, subExprNet200)
	createdJSON := fmt.Sprintf(`{
    "name": "e", "operator": "intersection", "key": 500,
    "sub_expression": [
        {"operator": "union", "key": 601, "sub_expression": [{"operator": "union", "ne_ref": ["%s"], "key": 701}]},
        {"operator": "union", "key": 602, "sub_expression": [{"operator": "union", "ne_ref": ["%s"], "key": 702}]}
    ]}`, subExprGroup100, subExprNet200)

	got := mergeExpressionJSON(t, configJSON, createdJSON)

	for i, want := range []int64{601, 602} {
		key := (*got.SubExpression)[i].Key
		if key.IsUnknown() || key.IsNull() || key.ValueInt64() != want {
			t.Errorf("sub_expression[%d].key: expected %d, got %v", i, want, key)
		}
		child := (*(*got.SubExpression)[i].SubExpression)[0].Key
		if child.IsUnknown() || child.IsNull() || child.ValueInt64() != want+100 {
			t.Errorf("sub_expression[%d].sub_expression[0].key: expected %d, got %v", i, want+100, child)
		}
	}
}

// TestExpressionGetSliceIdsNullOperator: a null/unknown operator yields no
// id, so two such entries cannot collide on a "|<refs>" signature.
func TestExpressionGetSliceIdsNullOperator(t *testing.T) {
	for name, op := range map[string]types.String{"null": types.StringNull(), "unknown": types.StringUnknown()} {
		m := schema.ExpressionResourceModel{Operator: op}
		if ids := m.GetSliceIds(context.Background()); ids != nil {
			t.Errorf("%s operator: expected nil ids, got %v", name, ids)
		}
	}
}

func TestMergeExpressionSubExpression(t *testing.T) {
	setupTest()

	configJSON := `
{
    "name": "not accor not guest not sw-admin",
    "operator": "intersection",
    "sub_expression": [
        {
            "ne_ref": ["http://localhost:8082/7.4/elements/group/100", "http://localhost:8082/7.4/elements/group/101"],
            "operator": "intersection"
        },
        {
            "ne_ref": ["http://localhost:8082/7.4/elements/network/200"],
            "operator": "exclusion"
        }
    ]
}`

	createdJSON := `
{
    "name": "not accor not guest not sw-admin",
    "operator": "intersection",
    "key": 500,
    "etag": "abc123",
    "sub_expression": [
        {
            "ne_ref": ["http://localhost:8082/7.4/elements/group/100", "http://localhost:8082/7.4/elements/group/101"],
            "operator": "intersection",
            "key": 501,
            "read_only": false,
            "system": false
        },
        {
            "ne_ref": ["http://localhost:8082/7.4/elements/network/200"],
            "operator": "exclusion",
            "key": 502,
            "read_only": false,
            "system": false
        }
    ]
}`

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var configData schema.ExpressionResourceModel
	var createdData schema.ExpressionResourceModel

	if err := apijson.UnmarshalRoot([]byte(configJSON), &configData); err != nil {
		t.Fatalf("Failed to unmarshal configData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(createdJSON), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}

	if err := resource.MergeResourceModels(ctx, &createdData, &configData); err != nil {
		t.Fatalf("MergeResourceModels failed: %v", err)
	}

	apiData, _ := apijson.MarshalRoot(configData)
	fmt.Printf("After merge, configData: %s\n", apiData)

	if configData.Key.IsUnknown() || configData.Key.IsNull() {
		t.Errorf("top-level expression.key is unknown/null after merge")
	}

	sub0 := (*configData.SubExpression)[0]
	if sub0.Key.IsUnknown() || sub0.Key.IsNull() {
		t.Errorf("sub_expression[0].key is unknown/null after merge (expected 501)")
	} else if sub0.Key.ValueInt64() != 501 {
		t.Errorf("sub_expression[0].key: expected 501, got %d", sub0.Key.ValueInt64())
	}

	sub1 := (*configData.SubExpression)[1]
	if sub1.Key.IsUnknown() || sub1.Key.IsNull() {
		t.Errorf("sub_expression[1].key is unknown/null after merge (expected 502)")
	} else if sub1.Key.ValueInt64() != 502 {
		t.Errorf("sub_expression[1].key: expected 502, got %d", sub1.Key.ValueInt64())
	}
}

// TestMergeSingleFwInterfacesAndVlanInterfaces documents an investigation
// for SMC-67247's OTHER apply error on the same ticket: "physical_interfaces
// [N].interfaces[0].single_node_interface.key/link unknown after apply",
// seen only on a physical_interface configured with BOTH a native
// `interfaces` block and `vlan_interfaces` on the same interface_id (a
// combination with no example or test coverage anywhere else in this
// provider - every existing usage keeps them mutually exclusive per
// physical interface).
//
// This test feeds MergeResourceModels a hand-built, COMPLETE API response
// (both "interfaces" and "vlanInterfaces" present, each with their computed
// key/link) for that exact shape, and it merges correctly - proving the
// generic reconciliation mechanism (mergeSlice/GetElementIds) has no bug
// for this combination. That means the real apply failure is not
// reproducible as a provider-side merge defect: it points instead at SMC's
// actual API response genuinely omitting (or not persisting) the native
// interfaces[] entry for a physical interface that also carries VLAN
// sub-interfaces - a backend/data-model question outside what this
// unit test (or this provider) can fix. See smcaas-infra's
// tests/terraform/many_engines_special_conf module, which avoids the
// combination for this reason.
func TestMergeSingleFwInterfacesAndVlanInterfaces(t *testing.T) {
	setupTest()

	configJSON := `
{
    "log_server_ref": "http://localhost:8082/7.4/elements/log_server/1441",
    "name": "myfw",
    "nodes": [
        {"firewall_node": {"name": "myfwnode", "nodeid": 1}}
    ],
    "physicalInterfaces": [
        {
            "physical_interface": {
                "interface_id": "2",
                "zone_ref": "http://localhost:8082/7.4/elements/interface_zone/1",
                "interfaces": [
                    {
                        "single_node_interface": {
                            "address": "10.0.64.10",
                            "network_value": "10.0.64.0/25",
                            "nicid": "2",
                            "nodeid": 1
                        }
                    }
                ],
                "vlanInterfaces": [
                    {
                        "interface_id": "2.5",
                        "name": "VLAN 2.5",
                        "interfaces": [
                            {
                                "single_node_interface": {
                                    "address": "196.207.244.226",
                                    "network_value": "196.207.244.224/30",
                                    "nicid": "2.5",
                                    "nodeid": 1
                                }
                            }
                        ]
                    }
                ]
            }
        }
    ]
}`

	// A hand-built, COMPLETE API response: both "interfaces" and
	// "vlanInterfaces" round-tripped with their computed key/link set,
	// exactly mirroring the shape TestMergeSingleFw already proves works
	// when vlan_interfaces is absent.
	createdJSON := `
{
    "key": 5437,
    "log_server_ref": "http://localhost:8082/7.4/elements/log_server/1441",
    "name": "myfw",
    "physicalInterfaces": [
        {
            "physical_interface": {
                "interface_id": "2",
                "key": 900,
                "link": [
                    {"href": "http://localhost:8082/7.4/elements/single_fw/5437/physical_interface/900", "rel": "self", "type": "physical_interface"}
                ],
                "interfaces": [
                    {
                        "single_node_interface": {
                            "address": "10.0.64.10",
                            "network_value": "10.0.64.0/25",
                            "nicid": "2",
                            "nodeid": 1,
                            "key": 901,
                            "link": [
                                {"href": "http://localhost:8082/7.4/elements/single_fw/5437/physical_interface/900/single_node_interface/901", "rel": "self", "type": "single_node_interface"}
                            ]
                        }
                    }
                ],
                "vlanInterfaces": [
                    {
                        "interface_id": "2.5",
                        "name": "VLAN 2.5",
                        "key": 902,
                        "link": [
                            {"href": "http://localhost:8082/7.4/elements/single_fw/5437/physical_interface/900/vlan_interface/902", "rel": "self", "type": "physical_vlan_interface"}
                        ],
                        "interfaces": [
                            {
                                "single_node_interface": {
                                    "address": "196.207.244.226",
                                    "network_value": "196.207.244.224/30",
                                    "nicid": "2.5",
                                    "nodeid": 1,
                                    "key": 903,
                                    "link": [
                                        {"href": "http://localhost:8082/7.4/elements/single_fw/5437/physical_interface/900/vlan_interface/902/single_node_interface/903", "rel": "self", "type": "single_node_interface"}
                                    ]
                                }
                            }
                        ]
                    }
                ]
            }
        }
    ]
}`

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var configData schema.SingleFirewallResourceModel
	var createdData schema.SingleFirewallResourceModel

	if err := apijson.UnmarshalRoot([]byte(configJSON), &configData); err != nil {
		t.Fatalf("Failed to unmarshal configData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(createdJSON), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}

	if err := resource.MergeResourceModels(ctx, &createdData, &configData); err != nil {
		t.Fatalf("MergeResourceModels failed: %v", err)
	}

	apiData, _ := apijson.MarshalRoot(configData)
	fmt.Printf("After merge, configData: %s\n", apiData)

	physIntf := (*configData.PhysicalInterfaces)[0].PhysicalInterface

	if physIntf.Key.IsUnknown() || physIntf.Key.IsNull() {
		t.Errorf("physical_interface.key is unknown/null after merge")
	} else if physIntf.Key.ValueInt64() != 900 {
		t.Errorf("physical_interface.key: expected 900, got %d", physIntf.Key.ValueInt64())
	}

	sni := (*physIntf.Interfaces)[0].SingleNodeInterface
	if sni.Key.IsUnknown() || sni.Key.IsNull() {
		t.Errorf("interfaces[0].single_node_interface.key is unknown/null after merge (expected 901)")
	} else if sni.Key.ValueInt64() != 901 {
		t.Errorf("interfaces[0].single_node_interface.key: expected 901, got %d", sni.Key.ValueInt64())
	}

	vlanSni := (*(*physIntf.VlanInterfaces)[0].Interfaces)[0].SingleNodeInterface
	if vlanSni.Key.IsUnknown() || vlanSni.Key.IsNull() {
		t.Errorf("vlan_interfaces[0].interfaces[0].single_node_interface.key is unknown/null after merge (expected 903)")
	} else if vlanSni.Key.ValueInt64() != 903 {
		t.Errorf("vlan_interfaces[0].interfaces[0].single_node_interface.key: expected 903, got %d", vlanSni.Key.ValueInt64())
	}
}
