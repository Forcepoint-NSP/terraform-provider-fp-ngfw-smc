# The two firewall engines, one per cloud site.
#
# Unlike the gcp-ngfw-ha reference stack these are *created* from scratch
# (no import {} blocks) because demo mode starts empty. Interfaces use
# static addresses instead of DHCP so the example applies cleanly.
#
#  nic0 - management (primary_mgt)
#  nic1 - WAN / internet-facing (carries the VPN endpoint + contact address)
#  nic2 - LAN / protected subnet
#
# On create, each engine binds a license and runs initial contact (see the
# actions in smc_actions.tf); the policy is uploaded afterwards, once every
# element exists (smc_policy_upload.tf).

resource "smc_single_fw" "ngfw" {
  for_each       = local.sites
  name           = each.value.name
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href

  domain_server_address {
    value = var.dns_server
    rank  = 1.0
  }

  nodes {
    firewall_node {
      name   = "${each.value.name} node 1"
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
          address       = local.mgmt_ip[each.key]
          network_value = each.value.mgmt_cidr
          primary_mgt   = true
        }
      }
    }
  }

  physical_interfaces {
    physical_interface {
      interface_id = 1
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 1
          address       = local.wan_ip[each.key]
          network_value = each.value.wan_cidr
        }
      }
    }
  }

  physical_interfaces {
    physical_interface {
      interface_id = 2
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 2
          address       = local.lan_ip[each.key]
          network_value = each.value.lan_cidr
        }
      }
    }
  }

  # On first create, bind a license to the node then run initial contact
  # (which writes the node's one-time-password config file locally). The
  # policy upload is deliberately NOT here: it must wait for the policy,
  # VPN and routing elements, so it runs from smc_policy_upload.tf.
  lifecycle {
    action_trigger {
      events = [after_create]
      actions = [
        action.smc_command.bind_license[each.key].id,
        action.smc_initial_contact.contact[each.key].id,
      ]
    }
  }
}
