---
page_title: "sync_parameter"
subcategory: ""
description: |-
  This represents the synchronization parameters for a cluster of engines, including sync mode, intervals, security level, and multicast IP addresses.
---

# sync_parameter (Nested-Attribute)

This represents the synchronization parameters for a cluster of engines, including sync mode, intervals, security level, and multicast IP addresses.




## Simple Attributes
- `full_sync_interval` (Number) The Full Synchronization Interface value (in milliseconds), which defines how frequently the full synchronizations are done. Do not set the values much higher or lower than their defaults (5000 ms for full). Caution! Adjusting the Sync Intervals has significant impact on the cluster's performance. Inappropriate settings seriously degrade the firewall's performance.
- `heartbeat_group_ip` (String) A valid Heartbeat Synchronization Group IP Address (IPv4 or IPv6) between '224.0.0.0 and 239.255.255.255' if you want to change the multicast IP addresses used for node-to-node communications (default: 225.1.1.1). This multicast IP address must not be used for other purposes on any of the network interfaces.
- `incr_sync_interval` (Number) The Incremental Synchronization Interface value (in milliseconds), which defines how frequently the incremental synchronizations are done. Do not set the values much higher or lower than their defaults (50 ms for incremental). Caution! Adjusting the Sync Intervals has significant impact on the cluster's performance. Inappropriate settings seriously degrade the firewall's performance.
- `statesync_group_ip` (String) A valid State Synchronization Group IP Address (IPv4 or IPv6) between '224.0.0.0 and 239.255.255.255' if you want to change the multicast IP addresses used for node-to-node communications (default: 225.1.1.2). This multicast IP address must not be used for other purposes on any of the network interfaces.
- `sync_mode` (String) The synchronization mode for the cluster, which defines how nodes exchange information about the traffic they process. 'sync_all': (recommended) both full and incremental synchronization messages are sent. This allows frequent updates without consuming resources excessively. Regular full synchronization ensures that all nodes stay synchronized even if some incremental messages are not delivered. 'sync_full': (not recommended) Only full synchronization messages are sent. Incremental updates are not sent in between, so nodes may not have the same information on connections unless the full sync interval is significantly reduced. 
- `sync_security` (String) The Synchronization Security Level, which defines the security features for synchronization. 'none': no security features. Do not select this option unless the heartbeat traffic uses a dedicated, secure network that does not handle other traffic. 'sign': (default) transmissions are authenticated to prevent outside injections of connection state information. 'encrypt': transmissions are authenticated and encrypted. This option increases the overhead compared to the default option, but is strongly recommended if the node-to-node communications are relayed through insecure networks (for example, if the backup heartbeat is configured on an interface that handles other traffic).


