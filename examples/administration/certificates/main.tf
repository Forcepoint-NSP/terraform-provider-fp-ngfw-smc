variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_tls_certificate_authority" "test_ca" {
  name        = "tf_test_ca"
  certificate = file("${path.module}/ca/ca.crt")
  comment     = var.resource_comment
}

output "test_ca_href" {
  value = smc_tls_certificate_authority.test_ca.link.self
}
