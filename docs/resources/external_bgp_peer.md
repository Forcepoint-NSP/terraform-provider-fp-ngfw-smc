---
page_title: "smc_external_bgp_peer"
subcategory: "routing"
description: |-
  This represents the External BGP Peer for Dynamic Routing Firewall functionality. It is used to configure external BGP peers in the firewall's dynamic routing capabilities.
---

# smc_external_bgp_peer (Resource)

This represents the External BGP Peer for Dynamic Routing Firewall functionality. It is used to configure external BGP peers in the firewall's dynamic routing capabilities.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/BGP/external_bgp_peers) for a complete minimal example

This example creates an external BGP peer object which references an Autonomous System.

```hcl
resource "smc_external_bgp_peer" "external_bgp_peer" {
  name = "tf_external_bgp_peer"
  neighbor_as = smc_autonomous_system.autonomous_system.id
  neighbor_ip = "10.10.14.10"
  neighbor_port = 179
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `neighbor_as` (String) This represents the Autonomous System for Dynamic Routing Firewall functionality. The BGP Version is 4/4+. It includes the AS Number in decimal notation, which is used to identify the autonomous system in BGP routing.
- `neighbor_ip` (String) The IP address of the external BGP peer.
- `neighbor_port` (Number) The port number for the external BGP peer. Default is 179.


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
