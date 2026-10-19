# Cloud NGFW site-to-site VPN (demo-mode example)

A self-contained example that builds a two-site NGFW topology entirely out
of **SMC elements** — engines, a firewall policy, and a policy-based IPsec
site-to-site VPN — so it can be applied against an **SMC running in demo
mode** without any real cloud infrastructure.

It is modelled on the real `gcp-ngfw-ha` deployment, but every value that
would normally come from AWS/GCP (public IPs, subnet CIDRs, DNS) is **mocked
as a static variable** in `variables.tf`. Nothing here calls a cloud API.

## What it creates

Two firewall "sites", A and B, each an `smc_single_fw` with three static
interfaces (mgmt / WAN / LAN):

| File                   | Elements                                                            |
| ---------------------- | ------------------------------------------------------------------- |
| `smc_engines.tf`       | The two engines (`cloud-a-ngfw`, `cloud-b-ngfw`)                     |
| `smc_policy.tf`        | LAN networks + shared north-south policy (site-to-site / internet / ssh / discard) |
| `smc_vpn.tf`           | VPN element, contact addresses, IPsec endpoints, gateway nodes, VPN sites |
| `smc_routing.tf`       | WAN networks, gateway routers, peer host routes                     |
| `smc_actions.tf`       | Per-engine license-bind + initial-contact actions                   |
| `smc_policy_upload.tf` | Per-engine policy install action                                    |
| `main.tf`              | The `sites` map (single source of truth) + shared data sources      |

Site A is the VPN **central** gateway, site B the **satellite**.

## Lifecycle actions

Applying the example also runs the standard engine bring-up actions, one per
engine:

1. **License bind** — `system/license_bind` for each node (`smc_actions.tf`).
2. **Initial contact** — writes each node's one-time-password config to
   `${initial_contact_dir}/<engine>_initial_contact.txt` (`smc_actions.tf`).
   Steps 1–2 fire from each engine's `after_create` trigger.
3. **Policy upload** — the `smc_engine_upload_policy` action installs the
   shared policy on each engine, filtered by engine name (`smc_policy_upload.tf`).
   It runs last, driven by per-engine `terraform_data` sync-point resources
   whose `depends_on` covers every policy, VPN and routing element (so the
   pushed policy reflects the full config) and whose `triggers_replace`
   re-pushes when the engine, policy, rules or routing change.

Individual actions can be re-invoked without a full apply, e.g.:

```bash
terraform apply -invoke 'action.smc_initial_contact.contact["a"]' -auto-approve
terraform apply -invoke 'action.smc_engine_upload_policy.upload["b"]' -auto-approve
```

## Run it against demo mode

SMC demo mode serves the API over HTTP and ships with a fixed API key.
Point the provider at your demo instance:

```bash
export TF_VAR_url="http://localhost:8082"          # demo-mode API URL
export TF_VAR_api_key="HuphG4Uwg4dN6TyvorTR0001"   # demo-mode API key
export TF_VAR_api_version="7.6"                     # match your SMC version
```

Then, from this directory, using the bundled Makefile (which uses the
locally-built provider):

```bash
make clean init-dev   # initialize with the local plugin
make plan             # review
make apply            # create everything
make destroy          # tear it all down
```

Or with the released provider from the registry:

```bash
make clean init
make apply
```

## Customising

Override any input in `variables.tf`, e.g. to change the mocked public
endpoints or the subnet layout:

```bash
export TF_VAR_site_a_public_ip="203.0.113.20"
export TF_VAR_site_b_lan_cidr="10.0.5.0/24"
```
