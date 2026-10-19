include ../../../makefile.common

# Path to the dev-overrides RC written by init-dev.
DEV_RC := $(CURDIR)/.dev.terraformrc

# Prepended to every apply/destroy/invoke so dev_overrides take effect after init-dev.
TF_ENV := $(if $(wildcard $(DEV_RC)),TF_CLI_CONFIG_FILE=$(DEV_RC),)

# Override init-dev: phase 1 uses filesystem_mirror so Terraform can find the
# local fp-ngfw-smc binary while fetching hashicorp/null from the registry.
# Phase 2 writes dev_overrides so subsequent applies pick up provider rebuilds.
init-dev:
	-@rm -rf .terraform .terraform.lock.hcl $(DEV_RC)
	@PROVIDER_DIR="$(PLUGIN_DIR)/registry.terraform.io/forcepoint-nsp/fp-ngfw-smc/$(PROVIDER_VERSION)/$(OS_ARCH)"; \
	ln -sfn "$$PROVIDER_DIR/terraform-provider-fp-ngfw-smc" \
	        "$$PROVIDER_DIR/terraform-provider-fp-ngfw-smc_v$(PROVIDER_VERSION)" 2>/dev/null || true; \
	if [ "$(TF_CLI)" = "tofu" ]; then \
		OTF_DIR="$(PLUGIN_DIR)/registry.opentofu.org/forcepoint-nsp/fp-ngfw-smc/$(PROVIDER_VERSION)/$(OS_ARCH)"; \
		mkdir -p "$$OTF_DIR"; \
		ln -sfn "$$PROVIDER_DIR/terraform-provider-fp-ngfw-smc" "$$OTF_DIR/terraform-provider-fp-ngfw-smc"; \
		ln -sfn "$$PROVIDER_DIR/terraform-provider-fp-ngfw-smc" "$$OTF_DIR/terraform-provider-fp-ngfw-smc_v$(PROVIDER_VERSION)"; \
	fi; \
	printf 'provider_installation {\n  filesystem_mirror {\n    path    = "%s"\n    include = ["registry.terraform.io/forcepoint-nsp/fp-ngfw-smc"]\n  }\n  direct {\n    exclude = ["registry.terraform.io/forcepoint-nsp/fp-ngfw-smc"]\n  }\n}\n' \
	    "$(PLUGIN_DIR)" > $(DEV_RC)
	TF_CLI_CONFIG_FILE=$(DEV_RC) $(TF_CLI) init $(TF_OPTS)
	@PROVIDER_DIR="$(PLUGIN_DIR)/registry.terraform.io/forcepoint-nsp/fp-ngfw-smc/$(PROVIDER_VERSION)/$(OS_ARCH)"; \
	printf 'provider_installation {\n  dev_overrides {\n    "registry.terraform.io/forcepoint-nsp/fp-ngfw-smc" = "%s"\n  }\n  direct {}\n}\n' \
	    "$$PROVIDER_DIR" > $(DEV_RC)

# Override apply/destroy to carry TF_CLI_CONFIG_FILE when dev_overrides are active.
# Keep the ensure-ca / ca cleanup hooks from makefile.common in sync here.
apply:
	$(call check-tool-support)
	$(call ensure-ca)
	$(TF_ENV) $(TF_CLI) apply -auto-approve $(TF_OPTS)

destroy:
	$(call check-tool-support)
	$(TF_ENV) $(TF_CLI) destroy -auto-approve $(TF_OPTS)
	@if [ -x ./ensure_ca.sh ]; then rm -rf ./ca; fi

# Step 2 — generate the CSR on the engine and export it.
generate-csr:
	$(TF_ENV) $(TF_CLI) apply -invoke action.smc_generate_vpn_certificate.generate_csr \
	    -auto-approve $(TF_OPTS)

# Step 3+4 — sign the CSR with the local test CA and import the result.
# Replacing wait_for_csr cascades to sign_cert (replace_triggered_by), which fires the import via action_trigger.
sign-cert:
	$(TF_ENV) $(TF_CLI) apply -replace=null_resource.wait_for_csr -auto-approve $(TF_OPTS)

# On initial apply the full certificate flow runs automatically (EP action_trigger → generate_csr →
# wait_for_csr → sign_cert → import). sign-cert regenerates the certificate on demand.
test-certificate: sign-cert

# Override test-update-comment to inject the certificate test into test-lifecycle.
test-update-comment: test-certificate
	TF_VAR_resource_comment="Updated by Terraform" $(TF_ENV) $(TF_CLI) apply -auto-approve $(TF_OPTS)

.PHONY: init-dev apply destroy generate-csr sign-cert test-certificate
