---
page_title: "firewall_node"
subcategory: ""
description: |-
  This represents an individual Engine Node in the Security Management Client, which is part of an Engine Cluster.
---

# firewall_node (Nested-Attribute)

This represents an individual Engine Node in the Security Management Client, which is part of an Engine Cluster.




## Simple Attributes
- `activate_test` (Boolean) Indicates whether test entries are activated on this engine node or not.
- `comment` (String) An optional comment for the element. This field is not required.
- `engine_version` (String) The version of the software running on the engine node.
- `name` (String) Name of the object.
- `nodeid` (Number) The unique identifier of the engine node in the cluster, which is a value between 1 and 16.
- `snmp_engine_id` (String) The SNMP Engine ID for the engine node, which is used for SNMP communication. If not specified, it will be automatically generated.
- `snmp_location` (String) The SNMP location for the engine node, which provides information about the physical location of the node.

## Nested Attributes
- `appliance_info` (Single Block, see [here](attr_appliance_info.md)) 
- `external_pki_certificate_settings` (Single Block, see [here](attr_certificate_settings.md)) 
- `loopback_node_dedicated_interface` (List of Blocks, see [here](attr_loopback_node_interface.md)) The loopback Node Dedicated IP addresses that allow the firewall to communicate with itself, not connected to any physical interface.
- `tests` (List of Blocks, see [here](attr_test_entry_wrapper.md)) The list of test entries associated with this engine node, which are used to monitor the health and performance of the node.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
