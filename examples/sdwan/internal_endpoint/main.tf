variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw"
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
}

resource "smc_internal_endpoint" "ep_192_168_100_14" {
  id        = "${smc_single_fw.tf_single_fw.link.internal_gateway}#*/internal_endpoint/192.168.100.14"
  enabled   = true
  ipsec_vpn = true
}
