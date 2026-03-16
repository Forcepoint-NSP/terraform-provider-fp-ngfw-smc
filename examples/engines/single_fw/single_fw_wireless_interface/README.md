# TODO: fix me

Current failure

```
Plan: 1 to add, 0 to change, 0 to destroy.
smc_single_fw.tf_single_fw: Creating...
╷
│ Error: Provider returned invalid result object after apply
│ 
│ After the apply operation, the provider still indicated an unknown value for smc_single_fw.tf_single_fw.physical_interfaces[1].wireless_interface.ssid_interfaces[0].interfaces[1].single_node_interface.key. All values
│ must be known after apply, so this is always a bug in the provider and should be reported in the provider's own repository. Terraform will still save the other known object values in the state.
╵
╷
│ Error: Provider returned invalid result object after apply
│ 
│ After the apply operation, the provider still indicated an unknown value for smc_single_fw.tf_single_fw.physical_interfaces[1].wireless_interface.ssid_interfaces[0].interfaces[1].single_node_interface.link. All
│ values must be known after apply, so this is always a bug in the provider and should be reported in the provider's own repository. Terraform will still save the other known object values in the state.
╵
╷
│ Error: Provider returned invalid result object after apply
│ 
│ After the apply operation, the provider still indicated an unknown value for smc_single_fw.tf_single_fw.physical_interfaces[1].wireless_interface.ssid_interfaces[0].interfaces[1].single_node_interface.lk. All values
│ must be known after apply, so this is always a bug in the provider and should be reported in the provider's own repository. Terraform will still save the other known object values in the state.
╵
make: *** [Makefile:17: apply] Error 1
```