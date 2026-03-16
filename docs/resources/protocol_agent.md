---
page_title: "smc_protocol"
subcategory: ""
description: |-
  This represents a Protocol Agent. It is a process on the engines that assists the engine in handling a particular Protocol. Protocol Agents ensure that related connections for a service are properly grouped and evaluated by the engine, as well as assisting the engine with content filtering or network address translation tasks.
---

# smc_protocol (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Protocol Agent. It is a process on the engines that assists the engine in handling a particular Protocol. Protocol Agents ensure that related connections for a service are properly grouped and evaluated by the engine, as well as assisting the engine with content filtering or network address translation tasks.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `deep_inspection_registration_number` (Number) Internal use: it generates specific registration number when deep inspection is used. This concerns only for fw version <= 5.2.x.
- `fallback_agent_id` (Number) Internal use: Id of the agent that should be used as a fallback.
- `identifier_situation` (Number) Internal use: Id of the situation (application) which can be used for identifying this protocol.
- `module_build` (String) Dummy module info: build of the sensor module.
- `module_max_version` (String) Dummy module info: Max Version (concatenation between ips and fw max version).
- `module_min_version` (String) Dummy module info: Min Version (concatenation between ips and fw min version).
- `module_name` (String) Dummy module info: name of the sensor module *.so binary file name.
- `module_upload_always` (Boolean) Dummy module info: upload always flag of the sensor module.
- `name` (String) Name of the object.
- `override_module_id` (Number) Overrides the registration number or system key as the module id, if set. Used for inspection module and sidewinder proxy stacked services.
- `protocol_matching_in_inspection` (Boolean) Flag to know if a protocol can be used in inspection policy, for column Protocol. Default value = true
- `registration_number` (Number) Internal use: used by the firewall to reference a protocol agent.
- `supported_protocol` (List of String) List of supported protocols: ethernet, ip-proto, tcp, udp, icmp, physical.
- `type` (String) The Protocol Agent type.

## Nested Attributes
- `link_selection` (Single Block, see [here](link_selection_value.md)) 
- `pa_parameters` (List of Blocks, see [here](attr_abstract_pa_parameter_wrapper.md)) The parameters belonging to the protocol agent.
- `pa_values` (List of Blocks, see [here](attr_abstract_pa_parameter_value_wrapper.md)) The values belonging to the protocol agent.

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
