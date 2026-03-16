variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ip_list" "tf_example_ip_list" {
  name    = "tf_example_ip_list"
  comment = var.resource_comment
}

resource "smc_list" "tf_example_ip_address_list" {
  # the list itself is implicitly created by "smc_ip_list" and always
  # exist this is why we use "id" here and not "from_ref"
  id = smc_ip_list.tf_example_ip_list.link.ip_address_list
  ip = ["104.209.194.165", "104.209.194.52", "104.209.233.92"]
}

