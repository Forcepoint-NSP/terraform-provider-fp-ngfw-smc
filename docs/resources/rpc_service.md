---
page_title: "smc_rpc_service"
subcategory: "services"
description: |-
  This represents a SUN-RPC service, which is used to define a service based on the Remote Procedure Call (RPC) protocol. It includes a program number, transport type, and optional RPC version for traffic identification.
---

# smc_rpc_service (Resource)

This represents a SUN-RPC service, which is used to define a service based on the Remote Procedure Call (RPC) protocol. It includes a program number, transport type, and optional RPC version for traffic identification.

## Examples

- [RPC Service Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/services/sun-rpc_service/main.tf)

Defines an RPC service by program number and protocol for firewall rules.

```hcl
resource "smc_rpc_service" "tf_rpc_service" {
  name           = "tf_rpc_service"
  comment        = var.resource_comment
  program_number = "14"
  rpc_version    = "2"
  transport      = "both"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `program_number` (String) The program number that the RPC traffic uses, which is a unique identifier for the RPC program.
- `protocol_agent_ref` (String) This represents a Protocol Agent. It is a process on the engines that assists the engine in handling a particular Protocol. Protocol Agents ensure that related connections for a service are properly grouped and evaluated by the engine, as well as assisting the engine with content filtering or network address translation tasks.
- `rpc_version` (String) The RPC version number for the service, which is optional. If not specified, it matches traffic of any version.
- `transport` (String) The transport type for the RPC service, which can be 'tcp', 'udp', or 'both'. If not specified, it defaults to 'both'.

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
