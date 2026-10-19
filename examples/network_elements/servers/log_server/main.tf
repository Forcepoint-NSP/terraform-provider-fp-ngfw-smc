variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_log_server" "tf_log_server" {
  name    = "tf_log_server"
  address = "10.0.0.1"
  comment = var.resource_comment
}
