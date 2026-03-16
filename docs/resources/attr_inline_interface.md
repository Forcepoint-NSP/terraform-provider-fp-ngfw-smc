---
page_title: "inline_interface"
subcategory: "interfaces"
description: |-
  This represents the interface that combines together two physical interfaces, enabling the traffic to be routed through as if the engine were an extension of the network cable, but allowing the engine to actively monitor packets and connections and stop them according to its Actions and Inspection Rules.
---

# inline_interface (Nested-Attribute)

This represents the interface that combines together two physical interfaces, enabling the traffic to be routed through as if the engine were an extension of the network cable, but allowing the engine to actively monitor packets and connections and stop them according to its Actions and Inspection Rules.




## Simple Attributes
- `address` (String) The IP Address (IPv4 or IPv6) of the interface. For dynamic interfaces, this will be null.
- `comment` (String) An optional comment for the element. This field is not required.
- `failure_mode` (String) The mode of operation for the Inline Interface when the IPS engine goes offline. Options are 'bypass' (traffic allowed without inspection) or 'normal' (traffic not allowed). Caution! Using the Bypass mode requires the IPS appliance to have a fail-open network interface card.
- `inspect_qinq` (Boolean) Indicates whether QinQ is inspected or not. This option allows the IPS engine to inspect traffic that uses QinQ encapsulation.
- `inspect_unspecified_vlans` (Boolean) Indicates whether unspecified VLANs are inspected or not. Deselect this option to make the IPS engine ignore traffic from VLANs that are not included in the IPS engine's interface configuration. We recommend that you keep this option selected if you do not have a specific reason to deselect it.
- `logical_interface_ref` (String) This represents a Logical Interface. It is an IPS Element used in the IPS policies to represent one or more physical network interfaces as defined in the IPS engine properties.
- `name` (String) Name of the object.
- `network_value` (String) The IP Network (IPv4 or IPv6) of the interface. For dynamic interfaces, this will be null.
- `nicid` (String) The Interface ID of the interface.
- `virtual_second_mapping` (String) The Second Virtual Mapping.
- `zone_ref` (String) This represents a Zone, which is used to group together network interfaces of Firewall, IPS, and Layer 2 Firewall engines. Zones can be used to specify receiving or sending interfaces in policies and automatically apply to new interfaces associated with the same Zone.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
