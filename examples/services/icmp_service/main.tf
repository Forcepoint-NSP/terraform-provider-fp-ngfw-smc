variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_icmp_service" "tf_icmp_service" {
  icmp_code = 144
  icmp_type = 14
  name      = "tf_icmp_service"
  comment   = var.resource_comment
}

resource "smc_icmp_ipv6_service" "tf_icmpv6_service" {
  comment   = var.resource_comment
  icmp_code = 254
  icmp_type = 41
  name      = "tf_icmpv6_service"
}

resource "smc_icmp_service_group" "tf_icmp_service_group" {
  name    = "tf_icmp_service_group"
  comment = var.resource_comment
  element = [
    smc_icmp_service.tf_icmp_service.id,
    smc_icmp_ipv6_service.tf_icmpv6_service.id,
  ]
}