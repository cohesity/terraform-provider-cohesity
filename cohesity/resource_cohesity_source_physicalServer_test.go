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

func TestAccSourcePhysicalServer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSourcePhysicalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSourcePhysicalServerConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourcePhysicalServerCreated(),
				),
			},
		},
	})
}

func testAccCheckSourcePhysicalServerCreated() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
			cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}

		var endpoint = "10.2.150.118"
		var environmentType = []models.EnvironmentListProtectionSourcesEnum{models.EnvironmentListProtectionSources_KPHYSICAL}
		var trueValue = true
		log.Printf("[INFO] Get list of Physical Server sources to verify %s creation", endpoint)

		log.Println("[INFO] Access to Global", ProtectionSourceID)
		// result, err := client.ProtectionSources().
		// 	ListProtectionSources(nil, nil, nil, &trueValue, &trueValue,
		// 		&trueValue, environmentType, nil, nil, nil, nil, nil)

		result, err := client.ProtectionSources().ListProtectionSources(&ProtectionSourceID, nil, nil, &trueValue, &trueValue, &trueValue, environmentType, nil, &trueValue, nil, nil, &trueValue)

		if err != nil {
			log.Printf(err.Error())
			return fmt.Errorf("Failed to get Physical Server protection sources")
		}

		for _, protectionSource := range result {
			if *protectionSource.ProtectionSource.Name == endpoint {
				log.Printf("[INFO] Found the Physical Server protection source %s on cohesity cluster", endpoint)
				//validate the Physical Server Source Registration
				if *protectionSource.LogicalSize == 0 {
					return fmt.Errorf("Failed to valaidate created physical server source")
				}
				// Check if registered
				// Check if deleted

				// 	*protectionSource.RegistrationInfo.AccessInfo.Endpoint != endpoint {
				// 	return fmt.Errorf("Failed to valaidate created physical server source")
				// }

				// if *protectionSource.RegistrationInfo.ThrottlingPolicy.IsEnabled != true ||
				// 	*protectionSource.RegistrationInfo.ThrottlingPolicy.EnforceMaxStreams != true ||
				// 	*protectionSource.RegistrationInfo.ThrottlingPolicy.MaxConcurrentStreams != 5 ||
				// 	*protectionSource.RegistrationInfo.ThrottlingPolicy.LatencyThresholds.NewTaskMsecs != 110 ||
				// 	*protectionSource.RegistrationInfo.ThrottlingPolicy.LatencyThresholds.ActiveTaskMsecs != 120 {
				// 	return fmt.Errorf("Failed to valaidate throttling policy parameters of created Vmware source")
				// }
				return nil
			}
		}
		return fmt.Errorf("Failed to create Physical Server protection source: %s", endpoint)
	}
}

func testAccCheckSourcePhysicalServerDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}

	var endpoint = "10.2.150.118"
	var environmentType = []models.EnvironmentListProtectionSourcesEnum{models.EnvironmentListProtectionSources_KPHYSICAL}
	var trueValue = true
	log.Printf("[INFO] Get list of Physical Server protection sources to verify %s deletion", endpoint)

	result, err := client.ProtectionSources().
		ListProtectionSources(nil, nil, nil, &trueValue, &trueValue,
			&trueValue, environmentType, nil, nil, nil, nil, nil)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to get Physical Server protection sources")
	}

	for _, protectionSource := range result {
		log.Printf("[INFO] Protection Name: %s", *protectionSource.ProtectionSource.Name)
		if *protectionSource.ProtectionSource.Name == endpoint {
			log.Printf("[INFO] Found the Physical Server protection source %s on cohesity cluster",
				*protectionSource.ProtectionSource.Name)
			return fmt.Errorf("Failed to unregister the Physical Server protection source %s", endpoint)
		}
	}
	return nil
}

const testAccSourcePhysicalServerConfig = `
provider "cohesity" {
	cluster_vip = "10.2.145.47"
	cluster_username = "admin"
	cluster_domain = "LOCAL"
}

resource "cohesity_source_physical_server" "physical_server"{
	endpoint = "10.2.150.118"
	force_register = true
	host_type = "kWindows"
	physical_type = "kHost"
}
`
