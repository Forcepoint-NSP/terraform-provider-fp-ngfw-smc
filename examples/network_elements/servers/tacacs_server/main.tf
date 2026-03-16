variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_authentication_service" "tf_tacacs_auth" {
  name = "tf_tacacs_auth"
  type = "tacacs"
}

resource "smc_tacacs_server" "tacacs" {
  address         = "mytacacs.example.com"
  clear_text      = false
  name            = "tf_tacacs_server"
  comment         = var.resource_comment
  port            = 49
  provided_method = [smc_authentication_service.tf_tacacs_auth.id]
  retries         = 2
  shared_secret   = "secret_shared00!!"
  timeout         = 10
}
