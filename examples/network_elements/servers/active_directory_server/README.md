# TODO: fix me or update smc ? 

```
smc_active_directory_server.ad: Creating...
╷
│ Error: Error creating resource
│ 
│   with smc_active_directory_server.ad,
│   on main.tf line 6, in resource "smc_active_directory_server" "ad":
│    6: resource "smc_active_directory_server" "ad" {
│ 
│ Could not create active_directory_server: create active_directory_server failed with status 400: {"details":["Invalid JSON format: At line 1 and column 547, long_user_id_attr is not recognized as JSON
│ attribute."],"status":0}
╵
make: *** [Makefile:17: apply] Error 1
```
