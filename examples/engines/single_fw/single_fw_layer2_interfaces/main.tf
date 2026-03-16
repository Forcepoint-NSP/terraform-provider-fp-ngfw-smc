variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

resource "smc_logical_interface" "capture_logical_interface" {
  name    = "tf_capture_logical_interface"
  comment = var.resource_comment
}

resource "smc_logical_interface" "inline_logical_interface" {
  name    = "tf_inline_logical_interface"
  comment = var.resource_comment
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw_layer2_interfaces"
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
      cvi_mode     = "packetdispatch"
      interface_id = "2"
      interfaces {
        inline_l2fw_interface {
          inspect_qinq              = true
          inspect_unspecified_vlans = true
          logical_interface_ref     = smc_logical_interface.inline_logical_interface.id
          nicid                     = "2-3"
        }
      }
      name = "Interface 2 - Interface 3 (Inline L2FW)"
    }
  }
  physical_interfaces {
    physical_interface {
      cvi_mode     = "packetdispatch"
      interface_id = "1"
      interfaces {
        capture_interface {
          inspect_qinq              = true
          inspect_unspecified_vlans = true
          logical_interface_ref     = smc_logical_interface.capture_logical_interface.id
          nicid                     = "1"
        }
      }
      name = "Interface 1 (Capture)"
    }
  }
}
