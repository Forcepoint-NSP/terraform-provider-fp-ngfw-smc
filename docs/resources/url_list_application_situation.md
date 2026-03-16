---
page_title: "smc_url_list_application"
subcategory: "situations"
description: |-
  This represents a URL List Application Situation, which is used for URL list filtering in URL filtering.
---

# smc_url_list_application (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a URL List Application Situation, which is used for URL list filtering in URL filtering.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `attacker` (String) The attacker involved in the situation, which can be 'none', 'packet_destination', 'connection_source', or 'connection_destination'. By default, it is set to 'none'.
- `comment` (String) An optional comment for the element. This field is not required.
- `description` (String) A long description of the situation, providing detailed information about its nature and implications.
- `display_name` (String) The internal display name of the situation, which is used for identification purposes within the system.
- `hidden` (Boolean) Indicates whether the situation is hidden from the user interface. If true, it is not displayed to users.
- `identifiable_with_tls_match` (Boolean) Indicates whether the application can be identified by a TLS match.
- `mime_type` (String) The MIME type associated with the situation, which indicates the type of content or data involved in the situation.
- `name` (String) Name of the object.
- `parameter_values` (List of String) URI of the associated parameter value.
- `parent_application_ref` (String) This represents a situation related to an application, including its parent application, protocol agent, TLS match, and associated ports.
- `protocol_agent_ref` (String) This represents a Protocol Agent. It is a process on the engines that assists the engine in handling a particular Protocol. Protocol Agents ensure that related connections for a service are properly grouped and evaluated by the engine, as well as assisting the engine with content filtering or network address translation tasks.
- `sanity_check` (Boolean) Indicates whether a sanity check is performed on the situation. If true, it ensures that the situation meets certain criteria before being processed.
- `severity` (Number) The severity of the situation, ranging from 1 (Info) to 10 (Critical). By default, it is set to 1 (Info).
- `situation_context_ref` (String) This represents a context for a situation, including its minimum and maximum supported engine versions and associated parameters.
- `ssm_mib` (String) The Sidewinder MIB (Management Information Base) associated with the situation, which is used for monitoring and management purposes.
- `tls_match_ref` (String) This represents a TLS Match Situation, which defines matching criteria for the use of the TLS protocol in traffic, allowing you to prevent specified traffic from being decrypted.
- `upcoming_event_enabled` (Boolean) Indicates whether the situation is enabled for upcoming events. If true, it means that the situation will be considered for upcoming event processing.
- `upcoming_event_threshold` (Number) The threshold in days for upcoming events related to the situation. If the situation is enabled for upcoming events, this value indicates how many days in advance the event should be considered.
- `url_entry` (List of String) The URLs.
- `victim` (String) The victim involved in the situation, which can be 'none', 'packet_destination', 'connection_source', or 'connection_destination'. By default, it is set to 'none'.
- `vulnerability_references` (List of String) A set of references to vulnerabilities associated with the situation, providing additional context and information about potential security issues.

## Nested Attributes
- `application_port` (List of Blocks, see [here](attr_application_port.md)) A list of application ports associated with this application situation.
- `pa_values` (List of Blocks, see [here](attr_abstract_pa_parameter_value_wrapper.md)) A list of parameter values associated with this application situation, allowing for dynamic configuration of inspection settings.

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
