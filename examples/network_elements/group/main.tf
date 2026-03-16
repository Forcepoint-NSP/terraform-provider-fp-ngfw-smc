variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_host" "tf_host1" {
  name      = "tf_host1"
  address   = "192.168.1.44"
  comment   = var.resource_comment
  secondary = ["212.20.1.1", "123.6.5.22"]
}

resource "smc_host" "tf_host2" {
  name    = "tf_host2"
  address = "192.168.1.88"
  comment = var.resource_comment
}

resource "smc_host" "tf_host3" {
  name    = "tf_host3"
  address = "192.168.1.22"
  comment = var.resource_comment
}

resource "smc_group" "tf_sub_group_example" {
  element = [
    smc_host.tf_host1.id,
    smc_host.tf_host2.id
  ]
  name    = "tf_sub_group_example"
  comment = var.resource_comment
}

resource "smc_group" "tf_group_example" {
  element = [
    smc_group.tf_sub_group_example.id,
    smc_host.tf_host3.id
  ]
  name    = "tf_group_example"
  comment = var.resource_comment
}
