package cohesity

import (
	"errors"
	"fmt"
	"log"
	"testing"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccJobRun(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJobRunDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccJobRunCreateConfig, jobRunCreateName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJobRunCreated(),
				),
			},
			{
				Config: fmt.Sprintf(testAccJobRunUpdateConfig, jobRunUpdateName, jobRunUpdateState),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJobRunUpdated(),
				),
			},
		},
	})
}

func testAccCheckJobRunDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}
	jobName := jobRunCreateName
	//get the protection job details
	log.Printf("[INFO] Get the protection job (%s) details", jobName)
	protectionJob, err := client.ProtectionJobs().GetProtectionJobs(nil, []string{jobName}, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		log.Printf(err.Error())
		return errors.New("Failed to find the protection job")
	}
	jobID := *protectionJob[0].Id
	//get the protection runs to check the the status of recent job run
	log.Printf("[INFO] Protection job %s found with ID %d ", jobName, jobID)
	return fmt.Errorf("Failed to delete the VMware protection job")
}

func testAccCheckJobRunCreated() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
			cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}
		jobName := jobRunCreateName
		//get the protection job details
		log.Printf("[INFO] Get the protection job (%s) details", jobName)
		protectionJob, err := client.ProtectionJobs().GetProtectionJobs(nil, []string{jobName}, nil,
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to find the protection job")
		}
		jobID := *protectionJob[0].Id
		//get the protection runs to check the the status of recent job run
		log.Printf("[INFO] Protection job %s found with ID %d ", jobName, jobID)
		return nil
	}
}

func testAccCheckJobRunUpdated() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}

const testAccJobRunCreateConfig = `
resource "cohesity_job_run" "terraform_vmware_job_run"{
	name = "%s"
	timestamp = "${formatdate("YYYYMMDD", timestamp())}"
}`

const testAccJobRunUpdateConfig = `
resource "cohesity_job_run" "terraform_vmware_job_run"{
	name = "%s"
	state = "%s"
}`
