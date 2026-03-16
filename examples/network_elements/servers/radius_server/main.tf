variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_smtp_server" "smtp" {
  address              = "10.1.1.123"
  email_sender_address = "sender@example.com"
  email_sender_name    = "example_sender"
  name                 = "tf_smtp"
  comment              = var.resource_comment
  port                 = 25
}
