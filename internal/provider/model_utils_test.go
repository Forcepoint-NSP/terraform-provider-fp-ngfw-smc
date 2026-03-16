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

// this test should be in package resource with model_utils.go, but we
// need struct from the schema to test and we would have a circular
// dependencies

package provider

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/terraform-providers/terraform-provider-smc/internal/apijson"
	"github.com/terraform-providers/terraform-provider-smc/internal/resource"
	"github.com/terraform-providers/terraform-provider-smc/internal/schema"
)

func setupTest() {
	resource.AddSpecialCopyHandlers("link", schema.CopyLinkField)
}

func TestGetElementId(t *testing.T) {
	setupTest()
	var config_data_json = `{
    "name": "mycluster",
    "nodes": [
        {
            "firewall_node": {
                "name": "mycluster node 1",
                "nodeid": 1
            }
        },
        {
            "firewall_node": {
                "name": "mycluster node 2",
                "nodeid": 2
            }
        }
    ]
}`

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	var configData schema.SingleFirewallResourceModel

	if err := apijson.UnmarshalRoot([]byte(config_data_json), &configData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}

	fwnode0 := (*configData.Nodes)[1]
	elementIds, err := resource.GetElementIds(ctx, reflect.ValueOf(fwnode0))
	if err != nil {
		t.Fatalf("GetElementIds failed: %v", err)
	}
	if elementIds[0] != "FirewallNode/2" {
		t.Errorf("Expected elementId to be 'FirewallNode/2', got '%s'", elementIds)
	}
}

func TestMergeHost(t *testing.T) {
	setupTest()

	var config_data_json = `{
		"address": "192.168.1.2",
		"admin_domain": "http://localhost:8082/7.3/elements/admin_domain/1",
		"comment": "Created via Terraform",
		"key": 5253,
		"link": [
			{
				"href": "http://localhost:8082/7.3/elements/host/5253",
				"rel": "self",
				"type": "host"
			}
		],
		"name": "AExampleHost",
		"secondary": [
			"212.20.1.1",
			"123.6.5.19"
		]
	}`

	var created_data_json = `{
		"address": "192.168.1.2",
		"admin_domain": "xxxx",
		"comment": "Created via Terraform",
		"name": "AExampleHost",
		"secondary": ["212.20.1.1", "123.6.5.19"]
	}`

	var data schema.HostResourceModel
	var createdData schema.HostResourceModel
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	if err := apijson.UnmarshalRoot([]byte(config_data_json), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(created_data_json), &data); err != nil {
		t.Fatalf("Failed to unmarshal data: %v", err)
	}

	err := resource.MergeResourceModels(ctx, &createdData, &data)
	if err != nil {
		t.Fatalf("resource.MergeResourceModels failed: %v", err)
	}

	fmt.Printf("After merge, data: %+v\n", data)
	// Verify the merge worked - should have link data from createdData

	// verify the "lk"
	if data.Lk.IsNull() || data.Lk.IsUnknown() {
		t.Errorf("Expected Lk to be set after merge, but it is null or unknown")
	}
	selfLink := data.Lk.Elements()["self"].(types.String).ValueString()
	if selfLink != "http://localhost:8082/7.3/elements/host/5253" {
		t.Errorf("Expected Lk.Elements['self'] to be 'http://localhost:8082/7.3/elements/host/5253', got '%s'", selfLink)
	}
}

func TestMergeCluster(t *testing.T) {
	setupTest()

	// nodes are out of order to test matching by nodeid
	var config_data_json = `{
    "name": "mycluster",
    "nodes": [
        {
            "firewall_node": {
                "name": "mycluster node 1",
                "nodeid": 1
            }
        },
        {
            "firewall_node": {
                "name": "mycluster node 2",
                "nodeid": 2
            }
        }
    ]
}`

	var created_data_json = `{
    "name": "mycluster",
    "nodes":     [
        {
            "firewall_node": {
                "disabled": false,
                "key": 5387,
                "link": [
                    {
                        "href": "http://localhost:8082/7.3/elements/fw_cluster/5385/firewall_node/5387",
                        "rel": "self",
                        "type": "firewall_node"
                    },
                ],
                "name": "mycluster node 2",
                "nodeid": 2,
            }
        },
        {
            "firewall_node": {
                "activate_test": false,
                "appliance_info": {
                    "first_upload_time": 0,
                    "initial_contact_time": 0,
                    "initial_license_remaining_days": 0
                },
                "disabled": false,
                "key": 5386,
                "link": [
                    {
                        "href": "http://localhost:8082/7.3/elements/fw_cluster/5385/firewall_node/5386",
                        "rel": "self",
                        "type": "firewall_node"
                    }                ],
                "name": "mycluster node 1",
                "nodeid": 1,
            }
        },
    ]
}`

	var configData schema.FirewallClusterResourceModel
	var createdData schema.FirewallClusterResourceModel

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	if err := apijson.UnmarshalRoot([]byte(config_data_json), &configData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(created_data_json), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal configData: %v", err)
	}

	err := resource.MergeResourceModels(ctx /*src*/, &createdData /*dest*/, &configData)
	if err != nil {
		t.Fatalf("resource.MergeResourceModels failed: %v", err)
	}

	apiData, err := apijson.MarshalRoot(configData)
	fmt.Printf("After merge, configData: %s\n", apiData)

	nodesLinks, _ := (*(*configData.Nodes)[0].FirewallNode).Link.AsStructSliceT(context.TODO())

	if nodesLinks[0].Href.ValueString() != "http://localhost:8082/7.3/elements/fw_cluster/5385/firewall_node/5386" {
		t.Errorf("Wrong Link Href after merge: %s", nodesLinks[0].Href.ValueString())
	}

}

