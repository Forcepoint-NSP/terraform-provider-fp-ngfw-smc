---
page_title: "smc_task_schedule"
subcategory: "admin"
description: |-
  This represents a schedule entry for a task element, defining when and how often the task should be executed.
---

# smc_task_schedule (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a schedule entry for a task element, defining when and how often the task should be executed.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `activated` (Boolean) Indicates whether the schedule is enabled (true) or disabled (false).
- `activation_date` (Number) The date when the schedule is activated, represented as a timestamp in milliseconds since epoch.
- `comment` (String) An optional comment for the element. This field is not required.
- `day_mask` (Number) The day mask for weekly schedules, where each bit represents a day of the week (1 for Sunday, 2 for Monday, etc.). For example, to repeat every Monday and Wednesday, the value would be 42 (2 + 8 + 32).
- `day_period` (String) The period of the schedule, which can be 'one_shot' for no repeat, 'daily' for every day, 'weekly' for every week, 'monthly' for every month, or 'yearly' for every year.
- `expired` (Boolean) Indicates whether the schedule has expired (true) or is still active (false). An expired schedule will not trigger any further actions.
- `final_action` (String) The final action to be taken at the end of the schedule, which can be 'ALERT_FAILURE' to send an alert if the task fails, 'ALERT' to send an alert at the end of the schedule, or 'NO_ACTION' to do nothing.
- `minute_period` (String) The period in minutes for daily schedules, representing the number of minutes between each execution. The maximum value is 86400 minutes (24 hours).
- `name` (String) Name of the object.
- `repeat_until_date` (Number) The date until which the task can be repeated, represented as a timestamp in milliseconds since epoch.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
