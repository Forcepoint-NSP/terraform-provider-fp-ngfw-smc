# Shared data sources and the per-site topology map.
#
# `sites` is the single source of truth for the two firewalls. Every
# resource that is symmetric across the two sites (engine, networks,
# routers, VPN endpoints, routing) iterates this map with for_each, keyed
# "a" and "b". The asymmetric parts of the VPN (which gateway is central
# vs. satellite) are declared explicitly in smc_vpn.tf.

locals {
  sites = {
    a = {
      name      = "cloud-a-ngfw"
      mgmt_cidr = var.site_a_mgmt_cidr
      wan_cidr  = var.site_a_wan_cidr
      lan_cidr  = var.site_a_lan_cidr
      public_ip = var.site_a_public_ip
      peer      = "b"
    }
    b = {
      name      = "cloud-b-ngfw"
      mgmt_cidr = var.site_b_mgmt_cidr
      wan_cidr  = var.site_b_wan_cidr
      lan_cidr  = var.site_b_lan_cidr
      public_ip = var.site_b_public_ip
      peer      = "a"
    }
  }

  # Interface addresses are recomputed from each site's CIDR layout, the
  # same cidrhost() convention the reference cloud stacks use.
  #  nic0 (mgmt) -> .14   nic1 (wan) -> .101   nic2 (lan) -> .1
  mgmt_ip   = { for k, s in local.sites : k => cidrhost(s.mgmt_cidr, 14) }
  wan_ip    = { for k, s in local.sites : k => cidrhost(s.wan_cidr, 101) }
  lan_ip    = { for k, s in local.sites : k => cidrhost(s.lan_cidr, 1) }
  wan_gw_ip = { for k, s in local.sites : k => cidrhost(s.wan_cidr, 1) }

  # JMESPath used to resolve an interface's contact-address link by the
  # single_node_interface address (see smc_vpn.tf).
  contact_addr_match = "%s/interfaces[?single_node_interface.address=='%s']|[0].single_node_interface.link[0].href"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "default_location" {
  type = "location"
  name = "Default"
}

data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

data "smc_href" "suite_b_gcm_128" {
  type = "vpn_profile"
  name = "Suite-B-GCM-128"
}

data "smc_href" "service_ssh" {
  type = "tcp_service"
  name = "SSH"
}
