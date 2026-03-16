---
page_title: "audit_object_expr"
subcategory: ""
description: |-
  This represents a reference expression for an audit object, used in filter expressions to match log entries based on specific audit object references.
---

# audit_object_expr (Nested-Attribute)

This represents a reference expression for an audit object, used in filter expressions to match log entries based on specific audit object references.




## Simple Attributes
- `audit_object_key` (Number) The key of the audit object reference, which is read-only and used for identification in filtering logic.
- `audit_object_name` (String) The name of the audit object reference, which is read-only and used for identification in filtering logic.
- `audit_object_ref` (String) This is a base class for objects that can be stored in the database. It includes properties such as key, comment, and links.
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
