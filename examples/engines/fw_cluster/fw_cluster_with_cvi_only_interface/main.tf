variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

locals {
  engine_name = "tf_fw_cluster_with_cvi_only_interface"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "Log Server"
}

resource "smc_fw_cluster" "tf_cluster_example" {
  name           = local.engine_name
  comment        = var.resource_comment
  cluster_mode   = "balancing"
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
    physical_interface {
      comment      = "CVI only"
      cvi_mode     = "packetdispatch"
      interface_id = "2"
      interfaces {
        cluster_virtual_interface {
          address       = "10.100.101.1"
          network_value = "10.100.101.0/24"
          nicid         = "2"
        }
      }
      macaddress = "00:00:5e:04:00:01"
      name       = "Interface 2"
    }
  }

}
