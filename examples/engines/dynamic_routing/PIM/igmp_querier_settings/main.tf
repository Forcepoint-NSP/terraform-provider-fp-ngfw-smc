variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_igmp_querier_settings" "ipgm_profile" {
  igmp_version   = "igmpv3"
  name           = "tf_igmp_querier_settings"
  query_interval = 180
  comment        = var.resource_comment
}
