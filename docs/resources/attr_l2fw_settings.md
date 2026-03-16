---
page_title: "l2fw_settings"
subcategory: "engines"
description: |-
  This represents the Layer 2 settings for NGFW engines, including interface policies, traffic bypass settings, and TCP tracking modes.
---

# l2fw_settings (Nested-Attribute)

This represents the Layer 2 settings for NGFW engines, including interface policies, traffic bypass settings, and TCP tracking modes.




## Simple Attributes
- `bypass_overload_traffic` (Boolean) Indicates whether traffic is bypassed during overload conditions on Inline IPS and Capture interfaces.
- `l2_interface_policy_ref` (String) This represents a Layer 2 Interface Policy, which is used to define the action and inspection rules for L2FW Interface Engines.
- `tracking_mode` (String) The connection tracking mode for Layer 2 Interfaces, which can be 'normal', 'strict', or 'loose'. Default is 'normal'.


