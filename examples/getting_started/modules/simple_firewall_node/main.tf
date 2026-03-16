
variable "name" {}
variable "nodeid" {}

output "value" {
  value = {
    firewall_node = {
      name   = var.name
      nodeid = var.nodeid
    }
  }
}

output "nodeid" {
  value = var.nodeid
}
