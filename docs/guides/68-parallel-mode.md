---
page_title: "Parallel Mode"
subcategory: ""
description: |-
  How to enable parallel request mode, manage resource exclusions, and handle concurrent update conflicts
---
# Parallel Mode

By default, terraform-provider-smc creates resources one at a time (serially). 
For workloads with many resources, parallel mode lets Terraform create 
multiple resources concurrently, reducing total apply time — see 
[Performance Impact](#performance-impact) below for typical speedups.

Parallel mode is **opt-in** because certain SMC resource types or specific 
configurations may experience concurrent update conflicts. This guide shows 
you how to enable it, identify problematic resources, and add them to an 
exclusion list if needed.

## Enabling Parallel Mode

Parallel mode is disabled by default and can be enabled globally, either via 
an environment variable or via a config file.

### Globally, via an environment variable

Set the environment variable before running Terraform:

```bash
export TFSMC_PARALLEL_CREATE=1   # enable parallel mode (1/0/true/false/yes/no/on/off)
terraform apply
```

If set, this environment variable takes precedence over `tfsmc.conf.json` — 
it's a convenient way to try parallel mode out, or to override a config file, 
without editing it.

### Via a config file

Create a `tfsmc.conf.json` file in the directory where you run Terraform. 
Unlike the environment variable, this fully enables parallel mode on its own 
— no environment variable required:

```json
{
  "parallel_mode": true
}
```

To exclude specific resource types and keep them serial, add a 
`parallel_mode_exclusion_list`:

```json
{
  "parallel_mode": true,
  "parallel_mode_exclusion_list": ["smc_routing_rule", "smc_fw_cluster"]
}
```

`parallel_mode_exclusion_list` only applies once parallel mode is enabled 
(by the config's `parallel_mode: true` or by the environment variable); it 
has no effect on its own. Some resource types are excluded unconditionally, 
regardless of this list — see [Always-Serial Resource Types](#always-serial-resource-types) 
below.

## Always-Serial Resource Types

The following resource types are **hardcoded to always run serially**, 
whatever `parallel_mode_exclusion_list` says (even `[]`) and whether parallel 
mode is on or off. This is not configurable — see 
`alwaysSerializedResourceTypes` in `internal/smc/parallel.go` for the 
authoritative, up-to-date list:

- Rank/position-ordered rules (SMC holds an exclusive per-policy write lock 
  on these, so concurrent creates fail with a 409/422): `smc_fw_ipv4_access_rule`, 
  `smc_fw_ipv4_nat_rule`, `smc_fw_ipv6_access_rule`, `smc_fw_ipv6_nat_rule`, 
  `smc_ips_ipv6_access_rule`, `smc_ips_ethernet_rule`, 
  `smc_l2_interface_ipv4_access_rule`, `smc_l2_interface_ipv6_access_rule`, 
  `smc_l2_interface_ethernet_rule`, `smc_layer2_ipv4_access_rule`, 
  `smc_layer2_ipv6_access_rule`, `smc_layer2_ethernet_rule`, 
  `smc_inspection_exception_rule`, `smc_inspection_global_rule`, 
  `smc_alert_chain_rule`, `smc_file_filtering_rule`, `smc_limit_rule`, 
  `smc_route_map_rule`, `smc_rule_validity_time`
- Firewall and VPN resources with complex bring-up dependencies: 
  `smc_single_fw`, `smc_fw_cluster`, `smc_vpn`, `smc_rbvpn_tunnel`, 
  `smc_rbvpn_tunnel_group`, `smc_gateway_node`, `smc_internal_endpoint`, 
  `smc_vpn_site`

## Concurrent Update Conflicts

Parallel mode runs multiple create operations simultaneously. The SMC API 
generally handles this well, but some operations or specific configurations may 
experience conflicts:

- **Firewall rules under one policy**: SMC takes an exclusive per-policy write 
  lock on rule creation and does not queue concurrent writers — this is part 
  of why rule types are in the [Always-Serial Resource Types](#always-serial-resource-types) 
  list above.

- **Custom fields and dependencies**: Resources with complex inter-dependencies 
  or custom field constraints may fail under concurrency.

- **SMC Management Server capacity**: this isn't a specific rate-limit setting 
  you configure or disable — it's the practical limit of your SMC Management 
  Server's own CPU, memory, and database connections. High parallelism on a 
  resource-constrained SMC instance can exhaust that capacity and cause 
  requests to fail or slow down, regardless of where the Terraform client 
  runs (including from a private network). `terraform apply --parallelism` 
  (recommended: 20, matching SMC's default DB connection limit) is the knob 
  that controls how much concurrent load your client generates against it.

On a `409 Conflict` or a `422` (e.g. "locked by..."), the provider 
automatically retries the create up to 4 times, waiting 5 seconds between 
attempts (about 20 seconds in total), before giving up. A genuine 
duplicate-name conflict cannot be told apart from a transient one, so in 
parallel mode it also fails only after the retries are exhausted, with a 
generic "after 4 conflict retries" error instead of an immediate one; check the 
wrapped SMC error for the duplicate name. If you still see conflicts after 
retries are exhausted:

1. **Identify the problematic resource type** from the error message
2. **Add it to the exclusion list** in `tfsmc.conf.json`
3. **Re-run apply** — the excluded type will use serial mode while others 
   remain parallel

## Managing the Exclusion List

### Starting point

The most common exclusions are already documented above. Start with an empty 
list and add types only if you encounter conflicts:

```json
{
  "parallel_mode": true,
  "parallel_mode_exclusion_list": []
}
```

### Troubleshooting

If an apply fails with a concurrent update error:

1. Note the resource type (e.g. `smc_routing_rule`)
2. Add it to the exclusion list:
   ```json
   {
     "parallel_mode": true,
     "parallel_mode_exclusion_list": ["smc_routing_rule"]
   }
   ```
3. Run `terraform apply` again — that type will now use serial mode
4. Once stable, you can try removing it and re-running if you want to 
   validate whether the issue was transient or configuration-specific

### Examples

Exclude multiple types:
```json
{
  "parallel_mode": true,
  "parallel_mode_exclusion_list": ["smc_routing_rule", "smc_custom_field_type"]
}
```

Enable parallel mode with no *additional* exclusions (the 
[always-serial types](#always-serial-resource-types) still apply — this list 
is empty by default):
```json
{
  "parallel_mode": true
}
```

Disable parallel mode (serial only):
```json
{
  "parallel_mode": false
}
```
(Or simply omit `parallel_mode` and unset `TFSMC_PARALLEL_CREATE`.)

## Performance Impact

Parallel mode is most effective when:
- Your configuration has **many resources** (50+) of independent types
- Your **SMC instance** has sufficient CPU, memory, and database connection 
  headroom to handle concurrent requests — see 
  [SMC Management Server capacity](#concurrent-update-conflicts) above

Typical speedup: **2-5x** for large workloads, depending on resource type 
mix and SMC instance capacity. Small configurations (< 20 resources) see 
little benefit, as the overhead of parallelism can dominate.
