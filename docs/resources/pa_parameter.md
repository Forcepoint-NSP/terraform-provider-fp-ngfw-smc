---
page_title: "smc_pa_parameter"
subcategory: ""
description: |-
  This represents a parameter for the Protocol Agent, allowing for detailed configuration of agent settings.
---

# smc_pa_parameter (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a parameter for the Protocol Agent, allowing for detailed configuration of agent settings.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `comment` (String) An optional comment for the element. This field is not required.
- `default_integer` (Number) The default integer value for this parameter, used if no value is provided by the user.
- `default_string` (String) The default string value for this parameter, used if no value is provided by the user.
- `description` (String) A brief description of the parameter, providing context and usage information.
- `explanation` (String) A detailed explanation of the parameter, including its purpose and how it should be used.
- `is_visible` (Boolean) Indicates whether the parameter is visible in the user interface.
- `max_integer` (Number) The maximum integer value allowed for this parameter, if applicable.
- `max_version` (String) The maximum version of the protocol or system that this parameter applies to.
- `min_integer` (Number) The minimum integer value allowed for this parameter, if applicable.
- `min_pa_version` (Number) The minimum Protocol Agent version required for this parameter to be applicable.
- `min_version` (String) The minimum version of the protocol or system that this parameter applies to.
- `name` (String) Name of the object.
- `rank` (Number) The rank of the parameter, used for ordering parameters in the user interface.
- `type` (String) The type of the parameter, which determines how it is processed and displayed.

## Nested Attributes
- `parent_group` (Single Block, see [here](attr_pa_parameter_group.md)) 

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
