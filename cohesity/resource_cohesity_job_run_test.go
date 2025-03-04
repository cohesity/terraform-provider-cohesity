package cohesity

import (
	"errors"
	"fmt"
	"log"
	"testing"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func TestAccJobRun_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJobRunDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccJobRunCreateConfig, jobVMwareSourceEndpoint, vmWareJobName, vmWareJobProtectVM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJobRunCreated(),
				),
			},
		},
	})
}

func testAccCheckJobRunDestroy(s *terraform.State) error {
	return nil
}

func testAccCheckJobRunCreated() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(utils.Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.ClusterVIP,
			cohesityConfig.ClusterUsername, cohesityConfig.ClusterPassword, cohesityConfig.ClusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}
		jobName := vmWareJobName
		//get the protection job details
		log.Printf("[INFO] Get the protection job (%s) details", jobName)
		protectionJob, err := client.ProtectionJobs().GetProtectionJobs(nil, []string{jobName}, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to find the protection job")
		}
		jobID := *protectionJob[0].Id
		//get the protection runs to check the status of recent job run
		log.Printf("[INFO] Protection job %s found with ID %d ", jobName, jobID)
		return nil
	}
}

const testAccJobRunCreateConfig = `
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
	name = "%s"
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
	timestamp = "18000000"
}`
