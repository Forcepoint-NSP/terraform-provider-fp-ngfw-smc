variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_interface_zone" "tf_zone_example" {
  name    = "Northbound"
  comment = var.resource_comment
}

