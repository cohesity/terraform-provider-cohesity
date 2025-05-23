---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cohesity_ngce_cluster Resource - terraform-provider-cohesity"
subcategory: ""
description: |-
  
---

# cohesity_ngce_cluster (Resource)



## Example Usage

```terraform
resource "cohesity_ngce_cluster" "cluster" {
  name                     = "cluster_name"
  node_ips                 = ["0.0.0.0", "0.0.0.1"]
  hostname                 = "hostname"
  subnet_gateway           = "subnetwork_gateway"
  subnet_mask              = "subnet_mask"
  dns_server_ips           = "dns-server-ips"
  ntp_servers              = "ntp-servers"
  domain_names             = "LOCAL"
  metadata_fault_tolerance = 1
  add_apps                 = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `add_apps` (Boolean)
- `dns_server_ips` (String)
- `domain_names` (String)
- `hostname` (String)
- `metadata_fault_tolerance` (Number)
- `name` (String)
- `node_ips` (List of String)
- `ntp_servers` (String)
- `subnet_gateway` (String)
- `subnet_mask` (String)

### Optional

- `apps_subnet` (String)
- `apps_subnet_mask` (String)

### Read-Only

- `id` (String) The ID of this resource.
- `node_ip_map` (List of Object) (see [below for nested schema](#nestedatt--node_ip_map))

<a id="nestedatt--node_ip_map"></a>
### Nested Schema for `node_ip_map`

Read-Only:

- `node_id` (Number)
- `node_ips` (List of String)
