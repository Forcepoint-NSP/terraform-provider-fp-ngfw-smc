---
page_title: "ospfv2_key_chain_entry"
subcategory: "routing"
description: |-
  This represents an OSPFv2 Key Chain Entry, which contains the Message Digest Key ID and Key. It includes fields such as sendKey, keyID, key, comment, and algorithm.
---

# ospfv2_key_chain_entry (Nested-Attribute)

This represents an OSPFv2 Key Chain Entry, which contains the Message Digest Key ID and Key. It includes fields such as sendKey, keyID, key, comment, and algorithm.




## Simple Attributes
- `algorithm` (String) The cryptographic algorithm used for OSPFv2 authentication, which can be 'md5'
- `comment` (String) An optional comment for the OSPF Key Chain Entry, providing additional context or information.
- `key` (String) The cryptographic key used for OSPFv2 authentication, limited to 16 characters.
- `key_id` (Number) The unique identifier for the key, which is used to reference the key in OSPFv2.
- `send_key` (Boolean) Indicates whether the key should be sent to the engine.


