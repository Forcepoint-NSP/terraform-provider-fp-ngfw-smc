---
page_title: "Importing Resources"
subcategory: ""
description: |-
  How to import existing resources into Terraform
---
# Importing Resources

You can import resources that already exist in the SMC. For example, you may want to import an engine that has already completed initial contact.

There are two methods to import resources:
- adding an `import` block
- using the cli `terraform import`

See
[here](https://developer.hashicorp.com/terraform/language/block/import)
and [here](https://developer.hashicorp.com/terraform/language/import)
for more info.

## Adding an Import Block

### First Step: Get the URL of the Resource to Import

You need to know the URL of the resource. You can use [smc-explorer](https://github.com/Forcepoint/fp-ngfw-smc-explorer) for this:

```bash
$ smc-explorer get-url single_fw/Plano
http://localhost:8082/7.4/elements/single_fw/268586794
```

### Second Step: Add an Import Block

Add a block like this to any Terraform file:

```hcl
# my_imported_fw.tf
import {
  id = "http://localhost:8082/7.4/elements/single_fw/268586794"
  to = smc_single_fw.tf_imported_fw

}
```

### Third Step: Generate the Imported File

The following command will create a file `generated.tf` containing the Terraform configuration of the single firewall you're importing:

```
terraform plan -generate-config-out=generated.tf
```

#### Check the Result:

The file `generated.tf` should contain a `resource "smc_single_fw" "tf_imported_fw"` block with all its attributes populated.

Note: There is a bug in Terraform import (at least up to version 14). You need to edit the generated file and remove the line `provider=fp-ngfw-smc`.

You may also need to make some manual changes to the imported resource:

- Replace bare URLs with query results (see `data "smc_href"`)
- Remove any default sections you do not need for clarity

### Fourth Step: Update the tfstate

To update the tfstate, simply run:

```bash
terraform apply
```

#### Check the Result:

The file `terraform.tfstate` should contain an entry for the imported firewall. In particular, check that the `id` attribute value matches the ID specified in the import block.

You can also use the command:

```bash
$ terraform state list
...
smc_single_fw.tf_imported_fw
...
```

## Using the "terraform import" Command

### First Step: Get the URL of the Resource to Import

You need to know the URL of the resource. You can use [smc-explorer](https://github.com/Forcepoint/fp-ngfw-smc-explorer) for this:

```bash
$ smc-explorer get-url single_fw/Plano
http://localhost:8082/7.4/elements/single_fw/268586794
```

### Second Step: Create a Resource Block for the Imported Resource

You can use [smc-explorer](https://github.com/Forcepoint/fp-ngfw-smc-explorer) to dump the resource you want to import, or create the resource block manually.

Example:

```hcl
# my_imported_fw.tf
resource "smc_single_fw" "tf_imported_fw" {
  name = "..."
  # rest of the configuration here...
}
```

### Third Step: Invoke the Import Command

```bash
terraform import smc_single_fw.tf_imported_fw http://localhost:8082/7.4/elements/single_fw/268586794
```

#### Check the Result

Open the file `terraform.tfstate` and verify that the single firewall is present (for example, look for the URL you used in the import command).

You can also use the command:

```bash
$ terraform state list
...
smc_single_fw.tf_imported_fw
...
```
