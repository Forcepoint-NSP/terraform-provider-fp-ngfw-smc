---
page_title: "lbfilter"
subcategory: ""
description: |-
  This represents the Load Balancing Filter, which includes IP descriptor, NAT enforcement, port usage in hash, IPSec usage, action type, node ID, and replace address.
---

# lbfilter (Nested-Attribute)

This represents the Load Balancing Filter, which includes IP descriptor, NAT enforcement, port usage in hash, IPSec usage, action type, node ID, and replace address.




## Simple Attributes
- `action` (String) The action type for the filter, which can be 'replace', 'node', or 'drop'.
- `ignore_other` (Boolean) Indicates whether other entries might not be concerned by this filter, allowing for more specific handling.
- `ip_descriptor` (String) The IP descriptor, which can be an IP network or an IP address range.
- `nat_enforce` (Boolean) Indicates whether NAT should enforce translated packet headers to the same hash value as the matching packet.
- `nodeid` (Number) The ID of the node to which this filter applies, used when the action type is 'node'.
- `replace_ip` (String) The address to replace with when the action type is 'replace'. This is the new IP address that will be used.
- `use_ipsec` (Boolean) Indicates whether this entry should be handled with special care because it is part of a VPN.
- `use_ports` (Boolean) Indicates whether to use port numbers when calculating the hash value for the packet.


