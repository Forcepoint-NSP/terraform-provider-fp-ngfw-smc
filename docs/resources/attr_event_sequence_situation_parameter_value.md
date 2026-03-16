---
page_title: "event_sequence_situation_parameter_value"
subcategory: "situations"
description: |-
  This represents an event sequence parameter value within a situation, allowing for the application of specific event sequences to the situation's parameters.
---

# event_sequence_situation_parameter_value (Nested-Attribute)

This represents an event sequence parameter value within a situation, allowing for the application of specific event sequences to the situation's parameters.




## Simple Attributes
- `binding_sets` (List of String) URI of the associated binding set.
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `order` (Number) The order of this parameter value within the situation, determining its sequence.
- `parameter_ref` (String) This represents a parameter within a situation, allowing for the configuration of specific parameters that can be used in the context of the situation.
- `primary_filters` (List of String) URI of the associated primary filter.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
