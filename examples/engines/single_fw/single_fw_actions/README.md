# Single Firewall Actions Example

This example demonstrates how to deploy a single firewall engine (single_fw) and automate a range of management actions on it using the SMC Terraform provider. It provides practical patterns for configuring, operating, and maintaining a Forcepoint NGFW appliance through Terraform—covering both resource provisioning and advanced post-deploy operations.

## What this example does

- **Deploys a single firewall:**
  - Creates a single firewall engine with a minimal node and management interface.
  - Assigns it to a monitored group.

- **Configures firewall policy:**
  - Uses a template to create a policy (`tf_aaa`), then adds an allow-all IPv4 rule.
  - Associates the policy to the deployed firewall.

- **Automates operational actions:**
  - **Initial contact:** Automates Firewall initial contact (SSH enablement, timezone/keyboard selection, root password setup, install to server).
  - **License binding:** Runs a command to bind the firewall license after creation.
  - **Upload or refresh policy/configuration:** Triggers upload or refresh operations via lifecycle hooks or on-demand via CLI.
  - **Dynamic state changes:** Changes the firewall node state (e.g. go_online, go_offline, reboot, power_off, etc.) based on a variable.
  - **Diagnostics and support:**
    - Generates and downloads an sginfo diagnostics package on demand.
    - Manages backup tasks and triggers backups programmatically.
    - Restarts web access service and demonstrates admin domain refresh.

## Usage Notes

- **Lifecycle action triggers** are used to invoke key actions (like license bind, initial contact, upload) immediately after resource creation.
- **On-demand action invocation** is supported via Terraform `action` blocks and CLI flags (see code comments for examples).
- **Custom variables** allow you to adjust operational state and other on-demand actions (e.g., `var.new_state`, `var.refresh`).
- Example covers backup management for management servers and demonstrates interaction with multiple SMC objects (log server, management server, admin domain).

## How to use this example

1. **Initialize and plan as usual**:
   ```sh
   terraform init
   terraform plan
   ```
2. **Apply normally to provision resources and trigger actions:**
   ```sh
   terraform apply
   ```
3. **Trigger ad-hoc actions with Terraform CLI:**
   - Run initial contact again:
     ```sh
     terraform apply -invoke action.smc_initial_contact.tf_single_fw_contact -auto-approve
     ```
   - Upload/refresh the policy configuration:
     ```sh
     terraform apply -invoke action.smc_command.upload -auto-approve
     terraform apply -invoke action.smc_command.upload -auto-approve -var refresh=true
     ```
   - Change firewall state:
     ```sh
     terraform apply -invoke action.smc_command.change_state -var new_state=go_online -auto-approve
     ```
   - Gather diagnostics (`sginfo`):
     ```sh
     terraform apply -invoke action.smc_command.sginfo -var sginfo_file="./sginfo_output.zip" -auto-approve
     ```
   - Run backup:
     ```sh
     terraform apply -invoke action.smc_command.backup_now -auto-approve
     ```
   - Restart web access service:
     ```sh
     terraform apply -invoke action.smc_command.restart_web_access -auto-approve
     ```
   - Refresh admin domain policies:
     ```sh
     terraform apply -invoke action.smc_command.refresh_admin_domain_policy -auto-approve
     ```
