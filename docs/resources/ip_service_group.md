---
page_title: "smc_ip_service_group"
subcategory: "services"
description: |-
  This represents an IP-proto Service Group, which is used to group a list of IP-proto Services. It can contain both individual IP-proto Services and other Service Groups.
---

# smc_ip_service_group (Resource)

This represents an IP-proto Service Group, which is used to group a list of IP-proto Services. It can contain both individual IP-proto Services and other Service Groups.

## Examples

- [IP Service Group Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/services/ip-proto_service/main.tf)

Groups multiple IP services for easier policy and rules management.

```hcl
resource "smc_ip_service_group" "tf_ip_proto_service_group" {
  name    = "tf_ip-proto_service_group"
  comment = var.resource_comment
  element = [
    smc_ip_service.tf_ip_proto_service.id,
    data.smc_href.smc_ip_service_TCP.id
  ]
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `element` (List of String) URI of the IP Service Entry belonging to this group.
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
