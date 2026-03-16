variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "igmp_querier_settings" {
  name = "IGMP v3"
  type = "igmp_querier_settings"
}

resource "smc_pim_ipv4_interface_settings" "pim_interface_settings" {
  dr_priority               = 14
  igmp_querier_settings_ref = data.smc_href.igmp_querier_settings.id
  name                      = "tf_pim_interface_settings"
  random_delay              = 10
  zbr                       = "224.0.1.1,239.255.255.250"
  comment                   = var.resource_comment
}
