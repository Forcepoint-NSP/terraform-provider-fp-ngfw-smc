---
page_title: "event_group_situation_parameter_value"
subcategory: "situations"
description: |-
  This represents an event group parameter value within a situation, allowing for the application of specific event groups to the situation's parameters.
---

# event_group_situation_parameter_value (Nested-Attribute)

This represents an event group parameter value within a situation, allowing for the application of specific event groups to the situation's parameters.




## Simple Attributes
- `binding_set_ref` (String) This represents a binding set, which is used to manage bindings in the system.
- `comment` (String) An optional comment for the element. This field is not required.
- `counts` (List of Number) A list of counts associated with this parameter value.
- `name` (String) Name of the object.
- `order` (Number) The order of this parameter value within the situation, determining its sequence.
- `parameter_ref` (String) This represents a parameter within a situation, allowing for the configuration of specific parameters that can be used in the context of the situation.
- `primary_filters` (List of String) URI of the associated primary filter.
- `secondary_context_members` (List of Number) A secondary context member associated with this parameter value.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
