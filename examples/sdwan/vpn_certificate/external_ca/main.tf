/*
  Generate and sign a VPN certificate for a firewall's internal gateway
  using an external (test) certificate authority.

  A single apply performs the full lifecycle:
    1. Create the firewall and VPN endpoint.
    2. Trigger: bind license, run initial contact, upload policy.
    3. Trigger (on VPN endpoint creation): generate CSR.
    4. Wait for the CSR file to appear, sign it with the local test CA,
       then import the signed certificate.

  To regenerate the certificate:
    make sign-cert   (replaces wait_for_csr which cascades to sign_cert)
*/

variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

locals {
  fw_name           = "tf_vpn_cert_ext_fw"
  fw_address        = "192.168.10.1"
  fw_network        = "192.168.10.0/24"
  tf_single_fw_name = local.fw_name
  tf_fw_ref         = smc_single_fw.fw
  tf_node_ref       = local.tf_fw_ref.nodes[0].firewall_node
  csr_file          = "/tmp/${local.fw_name}_csr.pem"
  signed_file       = "/tmp/${local.fw_name}_signed.pem"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

resource "smc_fw_policy" "policy" {
  name     = "tf_aaa_ext"
  template = data.smc_href.default_policy_template.id
  comment  = var.resource_comment
}

resource "smc_fw_ipv4_access_rule" "allow_all" {
  from_ref = smc_fw_policy.policy.link.fw_ipv4_access_rules
  name     = "aaa"
  sources { any = true }
  destinations { any = true }
  services { any = true }
  action { action = ["allow"] }
}

resource "smc_single_fw" "fw" {
  name           = local.fw_name
  comment        = var.resource_comment
  log_server_ref = data.smc_href.log_server.href

  # The imported gateway certificate (created under this firewall's internal
  # gateway) references the test CA. Depending on the CA reverses the destroy
  # order: the firewall — and with it the certificate — is deleted first, so
  # the CA is no longer referenced when its turn comes.
  depends_on = [smc_vpn_certificate_authority.test_ca]

  nodes {
    firewall_node {
      name   = "${local.fw_name}_node1"
      nodeid = 1
    }
  }

  # Interface 0 — management
  physical_interfaces {
    physical_interface {
      interface_id = 0
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 0
          address       = "192.168.1.1"
          network_value = "192.168.1.0/24"
          primary_mgt   = true
        }
      }
    }
  }

  # Interface 1 — VPN endpoint
  physical_interfaces {
    physical_interface {
      interface_id = 1
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 1
          address       = local.fw_address
          network_value = local.fw_network
        }
      }
    }
  }

  lifecycle {
    action_trigger {
      events = [after_create]
      actions = [
        action.smc_command.bind_license.id,
        action.smc_initial_contact.contact.id,
        action.smc_engine_upload_policy.upload.id,
      ]
    }
  }
}

action "smc_command" "bind_license" {
  config {
    command_href = "${var.url}/${var.api_version}/system/license_bind"
    parameters = {
      component_href = local.tf_node_ref.link.self
    }
  }
}

action "smc_initial_contact" "contact" {
  config {
    initial_contact_href = local.tf_node_ref.link.initial_contact
    enable_ssh           = true
    time_zone            = "CET"
    keyboard             = "us"
    root_password        = "root"
    install_on_server    = true
    export_to_base64     = false
    output_file          = "/tmp/${local.tf_single_fw_name}_initial_contact.txt"
  }
}

action "smc_engine_upload_policy" "upload" {
  config {
    href        = local.tf_fw_ref.link.upload
    policy_name = smc_fw_policy.policy.name
    wait        = true
  }
}

# Enable the VPN endpoint on interface 1 and automatically generate the CSR.
resource "smc_internal_endpoint" "vpn_ep" {
  id        = "${smc_single_fw.fw.link.internal_gateway}#*/internal_endpoint/${local.fw_address}"
  enabled   = true
  ipsec_vpn = true

  lifecycle {
    action_trigger {
      events  = [after_create]
      actions = [action.smc_generate_vpn_certificate.generate_csr.id]
    }
  }
}

# Resolve the internal gateway href (created implicitly with the firewall).
data "smc_sub_href" "internal_gateway" {
  from_ref = smc_single_fw.fw.link.internal_gateway
  match    = "${local.fw_name}*"
}

# Step 2 — generate a certificate request on the engine and export the CSR.
# The action resolves the gateway's generate_certificate link itself and
# automatically waits for the engine to be Online before proceeding.
action "smc_generate_vpn_certificate" "generate_csr" {
  config {
    gateway_href         = data.smc_sub_href.internal_gateway.id
    common_name          = "${local.fw_name} VPN certificate"
    organization         = "Forcepoint"
    public_key_algorithm = "rsa"
    signature_algorithm  = "rsa_sha_512"
    public_key_length    = 2048
    csr_output_file      = local.csr_file
  }
}

# Step 3a — wait for the CSR generated by the action_trigger on vpn_ep.
# A stale CSR from a previous run must not be signed, so instead of merely
# checking that the file exists, wait until the companion .href file (written
# by the generate action after the CSR export) references the CURRENT internal
# gateway — its element id changes every time the firewall is recreated.
# Replace this resource to re-trigger the sign+import chain: terraform apply -replace=null_resource.wait_for_csr
resource "null_resource" "wait_for_csr" {
  provisioner "local-exec" {
    command = "timeout 600 bash -c 'until [ -f ${local.csr_file} ] && grep -qs \"${data.smc_sub_href.internal_gateway.id}/\" ${local.csr_file}.href; do echo Waiting for CSR...; sleep 10; done' && echo CSR ready."
  }
}

# Step 3b — sign the CSR with the local test CA, then import via action_trigger.
resource "null_resource" "sign_cert" {
  depends_on = [null_resource.wait_for_csr, smc_vpn_certificate_authority.test_ca]

  provisioner "local-exec" {
    command = "./sign_cert.sh ${local.csr_file} ${local.signed_file}"
  }

  lifecycle {
    replace_triggered_by = [null_resource.wait_for_csr]
    action_trigger {
      events = [after_create]
      actions = [
        action.smc_import_vpn_certificate.import_certificate.id,
      ]
    }
  }
}

# Step 4a — register the test CA as a vpn_certificate_authority element in the
# SMC. certificate_import on a gateway_certificate_request builds the trust
# path against the VPN CA store, and the resulting gateway_certificate
# references the vpn_certificate_authority element as its signing CA — a
# tls_certificate_authority alone is not consulted and the import fails with
# "Impossible to find the signing certificate authority" (verified against the
# API: same CSR and signed certificate, 400 with only a TLS CA present, 201
# once a VPN CA holding the same certificate exists). The signed certificate
# must also carry the X509v3 Authority Key Identifier extension (sign_cert.sh).
resource "smc_vpn_certificate_authority" "test_ca" {
  name        = "tf_test_ext_ca"
  certificate = file("${path.module}/ca/ca.crt")
}

# Step 4b — import the signed certificate.
# The action navigates gateway → gateway_certificate_request → first item →
# certificate_import link internally, so no data source is needed.
action "smc_import_vpn_certificate" "import_certificate" {
  config {
    gateway_href  = data.smc_sub_href.internal_gateway.id
    cert_file     = local.signed_file
    csr_href_file = "${local.csr_file}.href"
  }
}
