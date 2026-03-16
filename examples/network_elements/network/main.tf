variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_network" "tf_sample_network" {
  name         = "tf_sample_network"
  comment      = var.resource_comment
  broadcast    = false
  ipv4_network = "10.20.30.0/24"
}
