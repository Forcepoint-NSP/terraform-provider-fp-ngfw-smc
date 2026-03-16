---
page_title: "smc_sub_ipv4_fw_policy"
subcategory: "policy"
description: |-
  This represents a Firewall Sub-Policy, which is used to group related Firewall IPv4 Access rules for better organization and processing efficiency.
---

# smc_sub_ipv4_fw_policy (Resource)

This represents a Firewall Sub-Policy, which is used to group related Firewall IPv4 Access rules for better organization and processing efficiency.

## Examples

- [sub_ipv4_fw_policy/main.tf](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/policies/sub_ipv4_fw_policy/main.tf): Creates a sub-IPv4 firewall policy.

This demonstrates the minimal configuration for the `smc_sub_ipv4_fw_policy` resource, used to build hierarchical firewall policy designs.

```hcl
resource "smc_sub_ipv4_fw_policy" "sub_policy" {
  name    = "tf_example_subpolicy"
  comment = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
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
