---
page_title: "dhcp_range_per_node"
subcategory: ""
description: |-
  This represents the address range to be used when configuring DHCP on an interface, including the IPv4 address range and the associated DHCP Node ID.
---

# dhcp_range_per_node (Nested-Attribute)

This represents the address range to be used when configuring DHCP on an interface, including the IPv4 address range and the associated DHCP Node ID.




## Simple Attributes
- `dhcp_address_range` (String) The IPv4 address range for the DHCP server, specified in the format X.X.X.X-Y.Y.Y.Y or as a single IPv4 address X.X.X.X.
- `node_id` (Number) The ID of the DHCP node. This ID is used to identify the node in the cluster.


