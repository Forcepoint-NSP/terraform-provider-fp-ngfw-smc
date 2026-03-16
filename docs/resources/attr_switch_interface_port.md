---
page_title: "switch_interface_port"
subcategory: "interfaces"
description: |-
  This represents a port on an engine that supports Switch Interfaces, including the port number, interface reference, and comment.
---

# switch_interface_port (Nested-Attribute)

This represents a port on an engine that supports Switch Interfaces, including the port number, interface reference, and comment.




## Simple Attributes
- `physical_switch_port_number` (String) The port number for a Physical-Switch, which is numbered from 0 to (N - 1), where N is the number of ports available on the Physical-Switch appliance.
- `soft_switch_pint_ref` (String) This represents the abstract physical interface used in the engine cluster, which includes various settings and configurations for network interfaces.
- `switch_interface_port_comment` (String) A comment describing the port, which can be used for additional information.


