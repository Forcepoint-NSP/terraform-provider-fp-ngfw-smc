variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_domain_name" "tfdomain_name" {
  domain_name_entry = ["test.example.com", "prod.example.com", "example.com"]
  name              = "tfdomain_name"
  comment           = var.resource_comment
}
