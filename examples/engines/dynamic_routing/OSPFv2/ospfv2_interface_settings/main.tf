variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ospfv2_interface_settings" "ospfv2_interface_settings" {
  authentication_type    = "none"
  dead_interval          = 1
  dead_multiplier        = 4
  hello_interval_type    = "fast_hello"
  interface_cost         = 14
  mtu_mismatch_detection = true
  name                   = "tf_ospfv2_interface_settings"
  password               = ""
  retransmit_interval    = 10
  router_priority        = 5
  transmit_delay         = 2
  comment                = var.resource_comment
}
