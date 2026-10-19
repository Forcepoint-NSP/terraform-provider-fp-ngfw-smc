# Per-engine create-time actions: license bind + initial contact.
# Both are triggered from each engine's after_create lifecycle block
# (smc_engines.tf). One instance per site, keyed by the local.sites map.

variable "root_password" {
  type        = string
  description = "Root password set on each engine node during initial contact"
  default     = "root"
  sensitive   = true
}

variable "time_zone" {
  type    = string
  default = "CET"
}

variable "keyboard" {
  type    = string
  default = "us"
}

variable "initial_contact_dir" {
  type        = string
  description = "Local directory where each node's initial-contact OTP file is written"
  default     = "/tmp"
}

# Bind a license to each engine's node. Mirrors the getting_started /
# single_fw_actions examples: POST system/license_bind with the node self
# href as component_href.
action "smc_command" "bind_license" {
  for_each = local.sites
  config {
    command_href = "${var.url}/${var.api_version}/system/license_bind"
    parameters = {
      component_href = smc_single_fw.ngfw[each.key].nodes[0].firewall_node.link.self
    }
  }
}

# Run initial contact on each node, writing the base64 OTP config locally.
# To re-run for a single site:
#   terraform apply -invoke 'action.smc_initial_contact.contact["a"]' -auto-approve
action "smc_initial_contact" "contact" {
  for_each = local.sites
  config {
    initial_contact_href = smc_single_fw.ngfw[each.key].nodes[0].firewall_node.link.initial_contact
    enable_ssh           = true
    time_zone            = var.time_zone
    keyboard             = var.keyboard
    root_password        = var.root_password
    install_on_server    = true
    export_to_base64     = false
    output_file          = "${var.initial_contact_dir}/${each.value.name}_initial_contact.txt"
  }
}
