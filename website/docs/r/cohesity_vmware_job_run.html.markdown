---
layout: "cohesity"
page_title: "Cohesity: cohesity_vmware_job_run"
description: |-
  Create, Update and Delete a Protection Job Run.
---

# cohesity\_virtual\_edition\_cluster

Create, Update and Delete a Protection Job Run.

## Example Usage
```
provider "cohesity" {
        cluster_vip = "10.2.33.199"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}

resource "cohesity_job_run" "terraform_vmware_job_run"{
	name = "protection_job_name"
    run_type = "Regular"
    state = "start"
    operation_timeout = 100
	timestamp = "${formatdate("YYYYMMDD", timestamp())}"
}
```

## Argument Reference

The following arguments are supported:
- name - (Required, string) Specifies the name of the Protection Job.
- run_type - (Optional, string) Specifies the type of backup.
- state - (Optional, string) Specifies whether to start or stop a protection job run.
- operation_timeout - (Optional, int) Specifies the time to wait in minutes for the protection job run to complete the run or stop the run. The default value is **120 minutes**
- timestamp - (Required, string) Specifies the current timestamp to trigger starting or stopping a job run.

#### Attributes Reference

The following attributes are exported:
- id - ID of the VMware protection Job
