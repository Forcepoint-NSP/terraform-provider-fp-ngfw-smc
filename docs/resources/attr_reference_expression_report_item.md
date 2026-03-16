---
page_title: "report_item_expr"
subcategory: ""
description: |-
  This represents a reference expression for a report item, used in filter expressions to match log entries based on specific report file references and item specifiers.
---

# report_item_expr (Nested-Attribute)

This represents a reference expression for a report item, used in filter expressions to match log entries based on specific report file references and item specifiers.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `item_specifier` (String) The item specifier used to identify the specific item within the report, which is used for filtering log entries based on item characteristics.
- `name` (String) Name of the object.
- `report_file_ref` (String) This represents a Report File, which contains the generated report data. It includes metadata such as creation time, period begin, and period end, along with sections of the report.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
