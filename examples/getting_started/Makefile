TF_CLI ?= terraform

init:
	-@rm -rf .terraform*
	$(TF_CLI) init

plan:
	$(TF_CLI) plan

apply:
	$(TF_CLI) apply -auto-approve

destroy:
	$(TF_CLI) destroy -auto-approve

clean:
	rm -rf .terraform* terraform*
