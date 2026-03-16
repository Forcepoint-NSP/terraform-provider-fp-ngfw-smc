variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "any_network" {
  type = "network"
  name = "Any network"
}

resource "smc_network" "tf_sample_network" {
  name         = "tf_sample_network"
  comment      = var.resource_comment
  ipv4_network = "192.168.100.0/24"

}

resource "smc_router" "tf_sample_router" {
  address = "192.168.100.254"
  name    = "tf_sample_router"
  comment = var.resource_comment
}

resource "smc_network" "tf_sample_network2" {
  name         = "tf_sample_network2"
  comment      = var.resource_comment
  ipv4_network = "192.168.101.0/24"

}

resource "smc_router" "tf_sample_router2" {
  address = "192.168.101.254"
  name    = "tf_sample_router2"
  comment = var.resource_comment
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw"
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href

  nodes {
    firewall_node {
      name   = "myfwnode"
      nodeid = 1
    }
  }

  physical_interfaces {
    physical_interface {
      interface_id = 0
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 0
          address       = "192.168.100.14"
          network_value = "192.168.100.00/24"
          primary_mgt   = true
        }
      }
    }
  }
}

resource "smc_routing_node" "tf_single_fw" {
  id    = smc_single_fw.tf_single_fw.link.routing
  level = "engine_cluster"
  routing_node {
    level = "interface"
    name  = "Interface 0"
    routing_node {
      level = "network"
      href  = smc_network.tf_sample_network.id
      routing_node {
        level = "gateway"
        href  = smc_router.tf_sample_router.id
        routing_node {
          level = "any"
          href  = data.smc_href.any_network.id
        }
      }
    }
  }
}
