---
page_title: "Relationship Between Resources"
subcategory: ""
description: |-
  Understanding parent/child relationships and resource linking
---
# Relationship Between Resources

In the SMC Terraform provider, resources often have relationships with one another. Understanding these relationships is essential for building correct and maintainable configurations.

## Parent and Child Resources

Some resources are considered top-level (such as firewalls or policies), while others are sub-resources that depend on a parent. For example, an access rule is a sub-resource of a firewall policy, and an interface is a sub-resource of a firewall engine.

## Discovering Resource Relationships

To correctly reference and link resources, you often need to know the unique identifier (URL) of a parent or related resource. The provider offers data sources to help with this:

- [smc_href](../data-sources/href): Use this data source to discover the ID (URL) of top-level or system resources. For example, to reference a default policy template or a specific engine.
- [smc_sub_href](../data-sources/sub_href): Use this data source to discover the ID (URL) of sub-resources that are children of another resource. For example, to find the id of an internal gateway

These data sources allow you to dynamically look up resource references, making your Terraform configuration more robust and less dependent on hardcoded URLs.

## Example: Refencing system elements

Get the reference of an existing policy template

```hcl
data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

resource "smc_fw_policy" "tf_aaa_policy" {
  name     = "tf_aaa"
  template = data.smc_href.default_policy_template.id
  comment  = var.resource_comment
}
```

## Example: Get the reference of an internal gateway

```hcl
data "smc_sub_href" "internal_gateway_1_ref" {
  from_ref = smc_single_fw.tf_single_fw1.link.internal_gateway
  match    = "tf_single_fw1*"
}

resource "smc_internal_gateway" "internal_gateway_1" {
  id         = data.smc_sub_href.internal_gateway_1_ref.id
  name       = "internal_gateway_1"
  # rest of the attributes
  # ...
}
```
