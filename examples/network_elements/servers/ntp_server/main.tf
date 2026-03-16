variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ntp_server" "tf_ntp_server" {
  address           = "10.2.3.1"
  ipv6_address      = "fc00::10"
  name              = "tf_ntp_server"
  comment           = var.resource_comment
  ntp_auth_key      = "theSuperKey!!"
  ntp_auth_key_id   = 655
  ntp_auth_key_type = "SHA256"
  ntp_host_name     = "myntp.example.com"
}
