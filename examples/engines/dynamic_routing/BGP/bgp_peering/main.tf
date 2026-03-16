variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ip_access_list" "ip_access_list" {
  comment = var.resource_comment
  entries {
    ip_access_list_entry {
      action = "deny"
      subnet = "172.16.16.0/21"
    }
  }
  entries {
    ip_access_list_entry {
      action = "permit"
      subnet = "192.168.100.0/24"
    }
  }
  name = "tf_ip_access_list"
}

resource "smc_as_path_access_list" "as_path_access_list" {
  comment = var.resource_comment
  entries {
    as_path_access_list_entry {
      action     = "permit"
      expression = "^65001_"
    }
  }
  entries {
    as_path_access_list_entry {
      action     = "deny"
      expression = "_65002$"
    }
  }
  name = "tf_as_path_access_list"
}

resource "smc_bgp_connection_profile" "bgp_connection_profile" {
  connect            = 150
  md5_password       = "TheSuperPassword00!!"
  name               = "tf_bgp_connection_profile"
  session_hold_timer = 190
  session_keep_alive = 70
  comment            = var.resource_comment
}

resource "smc_bgp_peering" "bgp_peering" {
  bfd_enabled               = false
  connected_check           = "disabled"
  connection_profile        = smc_bgp_connection_profile.bgp_connection_profile.id
  default_originate         = false
  dont_capability_negotiate = false
  inbound_aspath_filter     = smc_as_path_access_list.as_path_access_list.id
  inbound_ip_filter         = smc_ip_access_list.ip_access_list.id
  local_as_option           = "not_set"
  max_prefix_option         = "not_enabled"
  name                      = "tf_bgp_peering"
  next_hop_self             = true
  orf_option                = "disabled"
  outbound_aspath_filter    = smc_as_path_access_list.as_path_access_list.id
  outbound_ip_filter        = smc_ip_access_list.ip_access_list.id
  override_capability       = false
  remove_private_as         = false
  route_reflector_client    = true
  send_community            = "no"
  soft_reconfiguration      = true
  ttl_option                = "disabled"
  comment                   = var.resource_comment
}
