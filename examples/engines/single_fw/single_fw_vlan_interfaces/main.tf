variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw_vlan_interfaces"
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
      interface_id           = "1"
      name                   = "Interface 1"
      virtual_engine_vlan_ok = false
      vlan_interfaces {
        interface_id = "1.2000"
        interfaces {
          single_node_interface {
            address       = "fc00:2000:2000:2000::2"
            network_value = "fc00:2000:2000:2000::/64"
            nicid         = "1.2000"
            nodeid        = 1
          }
        }
        name = "VLAN 1.2000"
      }
      vlan_interfaces {
        interface_id = "1.100"
        interfaces {
          single_node_interface {
            address       = "100.100.100.100"
            network_value = "100.100.100.0/24"
            nicid         = "1.100"
            nodeid        = 1
          }
        }
        name = "VLAN 1.100"
      }
    }
  }
}
