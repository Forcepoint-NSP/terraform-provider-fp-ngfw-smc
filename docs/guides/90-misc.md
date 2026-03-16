---
page_title: "Misc information"
subcategory: ""
description: |-
  Misc information
---

# Misc information

## Using a forward proxy

In case the host running terraform cli needs to go through a forward proxy to contact the SMC host:
under linux, define the following env. var 

```
export HTTP_PROXY="http://smcproxy:8080"
```

## Experimental mode

Some resources are in experimental state and are unavailable by default.

To activate them, modify the `tfsmc.conf.json` config file:

```json
{
   "experimental": true
}
```
