package cohesity

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/cohesity/management-sdk-go/models"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccRestoreVmwareVM(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRestoreVmwareVMDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccRestoreVmwareVMCreateConfig, restoreVMSourceEndpoint, restoreVMProtectionJobInclude, restoreVMVMNames),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRestoreVmwareVMCreated(),
				),
			},
		},
	})
}

func testAccCheckRestoreVmwareVMDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}

	result, err := client.RestoreTasks().GetRestoreTasks(nil, nil, nil, nil, nil, models.EnvironmentGetRestoreTasks_KVMWARE)
	if err != nil || len(result) == 0 {
		return fmt.Errorf("Failed to create VMware Restore Task")
	}
	//validate protection job parameters with values from testAccJobVmwareCreateConfig
	//configuration
	if result[0].Name != "terraform_restore_vm" {
		return fmt.Errorf("Failed to validate the created VMware Restore Task")
	}
	return fmt.Errorf("Failed to delete the VMware Restore Task")
}

func testAccCheckRestoreVmwareVMCreated() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
			cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}

		result, err := client.RestoreTasks().GetRestoreTasks(nil, nil, nil, nil, nil, models.EnvironmentGetRestoreTasks_KVMWARE)
		if err != nil || len(result) == 0 {
			return fmt.Errorf("Failed to create VMware Restore Task")
		}
		//validate protection job parameters with values from testAccJobVmwareCreateConfig
		//configuration
		if result[0].Name != "terraform_restore_vm" {
			return fmt.Errorf("Failed to validate the created VMware Restore Task")
		}

		return nil
	}
}

const testAccRestoreVmwareVMCreateConfig = `
resource "cohesity_source_vmware" "terraform_vmware_source"{
	endpoint = "%s"
	username = "administrator"
	vmware_type = "VCenter"
	cap_streams_per_datastore = true
	number_of_streams = 5
	enable_latency_throttling = true
	new_task_latency = 110
	active_task_latency = 120
}

resource "cohesity_job_vmware" "terraform_vmware_protection_job" {
	name = "terraform_protect_vcenter"
	protection_source = cohesity_source_vmware.terraform_vmware_source.endpoint
	include = ["%s"]
	policy = "Bronze"
	storage_domain = "DefaultStorageDomain"
	delete_snapshots = true
	full_sla = 200
	incremental_sla = 140
	qos_type = "BackupSSD"
	priority = "Low"
}

resource "cohesity_job_run" "terraform_vmware_job_run"{
	name = cohesity_job_vmware.terraform_vmware_protection_job.name
	timestamp = "${formatdate("YYYYMMDD", timestamp())}"
}

resource "cohesity_restore_vmware_vm" "terraform_restore_vmware_vm"{
	name = "terraform_restore_vmware_vm"
	job_name = cohesity_job_run.terraform_vmware_job_run.name
	vm_names = [%s]
	vmware_parameters = {
		"prefix" = "terraform"
	}
}
`
