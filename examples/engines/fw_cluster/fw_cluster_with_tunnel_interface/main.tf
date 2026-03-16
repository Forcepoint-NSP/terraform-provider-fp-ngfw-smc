variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

locals {
  engine_name = "tf_fw_cluster_with_layer2_interface"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "Log Server"
}

resource "smc_fw_cluster" "tf_cluster_example" {
  name           = local.engine_name
  cluster_mode   = "balancing"
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href
  nodes {
    firewall_node {
      name   = "${local.engine_name} node 1"
      nodeid = 1
    }
  }
  nodes {
    firewall_node {
      name   = "${local.engine_name} node 2"
      nodeid = 2
    }
  }
  physical_interfaces {
    physical_interface {
      interface_id = "0"
      name         = "Interface 0"
      cvi_mode     = "packetdispatch"
      macaddress   = "00:00:5e:00:02:00"

      interfaces {
        node_interface {
          nicid             = "0"
          nodeid            = 1
          address           = "192.168.2.21"
          network_value     = "192.168.2.0/24"
          outgoing          = true
          primary_heartbeat = true
          primary_mgt       = true
        }
      }
      interfaces {
        node_interface {
          nicid             = "0"
          nodeid            = 2
          address           = "192.168.2.22"
          network_value     = "192.168.2.0/24"
          outgoing          = true
          primary_heartbeat = true
          primary_mgt       = true

        }
      }
      interfaces {
        cluster_virtual_interface {
          nicid           = "0"
          address         = "192.168.2.254"
          network_value   = "192.168.2.0/24"
          relayed_by_dhcp = false

        }
      }
    }
  }
  physical_interfaces {
    tunnel_interface {
      interface_id = "1011"
      interfaces {
        node_interface {
          address       = "fc00:12e:1243::121"
          network_value = "fc00:12e:1243::12/128"
          nicid         = "1011"
          nodeid        = 1
        }
      }
      interfaces {
        node_interface {
          address       = "fc00:12e:1243::122"
          network_value = "fc00:12e:1243::12/128"
          nicid         = "1011"
          nodeid        = 2
        }
      }
      interfaces {
        cluster_virtual_interface {
          address       = "fc00:12e:1243::12"
          network_value = "fc00:12e:1243::12/128"
          nicid         = "1011"
        }
      }
      mtu  = -1
      name = "Tunnel Interface 1011"
    }
  }
  physical_interfaces {
    tunnel_interface {
      interface_id = "1000"
      interfaces {
        node_interface {
          address       = "10.100.200.1"
          network_value = "10.100.200.0/24"
          nicid         = "1000"
          nodeid        = 1
        }
      }
      interfaces {
        node_interface {
          address       = "10.100.200.4"
          network_value = "10.100.200.0/24"
          nicid         = "1000"
          nodeid        = 2
        }
      }
      name      = "Tunnel Interface 1000"
      qos_limit = -1
    }
  }

}
