

variable "interface_id" {}
variable "nodeid" {}
variable "nicid" {}
variable "address" {}
variable "network" {}
variable "mgmt" {}

output "value" {
  value = {
    physical_interface = {
      interface_id = var.interface_id
      interfaces = [{
        single_node_interface = {
          nodeid        = var.nodeid
          nicid         = var.nicid
          address       = var.address
          network_value = var.network
          primary_mgt   = var.mgmt
        }
      }]
    }
  }
}
