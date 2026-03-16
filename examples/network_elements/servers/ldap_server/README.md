# TODO: fix me or update smc ? 

```
smc_ldap_server.ldap: Creating...
╷
│ Error: Error creating resource
│
│   with smc_ldap_server.ldap,
│   on main.tf line 1, in resource "smc_ldap_server" "ldap":
│    1: resource "smc_ldap_server" "ldap" {
│
│ Could not create ldap_server: create ldap_server failed with status 400: {"details":["Invalid JSON format: At line 1 and column 681, short_user_id_attr is not recognized as JSON attribute."],"status":0}
╵
make: *** [Makefile:17: apply] Error 1
```
