variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_dns_server" "tf_dns_server" {
  address         = "10.10.10.1"
  name            = "tf_dns_server"
  comment         = var.resource_comment
  time_to_live    = 10
  update_interval = 5
}
