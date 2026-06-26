---
page_title: "smc_gateway_tunnel"
subcategory: "vpn"
description: |-
  This represents a gateway tunnel, which is used to manage the logical tunnels between gateway nodes in the VPN topology.
---

# smc_gateway_tunnel (Sub-resource)

This represents a gateway tunnel, which is used to manage the logical tunnels between gateway nodes in the VPN topology.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `comment` (String) An optional comment for the element. This field is not required.
- `enabled` (Boolean) Indicates whether this logical tunnel is enabled or not.
- `gateway_node_1` (String) This represents a gateway node, which is used to manage the VPN topology and its nodes, including their usage and relationships.
- `gateway_node_2` (String) This represents a gateway node, which is used to manage the VPN topology and its nodes, including their usage and relationships.
- `hashed_preshared_key` (String) A hashed version of the pre-shared key, used for secure storage and comparison.
- `name` (String) Name of the object.
- `ppk_ref` (String) Base class for Post-Quantum Preshared Key (PPK) elements.
- `preshared_key` (String) The pre-shared key used for authentication in the gateway tunnel.
- `vpn_profile` (String) This represents a VPN Profile. It contains settings for IKE and IPsec lifetimes, keep-alive options, certificate authorities, and authentication methods.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
