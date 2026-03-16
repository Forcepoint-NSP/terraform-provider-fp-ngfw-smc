---
page_title: "statistics_item_expr"
subcategory: ""
description: |-
  This represents a reference expression for a statistics item, used in filter expressions to match log entries based on specific item identifiers.
---

# statistics_item_expr (Nested-Attribute)

This represents a reference expression for a statistics item, used in filter expressions to match log entries based on specific item identifiers.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `item_id` (Number) The unique identifier of the statistics item, which is read-only and used for filtering log entries.
- `item_name` (String) The name of the statistics item, which is read-only and used for identification in filtering logic.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
