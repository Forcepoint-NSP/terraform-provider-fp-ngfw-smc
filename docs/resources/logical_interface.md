---
page_title: "smc_logical_interface"
subcategory: "interfaces"
description: |-
  This represents a Logical Interface. It is an IPS Element used in the IPS policies to represent one or more physical network interfaces as defined in the IPS engine properties.
---

# smc_logical_interface (Resource)

This represents a Logical Interface. It is an IPS Element used in the IPS policies to represent one or more physical network interfaces as defined in the IPS engine properties.

## Examples

- [Logical Interface Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/engines/single_fw/single_fw_layer2_interfaces/main.tf)

Defines a logical interface (e.g., capture or inline) on a firewall engine.

```hcl
resource "smc_logical_interface" "capture_logical_interface" {
  name    = "tf_capture_logical_interface"
  comment = var.resource_comment
}

resource "smc_logical_interface" "inline_logical_interface" {
  name    = "tf_inline_logical_interface"
  comment = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `one_lan` (Boolean) Indicates if the logical interface is a single LAN interface.


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
