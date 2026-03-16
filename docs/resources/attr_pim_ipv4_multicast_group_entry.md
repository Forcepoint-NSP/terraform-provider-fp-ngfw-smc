---
page_title: "pim_multicast_group_entry"
subcategory: ""
description: |-
  This represents the PIM IPv4 Multicast Group Entry, which includes multicast IP network, multicast group, mode, and mapping.
---

# pim_multicast_group_entry (Nested-Attribute)

This represents the PIM IPv4 Multicast Group Entry, which includes multicast IP network, multicast group, mode, and mapping.




## Simple Attributes
- `mapping` (String) The Multicast group RP or Mapping, which can be an IPv4 Unicast address or a list of addresses. For 'pim_sm' mode, it represents the RP IPv4 Unicast address. For 'pim_ssm' mode, it can be a list of IPv4 Unicast addresses separated by commas or a DNS suffix.
- `mode` (String) The mode of the Multicast group, which can be pim_sm, pim_ssm, or pim_dm.
- `multicast_group_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `multicast_ip_network` (String) The manual Multicast IPv4 Network, if not using a Network element.


