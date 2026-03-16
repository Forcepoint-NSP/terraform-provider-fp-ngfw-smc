---
page_title: "members_list"
subcategory: ""
description: |-
  This represents a member of a Server Pool, which is used to define the internal IP address of each member in the Server Pool. The Firewall uses these addresses to select which server handles traffic arriving at the Server Pool's external address.
---

# members_list (Nested-Attribute)

This represents a member of a Server Pool, which is used to define the internal IP address of each member in the Server Pool. The Firewall uses these addresses to select which server handles traffic arriving at the Server Pool's external address.




## Simple Attributes
- `member` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `member_rank` (Number) The rank of the server pool member, which determines the order of the member in the Server Pool. A value of -1 indicates no specific order.


