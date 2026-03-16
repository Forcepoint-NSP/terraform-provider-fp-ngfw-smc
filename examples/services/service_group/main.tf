variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "tcp_service_bgp" {
  name = "BGP"
  type = "tcp_service"
}

data "smc_href" "udp_service_biff" {
  name = "Biff"
  type = "udp_service"
}

resource "smc_service_group" "tf_service_group" {
  name    = "tf_service_group"
  comment = var.resource_comment
  element = [
    data.smc_href.tcp_service_bgp.id,
    data.smc_href.udp_service_biff.id
  ]
}