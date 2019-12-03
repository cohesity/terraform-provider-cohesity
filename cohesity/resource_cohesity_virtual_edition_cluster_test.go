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

func TestAccVirtualEditionCluster(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVirtualEditionClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualEditionClusterConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccVirtualEditionClusterExists(),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "cluster_name",
						"AcceptanceTestTerraformVirtaulEditionCluster"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "enable_encryption", "true"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "cluster_gateway", "10.2.144.1"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "enable_fips_mode", "true"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "metadata_fault_tolerance", "0"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "virtual_ip_hostname", "test"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "cluster_subnet_mask", "255.255.240.0"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "encryption_keys_rotation_period", "1"),
				),
			},
		},
	})
}

func testAccCheckVirtualEditionClusterDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
		cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}
	result, infoErr := client.Cluster().GetBasicClusterInfo()
	if infoErr == nil && result.Name != nil {
		return fmt.Errorf("Test virtual edition cluster (%s) still exists", *result.Name)
	}
	return nil
}

func testAccVirtualEditionClusterExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.clusterVip,
			cohesityConfig.clusterUsername, cohesityConfig.clusterPassword, cohesityConfig.clusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}
		result, infoErr := client.Cluster().GetBasicClusterInfo()
		if infoErr == nil && result.Name == nil {
			return fmt.Errorf("Failed to create the test virtual edition cluster")
		}
		var fetchStats, fetchTimeSeriesSchema = false, false
		clusterDetails, clusterDetailsErr := client.Cluster().GetCluster(&fetchStats, &fetchTimeSeriesSchema)
		if clusterDetailsErr == nil {
			if *clusterDetails.Name != "AcceptanceTestTerraformVirtaulEditionCluster" {

				return fmt.Errorf("Failed to validate test virtual edition cluster name")
			}
		}
		return nil
	}
}

const testAccVirtualEditionClusterConfig = `
provider "cohesity" {
	cluster_vip = "10.2.145.27"
	cluster_username = "admin"
	cluster_domain = "LOCAL"
}

resource "cohesity_virtual_edition_cluster" "virtual"{
		cluster_name = "AcceptanceTestTerraformVirtaulEditionCluster"
		dns_servers = ["10.2.145.14"]
		ntp_servers = ["time.google.com"]
		domain_names = ["eng.cohesity.com"]
		cluster_subnet_mask = "255.255.240.0"
		cluster_gateway = "10.2.144.1"
		enable_encryption = true
		enable_fips_mode = true
		encryption_keys_rotation_period = 1
		metadata_fault_tolerance = 0
		virtual_ips = ["10.2.145.27"]
		virtual_ip_hostname = "test"
		node_ips = ["10.2.145.27"]
}
`