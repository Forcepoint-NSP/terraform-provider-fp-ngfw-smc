---
page_title: "smc_logical_operator_situation_parameter_value"
subcategory: "situations"
description: |-
  This represents a logical operator parameter value within a situation, allowing for the application of logical operations to the situation's parameters.
---

# smc_logical_operator_situation_parameter_value (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a logical operator parameter value within a situation, allowing for the application of logical operations to the situation's parameters.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `operator_type` (String) The type of logical operator used in this parameter value, such as 'or', 'not', 'and', or 'ordered_and'.
- `order` (Number) The order of this parameter value within the situation, determining its sequence.
- `parameter_ref` (String) This represents a parameter within a situation, allowing for the configuration of specific parameters that can be used in the context of the situation.
- `references` (List of String) URI of the storable.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
