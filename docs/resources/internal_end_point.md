---
page_title: "smc_internal_endpoint"
subcategory: "vpn"
description: |-
  This represents an internal endpoint for VPN connections, supporting both tunnel and portal types. It includes properties such as physical interface and deducted name.
---

# smc_internal_endpoint (Sub-resource)

This represents an internal endpoint for VPN connections, supporting both tunnel and portal types. It includes properties such as physical interface and deducted name.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.752.0/examples/sdwan/policy_based_vpn/central_gateway_node/main.tf) for a minimal example

This example creates an internal endpoint (used for VPN or routing) tied to a single firewall.

```hcl
resource "smc_internal_endpoint" "ep_212_20_1_1" {
  id = format("%s#%s", smc_single_fw.tf_single_fw1.link.internal_gateway,
  "*/internal_endpoint/212.20.1.1")
  enabled   = true
  ipsec_vpn = true
}
```

### Disabling all VPN endpoints of an engine

SMC creates a VPN internal endpoint for every VPN-eligible interface and enables
it by default. To disable them all while keeping the Terraform state consistent,
declare one `smc_internal_endpoint` per interface with `enabled = false`. Using
`for_each` avoids repeating the resource block:

```hcl
locals {
  # Static interfaces: key = IP address (SMC names the endpoint after the IP).
  static_interfaces = {
    "192.168.1.1" = { interface_id = 1, nicid = 1, network = "192.168.1.0/24" }
  }
  # Dynamic (DHCP) interfaces: key = physical_interface name (SMC names the
  # endpoint after the interface). The name set on the physical_interface block
  # must match this key exactly.
  dynamic_interfaces = {
    "First DHCP Interface"  = { interface_id = 2, nicid = 2, dynamic_index = 1 }
    "Second DHCP Interface" = { interface_id = 3, nicid = 3, dynamic_index = 2 }
  }
}

resource "smc_internal_endpoint" "disabled" {
  for_each = toset(concat(keys(local.static_interfaces), keys(local.dynamic_interfaces)))

  id        = format("%s#*/internal_endpoint/%s", smc_single_fw.tf_single_fw.link.internal_gateway, each.key)
  enabled   = false
  ipsec_vpn = true
}
```

This is the idiomatic Terraform approach: the desired state (`enabled = false`) is
declared once and Terraform enforces it on every subsequent apply, so no manual
re-disabling is needed if another attribute changes later.

For dynamic interfaces, the endpoint name in SMC is derived from the
`physical_interface.name` attribute. If the name format differs on your SMC
version (e.g. `"Interface 2 (DHCP)"`), check with a direct API call:
`GET <smc>/internal_gateway/<id>/internal_endpoint` and inspect the `name` field.

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.752.0/examples/sdwan/internal_endpoint_all_disabled/main.tf) for a full example.


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
- `ipsec_vpn` (Boolean) Indicates whether the endpoint supports IPSec VPN type.
- `name` (String) Name of the object.
- `nat_t` (Boolean) Indicates whether the endpoint supports NAT Traversal (NAT_T) for VPN connections.
- `physical_interface` (String) This represents the abstract physical interface used in the engine cluster, which includes various settings and configurations for network interfaces.
- `ssl_vpn_portal` (Boolean) Indicates whether the endpoint supports VPN portal type using Reverse Proxy.
- `ssl_vpn_tunnel` (Boolean) Indicates whether the endpoint supports VPN tunnel type using Tunneling.
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
