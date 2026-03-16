---
page_title: "smc_pdf_template"
subcategory: ""
description: |-
  This represents the PDF Template used as Style template for PDF reports. It contains settings for page indices, header and footer heights, font colors, and a reference to the content file.
---

# smc_pdf_template (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents the PDF Template used as Style template for PDF reports. It contains settings for page indices, header and footer heights, font colors, and a reference to the content file.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `back_font_color` (String) Font color of back page. It should be specified in hexadecimal format (e.g., #FFFFFF for white).
- `comment` (String) An optional comment for the element. This field is not required.
- `content_file` (String) This is an Internal File. It represents a file used internally within the system, typically for configuration or other internal purposes.
- `content_page` (Number) Index of content page. This is the page where the main content of the report starts.
- `cover_font_color` (String) Font color of cover page. It should be specified in hexadecimal format (e.g., #FFFFFF for white).
- `first_end_page` (Number) First index of end pages.
- `first_start_page` (Number) First index of starting pages.
- `footer_height` (Number) Footer height in mm. This is the height of the footer section on each page of the PDF report.
- `header_height` (Number) Header height in mm. This is the height of the header section on each page of the PDF report.
- `last_end_page` (Number) Last index of end pages.
- `last_start_page` (Number) Last index of starting pages.
- `name` (String) Name of the object.


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
