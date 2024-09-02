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

func TestAccPhysicalEditionCluster_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPhysicalEditionClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccPhysicalEditionClusterConfig, physicalEditionDNSServer,
					physicalEditionDomainName, physicalEditionClusterSubnetMask, physicalEditionClusterGateway,
					physicalEditionClusterVip, ipmiGateway, ipmiSubnetMask, physicalEditionClusterVip,
					ipmiUsername, physicalEditionClusterNodeIP1, physicalEditionClusterNodeIpmiIP1,
					physicalEditionClusterNodeIP2, physicalEditionClusterNodeIpmiIP2, physicalEditionClusterNodeIP3,
					physicalEditionClusterNodeIpmiIP3),
				Check: resource.ComposeTestCheckFunc(
					testAccPhysicalEditionClusterExists(),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "cluster_name",
						"AcceptanceTestTerraformPhysicalEditionCluster"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "enable_encryption", "true"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "cluster_gateway", physicalEditionClusterGateway),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "enable_fips_mode", "true"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "metadata_fault_tolerance", "0"),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "virtual_ip_hostname", physicalEditionClusterVip),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "cluster_subnet_mask", physicalEditionClusterSubnetMask),
					resource.TestCheckResourceAttr("cohesity_physical_edition_cluster.physical", "encryption_keys_rotation_period", "1"),
				),
			},
		},
	})
}

func testAccCheckPhysicalEditionClusterDestroy(s *terraform.State) error {
	var cohesityConfig = testAccProvider.Meta().(Config)
	client, err := CohesityManagementSdk.NewCohesitySdkClient(cohesityConfig.ClusterVIP,
		cohesityConfig.ClusterUsername, cohesityConfig.ClusterPassword, cohesityConfig.ClusterDomain)
	if err != nil {
		log.Printf(err.Error())
		return fmt.Errorf("Failed to authenticate with Cohesity")
	}
	result, infoErr := client.Cluster().GetBasicClusterInfo()
	if infoErr == nil && result.Name != nil {
		return fmt.Errorf("Test physical edition cluster (%s) still exists", *result.Name)
	}
	return nil
}

func testAccPhysicalEditionClusterExists() resource.TestCheckFunc {
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
			return fmt.Errorf("Failed to create the test physical edition cluster")
		}
		var fetchStats, fetchTimeSeriesSchema = false, false
		clusterDetails, clusterDetailsErr := client.Cluster().GetCluster(&fetchStats, &fetchTimeSeriesSchema)
		if clusterDetailsErr == nil {
			if *clusterDetails.Name != "AcceptanceTestTerraformPhysicalEditionCluster" {

				return fmt.Errorf("Failed to validate test physical edition cluster name")
			}
		}
		return nil
	}
}

const testAccPhysicalEditionClusterConfig = `
resource "cohesity_physical_edition_cluster" "physical"{
		cluster_name = "AcceptanceTestTerraformPhysicalEditionCluster"
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
		ipmi_gateway = "%s"
		ipmi_subnet_mask = "%s"
		virtual_ip_hostname = "%s"
		ipmi_username="%s"
		node_configs {
					node_ip="%s"
					node_ipmi_ip="%s"
					}
		node_configs {
					node_ip="%s"
					node_ipmi_ip="%s"
					}
		node_configs {
					node_ip="%s"
					node_ipmi_ip="%s"
					}
}
`
