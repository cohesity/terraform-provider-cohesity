---
layout: "cohesity"
page_title: "Cohesity: cohesity_physical_edition_cluster"
description: |-
  Create physical edition cluster, apply license key and destroy cluster.
---

# cohesity\_physical\_edition\_cluster

Create physical edition cluster, apply license key and destroy cluster

## Example Usage
```
provider "cohesity" {
        cluster_vip = "10.9.33.133"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}

resource "cohesity_virtual_edition_cluster" "physical"{
            cluster_name = "TerraformPhysicalEditionCluster"
            dns_servers = ["10.2.145.14"]
            ntp_servers = ["time.google.com"]
            domain_names = ["eng.cohesity.com"]
            cluster_subnet_mask = "255.255.240.0"
            cluster_gateway = "10.2.32.1"
            enable_encryption = true
            enable_fips_mode = true
            encryption_keys_rotation_period = 1
            metadata_fault_tolerance = 0
            virtual_ips = ["10.2.33.137"]
            ipmi_gateway = "10.2.144.54"
            ipmi_subnet_mask = "255.255.240.0"
            virtual_ip_hostname = "test"
            ipmi_username="cohesity"
            node_configs {
                        node_ip="10.9.33.133"
                        node_ipmi_ip="10.9.33.133"
                        }
            node_configs {
                        node_ip="10.9.33.134"
                        node_ipmi_ip="10.9.33.134"
                        }
            node_configs {
                        node_ip="10.9.33.135"
                        node_ipmi_ip="10.9.33.135"
                        }
}
```

## Argument Reference

The following arguments are supported:
- cluster_name - (Required, string) The name of the new Physical edition cluster
- licence_key - (Required, string) Cohesity licence key to apply after cluster creation. This can also be read from **PHYSICAL_COHESITY_CLUSTER_LICENCE_KEY** environment variable 
- metadata_fault_tolerance - (Optional, int) The metadata fault tolerance. Default value is **0**
- enable_encryption - (Optional, bool) Specifies whether or not to enable encryption. If encryption is enabled, all data on the cluster will be encrypted. Default value is true
- enable_fips_mode - (Optional, bool) Specifies whether or not to enable FIPS mode. This must be set to true in order to enable FIPS. Default value is **true**
- encryption_keys_rotation_period - (Optional, int) The rotation period for encryption keys in days. The default value is **1**
- cluster_gateway - (Required, string) The default gateway IP address for the cluster network
- cluster_subnet_mask - (Required, string) The subnet mask of the cluster network
- domain_names - (Required, set of strings) The domain names to configure on the cluster
- ntp_servers - (Required, set of strings) The NTP servers to configure on the cluster
- dns_servers - (Required, set of strings) The DNS servers to configure on the cluster
- virtual_ips - (Required, set of strings) The virtauls IPs for the new cluster
- operation_timeout - (Optional, int) The time to wait in minutes for cluster creation or destruction. The default value is **120 minutes**
- virtual_ip_hostname - (Required, string) The virtual IP hostname
- ipmi_username - (Required, string) The IPMI username. This can also be read from **PHYSICAL_COHESITY_CLUSTER_IPMI_USERNAME** environment variable 
- ipmi_password - (Required, string) The IPMI password. This can also be read from **PHYSICAL_COHESITY_CLUSTER_IPMI_PASSWORD** environment variable 
- ipmi_gateway - (Required, string) The default gateway IP address for the IPMI network
- ipmi_subnet_mask - (Required, string) The subnet mask for the IPMI network
- node_configs - (Required, block) node_configs is a block within the configuration to configure the nodes in the cluster. The block can be repeated to configure multiple nodes in the cluster. Each block supports the following:
   - node_ip - (Required, string) IP address of the node
   - node_ipmi_ip - (Required, string) IPMI IP for this node

#### Attributes Reference

The following attributes are exported:
- id - ID of the cluster