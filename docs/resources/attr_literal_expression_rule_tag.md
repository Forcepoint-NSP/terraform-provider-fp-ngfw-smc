---
page_title: "rule_tag_expr"
subcategory: ""
description: |-
  This represents a literal expression for a rule tag, allowing for filtering based on specific rule tags in log entries.
---

# rule_tag_expr (Nested-Attribute)

This represents a literal expression for a rule tag, allowing for filtering based on specific rule tags in log entries.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `fixed_rule_tag` (Number) The fixed part of the rule tag, which is a numeric identifier used to match specific rule tags in log entries.
- `name` (String) Name of the object.
- `rule_tag_name` (String) The name of the rule tag, which is read-only and used for identification in filtering logic.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
