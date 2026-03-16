---
page_title: "server_directories"
subcategory: ""
description: |-
  This represents the server directories used for log operations, including the server reference, archive directory mask, and whether only archive directories are considered.
---

# server_directories (Nested-Attribute)

This represents the server directories used for log operations, including the server reference, archive directory mask, and whether only archive directories are considered.




## Simple Attributes
- `archive_mask` (Number) The bit mask of selected archive directories, -1 for all.
- `only_archive` (Boolean) Indicates whether only archive directories are considered for log operations.
- `server` (String) This represents a server in the network. It is an abstract class that can be extended to define specific types of servers.


