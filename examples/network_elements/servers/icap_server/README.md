# TODO: fix me

Swagger json is not correct

```
commands will detect it and remind you to do so if necessary.
terraform apply -auto-approve
╷
│ Error: Invalid resource type
│
│   on main.tf line 1, in resource "smc_icap" "icap":
│    1: resource "smc_icap" "icap" {
│
│ The provider forcepoint/fp-ngfw-smc does not support resource type "smc_icap".
╵
make: *** [Makefile:17: apply] Error 1

```
