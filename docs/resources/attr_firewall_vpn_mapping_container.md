---
page_title: "vpn_mapping_entry"
subcategory: ""
description: |-
  This represents the container for VPN Mapping, which includes an internal gateway reference, a referenced VPN, and the usage of gateway nodes (satellite, central, mobile).
---

# vpn_mapping_entry (Nested-Attribute)

This represents the container for VPN Mapping, which includes an internal gateway reference, a referenced VPN, and the usage of gateway nodes (satellite, central, mobile).




## Simple Attributes
- `gateway_ref` (String) This represents the Internal Gateway, which is used for managing VPN connections and related settings.
- `vpn_ref` (String) This represents a Policy Based Virtual Private Network (VPN), which is used to establish secure connections over unsecured networks. It includes various configurations such as NAT rules, mobile VPN topology modes, and associated profiles.

## Nested Attributes
- `gateway_nodes_usage` (Single Block, see [here](attr_gateway_nodes_usage.md)) 

