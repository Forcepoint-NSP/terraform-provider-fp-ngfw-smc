variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_router" "tf_example_router" {
  address = "46.200.30.80"
  name    = "tf_example_router"
  comment = var.resource_comment
  third_party_monitoring {
    netflow   = false
    snmp_trap = false
  }
}
