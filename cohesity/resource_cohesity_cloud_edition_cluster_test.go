package cohesity

import (
	"errors"
	"fmt"
	"log"
	"testing"

	CohesityManagementSdk "github.com/cohesity/management-sdk-go/managementsdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudEditionCluster_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudEditionClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCloudEditionClusterConfig, cloudEditionDNSServer,
					cloudEditionDomainName, cloudEditionClusterSubnetMask, cloudEditionClusterGateway,
					cloudEditionClusterNodeIP1, cloudEditionClusterNodeIP2, cloudEditionClusterNodeIP3),
				Check: resource.ComposeTestCheckFunc(
					testAccCloudEditionClusterExists(),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "cluster_name",
						"AcceptanceTestTerraformCloudEditionCluster"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "enable_encryption", "true"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "cluster_gateway", cloudEditionClusterGateway),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "enable_fips_mode", "true"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "metadata_fault_tolerance", "0"),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "cluster_subnet_mask", cloudEditionClusterSubnetMask),
					resource.TestCheckResourceAttr("cohesity_cloud_edition_cluster.cloud", "encryption_keys_rotation_period", "1"),
				),
			},
		},
	})
}

func testAccCheckCloudEditionClusterDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.ClusterVIP,
		cohesityConfig.ClusterUsername, cohesityConfig.ClusterPassword, cohesityConfig.ClusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}
	result, infoErr := client.Cluster().GetBasicClusterInfo()
	if infoErr == nil && result.Name != nil {
		return fmt.Errorf("Test cloud edition cluster (%s) still exists", *result.Name)
	}
	return nil
}

func testAccCloudEditionClusterExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var cohesityConfig = testAccProvider.Meta().(Config)
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.ClusterVIP,
			cohesityConfig.ClusterUsername, cohesityConfig.ClusterPassword, cohesityConfig.ClusterDomain)
		if err != nil {
			log.Printf(err.Error())
			return errors.New("Failed to authenticate with Cohesity")
		}
		result, infoErr := client.Cluster().GetBasicClusterInfo()
		if infoErr == nil && result.Name == nil {
			return fmt.Errorf("Failed to create the test cloud edition cluster")
		}
		var fetchStats, fetchTimeSeriesSchema = false, false
		clusterDetails, clusterDetailsErr := client.Cluster().GetCluster(&fetchStats, &fetchTimeSeriesSchema)
		if clusterDetailsErr == nil {
			if *clusterDetails.Name != "AcceptanceTestTerraformCloudEditionCluster" {

				return fmt.Errorf("Failed to validate test cloud edition cluster name")
			}
		}
		return nil
	}
}

const testAccCloudEditionClusterConfig = `
resource "cohesity_cloud_edition_cluster" "cloud"{
		cluster_name = "AcceptanceTestTerraformCloudEditionCluster"
		dns_servers = ["%s"]
		ntp_servers = ["time.google.com"]
		domain_names = ["%s"]
		cluster_subnet_mask = "%s"
		cluster_gateway = "%s"
		enable_encryption = true
		enable_fips_mode = true
		encryption_keys_rotation_period = 1
		metadata_fault_tolerance = 0
		node_ips = ["%s", "%s", "%s"]        
}
`
