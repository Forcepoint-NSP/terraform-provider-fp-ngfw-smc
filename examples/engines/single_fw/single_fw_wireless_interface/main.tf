variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw_wireless_interface"
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
    wireless_interface {
      channel_width = "80"
      interface_id  = "1"
      mtu           = -1
      name          = "Wireless Interface 1"
      ssid_interfaces {
        interface_id = "1.1"
        interfaces {
          single_node_interface {
            address       = "192.168.10.1"
            network_value = "192.168.10.0/24"
            nicid         = "1.1"
            nodeid        = 1
          }
        }
        interfaces {
          single_node_interface {
            address       = "fc10::1"
            network_value = "fc10::/64"
            nicid         = "1.1"
            nodeid        = 1
          }
        }
        mac_filtering {
          mode = "disabled"
        }
        mtu            = -1
        name           = "SSID SSID"
        security_mode  = "wpa_personal"
        ssid           = "SSID"
        ssid_broadcast = false
        syn_mode       = "default"
        wpa_mode       = "CCMP"
        wpa_shared_key = "SsidSuperPassword00!"
      }
      syn_mode              = "default"
      transmit_power        = "15dbm"
      wireless_band         = "5"
      wireless_channel      = "32"
      wireless_country      = "US"
      wireless_network_mode = "ac"
    }
  }
}
