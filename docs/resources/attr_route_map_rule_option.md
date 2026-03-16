---
page_title: "route_entry_settings"
subcategory: "policy"
description: |-
  This represents the Route Map Rule Settings, which includes various options such as Local Preference, Weight, Metric, AS Path, Community, Extended Community, and Next Hop configurations.
---

# route_entry_settings (Nested-Attribute)

This represents the Route Map Rule Settings, which includes various options such as Local Preference, Weight, Metric, AS Path, Community, Extended Community, and Next Hop configurations.




## Simple Attributes
- `as_path_type` (String) The AS Path type, which determines how the AS Path is modified. It can be 'dont_modify' for nothing to modify, 'prepend' for prepending the AS number by this specified, or 'exclude' for excluding the AS number by this specified.
- `community_type` (String) The Community type, which determines how the Community is modified. It can be 'dont_modify' for nothing to modify, 'set_to_none' for setting to none the Community number, 'set_to' for setting the new Community number by this specified, 'additive' for adding the community number by this specified, or 'delete' for deleting the Community Access List element.
- `delete_community_ref` (String) This represents a Community Access List, which is used to define a list of communities for dynamic routing configurations.
- `extended_community_entry_type` (String) The Extended Community entry type, which specifies the type of Extended Community number. It can be 'soo' for Site of Origin or 'rt' for Route Target.
- `extended_community_type` (String) The Extended Community type, which determines how the Extended Community is modified. It can be 'dont_modify' for nothing to modify or 'set_to' for setting the new Extended Community number by this specified.
- `local_preference` (Number) The Local Preference value, which is an integer used to influence route selection.
- `metric` (Number) The Metric value, which is an integer used to influence route selection.
- `weight` (Number) The Weight value, which is an integer used to influence route selection.

## Nested Attributes
- `as_number` (List of Blocks, see [here](attr_as_path_expression.md)) The AS Path number value, which is used when 'prepend' or 'exclude' AS Path type is selected. This is a list of AS Path expressions.
- `community_number` (List of Blocks, see [here](attr_community_expression.md)) The Community number value, which is used when 'set_to' or 'additive' Community type is selected. This is a list of Community expressions.
- `extended_community_number` (List of Blocks, see [here](attr_extended_community_expression.md)) The Extended Community number values, which are used when 'set_to' Extended Community type is selected. This is a list of Extended Community expressions.
- `ipv4_next_hop` (Single Block, see [here](attr_next_hop.md)) 
- `ipv6_global_next_hop` (Single Block, see [here](attr_next_hop.md)) 
- `ipv6_local_next_hop` (Single Block, see [here](attr_next_hop.md)) 

