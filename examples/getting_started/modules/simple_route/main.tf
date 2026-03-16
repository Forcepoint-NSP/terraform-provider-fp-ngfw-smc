
variable "interface" {}
variable "network" {}
variable "gateway" {}
variable "destination" {}

output "value" {
  value = {
    level = "interface"
    name  = var.interface
    routing_node = [{
      level = "network"
      href  = var.network
      routing_node = [{
        level = "gateway"
        href  = var.gateway
        routing_node = [{
          level = "any"
          href  = var.destination
        }]
      }]
    }]
  }
}
