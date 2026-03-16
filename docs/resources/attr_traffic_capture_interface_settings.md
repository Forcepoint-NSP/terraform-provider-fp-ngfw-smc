---
page_title: "traffic_capture_interface_settings"
subcategory: ""
description: |-
  This represents the settings for capturing traffic on a specific interface, including the node, filter, and physical interface link.
---

# traffic_capture_interface_settings (Nested-Attribute)

This represents the settings for capturing traffic on a specific interface, including the node, filter, and physical interface link.




## Simple Attributes
- `filter` (String) The filter applied to the traffic capture, which can be an IP address or a specific filtering rule. By default, no filtering is applied.
- `node_ref` (String) This represents an individual Engine Node in the Security Management Client, which is part of an Engine Cluster.
- `pint_ref` (String) This represents the abstract physical interface used in the engine cluster, which includes various settings and configurations for network interfaces.


