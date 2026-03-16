---
page_title: "smc_elasticsearch_cluster"
subcategory: "network_elements"
description: |-
  This represents an Elasticsearch Cluster, which can be either Elasticsearch or OpenSearch. It includes attributes for product type, port, TLS profile, addresses, retention period, shard number, replica number, cluster sniffer, and authentication settings.
---

# smc_elasticsearch_cluster (Resource)

This represents an Elasticsearch Cluster, which can be either Elasticsearch or OpenSearch. It includes attributes for product type, port, TLS profile, addresses, retention period, shard number, replica number, cluster sniffer, and authentication settings.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/network_elements/servers/elasticsearch_cluster) for an example

This example creates an Elasticsearch cluster definition for SMC integration.

```hcl
resource "smc_elasticsearch_cluster" "elasticsearch_cluster" {
  addresses = ["10.3.3.3", "10.3.3.33"]
  name = "tf_elasticsearch_cluster"
  comment = var.resource_comment
  authentication_settings {
    login = "es_login"
    method = "basic"
    password = "superPassword00!!"
  }
  es_enable_cluster_sniffer = true
  es_replica_number = 1
  es_retention_period = 30
  es_shard_number = 0
  indexing_active = true
  port = 9200
  product = "elasticsearch"
  # tls_profile = data.smc_href.tls_profile.id
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `addresses` (List of String) The list of hostnames or IP addresses used to connect to the Elasticsearch Cluster. At least one address is required.
- `comment` (String) An optional comment for the element. This field is not required.
- `es_enable_cluster_sniffer` (Boolean) Indicates whether the cluster sniffer is enabled. When enabled, it allows automatic discovery of cluster nodes.
- `es_replica_number` (Number) The number of replicas for all indices, a positive number (default: 0). When changed, it will immediately reflect in the index settings. In status monitoring, the Elasticsearch cluster status will temporarily change (to orange) to reflect that some shards are pending allocation.
- `es_retention_period` (Number) The retention period for logs in the Elasticsearch Cluster, in days. Use -1 to retain all logs indefinitely.
- `es_shard_number` (Number) The number of shards for the fwlogsandalerts indices, a strictly positive number, or auto (default value). In auto mode, we will keep the number of shards set to the number of data nodes in the cluster. Changing this setting (or adapting it following a change in the number of data nodes) will take effect the next time a daily index gets created (technically speaking, we will only update the index template right away), i.e. for future indices. We won't update existing indices, as this may require to reindex them (the shrink API can only be used if the new number of shards is a divider of the previous one) and we would have to monitor all existing indices, etc. : users could do that via Kibana. 0 stands for auto
- `indexing_active` (Boolean) Indicates whether indexing is active in the Elasticsearch Cluster.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `port` (Number) The port number for the Elasticsearch Cluster.
- `product` (String) The product type of the Elasticsearch Cluster, either 'elasticsearch' or 'opensearch'.

## Nested Attributes
- `authentication_settings` (Single Block, see [here](attr_elasticsearch_authentication_settings.md)) 

## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
