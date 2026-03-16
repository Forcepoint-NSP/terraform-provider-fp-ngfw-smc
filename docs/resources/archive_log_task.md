---
page_title: "smc_archive_log_task"
subcategory: "admin"
description: |-
  This represents an Archive Log Task, which is used to archive logs in the system. It is a type of log task that can be scheduled and executed to ensure that logs are archived properly.
---

# smc_archive_log_task (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Archive Log Task, which is used to archive logs in the system. It is a type of log task that can be scheduled and executed to ensure that logs are archived properly.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `delete_source_data` (Boolean) Flag to know if after the archive operation, logs need to be deleted. If true, logs will be deleted after archiving.
- `end_time` (Number) The end time in ms used when the time_limit_type is set to 'absolute_time_range'.
- `file_format` (String) The file format for the export/archive mode: 'csv' for CSV format, 'xml' for XML format, 'zip' for ZIP format, 'cef' for CEF format, 'leef' for LEEF format, 'esm' for ESM format, 'snoop' for IPS recordings as SNOOP, 'pcap' for IPS recordings as PCAP.
- `file_name` (String) The path where the archive/export will be stored.
- `filter_for_delete` (String) This represents a Filter Expression Container, which is a top-level re-usable filter that contains filter expression nodes.
- `filter_for_export` (String) This represents a Filter Expression Container, which is a top-level re-usable filter that contains filter expression nodes.
- `for_alert_event_log` (Boolean) Flag to know if Alert event logs need to be considered for the task.
- `for_alert_log` (Boolean) Flag to know if Alert logs need to be considered for the task.
- `for_audit_log` (Boolean) Flag to know if Audit logs need to be considered for the task.
- `for_fw_log` (Boolean) Flag to know if firewall logs need to be considered for the task.
- `for_ips_log` (Boolean) Flag to know if IPS logs need to be considered for the task.
- `for_ips_recording_log` (Boolean) Flag to know if IPS Recording logs need to be considered for the task.
- `for_l2fw_log` (Boolean) Flag to know if L2FW logs need to be considered for the task.
- `for_third_party_log` (Boolean) Flag to know if Third Party logs need to be considered for the task.
- `ips_recordings` (List of Number) Optional IPS record IDs for PCAP/SNOOP IPS recording export logs.
- `is_local_location` (Boolean) Flag to know if the archive/export file will be stored on the log server (false) or locally (true).
- `name` (String) Name of the object.
- `overwrite_file_flag` (String) The overwrite options: 'append' for append option, 'overwrite' for overwrite option, 'use_number_in_file_name' for creating a unique file with number as name, 'fail_task' for failing the task.
- `relative_time_begin` (Number) The begin number of days when the time_limit_type is set to 'relative_time_range' or 'before'.
- `relative_time_end` (Number) The end number of days when the time_limit_type is set to 'relative_time_range' or 'before'.
- `resources` (List of String) URI of the resource.
- `script_name` (String) The script to execute after the task.
- `start_time` (Number) The start time in ms used when the time_limit_type is set to 'absolute_time_range' or 'after'.
- `time_limit_type` (String) Log task time limit type: Time range type. 'before' for 'relative_time_end' number of days, 'relative_time_range' for between 'relative_time_begin' number of days and 'relative_time_end' number of days, 'absolute_time_range' for between 'start_time' time in ms and 'end_time' time in ms. 'today' for today. 'yesterday' for yesterday. 'last_full_week_sun_sat' for last full week (sunday-saturday). 'last_full_week_mon_sun' for last full week (monday-sunday). 'last_full_month' for last full month. 'after' for after 'start_time' time in ms.

## Nested Attributes
- `server_directories` (List of Blocks, see [here](attr_server_directories.md)) The server directories used to retrieve logs.

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
