---
page_title: "smc_ospfv2_interface_settings"
subcategory: "routing"
description: |-
  This represents the OSPFv2 Interface Settings, which are used to configure OSPFv2 interfaces in the Dynamic Routing Firewall functionality.
---

# smc_ospfv2_interface_settings (Resource)

This represents the OSPFv2 Interface Settings, which are used to configure OSPFv2 interfaces in the Dynamic Routing Firewall functionality.

## Examples

- [OSPFv2 Interface Settings Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/OSPFv2/ospfv2_interface_settings/main.tf)

Defines per-interface OSPFv2 protocol settings.

```hcl
resource "smc_ospfv2_interface_settings" "ospfv2_interface_settings" {
  authentication_type    = "none"
  dead_interval          = 1
  dead_multiplier        = 4
  hello_interval_type    = "fast_hello"
  interface_cost         = 14
  mtu_mismatch_detection = true
  name                   = "tf_ospfv2_interface_settings"
  password               = ""
  retransmit_interval    = 10
  router_priority        = 5
  transmit_delay         = 2
  comment                = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `authentication_type` (String) The type of OSPF authentication used on the interface.
- `comment` (String) An optional comment for the element. This field is not required.
- `dead_interval` (Number) The Dead interval in seconds for OSPF, which is the time after which a neighbor is considered down if no Hello packets are received.
- `dead_multiplier` (Number) The multiplier for the Hello interval used to calculate the Dead interval when using 'minimal' Hello interval type.
- `hello_interval` (Number) The Hello interval in seconds for OSPF, which is the time between Hello packets sent on the interface.
- `hello_interval_type` (String) The type of Hello interval used in OSPF, which can be 'normal', 'minimal', or 'fast_hello'.
- `interface_cost` (Number) The cost of the OSPF interface, which affects the routing decisions made by OSPF.
- `key_chain_ref` (String) This represents an abstract OSPF Key Chain used as a Dynamic Routing element. It contains settings related to OSPF key chains.
- `mtu_mismatch_detection` (Boolean) Indicates whether MTU mismatch detection is enabled on the OSPF interface.
- `name` (String) Name of the object.
- `password` (String) The password used for OSPF authentication, required if the authentication type is 'password'.
- `retransmit_interval` (Number) The interval in seconds between retransmissions of OSPF packets.
- `router_priority` (Number) The router priority for OSPF, which is used to determine the designated router on a multi-access network.
- `transmit_delay` (Number) The delay in seconds before transmitting OSPF packets on the interface.


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
