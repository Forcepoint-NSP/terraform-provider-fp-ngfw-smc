---
page_title: "smc_rbvpn_tunnel"
subcategory: "vpn"
description: |-
  This represents a Route-Based VPN Tunnel. It defines the properties and configuration of a route-based VPN tunnel, including its sides, encryption mode, and other settings.
---

# smc_rbvpn_tunnel (Resource)

This represents a Route-Based VPN Tunnel. It defines the properties and configuration of a route-based VPN tunnel, including its sides, encryption mode, and other settings.

## Examples

- [RBVPN Tunnel Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/sdwan/route_based_vpn/ip_ip_mode_internal_gateway/main.tf)

Defines a Route-Based VPN tunnel between two endpoints in an SD-WAN design.

```hcl
resource "smc_rbvpn_tunnel" "gre_single_fw1_single_fw2" {
  depends_on     = [smc_internal_endpoint.ep_192_168_100_14, smc_internal_endpoint.ep_10_0_50_14]
  comment        = var.resource_comment
  enabled        = true
  mtu            = 0
  name           = "GRE-tf_single_fw1 to_tf_single_fw2"
  pmtu_discovery = true
  rbvpn_tunnel_side_a {
    endpoint_ref         = data.smc_sub_href.single_fw1_ep.id
    gateway_ref          = data.smc_sub_href.single_fw1_internal_gateway_ref.id
    tunnel_interface_ref = data.smc_sub_href.single_fw1_tunnel_interface_1000_ref.id
  }
  rbvpn_tunnel_side_b {
    endpoint_ref         = data.smc_sub_href.single_fw2_ep.id
    gateway_ref          = data.smc_sub_href.single_fw2_internal_gateway_ref.id
    tunnel_interface_ref = data.smc_sub_href.single_fw2_tunnel_interface_1234_ref.id
  }
  ttl               = 0
  tunnel_encryption = "no_encryption"
  tunnel_mode       = "ipip"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) A user-defined comment for the Route-Based VPN Tunnel. This field can be used to provide additional information or context about the tunnel.
- `enabled` (Boolean) Indicates whether the Route-Based VPN Tunnel is enabled or not.
- `gre_keepalive` (Boolean) Indicates whether GRE Keepalive is enabled for the Route-Based VPN Tunnel. This feature allows the tunnel to maintain its state by sending periodic keepalive messages.
- `gre_keepalive_period` (Number) The GRE Keepalive period for the Route-Based VPN Tunnel, in seconds. This value determines how often keepalive messages are sent to maintain the tunnel's state.
- `gre_keepalive_retry` (Number) The GRE Keepalive retry count for the Route-Based VPN Tunnel. This value determines how many times the tunnel will retry sending a keepalive message before considering the tunnel down.
- `hashed_preshared_key` (String) The hashed pre-shared key for the Route-Based VPN Tunnel. This is a one-way hash of the pre-shared key, used for secure storage and comparison.
- `mtu` (Number) The Maximum Transmission Unit (MTU) for the Route-Based VPN Tunnel. This value determines the largest packet size that can be transmitted over the tunnel.
- `name` (String) Name of the object.
- `pmtu_discovery` (Boolean) Indicates whether Path MTU Discovery is enabled for the Route-Based VPN Tunnel. This feature allows the tunnel to dynamically adjust its MTU based on the path characteristics.
- `ppk_ref` (String) Base class for Post-Quantum Preshared Key (PPK) elements.
- `preshared_key` (String) The pre-shared key used for the Route-Based VPN Tunnel. This key is used for authentication and encryption.
- `ttl` (Number) The Time To Live (TTL) for the Route-Based VPN Tunnel. This value determines how long the tunnel's packets are allowed to live in the network before being discarded.
- `tunnel_encryption` (String) The encryption mode of the Route-Based VPN Tunnel, indicating how the tunnel's data is encrypted.
- `tunnel_group_ref` (String) This represents a Route Based VPN Tunnel Group. It is used to group tunnels in the home/monitoring view. It contains a reference to the Link Usage Profile that defines the link usage for the tunnels in this group.
- `tunnel_mode` (String) The mode of the Route-Based VPN Tunnel, indicating how the tunnel is configured and operates.
- `tunnel_mode_vpn_ref` (String) This represents a Policy Based Virtual Private Network (VPN), which is used to establish secure connections over unsecured networks. It includes various configurations such as NAT rules, mobile VPN topology modes, and associated profiles.
- `vpn_profile_ref` (String) 

## Nested Attributes
- `geneve_settings` (Single Block, see [here](attr_rbvpn_geneve_settings.md)) 
- `rbvpn_tunnel_side_a` (Single Block, see [here](route_based_vpn_tunnel_side.md)) 
- `rbvpn_tunnel_side_b` (Single Block, see [here](route_based_vpn_tunnel_side.md)) 

## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
