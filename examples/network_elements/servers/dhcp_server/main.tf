variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_dhcp_server" "tf_example_dhcp_server" {
  address = "172.31.70.69"
  name    = "tf_example_dhcp_server"
  comment = var.resource_comment
}
