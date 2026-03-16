---
page_title: "smc_route_map_rule"
subcategory: "routing"
description: |-
  This represents a Route Map Rule for Route Map Policy. It can be used to define routing policies based on various conditions.
---

# smc_route_map_rule (Sub-resource)

This represents a Route Map Rule for Route Map Policy. It can be used to define routing policies based on various conditions.

## Examples

- [Route Map Rule Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/route_map/main.tf)

Defines a rule used within a routing map for policy-based routing adjustments.

```hcl
resource "smc_route_map_rule" "rule_1" {
  from_ref    = smc_route_map.route_map.link.route_map_rules
  action      = "permit"
  finish      = false
  is_disabled = false
  match_condition {
    metric = 10
    rank   = 1
    type   = "metric"
  }
  name = "route_map_rule_1"
  rank = 1.0
  route_entry_settings {
    as_path_type                  = "dont_modify"
    community_type                = "dont_modify"
    extended_community_entry_type = "rt"
    extended_community_type       = "dont_modify"
    ipv4_next_hop {
      next_hop_peer_address = false
      next_hop_ref          = smc_router.tf_example_router.id
    }
  }
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `action` (String) The action to be taken by the rule.
- `background_color` (String) The background color for the comment (rule section) and insert point rules. It is represented by its hexadecimal representation ('#RRGGBB').
- `call_route_map_ref` (String) This represents a Route Map Policy for the Dynamic Routing Firewall settings, which is used to control the routing behavior based on specific rules.
- `comment` (String) An optional comment for the element. This field is not required.
- `finish` (Boolean) Indicates whether this is the last rule in the Route Map.
- `goto` (String) This represents a Route Map Rule for Route Map Policy. It can be used to define routing policies based on various conditions.
- `is_disabled` (Boolean) Indicates whether the rule is disabled.
- `name` (String) Name of the object.
- `rank` (Number) The rank of the Rule within the Policy. If not specified, the Rule will be the first one in the Policy.
- `type` (String) The type of insert point for the rule. It can be 'normal' for a standard insert point or 'automatic' for an automatic rule insert point.

## Nested Attributes
- `match_condition` (List of Blocks, see [here](attr_route_map_matching_condition.md)) The match conditions for the Route Map rule. These conditions determine how the rule matches packets.
- `route_entry_settings` (Single Block, see [here](attr_route_map_rule_option.md)) 

## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `parent_insert_point` (String) This represents a rule in a policy, including its parent policy, parent insert point, rank, tag, background color, and insert point type.
- `parent_policy` (String) This represents a policy that can be applied to various elements in the system, such as network elements, inspection rules, etc.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `tag` (String) The unique identifier of the rule in this policy. It contains a static part that does not change when rules are added, removed, or moved, and a changing part that indicates the version of the rule.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
