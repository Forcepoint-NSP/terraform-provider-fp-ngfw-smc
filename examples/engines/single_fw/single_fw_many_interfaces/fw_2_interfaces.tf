locals {
  intf_list = [
    {
      interface_id = 0
      nodeid       = 1
      nicid        = 0
      address      = "192.168.100.14"
      network      = "192.168.100.0/24"
      primary_mgt  = true
    },
    {
      interface_id = 1
      nodeid       = 1
      nicid        = 1
      address      = "192.168.101.14"
      network      = "192.168.101.0/24"
      primary_mgt  = false
    }
  ]
}

resource "smc_single_fw" "tf_fw_2_interfaces" {
  name           = "tf_fw_2_interfaces"
  log_server_ref = data.smc_href.log_server.href

  nodes {
    firewall_node {
      name   = "tf_fw_2_interfaces_node"
      nodeid = 1
    }
  }

  dynamic "physical_interfaces" {
    for_each = local.intf_list
    content {
      physical_interface {
        interface_id = physical_interfaces.value.interface_id

        interfaces {
          single_node_interface {
            nodeid        = physical_interfaces.value.nodeid
            nicid         = physical_interfaces.value.nicid
            address       = physical_interfaces.value.address
            network_value = physical_interfaces.value.network
            primary_mgt   = physical_interfaces.value.primary_mgt
          }
        }
      }
    }
  }
}