func TestMergeSingleFw(t *testing.T) {
	setupTest()

	var config_data_json = `{
    "log_server_ref": "http://localhost:8082/7.3/elements/log_server/1441",
    "name": "myfw",
    "nodes": [
        {
            "firewall_node": {
                "name": "myfwnode",
                "nodeid": 1
            }
        }
    ],
    "physicalInterfaces": [
        {
            "physical_interface": {
                "interface_id": "0",
                "interfaces": [
                    {
                        "single_node_interface": {
                            "address": "192.168.100.14",
                            "network_value": "192.168.100.00/24",
                            "nicid": "0",
                            "nodeid": 1,
                            "primary_mgt": true
                        }
                    }
                ]
            }
        },
        {
            "physical_interface": {
                "interface_id": "1",
                "interfaces": [
                    {
                        "single_node_interface": {
                            "address": "192.168.101.14",
                            "network_value": "192.168.101.00/24",
                            "nicid": "1",
                            "nodeid": 1
                        }
                    }
                ]
            }
        }
    ]
}`

	var created_data_json = `{
    "key": 5437,
    "log_server_ref": "http://localhost:8082/7.3/elements/log_server/1441",
    "name": "myfw",
    "physicalInterfaces": [
        {
            "physical_interface": {
                "aggregate_mode": "none",
                "arp_entry": [],
                "cvi_mode": "none",
                "dhcp_server_on_interface": {
                    "default_lease_time": 7200,
                    "dhcp_range_per_node": []
                },
                "duplicate_address_detection": true,
                "include_prefix_info_option_flag": true,
                "interface_id": "0",
                "interfaces": [
                    {
                        "single_node_interface": {
                            "address": "192.168.100.14",
                            "auth_request": false,
                            "auth_request_source": false,
                            "automatic_default_route": false,
                            "backup_heartbeat": false,
                            "backup_mgt": false,
                            "domain_specific_dns_queries_source": false,
                            "dynamic": false,
                            "key": 1343,
                            "network_value": "192.168.100.0/24",
                            "nicid": "0",
                            "nodeid": 1,
                            "outgoing": false,
                            "pppoa": false,
                            "pppoe": false,
                            "primary_heartbeat": false,
                            "primary_mgt": true,
                            "relayed_by_dhcp": false,
                            "reverse_connection": false,
                            "vrrp": false,
                            "vrrp_id": -1,
                            "vrrp_priority": -1
                        }
                    }
                ],
                "key": 727,
                "link": [
                    {
                        "href": "http://localhost:8082/7.3/elements/single_fw/5437/physical_interface/727",
                        "rel": "self",
                        "type": "physical_interface"
                    }
                ],
                "lldp_mode": "disabled",
                "log_moderation": [
                    {
                        "burst": 1000,
                        "log_event": "antispoofing",
                        "rate": 100
                    },
                    {
                        "burst": 20000,
                        "log_event": "discard",
                        "rate": 5000
                    },
                    {
                        "burst": 80000,
                        "log_event": "allow",
                        "rate": 40000
                    }
                ],
                "managed_address_flag": false,
                "mtu": -1,
                "name": "Interface 0",
                "other_configuration_flag": false,
                "override_engine_settings": false,
                "override_log_moderation_settings": false,
                "qos_limit": 0,
                "qos_mode": "no_qos",
                "route_replies_back_mode": false,
                "router_advertisement": false,
                "set_autonomous_address_flag": true,
                "shared_interface": false,
                "syn_mode": "default",
                "sync_parameter": {
                    "full_sync_interval": 5000,
                    "heartbeat_group_ip": "224.0.0.221",
                    "incr_sync_interval": 50,
                    "statesync_group_ip": "224.0.0.222",
                    "sync_mode": "sync_all",
                    "sync_security": "sign"
                },
                "virtual_engine_vlan_ok": false,
                "virtual_resource_settings": [],
                "vlanInterfaces": []
            }
        }
    ],
}`

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	tflog.Info(ctx, "test log message", map[string]any{
		"key": "value",
	})

	var configData schema.SingleFirewallResourceModel
	var createdData schema.SingleFirewallResourceModel

	if err := apijson.UnmarshalRoot([]byte(config_data_json), &configData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(created_data_json), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal configData: %v", err)
	}

	err := resource.MergeResourceModels(ctx, &createdData, &configData)
	if err != nil {
		t.Fatalf("resource.MergeResourceModels failed: %v", err)
	}

	apiData, err := apijson.MarshalRoot(configData)
	fmt.Printf("After merge, configData: %s\n", apiData)

	if !configData.LocationRef.IsNull() && !configData.LocationRef.IsUnknown() {
		t.Errorf("Expected LocationRef to be null after merge, got: %v", configData.LocationRef)
	}

	expectedRootKey := int64(5437)
	actualRootKey := configData.Key.ValueInt64()
	if actualRootKey != expectedRootKey {
		t.Errorf("Expected root Key to be %d, got %d", expectedRootKey, actualRootKey)
	}

	physicalInterface0 := (*configData.PhysicalInterfaces)[0].PhysicalInterface
	expectedPhysInterfaceKey := int64(727)
	actualPhysInterfaceKey := physicalInterface0.Key.ValueInt64()
	if actualPhysInterfaceKey != expectedPhysInterfaceKey {
		t.Errorf("Expected PhysicalInterface Key to be %d, got %d", expectedPhysInterfaceKey, actualPhysInterfaceKey)
	}
	singleNodeInterface := (*physicalInterface0.Interfaces)[0].SingleNodeInterface
	expectedSingleNodeKey := int64(1343)
	actualSingleNodeKey := singleNodeInterface.Key.ValueInt64()
	if actualSingleNodeKey != expectedSingleNodeKey {
		t.Errorf("Expected SingleNodeInterface Key to be %d, got %d", expectedSingleNodeKey, actualSingleNodeKey)
	}
}

