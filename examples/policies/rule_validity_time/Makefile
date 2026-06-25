# -*- makefile -*-
REPO_DIR=$(shell git rev-parse --show-toplevel)
PLUGIN_DIR=$(REPO_DIR)/plugins
TF_OPTS ?= "-no-color"

init-dev:
	-@rm -rf .terraform*
	terraform init -plugin-dir $(PLUGIN_DIR) $(TF_OPTS)

init:
	-@rm -rf .terraform*
	terraform init $(TF_OPTS)

plan:
	terraform plan $(TF_OPTS)

apply:
	terraform apply -auto-approve $(TF_OPTS)

destroy:
	terraform destroy -auto-approve $(TF_OPTS)

clean:
	rm -rf .terraform* terraform*

.PHONY: update-src
update-src:
	patch main.tf <main.tf.update.diff

.PHONY: update
update: update-src apply

.PHONY: revert-src
revert-src:
	patch -R main.tf <main.tf.update.diff

.PHONY: revert
revert: revert-src apply

.PHONY: test-update-comment
test-update-comment:
	@echo "Testing update operation with comment change..."
	TF_VAR_resource_comment="Updated by Terraform" terraform apply -auto-approve $(TF_OPTS)

.PHONY: test-lifecycle
test-lifecycle:
	@echo "Starting full lifecycle test (init → apply → update → destroy)..."
	@$(MAKE) clean
	@$(MAKE) init-dev
	@$(MAKE) apply
	@echo "Waiting 3 seconds before update..."
	@sleep 3
	@$(MAKE) test-update-comment
	@echo "Waiting 3 seconds before destroy..."
	@sleep 3
	@$(MAKE) destroy
	@echo "Full lifecycle test complete"
