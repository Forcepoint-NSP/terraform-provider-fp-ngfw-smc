---
page_title: "smc_user_alert_check"
subcategory: ""
description: |-
  This represents a User Alert Check. It contains settings for the type of check, threshold values, filter, and associated alert.
---

# smc_user_alert_check (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a User Alert Check. It contains settings for the type of check, threshold values, filter, and associated alert.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `alert_ref` (String) This represents an abstract Alert, which is used to display messages when certain conditions are met.
- `check_type` (String) The type of check. It can be one of the following: bandwidth_check, web_content_check, access_check, file_transfers_check, attack_check, endpoint_check.
- `comment` (String) An optional comment for the element. This field is not required.
- `filter` (String) This represents a Query Data Filter, which is a local filter that contains filter expression nodes and is specific to the element or view in which it was created.
- `name` (String) Name of the object.
- `severity` (Number) The severity of the alert. It is a value between 1 (low) and 10 (critical). By default, it is set to 10 for high severity.
- `threshold_duration_unit` (String) The unit of time for the threshold duration. It can be one of the following: minutes, hours.
- `threshold_duration_value` (Number) The duration value for the threshold. It is used to specify how long the threshold should be monitored before triggering an alert.
- `threshold_type` (String) The type of threshold. It can be one of the following: single_occurrence, event_count, bandwidth_count.
- `threshold_unit` (String) The unit of measurement for the threshold value. It can be one of the following: events, kb, mb, gb, tb.
- `threshold_value` (Number) The threshold value. It can be the number of events or the amount of traffic in KB/MB/GB/TB, depending on the threshold type. The events quantity is only for bandwidth_check and web_content_check checks.


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
