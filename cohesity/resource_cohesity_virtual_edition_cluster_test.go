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

func TestAccVirtualEditionCluster_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVirtualEditionClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccVirtualEditionClusterConfig, virtualEditionDNSServer,
					virtualEditionDomainName, virtualEditionClusterSubnetMask, virtualEditionClusterGateway,
					virtualEditionClusterVip, virtualEditionClusterVip, virtualEditionClusterNodeIP),
				Check: resource.ComposeTestCheckFunc(
					testAccVirtualEditionClusterExists(),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "cluster_name",
						"AcceptanceTestTerraformVirtaulEditionCluster"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "enable_encryption", "true"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "cluster_gateway", virtualEditionClusterGateway),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "enable_fips_mode", "true"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "metadata_fault_tolerance", "0"),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "virtual_ip_hostname", virtualEditionClusterVip),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "cluster_subnet_mask", virtualEditionClusterSubnetMask),
					resource.TestCheckResourceAttr("cohesity_virtual_edition_cluster.virtual", "encryption_keys_rotation_period", "1"),
				),
			},
		},
	})
}

func testAccCheckVirtualEditionClusterDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.ClusterVIP,
		cohesityConfig.ClusterUsername, cohesityConfig.ClusterPassword, cohesityConfig.ClusterDomain)
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
		client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.ClusterVIP,
			cohesityConfig.ClusterUsername, cohesityConfig.ClusterPassword, cohesityConfig.ClusterDomain)
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
resource "cohesity_virtual_edition_cluster" "virtual"{
		cluster_name = "AcceptanceTestTerraformVirtaulEditionCluster"
		dns_servers = ["%s"]
		ntp_servers = ["time.google.com"]
		domain_names = ["%s"]
		cluster_subnet_mask = "%s"
		cluster_gateway = "%s"
		enable_encryption = true
		enable_fips_mode = true
		encryption_keys_rotation_period = 1
		metadata_fault_tolerance = 0
		virtual_ips = ["%s"]
		virtual_ip_hostname = "%s"
		node_ips = ["%s"]
}
`