func TestMergeSingleFwDynamicIntf(t *testing.T) {
	setupTest()

	var config_data_json = `
{
    "log_server_ref": "http://localhost:8082/7.4/elements/log_server/268581077",
    "name": "tf_single_fw_dynamic_interface",
    "nodes": [
        {
            "firewall_node": {
                "name": "myfwnode",
                "nodeid": 1
            }
        }
    ],
    "physicalInterfaces": [
        {
            "physical_interface": {
                "interface_id": "0",
                "interfaces": [
                    {
                        "single_node_interface": {
                            "address": "192.168.100.14",
                            "network_value": "192.168.100.00/24",
                            "nicid": "0",
                            "nodeid": 1,
                            "primary_mgt": true
                        }
                    }
                ]
            }
        },
        {
            "physical_interface": {
                "comment": "dynamic IPv4/IPv6 intertace",
                "interface_id": "1",
                "interfaces": [
                    {
                        "single_node_interface": {
                            "dynamic": true,
                            "dynamic_index": 1,
                            "nicid": "1",
                            "nodeid": 1
                        }
                    },
                    {
                        "single_node_interface": {
                            "dynamic": true,
                            "dynamic_ipv6_index": 2,
                            "nicid": "1",
                            "nodeid": 1
                        }
                    }
                ],
                "name": "Interface 1"
            }
        }
    ]
}
`

	var created_data_json = `
{
    "active_server_certificate_probing": true,
    "active_wait_time": "short",
    "admin_domain": "http://localhost:8082/7.4/elements/admin_domain/1",
    "alias_value": [],
    "allow_email_upn_lookup": false,
    "allow_long_userid_lookup": false,
    "allow_root_login": "disable_pwd_ssh",
    "antispoofing_node_ref": "http://localhost:8082/7.4/elements/single_fw/268583019/antispoofing/268616332",
    "antivirus": {
        "antivirus_enabled": false,
        "antivirus_http_proxy_enabled": false,
        "antivirus_proxy_port": 0,
        "antivirus_update": "never",
        "antivirus_update_day": "mo",
        "antivirus_update_time": 0,
        "virus_log_level": "undefined"
    },
    "auto_reboot_timeout": 10,
    "automatic_rules_settings": {
        "allow_auth_traffic": true,
        "allow_connections_to_dns_resolvers": true,
        "allow_connections_to_remote_dhcp_server": true,
        "allow_icmp_traffic_for_route_probing": true,
        "allow_listening_interfaces_to_dns_relay_port": true,
        "allow_no_nat": true,
        "log_level": "none"
    },
    "client_cert_identity_field": "",
    "connection_limit": 0,
    "connection_sync_mode": "disabled",
    "connection_timeout": [
        {
            "protocol": "tcp",
            "timeout": 1800
        },
        {
            "protocol": "udp",
            "timeout": 50
        },
        {
            "protocol": "icmp",
            "timeout": 5
        },
        {
            "protocol": "other",
            "timeout": 180
        }
    ],
    "contact_timeout": 120000,
    "control_plane_mode": "not_reserved",
    "custom_properties_profile": [],
    "default_nat": "false",
    "destination_server_certificate_cache_timeout": 1440,
    "disable_ahm": false,
    "discard_quic_if_cant_inspect": true,
    "dns_relay_interface": [],
    "domain_server_address": [],
    "dos_protection": "default_off",
    "dynamic_routing": {
        "antispoofing_ne_ref": [],
        "bgp": {
            "announced_ne_setting": [],
            "enabled": false
        },
        "ospfv2": {
            "enabled": false
        }
    },
    "enable_saml_for_bba": false,
    "enable_saml_for_ssl_vpn": false,
    "excluded_interface": -1,
    "file_reputation_settings": {
        "file_reputation_context": "disabled",
        "http_proxy": []
    },
    "granted_policy_ref": [],
    "ha_backup_unit_mode": "disabled",
    "icap_dlp_server_ref": [],
    "include_interfaces_for_control_plane": false,
    "inspection_cpu_balancing_mode": "default",
    "internal_gateway_ref": [
        "http://localhost:8082/7.4/elements/single_fw/268583019/internal_gateway/268439842"
    ],
    "ipv6_transition_mechanism": {
        "mode": "none"
    },
    "is_cert_auto_renewal": false,
    "is_config_encrypted": true,
    "is_fips_compatible_operating_mode": false,
    "is_fips_disable_engine_sginfo": false,
    "is_fips_disable_engine_upgrades": false,
    "is_icap_dlp_enabled": false,
    "is_loopback_tunnel_ip_address_enforced": false,
    "is_snort_enabled": false,
    "is_snort_file_defined": false,
    "is_virtual_defrag": false,
    "key": 268583019,
    "known_host_lists_ref": [],
    "l2fw_settings": {
        "bypass_overload_traffic": false,
        "tracking_mode": "normal"
    },
    "link": [
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019",
            "rel": "self",
            "type": "single_fw"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/export",
            "rel": "export"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/history",
            "rel": "history"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/lock",
            "rel": "lock"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/search_category_tags_from_element",
            "rel": "search_category_tags_from_element"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/contact_addresses",
            "rel": "contact_addresses"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/duplicate",
            "rel": "duplicate"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/refresh",
            "rel": "refresh"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/upload",
            "rel": "upload"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/remove_alternative_policies",
            "rel": "remove_alternative_policies"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/node",
            "rel": "nodes"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/interface",
            "rel": "interfaces"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/generate_snapshot",
            "rel": "generate_snapshot"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/add_route",
            "rel": "add_route"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/query_route",
            "rel": "query_route"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/upload_result",
            "rel": "upload_result"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/permissions",
            "rel": "permissions"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/block_list",
            "rel": "block_list"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/block_list/flush",
            "rel": "flush_block_list"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/join_domain",
            "rel": "join_domain"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/unjoin_domain",
            "rel": "unjoin_domain"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/alias_resolving",
            "rel": "alias_resolving"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/routing_monitoring",
            "rel": "routing_monitoring"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/neighbor_monitoring",
            "rel": "neighbor_monitoring"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/connection_monitoring",
            "rel": "connection_monitoring"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/blocklist_monitoring",
            "rel": "blocklist_monitoring"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/vpnsa_monitoring",
            "rel": "vpnsa_monitoring"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/user_monitoring",
            "rel": "user_monitoring"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/sslvpn_monitoring",
            "rel": "sslvpn_monitoring"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/web_auth_https/delete_certificate",
            "rel": "web_auth_https_delete_certificate"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/web_auth_https/delete_certificate_request",
            "rel": "web_auth_https_delete_certificate_request"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/web_auth_https/export_certificate",
            "rel": "web_auth_https_export_certificate"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/web_auth_https/export_certificate_request",
            "rel": "web_auth_https_export_certificate_request"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/web_auth_https/import_certificate",
            "rel": "web_auth_https_import_certificate"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/web_auth_https/generate_and_sign_certificate",
            "rel": "web_auth_https_generate_and_sign_certificate"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/web_auth_https/generate_certificate_request",
            "rel": "web_auth_https_generate_certificate_request"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/web_auth_https/sign_certificate_request",
            "rel": "web_auth_https_sign_certificate_request"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/pending_changes",
            "rel": "pending_changes"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/approve_all_changes",
            "rel": "approve_all_changes"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/disapprove_all_changes",
            "rel": "disapprove_all_changes"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/vpn_mapping",
            "rel": "vpn_mapping"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/ldap_replication",
            "rel": "ldap_replication"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/dynamic_routing_conf_preview",
            "rel": "dynamic_routing_conf_preview"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/validate_routing_antispoofing",
            "rel": "validate_routing_antispoofing"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/snort_configuration_file/export",
            "rel": "snort_configuration_file_export"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/snort_configuration_file/import",
            "rel": "snort_configuration_file_import"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/snort_configuration_file/delete",
            "rel": "snort_configuration_file_delete"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/internal_gateway",
            "rel": "internal_gateway",
            "type": "internal_gateway"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/routing/268571620",
            "rel": "routing",
            "type": "routing"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/antispoofing/268616332",
            "rel": "antispoofing",
            "type": "antispoofing"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/snapshot",
            "rel": "snapshots",
            "type": "snapshot"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/physical_interface",
            "rel": "physical_interface",
            "type": "physical_interface"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/tunnel_interface",
            "rel": "tunnel_interface",
            "type": "tunnel_interface"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/vpn_broker_interface",
            "rel": "vpn_broker_interface",
            "type": "vpn_broker_interface"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/modem_interface",
            "rel": "modem_interface",
            "type": "modem_interface"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/adsl_interface",
            "rel": "adsl_interface",
            "type": "adsl_interface"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/wireless_interface",
            "rel": "wireless_interface",
            "type": "wireless_interface"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/switch_interface",
            "rel": "switch_interface",
            "type": "switch_interface"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/ssh_host_key",
            "rel": "ssh_host_key",
            "type": "ssh_host_key"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node",
            "rel": "firewall_node",
            "type": "firewall_node"
        }
    ],
    "link_usage_exception_rules": [],
    "location_ref": "http://localhost:8082/7.4/elements/location/-1",
    "locked": false,
    "log_moderation": [
        {
            "burst": 1000,
            "log_event": "antispoofing",
            "rate": 100
        },
        {
            "burst": 20000,
            "log_event": "discard",
            "rate": 5000
        },
        {
            "burst": 80000,
            "log_event": "allow",
            "rate": 40000
        }
    ],
    "log_server_ref": "http://localhost:8082/7.4/elements/log_server/268581077",
    "log_spooling_policy": "discard",
    "multicast_routing_mode": "none",
    "name": "tf_single_fw_dynamic_interface",
    "nat_definition": [],
    "nodes": [
        {
            "firewall_node": {
                "activate_test": false,
                "appliance_info": {
                    "first_upload_time": 0,
                    "initial_contact_time": 0,
                    "initial_license_remaining_days": 0
                },
                "disabled": false,
                "key": 268583020,
                "link": [
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020",
                        "rel": "self",
                        "type": "firewall_node"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/fetch_license",
                        "rel": "fetch"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/bind_license",
                        "rel": "bind"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/unbind_license",
                        "rel": "unbind"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/cancel_unbind_license",
                        "rel": "cancel_unbind"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/initial_contact",
                        "rel": "initial_contact"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/appliance_status",
                        "rel": "appliance_status"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/status",
                        "rel": "status"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/go_online",
                        "rel": "go_online"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/go_offline",
                        "rel": "go_offline"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/go_standby",
                        "rel": "go_standby"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/lock_online",
                        "rel": "lock_online"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/lock_offline",
                        "rel": "lock_offline"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/reset_user_db",
                        "rel": "reset_user_db"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/diagnostic",
                        "rel": "diagnostic"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/send_diagnostic",
                        "rel": "send_diagnostic"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/reboot",
                        "rel": "reboot"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/power_off",
                        "rel": "power_off"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/reset_to_factory",
                        "rel": "reset_to_factory"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/sginfo",
                        "rel": "sginfo"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/ssh",
                        "rel": "ssh"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/change_ssh_pwd",
                        "rel": "change_ssh_pwd"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/time_sync",
                        "rel": "time_sync"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/certificate_info",
                        "rel": "certificate_info"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/dynamic_element_update",
                        "rel": "dynamic_element_update"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/pki_export_certificate_request",
                        "rel": "pki_export_certificate_request"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/pki_import_certificate",
                        "rel": "pki_import_certificate"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/pki_delete_certificate_request",
                        "rel": "pki_delete_certificate_request"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/pki_abort_certificate_request",
                        "rel": "pki_abort_certificate_request"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/pki_certificate_info",
                        "rel": "pki_certificate_info"
                    },
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/firewall_node/268583020/pki_start_certificate_renewal",
                        "rel": "pki_start_certificate_renewal"
                    }
                ],
                "loopback_node_dedicated_interface": [],
                "name": "myfwnode",
                "nodeid": 1,
                "tests": []
            }
        }
    ],
    "nondecrypted_ca_certificate_ref": [],
    "nondecrypted_tls_server_credentials_ref": [],
    "ntp_settings": {
        "ntp_enable": false,
        "ntp_server_ref": []
    },
    "opcua_client_x509_credentials": [],
    "opcua_decryption_mode": "none",
    "opcua_server_x509_credentials": [],
    "passive_discard_mode": false,
    "physicalInterfaces": [
        {
            "physical_interface": {
                "aggregate_mode": "none",
                "arp_entry": [],
                "comment": "dynamic IPv4/IPv6 intertace",
                "cvi_mode": "none",
                "dhcp_server_on_interface": {
                    "default_lease_time": 7200,
                    "dhcp_range_per_node": []
                },
                "duplicate_address_detection": true,
                "include_prefix_info_option_flag": true,
                "interface_id": "1",
                "interfaces": [
                    {
                        "single_node_interface": {
                            "auth_request": false,
                            "auth_request_source": false,
                            "automatic_default_route": true,
                            "backup_heartbeat": false,
                            "backup_mgt": false,
                            "domain_specific_dns_queries_source": false,
                            "dynamic": true,
                            "dynamic_index": 1,
                            "key": 268543713,
                            "nicid": "1",
                            "nodeid": 1,
                            "outgoing": false,
                            "pppoa": false,
                            "pppoe": false,
                            "primary_heartbeat": false,
                            "primary_mgt": false,
                            "relayed_by_dhcp": false,
                            "reverse_connection": false,
                            "vrrp": false,
                            "vrrp_id": -1,
                            "vrrp_priority": -1
                        }
                    },
                    {
                        "single_node_interface": {
                            "auth_request": false,
                            "auth_request_source": false,
                            "automatic_default_route": true,
                            "backup_heartbeat": false,
                            "backup_mgt": false,
                            "domain_specific_dns_queries_source": false,
                            "dynamic": true,
                            "dynamic_ipv6_index": 2,
                            "key": 268543714,
                            "nicid": "1",
                            "nodeid": 1,
                            "outgoing": false,
                            "pppoa": false,
                            "pppoe": false,
                            "primary_heartbeat": false,
                            "primary_mgt": false,
                            "relayed_by_dhcp": false,
                            "reverse_connection": false,
                            "vrrp": false,
                            "vrrp_id": -1,
                            "vrrp_priority": -1
                        }
                    }
                ],
                "key": 268496780,
                "link": [
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/physical_interface/268496780",
                        "rel": "self",
                        "type": "physical_interface"
                    }
                ],
                "lldp_mode": "disabled",
                "log_moderation": [
                    {
                        "burst": 1000,
                        "log_event": "antispoofing",
                        "rate": 100
                    },
                    {
                        "burst": 20000,
                        "log_event": "discard",
                        "rate": 5000
                    },
                    {
                        "burst": 80000,
                        "log_event": "allow",
                        "rate": 40000
                    }
                ],
                "managed_address_flag": false,
                "mtu": -1,
                "name": "Interface 1",
                "other_configuration_flag": false,
                "override_engine_settings": false,
                "override_log_moderation_settings": false,
                "qos_limit": 0,
                "qos_mode": "no_qos",
                "route_replies_back_mode": false,
                "router_advertisement": false,
                "set_autonomous_address_flag": true,
                "shared_interface": false,
                "syn_mode": "default",
                "sync_parameter": {
                    "full_sync_interval": 5000,
                    "heartbeat_group_ip": "224.0.0.221",
                    "incr_sync_interval": 50,
                    "statesync_group_ip": "224.0.0.222",
                    "sync_mode": "sync_all",
                    "sync_security": "sign"
                },
                "virtual_engine_vlan_ok": false,
                "virtual_resource_settings": [],
                "vlanInterfaces": []
            }
        },
        {
            "physical_interface": {
                "aggregate_mode": "none",
                "arp_entry": [],
                "cvi_mode": "none",
                "dhcp_server_on_interface": {
                    "default_lease_time": 7200,
                    "dhcp_range_per_node": []
                },
                "duplicate_address_detection": true,
                "include_prefix_info_option_flag": true,
                "interface_id": "0",
                "interfaces": [
                    {
                        "single_node_interface": {
                            "address": "192.168.100.14",
                            "auth_request": false,
                            "auth_request_source": false,
                            "automatic_default_route": false,
                            "backup_heartbeat": false,
                            "backup_mgt": false,
                            "domain_specific_dns_queries_source": false,
                            "dynamic": false,
                            "key": 268543715,
                            "network_value": "192.168.100.0/24",
                            "nicid": "0",
                            "nodeid": 1,
                            "outgoing": false,
                            "pppoa": false,
                            "pppoe": false,
                            "primary_heartbeat": false,
                            "primary_mgt": true,
                            "relayed_by_dhcp": false,
                            "reverse_connection": false,
                            "vrrp": false,
                            "vrrp_id": -1,
                            "vrrp_priority": -1
                        }
                    }
                ],
                "key": 268496781,
                "link": [
                    {
                        "href": "http://localhost:8082/7.4/elements/single_fw/268583019/physical_interface/268496781",
                        "rel": "self",
                        "type": "physical_interface"
                    }
                ],
                "lldp_mode": "disabled",
                "log_moderation": [
                    {
                        "burst": 1000,
                        "log_event": "antispoofing",
                        "rate": 100
                    },
                    {
                        "burst": 20000,
                        "log_event": "discard",
                        "rate": 5000
                    },
                    {
                        "burst": 80000,
                        "log_event": "allow",
                        "rate": 40000
                    }
                ],
                "managed_address_flag": false,
                "mtu": -1,
                "name": "Interface 0",
                "other_configuration_flag": false,
                "override_engine_settings": false,
                "override_log_moderation_settings": false,
                "qos_limit": 0,
                "qos_mode": "no_qos",
                "route_replies_back_mode": false,
                "router_advertisement": false,
                "set_autonomous_address_flag": true,
                "shared_interface": false,
                "syn_mode": "default",
                "sync_parameter": {
                    "full_sync_interval": 5000,
                    "heartbeat_group_ip": "224.0.0.221",
                    "incr_sync_interval": 50,
                    "statesync_group_ip": "224.0.0.222",
                    "sync_mode": "sync_all",
                    "sync_security": "sign"
                },
                "virtual_engine_vlan_ok": false,
                "virtual_resource_settings": [],
                "vlanInterfaces": []
            }
        }
    ],
    "policy_route": [],
    "quic_enabled": true,
    "read_only": false,
    "rollback_timeout": 60,
    "routing_node_ref": "http://localhost:8082/7.4/elements/single_fw/268583019/routing/268571620",
    "saml_bba_clock_skew": 0,
    "saml_settings": [],
    "saml_ssl_vpn_clock_skew": 0,
    "sandbox_type": "none",
    "scan_detection": {
        "log_level": "stored",
        "scan_detection_icmp_events": 220,
        "scan_detection_icmp_timewindow": 1,
        "scan_detection_icmp_unit": "minute",
        "scan_detection_tcp_events": 220,
        "scan_detection_tcp_timewindow": 1,
        "scan_detection_tcp_unit": "minute",
        "scan_detection_type": "default off",
        "scan_detection_udp_events": 220,
        "scan_detection_udp_timewindow": 1,
        "scan_detection_udp_unit": "minute"
    },
    "server_credential": [],
    "sidewinder_proxy_enabled": false,
    "slow_request_block_list_timeout": 0,
    "slow_request_sensitivity": "low",
    "snmp_interface": [],
    "ssh_host_key": [],
    "ssh_passwordless_login": "allow",
    "ssm_advanced_setting": [],
    "static_multicast_route": [],
    "strict_tcp_mode": false,
    "syn_flood_sensitivity": "medium",
    "syn_mode": "off",
    "system": false,
    "system_key": -1,
    "tcp_reset_sensitivity": "off",
    "tester_parameters": {
        "alert_interval": 3600,
        "auto_recovery": true,
        "boot_delay": 30,
        "boot_recovery": true,
        "restart_delay": 5,
        "status_delay": 5
    },
    "tests": [],
    "tls_crl_checks": false,
    "tls_deny_decrypting": true,
    "tracking_mode": "normal",
    "trashed": false,
    "ts_settings": {
        "http_proxy": [],
        "local_url_categorization_enabled": false,
        "ts_enabled": false
    }
}

`

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	tflog.Info(ctx, "test log message", map[string]any{
		"key": "value",
	})

	var configData schema.SingleFirewallResourceModel
	var createdData schema.SingleFirewallResourceModel

	if err := apijson.UnmarshalRoot([]byte(config_data_json), &configData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(created_data_json), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal configData: %v", err)
	}

	err := resource.MergeResourceModels(ctx, &createdData, &configData)
	if err != nil {
		t.Fatalf("resource.MergeResourceModels failed: %v", err)
	}

	apiData, err := apijson.MarshalRoot(configData)
	fmt.Printf("After merge, configData: %s\n", apiData)

	if !configData.LocationRef.IsNull() && !configData.LocationRef.IsUnknown() {
		t.Errorf("Expected LocationRef to be null after merge, got: %v", configData.LocationRef)
	}

	physicalInterface1 := (*configData.PhysicalInterfaces)[1].PhysicalInterface
	singleNodeInterface0 := (*physicalInterface1.Interfaces)[0].SingleNodeInterface
	expectedSingleNodeKey0 := int64(268543713)
	actualSingleNodeKey0 := singleNodeInterface0.Key.ValueInt64()
	if actualSingleNodeKey0 != expectedSingleNodeKey0 {
		t.Errorf("Expected SingleNodeInterface Key to be %d, got %d", expectedSingleNodeKey0, actualSingleNodeKey0)
	}
	singleNodeInterface1 := (*physicalInterface1.Interfaces)[1].SingleNodeInterface
	expectedSingleNodeKey1 := int64(268543714)
	actualSingleNodeKey1 := singleNodeInterface1.Key.ValueInt64()
	if actualSingleNodeKey1 != expectedSingleNodeKey1 {
		t.Errorf("Expected SingleNodeInterface Key to be %d, got %d", expectedSingleNodeKey1, actualSingleNodeKey1)
	}
}

