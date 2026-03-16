variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_alias" "tf_alias_test" {
  name    = "$ tf test alias"
  comment = var.resource_comment
}
