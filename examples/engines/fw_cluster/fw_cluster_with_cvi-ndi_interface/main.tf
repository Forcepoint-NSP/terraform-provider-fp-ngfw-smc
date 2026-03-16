variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "Log Server"
}

data "smc_href" "internal_zone" {
  type = "interface_zone"
  name = "Internal"
}

resource "smc_fw_cluster" "tf_cluster_example" {
  name           = "tf_cluster_example"
  cluster_mode   = "balancing"
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href
  nodes {
    firewall_node {
      name   = "tf_cluster_example node 1"
      nodeid = 1
    }
  }
  nodes {
    firewall_node {
      name   = "tf_cluster_example node 2"
      nodeid = 2
    }
  }
  physical_interfaces {
    physical_interface {
      zone_ref     = data.smc_href.internal_zone.href
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
}