func TestMergeIpAccessList(t *testing.T) {
	setupTest()

	// terraform config
	var config_data_json = `
{
    "entries": [
        {
            "ip_access_list_entry": {
                "action": "deny",
                "subnet": "172.16.16.0/21"
            }
        },
        {
            "ip_access_list_entry": {
                "action": "permit",
                "subnet": "192.168.100.0/24"
            }
        }
    ],
    "name": "tf_ip_access_list"
}
`

	// from the smc
	var created_data_json = `
{
    "entries": [
        {
            "ip_access_list_entry": {
                "action": "permit",
                "key": 268435477,
                "subnet": "192.168.100.0/24"
            }
        },
        {
            "ip_access_list_entry": {
                "action": "deny",
                "key": 268435476,
                "subnet": "172.16.16.0/21"
            }
        }
    ],
    "key": 268435466,
    "name": "tf_ip_access_list"
}
`

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var configData schema.IpAccessListResourceModel
	var createdData schema.IpAccessListResourceModel

	if err := apijson.UnmarshalRoot([]byte(config_data_json), &configData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(created_data_json), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal configData: %v", err)
	}

	err := resource.MergeResourceModels(ctx /*src=*/, &createdData /*dest=*/, &configData)
	if err != nil {
		t.Fatalf("resource.MergeResourceModels failed: %v", err)
	}

	apiData, err := apijson.MarshalRoot(configData)
	fmt.Printf("After merge, configData: %s\n", apiData)
	if string(apiData) != `{"entries":[{"ip_access_list_entry":{"action":"deny","key":268435476,"subnet":"172.16.16.0/21"}},{"ip_access_list_entry":{"action":"permit","key":268435477,"subnet":"192.168.100.0/24"}}],"key":268435466,"name":"tf_ip_access_list"}` {
		t.Errorf("Merged data does not match expected output")
	}
}

