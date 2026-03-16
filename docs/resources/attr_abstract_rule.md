---
page_title: "abstract_rule"
subcategory: ""
description: |-
  This represents a rule in a policy, including its parent policy, parent insert point, rank, tag, background color, and insert point type.
---

# abstract_rule (Nested-Attribute)

This represents a rule in a policy, including its parent policy, parent insert point, rank, tag, background color, and insert point type.




## Simple Attributes
- `background_color` (String) The background color for the comment (rule section) and insert point rules. It is represented by its hexadecimal representation ('#RRGGBB').
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `rank` (Number) The rank of the Rule within the Policy. If not specified, the Rule will be the first one in the Policy.
- `type` (String) The type of insert point for the rule. It can be 'normal' for a standard insert point or 'automatic' for an automatic rule insert point.


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
