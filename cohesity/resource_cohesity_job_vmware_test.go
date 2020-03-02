package cohesity

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccJobVmware(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJobVmwareDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccJobVmwareCreateConfig, jobVMwareCreateSourceEndpoint),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJobVmwareCreated(),
				),
			},
			{
				Config: fmt.Sprintf(testAccJobVmwareUpdateConfig, jobVMwareUpdateSourceEndpoint),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJobVmwareUpdated(),
				),
			},
		},
	})
}

func testAccCheckJobVmwareDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}

	resourceName := "cohesity_job_vmware.job1"
	resource, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return fmt.Errorf("Can't find the resource: %s", resourceName)
	}
	if resource.Primary.ID == "" {
		return fmt.Errorf("The ID for resoource %s for is not set", resourceName)
	}
	//parse a decimal string of base 10 and converts into int64
	jobID, _ := strconv.ParseInt(resource.Primary.ID, 10, 64)
	result, err := client.ProtectionJobs().GetProtectionJobs([]int64{jobID},
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil || len(result) == 0 {
		return nil
	}
	return fmt.Errorf("Failed to delete the VMware protection job")
}

func testAccCheckJobVmwareCreated() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
			cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}
		resourceName := "cohesity_job_vmware.job1"
		resource, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Can't find the resource: %s", resourceName)
		}
		if resource.Primary.ID == "" {
			return fmt.Errorf("The ID for resource %s is not set", resourceName)
		}
		//parse a decimal string of base 10 and converts into int64
		jobID, _ := strconv.ParseInt(resource.Primary.ID, 10, 64)
		result, err := client.ProtectionJobs().GetProtectionJobs([]int64{jobID},
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
		if err != nil || len(result) == 0 {
			return fmt.Errorf("Failed to create VMware protection job")
		}
		//validate protection job parameters with values from testAccJobVmwareCreateConfig
		//configuration
		if result[0].Name != "protect_vcenter" || *result[0].FullProtectionSlaTimeMins != 200 ||
			*result[0].IncrementalProtectionSlaTimeMins != 140 {
			return fmt.Errorf("Failed to validate the created VMware protection job")
		}
		return nil
	}
}

func testAccCheckJobVmwareUpdated() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
			cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}
		resourceName := "cohesity_job_vmware.job1"
		resource, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Can't find the resource: %s", resourceName)
		}
		if resource.Primary.ID == "" {
			return fmt.Errorf("The ID for resoource %s for is not set", resourceName)
		}
		//parse a decimal string of base 10 and converts into int64
		jobID, _ := strconv.ParseInt(resource.Primary.ID, 10, 64)
		result, err := client.ProtectionJobs().GetProtectionJobs([]int64{jobID},
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
		if err != nil || len(result) == 0 {
			return fmt.Errorf("Failed to find the VMware protection job")
		}
		//validate protection job parameters with values from testAccJobVmwareUpdateConfig
		//configuration
		if result[0].Name != "protect_vmware" || *result[0].FullProtectionSlaTimeMins != 180 ||
			*result[0].IncrementalProtectionSlaTimeMins != 120 {
			return fmt.Errorf("Failed to validate the updated VMware protection job")
		}
		return nil
	}
}

const testAccJobVmwareCreateConfig = `
resource "cohesity_source_vmware" "vmware_source"{
	endpoint = "%s"
	username = "administrator"
	vmware_type = "VCenter"
	cap_streams_per_datastore = true
	number_of_streams = 5
	enable_latency_throttling = true
	new_task_latency = 110
	active_task_latency = 120
}

resource "cohesity_job_vmware" "terraform_vmware_job" {
	name = "protect_vcenter"
	protection_source = cohesity_source_vmware.vmware_source.endpoint
	policy = "Bronze"
	storage_domain = "DefaultStorageDomain"
	delete_snapshots = true
	full_sla = 200
	incremental_sla = 140
	qos_type = "BackupSSD"
	priority = "Low"
}
`

const testAccJobVmwareUpdateConfig = `
resource "cohesity_source_vmware" "vmware_source"{
	endpoint = "%s"
	username = "administrator"
	vmware_type = "VCenter"
	cap_streams_per_datastore = true
	number_of_streams = 5
	enable_latency_throttling = true
	new_task_latency = 110
	active_task_latency = 120
}

resource "cohesity_job_vmware" "terraform_vmware_job" {
	name = "protect_vmware"
	protection_source = cohesity_source_vmware.vmware_source.endpoint
	policy = "Bronze"
	storage_domain = "DefaultStorageDomain"
	delete_snapshots = true
	full_sla = 180
	incremental_sla = 120
	qos_type = "BackupSSD"
	priority = "Low"
}
`
