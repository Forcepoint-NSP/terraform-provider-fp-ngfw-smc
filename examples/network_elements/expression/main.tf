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

resource "smc_group" "tf_group1" {
  element = [
    smc_host.tf_host1.id,
    smc_host.tf_host2.id
  ]
  name    = "tf_group1"
  comment = var.resource_comment
}

resource "smc_group" "tf_group2" {
  element = [
    smc_host.tf_host3.id
  ]
  name    = "tf_group2"
  comment = var.resource_comment
}

resource "smc_expression" "tf_expression_example" {
  name     = "tf_expression_example"
  comment  = var.resource_comment
  ne_ref   = [smc_group.tf_group1.id]
  operator = "intersection"
  sub_expression {
    name     = "tf_sub_expression"
    ne_ref   = [smc_group.tf_group2.id]
    operator = "exclusion"
  }
}
