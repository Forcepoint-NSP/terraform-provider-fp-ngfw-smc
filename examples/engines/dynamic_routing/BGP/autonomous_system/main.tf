variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_autonomous_system" "autonomous_system" {
  as_number = "578"
  name      = "tf_autonomous_system"
  comment   = var.resource_comment
}
