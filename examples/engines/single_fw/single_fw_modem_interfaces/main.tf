variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw_modem_interface"
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
  # SMC BUG ==> https://forcepoint.atlassian.net/browse/SMC-64143
  # physical_interfaces {
  #   modem_interface {
  #     interface_id = "1"
  #     interfaces {
  #       single_node_interface {
  #         dynamic = true
  #         dynamic_index = 2
  #         nicid = "1"
  #         nodeid = 1
  #         pin_code = "321654"
  #         pppoe_password = "*****"
  #         pppoe_servicename = "Orange"
  #         pppoe_username = "myusername"
  #       }
  #     }
  #     modem_interface_type = "3G"
  #     name = "Modem 1"
  #   }
  # }
  physical_interfaces {
    modem_interface {
      interface_id = "0"
      interfaces {
        single_node_interface {
          dynamic       = true
          dynamic_index = 1
          nicid         = "0"
          nodeid        = 1
          pin_code      = "123456"
        }
      }
      modem_interface_type = "LTE"
      name                 = "Modem 0"
    }
  }
}
