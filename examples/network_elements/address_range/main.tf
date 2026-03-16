variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_address_range" "tf_address_range_example" {
  ip_range = "10.20.30.40-10.20.30.50"
  name     = "tf_address_range_example"
  comment  = var.resource_comment
}
