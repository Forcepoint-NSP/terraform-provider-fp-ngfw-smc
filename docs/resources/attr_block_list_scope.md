---
page_title: "block_list_scope"
subcategory: ""
description: |-
  This represents a block list scope, which is a sub part of Rule Action Options used to configure the block list settings of the Rule. It allows you to block traffic manually on Firewalls, IPS engines, and Layer 2 Firewalls.
---

# block_list_scope (Nested-Attribute)

This represents a block list scope, which is a sub part of Rule Action Options used to configure the block list settings of the Rule. It allows you to block traffic manually on Firewalls, IPS engines, and Layer 2 Firewalls.




## Simple Attributes
- `block_traffic` (Boolean) Indicates whether to block traffic between endpoints. If true, the block list entry will block traffic between the specified endpoints.
- `drop_connection` (Boolean) Indicates whether to terminate the single connection. If true, the block list entry will only cut the current connection.
- `duration` (Number) The duration in seconds for which the block list entry will be kept. If set to 0, the entry only cuts the current connections.
- `include_observer` (Boolean) Indicates whether to include the original observer in the block list entry. If true, the engine that detects the situation will be included in the block list executors.
- `service_ip_ref` (String) This represents an IP-proto service, which is used to define a service based on the IP protocol number. It includes a protocol number that specifies the protocol used by the traffic.

## Nested Attributes
- `end_point1` (Single Block, see [here](attr_block_list_end_point.md)) 
- `end_point2` (Single Block, see [here](attr_block_list_end_point.md)) 

