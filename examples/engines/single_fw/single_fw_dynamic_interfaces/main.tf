variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw_dynamic_interface"
  log_server_ref = data.smc_href.log_server.href
  comment        = var.resource_comment
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
  physical_interfaces {
    physical_interface {
      name         = "Interface 1"
      comment      = "dynamic IPv4/IPv6 intertace"
      interface_id = "1"
      interfaces {
        single_node_interface {
          dynamic       = true
          dynamic_index = 1
          nicid         = "1"
          nodeid        = 1
        }
      }
      interfaces {
        single_node_interface {
          dynamic            = true
          dynamic_ipv6_index = 2
          nicid              = "1"
          nodeid             = 1
        }
      }
    }
  }

}
