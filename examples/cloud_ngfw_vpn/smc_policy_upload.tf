# Policy install actions — one per engine. Fired during the initial apply
# via the sync-point resources below, which depend on every SMC element so
# the upload runs last. wait = true: the apply waits for the upload to
# complete, which is what reproduces the "VPN has no valid tunnel" error
# this example was built to exercise (see repo commit history).
#
# Modelled directly on the gcp-ngfw-ha smc_policy_upload.tf. Upload is driven
# from the policy, not the engine: the action POSTs to the policy's upload
# link (`href`) with the engine name as the `filter` query parameter
# (`policy_name` is sent verbatim as `filter`). This mirrors smc-python's
# FirewallPolicy.upload(engine) (params={"filter": engine}), the documented
# path for a single_fw.
#
# The reference uses null_resource as the sync point; here the built-in
# terraform_data resource plays that role, so no extra provider is needed
# under `terraform init -plugin-dir`. Its triggers_replace re-runs the
# after_create upload whenever the engine, policy, rules or routing change —
# action_trigger only fires on its listed events, so without this an
# already-created sync point would never re-push after an edit.

action "smc_engine_upload_policy" "upload" {
  for_each = local.sites
  config {
    href        = smc_fw_policy.north_south.link.upload
    policy_name = smc_single_fw.ngfw[each.key].name
    wait        = true
  }
}

# Synchronization points — one per engine, because action_trigger.actions
# only accepts a static list literal. Each resource's depends_on covers every
# SMC element that must be settled before the policy push.
resource "terraform_data" "upload" {
  for_each = local.sites

  # Re-push when the engine, the policy, any rule, or the routing changes.
  triggers_replace = {
    engine = smc_single_fw.ngfw[each.key].name
    policy = smc_fw_policy.north_south.id
    rules = join(",", [
      smc_fw_ipv4_access_rule.site_to_site.id,
      smc_fw_ipv4_access_rule.allow_internet.id,
      smc_fw_ipv4_access_rule.allow_ssh.id,
      smc_fw_ipv4_access_rule.discard_all.id,
    ])
    routing = smc_routing_node.wan[each.key].id
  }

  depends_on = [
    smc_single_fw.ngfw,
    smc_fw_policy.north_south,
    smc_fw_ipv4_access_rule.site_to_site,
    smc_fw_ipv4_access_rule.allow_internet,
    smc_fw_ipv4_access_rule.allow_ssh,
    smc_fw_ipv4_access_rule.discard_all,
    smc_vpn.site_to_site,
    smc_gateway_node.central,
    smc_gateway_node.satellite,
    smc_vpn_site.site,
    smc_internal_endpoint.wan,
    smc_contact_address.wan,
    # Static routes must be in the engine config before it is uploaded.
    smc_routing_node.wan,
  ]

  lifecycle {
    action_trigger {
      events  = [after_create]
      actions = [action.smc_engine_upload_policy.upload[each.key].id]
    }
  }
}
