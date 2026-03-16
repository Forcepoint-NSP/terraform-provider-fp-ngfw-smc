---
page_title: "smc_delete_counter_data_task"
subcategory: "admin"
description: |-
  This represents a Delete Counter Data Task, which is used to delete counter data in the system. It is a type of log task that can be scheduled and executed to ensure that counter data is deleted properly.
---

# smc_delete_counter_data_task (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Delete Counter Data Task, which is used to delete counter data in the system. It is a type of log task that can be scheduled and executed to ensure that counter data is deleted properly.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `end_time` (Number) The end time in ms used when the time_limit_type is set to 'absolute_time_range'.
- `for_application_health_counter_type` (Boolean) Flag to know if application health-type counter data need to be deleted by the task.
- `for_default_counter_type` (Boolean) Flag to know if default-type counter data need to be deleted by the task.
- `for_elasticsearch_counter_type` (Boolean) Flag to know if Elasticsearch-type counter data need to be deleted by the task.
- `for_netlink_status_counter_type` (Boolean) Flag to know if NetLink status-type counter data need to be deleted by the task.
- `for_status_counter_type` (Boolean) Flag to know if status-type counter data need to be deleted by the task.
- `name` (String) Name of the object.
- `relative_time_begin` (Number) The begin number of days when the time_limit_type is set to 'relative_time_range' or 'before'.
- `relative_time_end` (Number) The end number of days when the time_limit_type is set to 'relative_time_range' or 'before'.
- `resources` (List of String) URI of the resource.
- `start_time` (Number) The start time in ms used when the time_limit_type is set to 'absolute_time_range' or 'after'.
- `time_limit_type` (String) Log task time limit type: Time range type. 'before' for 'relative_time_end' number of days, 'relative_time_range' for between 'relative_time_begin' number of days and 'relative_time_end' number of days, 'absolute_time_range' for between 'start_time' time in ms and 'end_time' time in ms. 'today' for today. 'yesterday' for yesterday. 'last_full_week_sun_sat' for last full week (sunday-saturday). 'last_full_week_mon_sun' for last full week (monday-sunday). 'last_full_month' for last full month. 'after' for after 'start_time' time in ms.


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
