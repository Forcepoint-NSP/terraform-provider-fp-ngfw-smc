---
page_title: "smc_ei_application_situation"
subcategory: "situations"
description: |-
  This represents an ECA Application Situation, which is used for managing ECA applications and their associated executable files.
---

# smc_ei_application_situation (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an ECA Application Situation, which is used for managing ECA applications and their associated executable files.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `attacker` (String) The attacker involved in the situation, which can be 'none', 'packet_destination', 'connection_source', or 'connection_destination'. By default, it is set to 'none'.
- `comment` (String) An optional comment for the element. This field is not required.
- `description` (String) A long description of the situation, providing detailed information about its nature and implications.
- `display_name` (String) The internal display name of the situation, which is used for identification purposes within the system.
- `eca_custom_situation_type` (String) The type of the ECA custom endpoint application.
- `file_name` (String) The name of the file associated with the ECA application.
- `hidden` (Boolean) Indicates whether the situation is hidden from the user interface. If true, it is not displayed to users.
- `mime_type` (String) The MIME type associated with the situation, which indicates the type of content or data involved in the situation.
- `name` (String) Name of the object.
- `parameter_values` (List of String) URI of the associated parameter value.
- `product_name` (String) The name of the product associated with the ECA application.
- `sanity_check` (Boolean) Indicates whether a sanity check is performed on the situation. If true, it ensures that the situation meets certain criteria before being processed.
- `severity` (Number) The severity of the situation, ranging from 1 (Info) to 10 (Critical). By default, it is set to 1 (Info).
- `signer_name` (String) The name of the signer associated with the ECA application.
- `situation_context_ref` (String) This represents a context for a situation, including its minimum and maximum supported engine versions and associated parameters.
- `ssm_mib` (String) The Sidewinder MIB (Management Information Base) associated with the situation, which is used for monitoring and management purposes.
- `upcoming_event_enabled` (Boolean) Indicates whether the situation is enabled for upcoming events. If true, it means that the situation will be considered for upcoming event processing.
- `upcoming_event_threshold` (Number) The threshold in days for upcoming events related to the situation. If the situation is enabled for upcoming events, this value indicates how many days in advance the event should be considered.
- `version_number` (String) The version number of the ECA application.
- `victim` (String) The victim involved in the situation, which can be 'none', 'packet_destination', 'connection_source', or 'connection_destination'. By default, it is set to 'none'.
- `vulnerability_references` (List of String) A set of references to vulnerabilities associated with the situation, providing additional context and information about potential security issues.

## Nested Attributes
- `eca_executable` (List of Blocks, see [here](attr_eca_executable.md)) List of ECA Executable files associated with the ECA application.

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
