---
page_title: "smc_ip_service"
subcategory: "services"
description: |-
  This represents an IP-proto service, which is used to define a service based on the IP protocol number. It includes a protocol number that specifies the protocol used by the traffic.
---

# smc_ip_service (Resource)

This represents an IP-proto service, which is used to define a service based on the IP protocol number. It includes a protocol number that specifies the protocol used by the traffic.

## Examples

- [IP Service Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/services/ip-proto_service/main.tf)

Creates a custom IP protocol service by its protocol number for firewall or policy use.

```hcl
resource "smc_ip_service" "tf_ip_proto_service" {
  comment         = var.resource_comment
  name            = "tf_ip-proto_service"
  protocol_number = 125
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `protocol_agent_ref` (String) This represents a Protocol Agent. It is a process on the engines that assists the engine in handling a particular Protocol. Protocol Agents ensure that related connections for a service are properly grouped and evaluated by the engine, as well as assisting the engine with content filtering or network address translation tasks.
- `protocol_number` (Number) The numerical protocol code that the traffic uses.

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
