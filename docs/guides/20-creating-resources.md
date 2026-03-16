---
page_title: "Creating Resources"
subcategory: ""
description: |-
  How to create top-level resources, sub-resources, and implicit sub-resources
---
# Creating resources and sub-resources

There are 3 different scenarios regarding the creation of resources

## Creation of top-level resources

This is the simplest case: these resources are self-contained and
they do not rely on any pre-existing resource.

For example smc_single_fw or smc_fw_policy.

To create such resources, simply add a "resource" declaration in your
terraform file with the proper parameters. Example

```hcl
resource "smc_host" "tf_host1" {
  name      = "tf_host1"
  address   = "192.168.1.44"
}
```

After successful creation (terraform apply), the `id` field will be
populated with the url of the resource.

## Creation of sub-resources

Sub-resources are resources that cannot exist independently; they require a parent resource to be present.

For example, access rules (e.g., smc_fw_ipv4_access_rule) depend on the existence of a policy.

All sub-resources include an additional attribute called `from_ref`, which allows you to specify the URL of the parent resource. 
The `from_ref` attribute typically points to one of the `link` attributes of the parent resource. 
You can use [smc-explorer](https://github.com/Forcepoint/fp-ngfw-smc-explorer) to dynamically discover the list of links for a given resource.

Example: Creating an access rule as a sub-resource of a policy.

```hcl
resource "smc_fw_policy" "example" {
  name     = "tf_example_access_rules_policy2"
  template = data.smc_href.default_policy_template.id
}

resource "smc_fw_ipv4_access_rule" "allow_https" {
  from_ref = smc_fw_policy.example.link.fw_ipv4_access_rules
  name     = "allow-https-from-any-to-any"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.https_service.id]
  }
  action { action = ["allow"] }
}
```

## <a name="implicit-sub-resources"></a>Creation of implicit sub-resources

Implicit sub-resources are created automatically by the SMC when a
top-level resource is provisioned. These sub-resources are not
directly created by Terraform.

For example, when an engine (such as a single firewall) is created,
the SMC automatically creates an internal gateway and a routing node
resource.

The challenge is that you cannot create these resources from
Terraform, because they already exist in the SMC.

To manage these sub-resources with Terraform, you must know the URL of
the resource and explicitly set the `id` field in your configuration.

When a resource specifies an `id`, several things occur:
- The creation phase becomes an update request at the SMC API level (HTTP PUT operation).
- Since Terraform did not create the sub-resource, it cannot destroy it. As a result, running `terraform destroy` on these sub-resources is essentially a no-op.

Example: Managing a routing table for a single firewall

```hcl
resource "smc_single_fw" "tf_single_fw" {
  name = "tf_single_fw"
  # other attributes
  ...
}

resource "smc_routing_node" "tf_single_fw" {
  id = smc_single_fw.tf_single_fw.link.routing
  # other attributes
  ...
}
```

## Sub-Resources vs. Implicit Sub-Resources

To summarize the previous sections:

- To create a sub-resource, use the `from_ref` attribute to reference the parent resource's URL. This explicitly links the new resource to its parent.
- To manage an implicit sub-resource, use the `id` attribute to reference a resource that was automatically created by the SMC. This allows Terraform to manage an existing sub-resource that was not created directly by Terraform.

Some resources can be managed as either sub-resources or implicit
sub-resources, depending on your use case.

For example, an internal gateway is implicitly created when you create
a firewall (and can be referenced using `id`). However, you can also
create additional internal gateways explicitly by using the `from_ref`
attribute (e.g., `smc_single_fw.tf_single_fw1.link.internal_gateway`).
