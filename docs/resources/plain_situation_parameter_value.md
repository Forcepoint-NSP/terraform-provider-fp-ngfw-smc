---
page_title: "smc_plain_situation_situation_parameter_value"
subcategory: "situations"
description: |-
  This represents a plain situation parameter value within a situation, allowing for the application of specific plain situations to the situation's parameters.
---

# smc_plain_situation_situation_parameter_value (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a plain situation parameter value within a situation, allowing for the application of specific plain situations to the situation's parameters.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `order` (Number) The order of this parameter value within the situation, determining its sequence.
- `parameter_ref` (String) This represents a parameter within a situation, allowing for the configuration of specific parameters that can be used in the context of the situation.
- `value` (String) This represents a situation that can occur in the system, including its severity, attacker, victim, and other related attributes.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
