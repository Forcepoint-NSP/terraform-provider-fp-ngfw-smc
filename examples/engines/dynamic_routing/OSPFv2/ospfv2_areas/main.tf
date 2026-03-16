variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "ospfv2_interface_settings" {
  name = "Default OSPFv2 Interface Settings"
  type = "ospfv2_interface_settings"
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

resource "smc_network" "tf_sample_network" {
  name         = "tf_sample_network"
  broadcast    = false
  ipv4_network = "10.20.30.0/24"
  comment      = var.resource_comment
}

resource "smc_ospfv2_area" "ospfv2_areas" {
  area_id                = 147
  area_type              = "normal"
  inbound_filters_ref    = [smc_ip_access_list.ip_access_list.id]
  interface_settings_ref = data.smc_href.ospfv2_interface_settings.id
  name                   = "tf_ospfv2_area"
  comment                = var.resource_comment
  ospf_abr_substitute_container {
    subnet_ref      = smc_network.tf_sample_network.id
    substitute_type = "aggregate"
  }
  ospfv2_virtual_links_endpoints_container {
    interface_settings_ref = data.smc_href.ospfv2_interface_settings.id
    router_id_endpoint_a   = "192.168.10.14"
    router_id_endpoint_b   = "198.168.10.254"
  }
  outbound_filters_ref  = [smc_ip_access_list.ip_access_list.id]
  shortcut_capable_area = true
}
