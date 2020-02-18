---
layout: "cohesity"
page_title: "Cohesity: cohesity_vmware_source"
description: |-
  Register, unregister and update a VMware protection source.
---

# cohesity\_virtual\_edition\_cluster

Register, unregister and update a VMware protection source

#### Example usage
```
provider "cohesity" {
	cluster_vip = "10.22.135.47"
	cluster_username = "admin"
	cluster_domain = "LOCAL"
}

resource "cohesity_source_vmware" "source1" {
	endpoint = "vcenter.lab.com"
	username = "admin"
	vmware_type = "VCenter"
	cap_streams_per_datastore = true
	number_of_streams = 5
	enable_latency_throttling = true
	new_task_latency = 110
	active_task_latency = 120
}
```
### Argument Reference
The following arguments are supported:
- endpoint - (Required, string) The network endpoint of the protection source where it is reachable. It could be an URL or hostname or an IP address of the protection source
- vmware_type - (Optional, string) The VMware entity type. The default vaule is **VCenter**. Supported types include **VCenter**, **VCloudDirector** and **HostSystem** (Standalone ESXi Host)
- username - (Required, string) The username to access the target source
- password - (Required, string) The password of the username to access the target source. This can also be read from **COHESITY_SOURCE_VMWARE_PASSWORD** environment variable
- enable_ssl_verification - (Optional, bool) Specifies whether SSL verification should be performed or not. The default vaule is **false**
- ca_certificate - (Optional, bool) The contents of CA certificate. Required when **enable_ssl_verification** is **true**
- cap_streams_per_datastore - (Optional, bool) Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set
				to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for **number_of_streams**. If
				not set or set to false, there is no max limit for the number of concurrent streams. The default vaule is **false**
- number_of_streams - (Optional, int) Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores
				of the registered entity. This limit is enforced only when the **cap_streams_per_datastore** is set to **true**. The default value is **1**
- enable_latency_throttling - (Optional, bool) Indicates whether read operations to the datastores which are part of the registered Protection Source are throttled or not.
The default vaule is **false**
- new_task_latency - (Optional, int) If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started. **enable_latency_throttling** must be set to **true** to apply this configuration. The default value is **30**
- active_task_latency - (Optional, int) If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled. **enable_latency_throttling** must be set to **true** to apply this configuration. The default value is **30**

#### Attributes Reference
The following attributes are exported:
- id - ID of the VMware protection source