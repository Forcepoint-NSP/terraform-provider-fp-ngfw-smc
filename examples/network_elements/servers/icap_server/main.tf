variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_icap_server" "tf_icap_server" {
  address            = "10.20.20.11"
  name               = "tf_icap_server"
  comment            = var.resource_comment
  icap_include_xhdrs = true
  icap_path          = "myicap"
  icap_secure        = false
  # icap_xhdr_clientip = "X-Client-IP"
  # icap_xhdr_serverip = "X-Server-IP"
  # icap_xhdr_username = "X-Authenticated-User"
  location_ref       = "http://localhost:18082/7.4/elements/location/-1"
}