func TestMergeRoutingNode(t *testing.T) {
	setupTest()

	var created_data_json = `
{
    "href": "http://localhost:8082/7.4/elements/single_fw/268581978",
    "key": 268570444,
    "level": "engine_cluster",
    "link": [
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268581978/routing/268570444",
            "rel": "self",
            "type": "routing"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268581978/routing/268570444/history",
            "rel": "history"
        },
        {
            "href": "http://localhost:8082/7.4/elements/single_fw/268581978/routing/268570444/lock",
            "rel": "lock"
        }
    ],
    "name": "tf_single_fw",
    "probe_ecmp": 0,
    "probe_interval": 0,
    "probe_metric": 0,
    "probe_retry_count": 0,
    "probe_test": "not_enabled",
    "read_only": false,
    "related_element_type": "single_fw",
    "routing_node": [
        {
            "exclude_from_ip_counting": false,
            "href": "http://localhost:8082/7.4/elements/single_fw/268581978/physical_interface/268496383",
            "key": 268570445,
            "level": "interface",
            "link": [
                {
                    "href": "http://localhost:8082/7.4/elements/single_fw/268581978/routing/268570445",
                    "rel": "self",
                    "type": "routing"
                }
            ],
            "name": "Interface 0",
            "nic_id": "0",
            "probe_ecmp": 0,
            "probe_interval": 0,
            "probe_metric": 0,
            "probe_retry_count": 0,
            "probe_test": "not_enabled",
            "read_only": false,
            "related_element_type": "physical_interface",
            "routing_node": [
                {
                    "href": "http://localhost:8082/7.4/elements/network/268581977",
                    "invalid": false,
                    "ip": "192.168.100.0/24",
                    "key": 268570446,
                    "level": "network",
                    "link": [
                        {
                            "href": "http://localhost:8082/7.4/elements/single_fw/268581978/routing/268570446",
                            "rel": "self",
                            "type": "routing"
                        }
                    ],
                    "name": "tf_sample_network",
                    "probe_ecmp": 0,
                    "probe_interval": 0,
                    "probe_metric": 0,
                    "probe_retry_count": 0,
                    "probe_test": "not_enabled",
                    "read_only": false,
                    "related_element_type": "network",
                    "routing_node": [
                        {
                            "href": "http://localhost:8082/7.4/elements/router/268581976",
                            "invalid": false,
                            "ip": "192.168.100.254",
                            "key": 268570447,
                            "level": "gateway",
                            "link": [
                                {
                                    "href": "http://localhost:8082/7.4/elements/single_fw/268581978/routing/268570447",
                                    "rel": "self",
                                    "type": "routing"
                                }
                            ],
                            "name": "tf_sample_router",
                            "probe_ecmp": 0,
                            "probe_interval": 0,
                            "probe_metric": 0,
                            "probe_retry_count": 0,
                            "probe_test": "not_enabled",
                            "read_only": false,
                            "related_element_type": "router",
                            "routing_node": [
                                {
                                    "href": "http://localhost:8082/7.4/elements/network/2",
                                    "invalid": false,
                                    "ip": "0.0.0.0/0",
                                    "key": 268570448,
                                    "level": "any",
                                    "link": [
                                        {
                                            "href": "http://localhost:8082/7.4/elements/single_fw/268581978/routing/268570448",
                                            "rel": "self",
                                            "type": "routing"
                                        }
                                    ],
                                    "name": "Any network",
                                    "probe_ecmp": 0,
                                    "probe_interval": 0,
                                    "probe_metric": 0,
                                    "probe_retry_count": 0,
                                    "probe_test": "not_enabled",
                                    "read_only": false,
                                    "related_element_type": "network",
                                    "routing_node": [],
                                    "system": false
                                }
                            ],
                            "system": false
                        }
                    ],
                    "system": false,
                    "to_delete": false
                }
            ],
            "system": false,
            "to_delete": false
        }
    ],
    "system": false,
    "to_delete": false
}


`

	var config_data_json = `
{
    "level": "engine_cluster",
    "routing_node": [
        {
            "level": "interface",
            "name": "Interface 0",
            "routing_node": [
                {
                    "href": "http://localhost:8082/7.4/elements/network/268581977",
                    "level": "network",
                    "routing_node": [
                        {
                            "href": "http://localhost:8082/7.4/elements/router/268581976",
                            "level": "gateway",
                            "routing_node": [
                                {
                                    "href": "http://localhost:8082/7.4/elements/network/2",
                                    "level": "any"
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    ]
}
`

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	var configData schema.RoutingNodeResourceModel
	var createdData schema.RoutingNodeResourceModel

	if err := apijson.UnmarshalRoot([]byte(config_data_json), &configData); err != nil {
		t.Fatalf("Failed to unmarshal createdData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(created_data_json), &createdData); err != nil {
		t.Fatalf("Failed to unmarshal configData: %v", err)
	}

	err := resource.MergeResourceModels(ctx, &createdData, &configData)
	if err != nil {
		t.Fatalf("resource.MergeResourceModels failed: %v", err)
	}

	apiData, err := apijson.MarshalRoot(configData)
	fmt.Printf("After merge, configData: %s\n", apiData)

	if configData.Key.ValueInt64() != int64(268570444) {
		t.Errorf("Expected key to be 268570444")
	}
	if len(*configData.RoutingNode) != 1 {
		t.Errorf("Expected 1 routing_node at root level after merge")
	}

	if (*configData.RoutingNode)[0].Key.ValueInt64() != int64(268570445) {
		t.Errorf("Expected first routing_node key to be 268570445")
	}

	if len(*(*configData.RoutingNode)[0].RoutingNode) != 1 {
		t.Errorf("Expected 1 routing_node at interface level after merge")
	}

	if (*(*configData.RoutingNode)[0].RoutingNode)[0].Key.ValueInt64() != int64(268570446) {
		t.Errorf("Expected first routing_node at network level key to be 268570446")
	}

	if len(*(*(*configData.RoutingNode)[0].RoutingNode)[0].RoutingNode) != 1 {
		t.Errorf("Expected 1 routing_node at gateway level after merge")
	}
	if (*(*(*configData.RoutingNode)[0].RoutingNode)[0].RoutingNode)[0].Key.ValueInt64() != int64(268570447) {
		t.Errorf("Expected first routing_node at gateway level key to be 268570447")
	}
	if len(*(*(*(*configData.RoutingNode)[0].RoutingNode)[0].RoutingNode)[0].RoutingNode) != 1 {
		t.Errorf("Expected 1 routing_node at any level after merge")
	}
	if (*(*(*(*configData.RoutingNode)[0].RoutingNode)[0].RoutingNode)[0].RoutingNode)[0].Key.ValueInt64() != int64(268570448) {
		t.Errorf("Expected first routing_node at any level key to be 268570448")
	}

}

// skipped for now. merge full does not work (not needed currently)
func TestMergeHostFull(t *testing.T) {
	setupTest()

	t.Skip("merge full does not work")
	var smc_data_json = `{
		"address": "192.168.1.2",
		"admin_domain": "http://localhost:8082/7.3/elements/admin_domain/1",
		"comment": "Comment from smc",
		"key": 5253,
		"link": [
			{
				"href": "http://localhost:8082/7.3/elements/host/5253",
				"rel": "self",
				"type": "host"
			}
		],
		"name": "AExampleHost",
		"secondary": [
			"212.20.1.1",
			"123.6.5.19"
		]
	}`

	var tf_data_json = `{
		"address": "192.168.1.3",
		"comment": "Comment from Terraform",
	}`

	var tfData schema.HostResourceModel
	var smcData schema.HostResourceModel
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	if err := apijson.UnmarshalRoot([]byte(smc_data_json), &smcData); err != nil {
		t.Fatalf("Failed to unmarshal smcData: %v", err)
	}
	if err := apijson.UnmarshalRoot([]byte(tf_data_json), &tfData); err != nil {
		t.Fatalf("Failed to unmarshal data: %v", err)
	}

	err := resource.MergeResourceFull(ctx /*src*/, &smcData /*dest*/, &tfData)
	if err != nil {
		t.Fatalf("resource.MergeResourceModels failed: %v", err)
	}

	fmt.Printf("After merge, data: %+v\n", tfData)
}
