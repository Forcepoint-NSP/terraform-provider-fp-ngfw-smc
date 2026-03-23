---
page_title: "smc_match_expression"
subcategory: "network_elements"
description: |-
  This element represents the Rule Source/Destination/Service Match expression, which defines how traffic is matched against rules in a policy.
---

# smc_match_expression (Resource)

This element represents the Rule Source/Destination/Service Match expression, which defines how traffic is matched against rules in a policy.

## Examples

- [fw_policy_with_match_expression/main.tf](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.741.0/examples/policies/fw_policy_with_match_expression/main.tf): Builds a match expression referencing multiple resources in a firewall policy.

This snippet shows an `smc_match_expression` referencing a host and a network, for use in access rules for refined traffic matching.

```hcl
resource "smc_match_expression" "match_expression_host1_net1" {
  name = "tf_match_expression_host1_net1"
  ref  = [smc_network.network_1.id, smc_host.host_1.id]
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `ref` (List of String) URI of the Match Element.
- `service_port_selection` (String) The port selection method used in the Match Expression, which can be 'standard', 'any', or 'automatic'.


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
