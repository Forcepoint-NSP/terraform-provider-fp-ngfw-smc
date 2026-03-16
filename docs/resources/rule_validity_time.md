---
page_title: "smc_rule_validity_time"
subcategory: "policy"
description: |-
  This represents the validity time of a rule, which defines when the rule is active. It includes settings for time zones, start and end times, and repeat modes.
---

# smc_rule_validity_time (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents the validity time of a rule, which defines when the rule is active. It includes settings for time zones, start and end times, and repeat modes.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `month_day` (List of Number) 
- `name` (String) Name of the object.
- `rule_time_active` (String) The repeat mode for the rule validity time. Possible values are 'always' (default), 'time_of_day', 'day_of_week', 'day_of_month', and 'date_of_year'.
- `rule_time_end` (String) The end time of the rule validity period, in the format 'yyyy-MM-dd HH:mm'.
- `rule_time_repeat_end` (String) The end time of the day for the repeat mode, in the format 'HH:mm'.
- `rule_time_repeat_start` (String) The start time of the day for the repeat mode, in the format 'HH:mm'.
- `rule_time_start` (String) The start time of the rule validity period, in the format 'yyyy-MM-dd HH:mm'.
- `rule_time_zone` (String) The time zone mode for the rule validity time. Possible values are 'utc' (default) and 'engine_time'.
- `week_day` (List of String) 
- `yearly_day_end` (Number) The end day of the month for the yearly repeat mode. This should be an integer between 1 and 31, representing the day of the month.
- `yearly_day_start` (Number) The start day of the month for the yearly repeat mode. This should be an integer between 1 and 31, representing the day of the month.
- `yearly_month_end` (String) The end month for the yearly repeat mode. This should be a valid month abbreviation (e.g., 'jan', 'feb', etc.).
- `yearly_month_start` (String) The start month for the yearly repeat mode. This should be a valid month abbreviation (e.g., 'jan', 'feb', etc.).


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
