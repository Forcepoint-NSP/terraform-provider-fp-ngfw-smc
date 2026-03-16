---
page_title: "Deleting Resources"
subcategory: ""
description: |-
  How to safely delete resources, handle dependency issues, and troubleshoot deletion failures in the SMC Terraform provider.
---

# Deleting Resources

Deleting resources in the SMC Terraform provider can be complicated by
implicit dependencies between resources. 

Terraform may not always be aware of all relationships, which can lead
to deletion errors or resources being deleted in the wrong order.

## Handling Deletion Failures with the Pending Deletion List

As explained in the section on [implicit
sub-resources](20-creating-resources.md#implicit-sub-resources),
Terraform may attempt to delete resources in an order that the SMC
does not allow. 

When a resource cannot be deleted (for example, due to a dependency or
a lock in the SMC GUI), the provider uses a pending deletion
mechanism:

- Resources that fail to delete are added to a "pending delete" list in the file `tfsmc.state.json`.
- A warning is emitted to notify you of the deletion error.
- Terraform continues processing deletions for other resources.
- Before exiting, the provider retries deletion for all resources in the pending list.

If a resource is eventually deleted, it is removed from the
list. However, some resources may remain undeletable (e.g., if they
are locked in the SMC GUI).

**Action required:**

If you see a "failed to delete" warning, check the contents of
`tfsmc.state.json` and attempt manual deletion of the listed
resources.

You can use [smc-explorer](https://github.com/Forcepoint/fp-ngfw-smc-explorer) to delete resources from the command line. For example:

```sh
smc-explorer delete http://mysmc:18082/7.5/elements/network/123456
# or
smc-explorer delete network/mynetwork
```

done with the `depends_on` statement.

## Using the `depends_on` Terraform Feature

To help Terraform understand resource dependencies and control the
order of deletion, you can use the `depends_on` argument. This
explicitly declares that a resource depends on others, ensuring
Terraform deletes resources in the correct order.

Example:

```hcl
resource "smc_example" "dependent" {
  # ...
  depends_on = [
    smc_router.tf_sample_router,
    smc_network.tf_sample_network,
    smc_fw_policy.tf_aaa_policy
  ]
}
```

This approach is especially useful when you know certain resources
must be deleted after others to avoid dependency errors.


