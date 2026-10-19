# Topology inputs.
#
# This example is modelled on a real two-cloud NGFW deployment (see the
# gcp-ngfw-ha reference stack), but it is meant to run against an SMC in
# *demo mode* where no real cloud infrastructure exists. Everything that
# would normally come from AWS/GCP (public IP addresses, subnet CIDRs, the
# DNS resolver address) is therefore MOCKED here as a plain static value.
#
# Only SMC elements are created: two firewall engines, a shared firewall
# policy, a policy-based site-to-site VPN between them, and the supporting
# network/host/router/routing elements. Nothing here talks to a cloud API,
# so the defaults below are safe placeholders you can apply as-is.

# --- Mocked "public" WAN addresses --------------------------------------
# In a real deployment these are the cloud reserved/elastic public IPs of
# each firewall's WAN interface. Here they are static RFC 5737 documentation
# addresses (TEST-NET) so nothing collides with a routable network.

variable "site_a_public_ip" {
  type        = string
  description = "Mocked public IP of site A's WAN interface (VPN peer endpoint)"
  default     = "203.0.113.10"
}

variable "site_b_public_ip" {
  type        = string
  description = "Mocked public IP of site B's WAN interface (VPN peer endpoint)"
  default     = "198.51.100.10"
}

# --- Per-site subnet layout ---------------------------------------------
# mgmt = management interface (nic0), wan = internet-facing (nic1),
# lan = protected subnet behind the firewall (nic2).

variable "site_a_mgmt_cidr" {
  type    = string
  default = "10.1.0.0/24"
}

variable "site_a_wan_cidr" {
  type    = string
  default = "10.1.100.0/24"
}

variable "site_a_lan_cidr" {
  type    = string
  default = "10.1.1.0/24"
}

variable "site_b_mgmt_cidr" {
  type    = string
  default = "10.0.0.0/24"
}

variable "site_b_wan_cidr" {
  type    = string
  default = "10.0.100.0/24"
}

variable "site_b_lan_cidr" {
  type    = string
  default = "10.0.1.0/24"
}

# --- DNS resolver used by the engines -----------------------------------

variable "dns_server" {
  type    = string
  default = "8.8.8.8"
}

variable "resource_comment" {
  type    = string
  default = "Created by Terraform (cloud_ngfw_vpn)"
}
