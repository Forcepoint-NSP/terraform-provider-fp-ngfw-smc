variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ip_service" "tf_ip_proto_service" {
  comment         = var.resource_comment
  name            = "tf_ip-proto_service"
  protocol_number = 125
}

data "smc_href" "smc_ip_service_TCP" {
  name = "TCP"
  type = "ip_service"
}

resource "smc_ip_service_group" "tf_ip_proto_service_group" {
  name    = "tf_ip-proto_service_group"
  comment = var.resource_comment
  element = [
    smc_ip_service.tf_ip_proto_service.id,
    data.smc_href.smc_ip_service_TCP.id
  ]
}