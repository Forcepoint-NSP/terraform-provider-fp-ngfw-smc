---
page_title: "smc_service_ethernet"
subcategory: "services"
description: |-
  This represents an Ethernet service, which is used to define a service based on Ethernet protocols. It includes a frame type and optional MAC values for traffic identification.
---

# smc_service_ethernet (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Ethernet service, which is used to define a service based on Ethernet protocols. It includes a frame type and optional MAC values for traffic identification.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `frame_type` (String) The frame type of the Ethernet service, which can be 'eth2', 'ipx', 'llc', or 'snap'. This is required to specify the type of Ethernet traffic being handled.
- `name` (String) Name of the object.
- `protocol_agent_ref` (String) This represents a Protocol Agent. It is a process on the engines that assists the engine in handling a particular Protocol. Protocol Agents ensure that related connections for a service are properly grouped and evaluated by the engine, as well as assisting the engine with content filtering or network address translation tasks.
- `value1` (String) The first MAC value associated with the Ethernet service, which varies based on the frame type. For 'eth2', it is the EtherType code; for 'llc', it is the SSAP address; for 'snap', it is the OUI.
- `value2` (String) The second MAC value associated with the Ethernet service, which varies based on the frame type. For 'llc', it is the DSAP address; for 'snap', it is the type.

## Nested Attributes
- `pa_values` (List of Blocks, see [here](attr_abstract_pa_parameter_value_wrapper.md)) List of parameter values associated with this service. These values are used to configure the protocol agent for this service.

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
