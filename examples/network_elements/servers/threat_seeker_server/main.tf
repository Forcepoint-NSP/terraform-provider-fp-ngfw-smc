variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_threatseeker_server" "tf_threatseeker_server" {
  name    = "tf_threatseeker_server"
  dds_url = "https://tss.example.com/api"
  comment = var.resource_comment
}
