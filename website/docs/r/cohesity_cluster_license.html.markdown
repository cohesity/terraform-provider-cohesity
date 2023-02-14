---
layout: "cohesity"
page_title: "Cohesity: cohesity_cluster_license"
description: |-
  Apply license to a Cohesity cluster.
---

# cohesity\_cluster\_license

Apply license to a Cohesity cluster

## Example Usage
```
provider "cohesity" {
        cluster_vip = "10.2.45.143"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}


resource "cohesity_cluster_license" "cs1" {
        license_key  = "XXXX-XXXX-XXXX-XXXX-XXXX-XXXX"
}
```

### Argument Reference
The following arguments are supported:

- license_key - (Required, string) Cohesity license key to apply.

#### Attributes Reference
The following attributes are exported:

- id - ID used to identify a license key detail
