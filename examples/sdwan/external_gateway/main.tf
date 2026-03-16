variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

locals {
  firewalls = {
    "fw1" = {
      address = "192.168.101.1"
      network = "192.168.101.0/24"
    }
    "fw2" = {
      address = "192.168.102.1"
      network = "192.168.102.0/24"
    }
  }

  # All interfaces per firewall (primary + additional)
  all_interfaces = {
    "fw1" = {
      "0" = {
        address     = "192.168.101.1"
        network     = "192.168.101.0/24"
        primary_mgt = true
      }
      "1" = {
        address     = "192.168.201.1"
        network     = "192.168.201.0/24"
        primary_mgt = false
      }
      "2" = {
        address     = "192.168.202.1"
        network     = "192.168.202.0/24"
        primary_mgt = false
      }
    }
    "fw2" = {
      "0" = {
        address     = "192.168.102.1"
        network     = "192.168.102.0/24"
        primary_mgt = true
      }
      "1" = {
        address     = "192.168.211.1"
        network     = "192.168.211.0/24"
        primary_mgt = false
      }
    }
  }
}

data "smc_href" "suite_b_gcm_128" {
  type = "vpn_profile"
  name = "Suite-B-GCM-128"
}

data "smc_href" "gateway_profile" {
  type = "gateway_profile"
  name = "Default (all capabilities)"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_sub_href" "internal_gateway" {
  for_each = local.firewalls
  from_ref = smc_single_fw.tf_single_fw[each.key].link.internal_gateway
  match    = "tf_single_fw*"
}


resource "smc_single_fw" "tf_single_fw" {
  for_each       = local.firewalls
  name           = "tf_single_${each.key}"
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href

  nodes {
    firewall_node {
      name   = "node ${each.key}"
      nodeid = 1
    }
  }

  dynamic "physical_interfaces" {
    for_each = local.all_interfaces[each.key]
    content {
      physical_interface {
        interface_id = tonumber(physical_interfaces.key)
        interfaces {
          single_node_interface {
            nodeid        = 1
            nicid         = tonumber(physical_interfaces.key)
            address       = physical_interfaces.value.address
            network_value = physical_interfaces.value.network
            primary_mgt   = physical_interfaces.value.primary_mgt
          }
        }
      }
    }
  }
}

resource "smc_vpn" "tf_sample_vpn" {
  mobile_vpn_topology_mode = "None"
  name                     = "tf_sample_vpn"
  nat                      = true
  vpn_profile              = data.smc_href.suite_b_gcm_128.id
  comment                  = var.resource_comment
}

resource "smc_gateway_node" "tf_sample_central_gateway" {
  from_ref   = smc_vpn.tf_sample_vpn.link.central_gateway_node
  name       = "tf_sample_central_gateway"
  node_usage = "central"
  gateway    = data.smc_sub_href.internal_gateway["fw1"].id
}

resource "smc_gateway_node" "tf_sample_satellite_gateway2" {
  from_ref   = smc_vpn.tf_sample_vpn.link.satellite_gateway_node
  name       = "tf_sample_satellite_external_gw"
  node_usage = "satellite"
  gateway    = smc_external_gateway.external_gateway.id
}

resource "smc_internal_endpoint" "ep_fw1" {
  id = format("%s#%s", smc_single_fw.tf_single_fw["fw1"].link.internal_gateway,
  "*/internal_endpoint/192.168.101.1")
  enabled   = true
  ipsec_vpn = true
}

resource "smc_internal_endpoint" "ep_fw2" {
  id = format("%s#%s", smc_single_fw.tf_single_fw["fw2"].link.internal_gateway,
  "*/internal_endpoint/192.168.102.1")
  enabled        = true
  ipsec_vpn      = true
  ssl_vpn_tunnel = true
}

resource "smc_network" "nw1" {
  name         = "host1"
  comment      = var.resource_comment
  ipv4_network = "192.168.104.0/24"
}

resource "smc_network" "host2" {
  name         = "host2"
  comment      = var.resource_comment
  ipv4_network = "192.168.150.0/24"
}

resource "smc_vpn_site" "site_fw1" {
  from_ref     = format("%s#%s", smc_single_fw.tf_single_fw["fw1"].link.internal_gateway, "*/vpn_site")
  name         = "site_fw1"
  site_element = [smc_network.nw1.id]
  comment      = var.resource_comment
}

resource "smc_external_gateway" "external_gateway" {
  gateway_profile                 = data.smc_href.gateway_profile.id
  name                            = "tf_external_gateway"
  trust_all_cas                   = true
  trusted_certificate_authorities = []
}

resource "smc_vpn_site" "site_external" {
  from_ref     = smc_external_gateway.external_gateway.link.vpn_site
  name         = "external_site"
  site_element = [smc_network.host2.id]
}

resource "smc_external_endpoint" "ep_external" {
  from_ref  = smc_external_gateway.external_gateway.link.external_endpoint
  enabled   = true
  address   = "10.25.25.15"
  ipsec_vpn = true
}
