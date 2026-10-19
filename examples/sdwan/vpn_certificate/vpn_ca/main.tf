variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_vpn_certificate_authority" "test_vpn_ca" {
  name        = "tf_test_vpn_ca"
  certificate = file("${path.module}/ca/ca.crt")
  comment     = var.resource_comment
}

output "test_vpn_ca_href" {
  value = smc_vpn_certificate_authority.test_vpn_ca.link.self
}
