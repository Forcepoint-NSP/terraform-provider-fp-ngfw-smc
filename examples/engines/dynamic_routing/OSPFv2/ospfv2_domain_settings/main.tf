variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ospfv2_domain_settings" "ospfv2_domain_settings" {
  abr_type                = "cisco"
  auto_cost_bandwidth     = 100
  deprecated_algorithm    = false
  initial_delay           = 200
  initial_hold_time       = 1000
  max_hold_time           = 10000
  name                    = "tf_ospfv2_domain_settings"
  shutdown_max_metric_lsa = 0
  startup_max_metric_lsa  = 0
  comment                 = var.resource_comment
}
