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
      name   = "Primary Node"
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

resource "smc_internal_endpoint" "ep_192_168_100_14" {
  id        = "${smc_single_fw.tf_single_fw.link.internal_gateway}#*/internal_endpoint/192.168.100.14"
  enabled   = true
  ipsec_vpn = true
}

# Endpoint looked up by a name that contains a space.
# Before the fix, NormalizeHref encoded the space to %20, causing Terraform to
# report "Provider produced inconsistent result after apply".
resource "smc_internal_endpoint" "ep_by_node_name" {
  id        = "${smc_single_fw.tf_single_fw.link.internal_gateway}#*/internal_endpoint/First DHCP*"
  enabled   = false
  ipsec_vpn = false
  ike_phase1_id_type  = "0"  # 0 = DNS_NAME
  ike_phase1_id_value = "test.tfsmc.com"
}
