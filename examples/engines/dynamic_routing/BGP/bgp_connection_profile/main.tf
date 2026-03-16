variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_bgp_connection_profile" "bgp_connection_profile" {
  connect            = 150
  md5_password       = "TheSuperPassword00!!"
  name               = "tf_bgp_connection_profile"
  session_hold_timer = 190
  session_keep_alive = 70
  comment            = var.resource_comment
}
