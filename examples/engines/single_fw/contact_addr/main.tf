variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

locals {
  physical_interfaces = [
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

  contact_address_match_expr = "%s/interfaces[?single_node_interface.address=='%s']|[0].single_node_interface.link[0].href"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "default_location" {
  type = "location"
  name = "Default"
}

data "smc_sub_href" "contact_address_intf1" {
  from_ref = smc_single_fw.tf_fw_contact.link.physical_interface
  match    = format(local.contact_address_match_expr, "Interface 1", "192.168.101.14")
}

output "contact_address_intf1" {
  value = data.smc_sub_href.contact_address_intf1.id
}

resource "smc_single_fw" "tf_fw_contact" {
  name           = "tf_fw_contact"
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href

  nodes {
    firewall_node {
      name   = "tf_fw_contact node"
      nodeid = 1
    }
  }

  dynamic "physical_interfaces" {
    for_each = local.physical_interfaces
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

resource "smc_contact_address" "tf_sample_contact_address" {
  id = data.smc_sub_href.contact_address_intf1.id
  contact_addresses {
    address      = "12.12.12.21"
    dynamic      = false
    location_ref = data.smc_href.default_location.id
  }
}

