# initial_contact is intentionally kept as an action rather than a resource:
# an action can be re-invoked on demand (see below), which a resource cannot
# do as cleanly. The trade-off is that actions have no `depends_on` and no
# computed attributes, so anything downstream that needs the resulting
# config blob must be ordered via a resource attribute reference instead.
#
# terraform_data's `before_create` action_trigger guarantees the action has
# finished before the terraform_data resource itself is considered created -
# so data.local_file, which reads the proxy's `output` attribute, can only
# run after the blob has been written. A plain `depends_on` the firewall
# node would NOT give this guarantee: Terraform has no way to depend on an
# action's completion, only on a resource's.
action "smc_initial_contact" "contact" {
  config {
    initial_contact_href = smc_single_fw.example.nodes[0].firewall_node.link.initial_contact
    enable_ssh           = true
    time_zone            = "CET"
    keyboard             = "us"
    install_on_server    = true
    export_to_base64     = false
    output_file          = "/tmp/example_initial_contact.txt"
  }
}

resource "terraform_data" "contact_ready" {
  input = "/tmp/example_initial_contact.txt"
  triggers_replace = [
    smc_single_fw.example.nodes[0].firewall_node.link.initial_contact,
  ]
  depends_on = [smc_single_fw.example]

  lifecycle {
    action_trigger {
      events  = [before_create]
      actions = [action.smc_initial_contact.contact.id]
    }
  }
}

data "local_file" "contact" {
  filename = terraform_data.contact_ready.output
}

# To re-mint the contact package at any time (e.g. after a factory reset):
# terraform apply -invoke action.smc_initial_contact.contact -auto-approve
