---
layout: "cohesity"
page_title: "Cohesity: cohesity_cloud_edition_cluster"
description: |-
  Create cloud edition cluster, apply license key and destroy cluster.
---

# cohesity\_cloud\_edition\_cluster

Create cloud edition cluster, apply license key and destroy cluster

## Example Usage
```
provider "cohesity" {
        cluster_vip = "10.2.45.143"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}

resource "cohesity_cloud_edition_cluster" "cloud"{
            cluster_name = "TerraformCloudEditionCluster"
            dns_servers = ["10.2.145.14"]
            ntp_servers = ["time.google.com"]
            domain_names = ["eng.cohesity.com"]
            cluster_subnet_mask = "255.255.240.0"
            cluster_gateway = "10.2.32.1"
            enable_encryption = true
            enable_fips_mode = true
            encryption_keys_rotation_period = 1
            metadata_fault_tolerance = 0
            node_ips = ["10.2.45.143", "10.2.45.144", "10.2.45.145"]           
}
```

## Argument Reference

The following arguments are supported:
- cluster_name - (Required, string) The name of the new Cloud edition cluster
- licence_key - (Required, string) Cohesity licence key to apply after cluster creation. This can also be read from **CLOUD_COHESITY_CLUSTER_LICENCE_KEY** environment variable 
- metadata_fault_tolerance - (Optional, int) The metadata fault tolerance. Default value is **0**
- enable_encryption - (Optional, bool) Specifies whether or not to enable encryption. If encryption is enabled, all data on the cluster will be encrypted. Default value is **true**
- enable_fips_mode - (Optional, bool) Specifies whether or not to enable FIPS mode. This must be set to true in order to enable FIPS. Default value is **true**
- encryption_keys_rotation_period - (Optional, int) The rotation period for encryption keys in days. The default value is **1**
- cluster_gateway - (Required, string) The default gateway IP address for the cluster network
- cluster_subnet_mask - (Required, string) The subnet mask of the cluster network
- domain_names - (Required, set of strings) The domain names to configure on the cluster
- ntp_servers - (Required, set of strings) The NTP servers to configure on the cluster
- dns_servers - (Required, set of strings) The DNS servers to configure on the cluster
- node_ips - (Required, set of strings) IP addresses of the nodes in the cluster
- operation_timeout - (Optional, int) The time to wait in minutes for cluster creation or destruction. The default value is **120 minutes**

#### Attributes Reference

The following attributes are exported:
- id - ID of the cluster