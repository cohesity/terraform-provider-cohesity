provider "cohesity" {
	cluster_vip = "10.10.10.10"
	cluster_username = "admin"
	cluster_password= "admin"
	cluster_domain = "LOCAL"
}

resource "cohesity_restore_vmware_vm" "cohesity_restore_vm01" {
	name = "restore-terraform01"
	job_name = "terraform-job2"
	vm_names = ["chandra-pwsh01"]
	vmware_parameters = {
    "prefix" = "terra-"
	}
}
