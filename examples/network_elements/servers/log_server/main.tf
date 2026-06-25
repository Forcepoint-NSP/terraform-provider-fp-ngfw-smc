variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_log_server" "tf_log_server" {
  name    = "tf_log_server"
  address = "10.0.0.1"
  uiid    = "550e8400-e29b-41d4-a716-446655440000"
  comment = var.resource_comment
}
