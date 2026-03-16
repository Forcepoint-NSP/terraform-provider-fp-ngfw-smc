---
page_title: "smc_internal_endpoint"
subcategory: "vpn"
description: |-
  This represents an internal endpoint for VPN connections, supporting both tunnel and portal types. It includes properties such as physical interface and deducted name.
---

# smc_internal_endpoint (Sub-resource)

This represents an internal endpoint for VPN connections, supporting both tunnel and portal types. It includes properties such as physical interface and deducted name.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/sdwan/policy_based_vpn/central_gateway_node/main.tf) for a minimal example

This example creates an internal endpoint (used for VPN or routing) tied to a single firewall.

```hcl
resource "smc_internal_endpoint" "ep_212_20_1_1" {
  id = format("%s#%s", smc_single_fw.tf_single_fw1.link.internal_gateway,
  "*/internal_endpoint/212.20.1.1")
  enabled   = true
  ipsec_vpn = true
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `address` (String) The static IP address of the endpoint, required for static endpoints.
- `comment` (String) An optional comment for the element. This field is not required.
- `connection_type_ref` (String) This represents a connection type used in endpoints, which defines the connectivity group, mode, and link type for VPN connections.
- `deducted_name` (String) The deducted name of the internal endpoint, which is required for identification.
- `dynamic` (Boolean) Indicates whether the endpoint is dynamic (true) or static (false).
- `enabled` (Boolean) Indicates whether the endpoint is enabled (true) or disabled (false).
- `force_nat_t` (Boolean) Indicates whether the endpoint forces NAT Traversal (NAT_T) for VPN connections.
- `ike_phase1_id_type` (String) The type of Phase-1 ID used for the endpoint, represented as a single digit. '0' for DNS_NAME, '1' for EMAIL, '2' for DISTINGUISHED_NAME, and '3' for IP_ADDRESS.
- `ike_phase1_id_value` (String) The value of the Phase-1 ID for the endpoint, which must match the specified Phase-1 ID type.
- `ipsec_vpn` (Boolean) 
- `name` (String) Name of the object.
- `nat_t` (Boolean) Indicates whether the endpoint supports NAT Traversal (NAT_T) for VPN connections.
- `physical_interface` (String) This represents the abstract physical interface used in the engine cluster, which includes various settings and configurations for network interfaces.
- `ssl_vpn_portal` (Boolean) 
- `ssl_vpn_tunnel` (Boolean) 
- `udp_encapsulation` (Boolean) Indicates whether the endpoint supports UDP encapsulation for VPN tunnels.


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
