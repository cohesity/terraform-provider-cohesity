package cohesity

import (
	"errors"
	"fmt"
	"log"
	"testing"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccSourceVmware(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSourceVmwareDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSourceVmwareCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceVmwareCreated(),
				),
			},
			{
				Config: testAccSourceVmwareUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceVmwareUpdated(),
				),
			},
		},
	})
}

func testAccCheckSourceVmwareDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}
	var endpoint = "vc-67.eco.eng.cohesity.com"
	var environmentType = []models.EnvironmentListProtectionSourcesEnum{models.
		EnvironmentListProtectionSources_KVMWARE}
	var trueValue = true
	log.Printf("[INFO] Get list of VMware protection sources to verify %s deletion", endpoint)
	result, err := client.ProtectionSources().
		ListProtectionSources(nil, nil, nil, &trueValue, &trueValue,
			&trueValue, environmentType, nil, nil, nil, nil, nil)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to get VMware protection sources")
	}
	for _, protectionSource := range result {
		if *protectionSource.ProtectionSource.Name == endpoint {
			log.Printf("[INFO] Found the VMware protection source %s on cohesity cluster",
				*protectionSource.ProtectionSource.Name)
			return fmt.Errorf("Failed to unregister the Vmware protection source %s", endpoint)
		}
	}
	return nil
}

func testAccCheckSourceVmwareCreated() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
			cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}
		var endpoint = "vc-67.eco.eng.cohesity.com"
		var environmentType = []models.EnvironmentListProtectionSourcesEnum{models.
			EnvironmentListProtectionSources_KVMWARE}
		var trueValue = true
		log.Printf("[INFO] Get list of VMware protection sources to verify %s creation",
			endpoint)
		result, err := client.ProtectionSources().
			ListProtectionSources(nil, nil, nil, &trueValue, &trueValue,
				&trueValue, environmentType, nil, nil, nil, nil, nil)
		if err != nil {
			log.Printf(err.Error())
			return fmt.Errorf("Failed to get VMware protection sources")
		}
		for _, protectionSource := range result {
			if *protectionSource.ProtectionSource.Name == endpoint {
				log.Printf("[INFO] Found the VMware protection source %s on cohesity cluster",
					endpoint)
				//validate the throttling policy parameters with values from testAccSourceVmwareCreateConfig
				//configuration
				if *protectionSource.RegistrationInfo.ThrottlingPolicy.IsEnabled != true ||
					*protectionSource.RegistrationInfo.ThrottlingPolicy.EnforceMaxStreams != true ||
					*protectionSource.RegistrationInfo.ThrottlingPolicy.MaxConcurrentStreams != 5 ||
					*protectionSource.RegistrationInfo.ThrottlingPolicy.LatencyThresholds.NewTaskMsecs != 110 ||
					*protectionSource.RegistrationInfo.ThrottlingPolicy.LatencyThresholds.ActiveTaskMsecs != 120 {
					return fmt.Errorf("Failed to valaidate throttling policy parameters of created Vmware source")
				}
				return nil
			}
		}
		return fmt.Errorf("Failed to create VMware protection source: %s", endpoint)
	}
}

func testAccCheckSourceVmwareUpdated() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
			cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}
		var endpoint = "vc-67.eco.eng.cohesity.com"
		var environmentType = []models.EnvironmentListProtectionSourcesEnum{models.
			EnvironmentListProtectionSources_KVMWARE}
		var trueValue = true
		log.Printf("[INFO] Get list of VMware protection sources to verify %s update",
			endpoint)
		result, err := client.ProtectionSources().
			ListProtectionSources(nil, nil, nil, &trueValue, &trueValue,
				&trueValue, environmentType, nil, nil, nil, nil, nil)
		if err != nil {
			log.Printf(err.Error())
			return fmt.Errorf("Failed to get VMware protection sources")
		}
		for _, protectionSource := range result {
			if *protectionSource.ProtectionSource.Name == endpoint {
				log.Printf("[INFO] Found the VMware protection source %s on cohesity cluster",
					endpoint)
				//validate the throttling policy parameters with values from testAccSourceVmwareUpdateConfig
				//configuration
				if *protectionSource.RegistrationInfo.ThrottlingPolicy.IsEnabled != true ||
					*protectionSource.RegistrationInfo.ThrottlingPolicy.EnforceMaxStreams != true ||
					*protectionSource.RegistrationInfo.ThrottlingPolicy.MaxConcurrentStreams != 20 ||
					*protectionSource.RegistrationInfo.ThrottlingPolicy.LatencyThresholds.NewTaskMsecs != 150 ||
					*protectionSource.RegistrationInfo.ThrottlingPolicy.LatencyThresholds.ActiveTaskMsecs != 100 {
					return fmt.Errorf("Failed to valaidate throttling policy parameters of updated Vmware source")
				}
				return nil
			}
		}
		return fmt.Errorf("Failed to find VMware protection source: %s", endpoint)
	}
}

const testAccSourceVmwareCreateConfig = `
provider "cohesity" {
	cluster_vip = "10.2.145.47"
	cluster_username = "admin"
	cluster_domain = "LOCAL"
}

resource "cohesity_source_vmware" "source1" {
	endpoint = "vc-67.eco.eng.cohesity.com"
	username = "administrator"
	vmware_type = "VCenter"
	cap_streams_per_datastore = true
	number_of_streams = 5
	enable_latency_throttling = true
	new_task_latency = 110
	active_task_latency = 120
}
`

const testAccSourceVmwareUpdateConfig = `
provider "cohesity" {
	cluster_vip = "10.2.145.47"
	cluster_username = "admin"
	cluster_domain = "LOCAL"
}

resource "cohesity_source_vmware" "source1" {
	endpoint = "vc-67.eco.eng.cohesity.com"
	username = "administrator"
	vmware_type = "VCenter"
	cap_streams_per_datastore = true
	number_of_streams = 20
	enable_latency_throttling = true
	new_task_latency = 150
	active_task_latency = 100
}
`
