variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_location" "tf_location_example" {
  name    = "tf_location_example"
  comment = var.resource_comment
}

