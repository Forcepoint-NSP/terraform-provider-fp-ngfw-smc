---
page_title: "smc_situation_parameter"
subcategory: "situations"
description: |-
  This represents a parameter within a situation, allowing for the configuration of specific parameters that can be used in the context of the situation.
---

# smc_situation_parameter (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a parameter within a situation, allowing for the configuration of specific parameters that can be used in the context of the situation.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `bounds` (String) Hints about the minimum and maximum size for integers, floats, etc., and string lengths for string parameters. For example, '0|100' means the minimum is 0 and the maximum is 100, '|100' means the minimum is not set, and '0|' means that the maximum is not set.
- `comment` (String) An optional comment for the element. This field is not required.
- `conn_info_type` (String) The type of connection information associated with this parameter, such as 'none', 'logserver', 'alertserver', 'block_list', or 'sendanalyzator'.
- `default_value` (String) The default value for this parameter, which is used if no value is provided.
- `display_name` (String) The display name of the parameter in the GUI, providing a user-friendly label.
- `enumeration` (String) Hints for the GUI to show the enumeration for this parameter's value setting. For example, 'Pass:1-Drop:2-Reject:3' would mean a pull-down menu for three integer values showing Pass, Drop, and Reject to the user.
- `max_version` (String) The maximum supported engine version for this parameter.
- `min_version` (String) The minimum supported engine version for this parameter.
- `name` (String) Name of the object.
- `optional` (Boolean) Indicates whether this parameter is optional. If true, the parameter can be omitted.
- `order` (Number) The order in which this parameter appears in the situation, allowing for custom ordering of parameters.
- `type` (String) The type of the parameter, such as 'string', 'integer', 'boolean', etc.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
