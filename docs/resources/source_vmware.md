---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cohesity_source_vmware Resource - terraform-provider-cohesity"
subcategory: ""
description: |-
  
---

# cohesity_source_vmware (Resource)



## Example Usage

```terraform
resource "cohesity_source_vmware" "source1" {
  endpoint    = "vc.xyz.com"
  username    = "administrator"
  password    = "password"
  vmware_type = "VCenter"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `endpoint` (String) Specifies the network endpoint of the Protection
				Source where it is reachable. It could be an URL or hostname or
				an IP address of the Protection Source
- `username` (String) Specifies username to access the target source

### Optional

- `active_task_latency` (Number) If the latency of a datastore is above this value,
				existing backup tasks using the datastore are throttled.
- `ca_certificate` (String, Sensitive) The contents of CA certificate
- `cap_streams_per_datastore` (Boolean) Specifies whether datastore streams are configured
				for all datastores that are part of the registered entity. If set
				to true, number of streams from Cohesity cluster to the registered
				entity will be limited to the value set for number_of_streams. If
				not set or set to false, there is no max limit for the number of 
				concurrent streams.
- `enable_latency_throttling` (Boolean) Indicates whether read operations to the datastores,
				which are part of the registered Protection Source, are throttled.
- `enable_ssl_verification` (Boolean) Specifies whether SSL verification should be performed.
- `new_task_latency` (Number) If the latency of a datastore is above this value,
				then new backup tasks using the datastore will not be started.
- `number_of_streams` (Number) Specifies the limit on the number of streams
				Cohesity cluster will make concurrently to the datastores
				of the registered entity. This limit is enforced only when the
				cap_streams_per_datastore is set to true
- `password` (String, Sensitive) Specifies password of the username to access the target source
- `vmware_type` (String) Specifies the VMware entity type

### Read-Only

- `id` (String) The ID of this resource.
