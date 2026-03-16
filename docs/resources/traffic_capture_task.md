---
page_title: "smc_traffic_capture_task"
subcategory: "admin"
description: |-
  This represents a Traffic Capture Task, which is used to capture network traffic data. It is a type of task that can be scheduled and executed to monitor and log network traffic for analysis.
---

# smc_traffic_capture_task (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Traffic Capture Task, which is used to capture network traffic data. It is a type of task that can be scheduled and executed to monitor and log network traffic for analysis.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `capture_headers_only` (Boolean) Flag to indicate whether to capture only headers in the traffic capture task. If true, only the headers of the traffic will be captured, excluding the payload.
- `comment` (String) An optional comment for the element. This field is not required.
- `description` (String) Description of the traffic capture task, providing details about its purpose or configuration.
- `destination_file` (String) The destination file path where the captured traffic data will be stored. This is the file location on the server or client where the traffic capture results will be saved.
- `duration` (Number) The duration in seconds for which the traffic capture task will run. This defines how long the task will capture network traffic data.
- `max_file_size` (Number) The maximum size of the file for the traffic capture task, in bytes.
- `name` (String) Name of the object.
- `resources` (List of String) URI of the resource.
- `sg_info_option` (Boolean) Flag to indicate whether to include SGInfo in the traffic capture task. If true, SGInfo data will be included in the captured traffic information.
- `traffic_capture_comment` (String) Comment for the traffic capture task, describing its purpose or details.

## Nested Attributes
- `interface_setting` (List of Blocks, see [here](attr_traffic_capture_interface_settings.md)) The interface setting for capturing traffic on a specific network interface. This defines how traffic should be captured on a particular interface.

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
