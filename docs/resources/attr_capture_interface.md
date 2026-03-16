---
page_title: "capture_interface"
subcategory: "interfaces"
description: |-
  This represents an IPS Engine interface that can listen to traffic passing in the network, but which is not used for routing traffic through the engine.
---

# capture_interface (Nested-Attribute)

This represents an IPS Engine interface that can listen to traffic passing in the network, but which is not used for routing traffic through the engine.




## Simple Attributes
- `address` (String) The IP Address (IPv4 or IPv6) of the interface. For dynamic interfaces, this will be null.
- `comment` (String) An optional comment for the element. This field is not required.
- `inspect_qinq` (Boolean) Indicates whether QinQ is inspected or not. This option allows the IPS engine to inspect traffic that uses QinQ encapsulation.
- `inspect_unspecified_vlans` (Boolean) Indicates whether unspecified VLANs are inspected or not. Deselect this option to make the IPS engine ignore traffic from VLANs that are not included in the IPS engine's interface configuration. We recommend that you keep this option selected if you do not have a specific reason to deselect it.
- `logical_interface_ref` (String) This represents a Logical Interface. It is an IPS Element used in the IPS policies to represent one or more physical network interfaces as defined in the IPS engine properties.
- `name` (String) Name of the object.
- `network_value` (String) The IP Network (IPv4 or IPv6) of the interface. For dynamic interfaces, this will be null.
- `nicid` (String) The Interface ID of the interface.
- `reset_interface_nicid` (Number) The Reset Interface to specify the interface through which TCP connection resets are sent when Reset responses are used in your IPS policy.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
