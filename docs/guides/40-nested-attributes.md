---
page_title: "Representation of Nested Attributes"
subcategory: ""
description: |-
  Syntax and best practices for nested attributes in the SMC Terraform provider
---
# Representation of Nested Attributes

Terraform provides two options for representing nested attributes:

- Legacy syntax (HCL1), which uses blocks
- New syntax (HCL2), which uses JSON-like constructs

The following examples show how to represent a list of two objects using both syntaxes:

```hcl
# Block syntax (HCL1)
resource "myresource" "myexample" {
  mylist {
    myindex = 1
  }
  mylist {
    myindex = 2
  }
}
```

```hcl
# JSON-like syntax (HCL2)
resource "myresource" "myexample" {
  mylist = [
    {
      myindex = 1
      mycomment = "this is hcl2"
    },
    {
      myindex = 2
    }
  ]
}
```

### Pros and Cons of Both Approaches

#### HCL1

This is the legacy syntax. It is slightly more compact, but it can be
harder to recognize that you are working with a list. Additionally, if
you need to iterate, you must use the "dynamic" keyword, which is not
very user-friendly.

#### HCL2

The HCL2 syntax:

- Is more intuitive, as lists are clearly expressed with brackets
- Allows you to assign the entire list to a variable, which can simplify complex configurations.

But the HCL2 syntax has a significant drawback: 

Terraform will silently discard any optional attribute value that is
not part of the schema. In other words, if you make a typo (for
example, "mycommnet" instead of "mycomment"), you will not receive a
warning, and the created resource will not match your Terraform
configuration.

<span style="color:red">
__This is especially problematic for security equipment, as silently
omitting part of a policy rule due to a typo can lead to unintended
access.__
</span>

See related discussions on the Terraform GitHub:
[here](https://github.com/hashicorp/terraform/issues/33570) and
[here](https://github.com/hashicorp/terraform-plugin-framework/issues/810).

### Nested Attributes in the Terraform SMC Provider

Due to the lack of validation in HCL2 syntax, the Terraform SMC
provider uses the legacy HCL1 syntax by default.

However, if you are confident in your configuration, you can use HCL2
syntax for all resources or for a selected subset.

This approach is illustrated in the [getting started](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/getting_started) example:

- Policies and rules use the safer HCL1 syntax
- Firewalls use the more convenient HCL2 syntax to simplify some parts of the configuration

To enable HCL2, create a configuration file named `tfsmc.conf.json` at the root of your module:

#### Using HCL2 Syntax Everywhere

`$ cat tfsmc.conf.json`

```json
{
  "hcl2": true
}
```

#### Using HCL2 for Selected Resources Only

`$ cat tfsmc.conf.json`

```json
{
  "hcl2": ["smc_single_fw", "smc_routing_node"]
}
```

#### Mitigating the Risks of Using "Unsafe" HCL2 Syntax

- Always edit Terraform files with an editor that supports the SMC Terraform provider syntax (e.g., VS Code with the Terraform plugin). The editor will highlight unknown attributes in red.
- Use Terraform modules to create parts or complete resources and enforce type checking of input variables.
