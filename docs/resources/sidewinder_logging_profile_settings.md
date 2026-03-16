---
page_title: "smc_sidewinder_logging_profile_setting"
subcategory: ""
description: |-
  This represents the Sidewinder Logging Profile Settings, which includes configuration for logging elements, activation status, threshold, and interval settings.
---

# smc_sidewinder_logging_profile_setting (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents the Sidewinder Logging Profile Settings, which includes configuration for logging elements, activation status, threshold, and interval settings.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `element` (String) This is the base class for all storable elements.
- `enable` (String) Indicates whether the logging profile settings are: '0' not enabled; '1' always enabled; '2' limited.
- `interval` (Number) The interval value in seconds for the logging profile settings, which is applicable when the activation is limited.
- `threshold` (Number) The threshold value for the logging profile settings, which is applicable when the activation is limited.


