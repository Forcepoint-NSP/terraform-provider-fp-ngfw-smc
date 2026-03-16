---
page_title: "dhcp_relay"
subcategory: ""
description: |-
  This represents the definition of DHCP Relay on an interface, allowing you to select which firewall interfaces perform DHCP relay and activate it towards the DHCP clients.
---

# dhcp_relay (Nested-Attribute)

This represents the definition of DHCP Relay on an interface, allowing you to select which firewall interfaces perform DHCP relay and activate it towards the DHCP clients.




## Simple Attributes
- `element` (List of String) URI of the DHCP Server.
- `enabled` (Boolean) Indicates whether the DHCP Relay is enabled on this interface.
- `max_packet_size` (Number) The maximum packet size for DHCP Relay, which is required if DHCP Relay is enabled.
- `trusted_circuit` (Boolean) Indicates whether the DHCP Relay is configured for a trusted circuit.


