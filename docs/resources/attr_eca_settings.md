---
page_title: "eca_settings"
subcategory: ""
description: |-
  This represents the definition of ECA settings feature, which includes client and server networks, listened zones, listened interfaces, and listening port.
---

# eca_settings (Nested-Attribute)

This represents the definition of ECA settings feature, which includes client and server networks, listened zones, listened interfaces, and listening port.




## Simple Attributes
- `eca_client_config` (String) This represents a ECA Client Configuration, which is used to manage the client-side configuration for ECA (Endpoint Compliance Agent). It includes settings such as trusted certificate authorities and auto-discovery options.
- `eca_client_network_ref` (List of String) URI of the Network representing the client network for this ECA settings.
- `eca_server_network_ref` (List of String) URI of the Network representing the server network for this ECA settings.
- `listened_zone_ref` (List of String) URI of the Interface Zone representing the listened zone for this ECA settings.
- `listening_port` (Number) The port on which the ECA settings will listen for incoming connections.

## Nested Attributes
- `enabled_interface` (List of Blocks, see [here](attr_enabled_interface_entry.md)) The listened interfaces associated with these ECA settings.

