---
layout: "cohesity"
page_title: "Cohesity: cohesity_vmware_restore_vm"
description: |-
  Create Protection source, protection job, restore task, and destroy them in chronological order.
---

# cohesity\_vmware\_restore\_vm

Create Protection source, protection job, restore task, and destroy them in chronological order.

## Example Usage
```
provider "cohesity" {
        cluster_vip = "10.2.33.199"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}

resource "cohesity_restore_vmware_vm" "terraform_restore_vmware_vm"{
	name = "terraform_restore_vmware_vm"
	job_name = cohesity_job_run.terraform_vmware_job_run.name
	vm_names = [protection_vm_names]
	vmware_parameters = {
		"prefix" = "terraform"
	}
}
```

## Argument Reference

The following arguments are supported:
- name - (Required, string) Specifies the name of the restore task.
- job_name - (Required, string) Specifies the name of the protection job that backed up the objects to be restored.
- backup_timestamp - (Optional, string) Specifies the time of the protection job run. Should be in the format YYYY-MM-DD HH:MM Area/Location.
- vm_names - (Optional, string) Specifies the names of the virtual machines to restore.
- operation_timeout - (Optional, int) The time to wait in minutes for restore task start/stop operation.
- state - (Optional, string) Specifies whether to start or stop a restore task.
- vmware_parameters - (Optional, map) Specifies vmware parameters for the restore task.

#### Attributes Reference

The following attributes are exported:
- id - ID of the VMware protection Job