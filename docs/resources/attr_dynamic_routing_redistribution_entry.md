---
page_title: "redistribution_entry"
subcategory: ""
description: |-
  This represents a Dynamic Routing Redistribution Entry, which defines how routes are redistributed between different routing protocols or from the kernel, static, or connected routes.
---

# redistribution_entry (Nested-Attribute)

This represents a Dynamic Routing Redistribution Entry, which defines how routes are redistributed between different routing protocols or from the kernel, static, or connected routes.




## Simple Attributes
- `enabled` (Boolean) Indicates whether the redistribution entry is enabled or not.
- `filter_type` (String) The type of filter applied to the redistribution, such as none, access_list, route_map_policy, always_on, or never.
- `metric` (Number) The seed metric value applied to the redistributed routes, which can be set null for Automatic or a specific value.
- `metric_type` (String) The type of metric used for the redistributed routes, such as external_1 or external_2.
- `redistribution_filter_ref` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `redistribution_rm_ref` (String) This represents a Route Map Policy for the Dynamic Routing Firewall settings, which is used to control the routing behavior based on specific rules.
- `type` (String) The type of redistribution, such as kernel, static, connected, bgp, ospfv2, or default_originate.


