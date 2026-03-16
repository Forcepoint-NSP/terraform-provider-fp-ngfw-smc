---
page_title: "smc_eca_endpoint_settings"
subcategory: ""
description: |-
  This represents the ECA Endpoint Settings, which includes configurations for operating systems, client AV settings, OS update times, and local firewall settings.
---

# smc_eca_endpoint_settings (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents the ECA Endpoint Settings, which includes configurations for operating systems, client AV settings, OS update times, and local firewall settings.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `client_av_disabled` (Boolean) Indicates if client AV is disabled.
- `client_av_enabled` (Boolean) Indicates if client AV is enabled.
- `client_av_unknown` (Boolean) Indicates if the client AV status is unknown.
- `comment` (String) An optional comment for the element. This field is not required.
- `eca_os` (List of String) URI of the ECA Operating System Situation.
- `local_firewall_disabled` (Boolean) Indicates if the local firewall is disabled.
- `local_firewall_enabled` (Boolean) Indicates if the local firewall is enabled.
- `local_firewall_unknown` (Boolean) Indicates if the local firewall status is unknown.
- `name` (String) Name of the object.
- `os_update_time_days` (Number) The number of days for OS update time, which should be a positive integer.
- `os_update_time_enabled` (Boolean) Indicates if OS update time is enabled.
- `os_update_time_operator` (String) The operator for OS update time, such as 'less_than' or 'more_than'.
- `os_update_unknown` (Boolean) Indicates if the OS update status is unknown.


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
