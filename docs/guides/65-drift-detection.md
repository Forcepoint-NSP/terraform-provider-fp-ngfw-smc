---
page_title: "Drift Detection"
subcategory: ""
description: |-
  How Terraform detects changes made outside of Terraform, how to enable it, and its limitations
---
# Drift Detection

"Drift" means someone (or something) changed a resource in SMC outside
of Terraform — for example by editing it directly in the SMC
Management Client. By default, Terraform has no way of knowing about
such a change, since `terraform plan` normally only compares your
`.tf` configuration against the last known state, not against SMC
itself.

Drift detection makes `terraform plan` read the current object from
SMC and compare it against your configuration, so out-of-band changes
to attributes you manage in Terraform show up as a plan diff, the same
way a change to your `.tf` file would.

## Enabling Drift Detection

Drift detection is disabled by default, and can be enabled per
resource type or globally.

### Per-resource, via a config file

Create a `tfsmc.conf.json` file in the directory where you run Terraform:

```json
{
  "drift_detection": ["smc_host", "smc_single_fw"]
}
```

This enables drift detection only for `smc_host` and `smc_single_fw` resources. To enable it for every resource type instead, use `true`:

```json
{
  "drift_detection": true
}
```

### Globally, via an environment variable

You can also enable (or disable) drift detection for every resource type without a config file, by setting:

```bash
export TFSMC_DRIFT_DETECTION=1   # enable for all resources
```

If set, this environment variable takes precedence over
`tfsmc.conf.json` — it's a convenient way to try drift detection out,
or to override a config file, without editing it.

## What Gets Checked

Once enabled for a resource type, `terraform plan` will show a diff
whenever an attribute **you have explicitly set in your
configuration** no longer matches SMC's current value.

for example, if a `comment` you set in your `.tf` file was changed
through the SMC GUI, the next `terraform plan` will show that change.

## Limitation: Attributes You Never Set Are Not Monitored

Drift detection only watches attributes you've actually given a value
to in your Terraform configuration. If you leave an attribute unset,
Terraform won't notice or report any changes SMC makes to it
afterwards — including a value set for the very first time through the
SMC GUI.

**This is expected, not a bug, and it's not something you need to work
around.** 

If you never gave Terraform an opinion about an attribute, Terraform
correctly has nothing to compare against — and importantly, **the
remote change is not lost or overwritten**. 

SMC keeps whatever value was set, whether by you, another tool, or the
GUI; Terraform simply doesn't manage or report on it. The moment you
do add that attribute to your configuration, Terraform starts tracking
it from then on, including detecting future drift on it.

__The practical rule of thumb__: **only attributes present in your `.tf`
files are Terraform's responsibility.** Anything you leave out is left
entirely to SMC, exactly as if drift detection didn't exist.

## A Note on Nested Lists

For resources with a nested list of sub-objects
(e.g. `physical_interface`), drift detection identifies list entries
by their identity (such as an interface ID), not by position, so
reordering or adding entries out-of-band is generally handled
cleanly. 

Removing an entry from the middle of such a list, however, can make
later entries show up as "changed" in the plan even though only one
entry actually disappeared — Terraform diffs these lists position by
position, and there is no way to represent "one entry removed" other
than shifting everything after it. 

If you see this, check whether an item further down the list was
actually removed rather than modified.

## Terraform Configuration Is the Source of Truth

Drift detection does not mean Terraform adopts SMC's value. For any
attribute you manage, **your `.tf` configuration is always the source
of truth** — detecting drift only means Terraform noticed SMC
disagrees with it, not that Terraform will follow SMC's lead.

This shows up directly in the `terraform plan` output. If you set
`comment = "Created by Terraform"` and someone changes it in the SMC
GUI to `"Created by someone"`, the plan looks like this:

```
~ comment = "Created by someone" -> "Created by Terraform"
```

__Read that carefully__: 

the left-hand side is what SMC currently has (the drift), and the
right-hand side — your configured value — is what `terraform apply`
will **push back to SMC, discarding the GUI change**. 

The plan is not proposing to keep the remote edit; it's telling you
Terraform is about to undo it.

**Reconciling the two is entirely up to you**, and Terraform will not
decide for you:

- Run `terraform apply` to enforce your configuration and discard the
  out-of-band change, or
- Edit your `.tf` file to match SMC's current value first, then apply
  (a no-op, since they'll now agree), if you actually want to keep the
  change made outside Terraform.

Always read a drift-related plan before applying it — it's easy to
assume Terraform will "sync up" with whatever SMC has, but for
anything Terraform manages, applying always moves SMC toward your
configuration, never the other way around.
