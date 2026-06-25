---
page_title: "Tips for Writing Terraform Configuration"
subcategory: ""
description: |-
  Practical tips and tools to help you write, Terraform configurations for the SMC provider.
---

# Tips for Writing Terraform Configuration

## Study the examples provided in the github repository

The [github repository](https://github.com/Forcepoint/terraform-provider-fp-ngfw-smc)
contains many examples to help you write your own configurations.

Here is the list of examples we provide:

```
├── getting_started
├── engines
│   ├── dynamic_routing
│   │   ├── BGP
│   │   │   ├── as_path_access_list
│   │   │   ├── autonomous_system
│   │   │   ├── bgp_connection_profile
│   │   │   ├── bgp_peering
│   │   │   ├── bgp_profile
│   │   │   ├── community_access_list
│   │   │   ├── exernal_bgp_peers
│   │   │   └── external_bgp_peers
│   │   ├── ip_access_list
│   │   ├── ip_prefix_list
│   │   ├── ipv6_access_list
│   │   ├── ipv6_prefix_list
│   │   ├── OSPFv2
│   │   │   ├── ospfv2_areas
│   │   │   ├── ospfv2_domain_settings
│   │   │   ├── ospfv2_interface_settings
│   │   │   ├── ospfv2_key_chain
│   │   ├── PIM
│   │   │   ├── igmp_querier_settings
│   │   │   ├── pim_ipv4_interface_settings
│   │   │   └── pim_profiles
│   │   ├── route_map
│   │   └── routing_node_bgp
│   ├── fw_cluster
│   │   ├── fw_cluster_with_cvi-ndi_interface
│   │   ├── fw_cluster_with_cvi_only_interface
│   │   ├── fw_cluster_with_layer2_interface
│   │   ├── fw_cluster_with_ndi_only_interface
│   │   ├── fw_cluster_with_tunnel_interface
│   │   └── fw_cluster_with_vpn_broker_interface
│   ├── location
│   ├── routing_node
│   ├── routing_node_with_netlinks_and_tunnels
│   └── single_fw
│       ├── contact_addr
│       ├── custom_domain
│       ├── single_fw_actions
│       ├── single_fw_dynamic_interfaces
│       ├── single_fw_layer2_interfaces
│       ├── single_fw_many_interfaces
│       ├── single_fw_modem_interfaces
│       ├── single_fw_tunnel_interfaces
│       ├── single_fw_vlan_interfaces
│       ├── single_fw_vpn_broker_interfaces
│       └── single_fw_wireless_interface
├── network_elements
│   ├── address_range
│   ├── alias
│   ├── dhcp_server
│   ├── domain_name
│   ├── expression
│   ├── group
│   ├── host
│   ├── interface_zone
│   ├── ip_list
│   ├── netlink
│   ├── network
│   ├── router
│   └── servers
│       ├── active_directory_server
│       ├── dhcp_server
│       ├── elasticsearch_cluster
│       ├── external_dns_server
│       ├── icap_server
│       ├── ldap_server
│       ├── ntp_server
│       ├── proxy_server
│       ├── radius_server
│       ├── smtp_server
│       └── tacacs_server
├── policies
│   ├── fw_ipv4_access_rule
│   ├── fw_ipv4_access_rule_rank
│   ├── fw_ipv4_access_rule_section
│   ├── fw_ipv4_nat_rules
│   ├── fw_policy
│   ├── fw_policy_logging_options
│   ├── fw_policy_vpn
│   ├── fw_policy_with_match_expression
│   ├── fw_template_policy
│   └── sub_ipv4_fw_policy
├── sdwan
│   ├── central_gateway_node
│   ├── external_gateway
│   ├── internal_endpoint
│   ├── policy_based_vpn
│   │   ├── central_gateway_node
│   │   ├── satellite_gateway_node
│   ├── route_based_vpn
│   │   ├── geneve_mode
│   │   ├── gre_tunnel_no_encryption
│   │   ├── gre_tunnel_transport_mode
│   │   ├── gre_tunnel_tunnel_mode
│   │   ├── ip_ip_mode_external_gateway
│   │   ├── ip_ip_mode_internal_gateway
│   │   ├── site_mode_internal_gateway
│   │   └── vpn_mode
│   └── vpn
├── services
│   ├── icmp_service
│   ├── ip-proto_service
│   ├── service_group
│   ├── sun-rpc_service
│   ├── tcp_service
│   └── udp_service
```


## Using smc-explorer

[smc-explorer](https://github.com/Forcepoint/fp-ngfw-smc-explorer) is
a command-line tool for exploring, querying, and exporting SMC
elements. It helps you:

- List resource types and their instances
- Show resources or sub-resources in HCL1, HCL2 or JSON formats
- Discover resource URLs for use with `from_ref` and `id`
- Filter or clean attributes for clarity

Common usage:
```sh
smc-explorer list                                            # List all resource types
smc-explorer list fw_policy                                  # List all policies
smc-explorer show single_fw/myfw                             # Show resource in HCL
smc-explorer show 'fw_policy/Lab FW1/fw_ipv4_access_rules/Rule @2097357.0' # show sub-resources via "link"
smc-explorer show single_fw/Plano -f json --raw              # Show resource in JSON.
smc-explorer get-url host/AExampleHost                       # Get resource URL by name
smc-explorer delete host/AExampleHost                        # Delete a resource
smc-explorer show single_fw/myfw -x -s alias_value,antivirus # Skip attributes
```

## Using IDE Smart Completion

Modern IDEs such as Visual Studio Code, with the [Terraform
extension](https://marketplace.visualstudio.com/items?itemName=HashiCorp.terraform),
provide smart completion (IntelliSense) for Terraform files. This
helps you:

- Discover available resource types and attributes
- Avoid typos and schema errors
- Navigate documentation and examples directly in the editor

Enable the Terraform extension in your IDE for a more productive
authoring experience.

## Prefer attribute references over hardcoded values

When a resource attribute (such as a firewall interface's `network_value`) mirrors a
value that belongs to another resource, reference the attribute directly instead of
copying the string:

```hcl
# Preferred — creates an explicit dependency, update ordering is guaranteed
network_value = smc_network.my_net.ipv4_network

# Avoid — no ordering guarantee when both resources are updated together
# network_value = "10.100.100.0/24"
```

This matters during `terraform apply` when both resources are modified at the same time.
Without an attribute reference, some tools may update them concurrently, which can cause
HTTP 409 conflicts on the SMC API (ETag mismatch). With an attribute reference, the
network resource is always fully updated before the firewall update begins.

See [Relationship Between Resources — Explicit attribute references during updates](./30-relationship-between-resources.md) for a detailed explanation.

## Dumping the Provider Schema

To inspect the full schema of the SMC Terraform provider, you can use
the following command:

```sh
terraform providers schema -json | jq .
```

This outputs the provider's schema in JSON format, which you can
search or filter using `jq`. This is helpful for advanced users who
want to inspect available resources, data sources, and attributes.
