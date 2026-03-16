---
page_title: "smc_report_operation"
subcategory: ""
description: |-
  This represents a report operation, which is a task that runs a report. It can be one-time or recurrent. It includes options for report design, launch time, overriding duration, filters, email notifications, and export formats.
---

# smc_report_operation (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a report operation, which is a task that runs a report. It can be one-time or recurrent. It includes options for report design, launch time, overriding duration, filters, email notifications, and export formats.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `email` (String) The email address to which the report will be sent as an attachment. If not provided, no email will be sent.
- `export_in_html` (Boolean) Flag indicating whether the report will be generated in HTML format. Optional, if true, the report will be exported in HTML format.
- `export_in_pdf` (Boolean) Flag indicating whether the report will be generated in PDF format. Default choice if no other export formats are specified.
- `export_in_txt` (Boolean) Flag indicating whether the report will be generated in TXT format. Optional, if true, the report will be exported in TXT format.
- `filter_ref` (String) This represents a container for filter expressions, which can be used to define complex filtering rules. It contains a root node that holds the main filter expression.
- `launch_time` (Number) The end time of the report operation, specified in UTC milliseconds from 1970-01-01. If not provided, the current time will be used.
- `name` (String) Name of the object.
- `overriding_duration` (Number) The overriding duration of the report operation in seconds. If not provided, the report design's period duration will be used. Caution: This option is not compatible with the repeat option for recurrent tasks.
- `post_processed` (Boolean) Flag indicating whether the report will be post-processed after generation. If true, additional processing will be applied to the report.
- `repeat` (Boolean) Flag indicating whether the report operation should be repeated according to the period defined in the report design. If true, the report will be generated repeatedly.
- `report_design_ref` (String) This represents a Report Design, which defines the structure and content of reports generated within the system. It includes sections and configurations for report generation.
- `report_per_sender` (Boolean) Flag indicating whether to generate one report per sender. If true, a separate report will be generated for each sender included in the filter.
- `stored` (Boolean) Flag indicating whether the report will be stored. If true, the report will be stored for later retrieval. Default is false.
- `use_elasticsearch` (Boolean) Flag indicating whether to use Elasticsearch for report generation. If true, Elasticsearch will be used; otherwise, local storage will be used.


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
