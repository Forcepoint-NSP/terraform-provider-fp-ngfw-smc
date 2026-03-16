variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw_tunnel_interfaces"
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
    tunnel_interface {
      interface_id = "6483"
      interfaces {
        single_node_interface {
          address       = "959b:4bb4:a752:a0ef:991b:45e4:a165:e0ee"
          network_value = "959b:4bb4:a752:a0ef:991b:45e4:a165:e0ee/128"
          nicid         = "6483"
          nodeid        = 1
        }
      }
      name = "Tunnel Interface 6483"
    }
  }
  physical_interfaces {
    tunnel_interface {
      interface_id = "1000"
      interfaces {
        single_node_interface {
          address       = "192.168.123.1"
          network_value = "192.168.123.1/32"
          nicid         = "1000"
          nodeid        = 1
        }
      }
      name = "Tunnel Interface 1000"
    }
  }
}
