variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_autonomous_system" "autonomous_system" {
  as_number = "578"
  name      = "tf_autonomous_system"
  comment   = var.resource_comment
}

resource "smc_external_bgp_peer" "external_bgp_peer" {
  name          = "tf_external_bgp_peer"
  neighbor_as   = smc_autonomous_system.autonomous_system.id
  neighbor_ip   = "10.10.14.10"
  neighbor_port = 179
  comment       = var.resource_comment
}
