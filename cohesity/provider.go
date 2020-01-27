package cohesity

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider represents a resource provider in terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cluster_vip": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_IP", ""),
				Description: "IP or hostname of Cohesity cluster node",
			},
			"cluster_username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_USERNAME", ""),
				Description: "Cohesity cluster username",
			},
			"cluster_password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_PASSWORD", ""),
				Description: "Cohesity cluster password",
			},
			"cluster_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "LOCAL",
				Description: "The domain name of cohesity user",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cohesity_cloud_edition_cluster":    resourceCohesityCloudEditionCluster(),
			"cohesity_virtual_edition_cluster":  resourceCohesityVirtualEditionCluster(),
			"cohesity_physical_edition_cluster": resourceCohesityPhysicalEditionCluster(),
			"cohesity_source_vmware":            resourceCohesitySourceVMware(),
			"cohesity_job_vmware":               resourceCohesityJobVMware(),
			"cohesity_job_run":                  resourceCohesityJobRun(),
			"cohesity_restore_vmware_vm":        resourceCohesityRestoreVMwareVM(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		clusterVip:      d.Get("cluster_vip").(string),
		clusterUsername: d.Get("cluster_username").(string),
		clusterPassword: d.Get("cluster_password").(string),
		clusterDomain:   d.Get("cluster_domain").(string),
	}
	return config, nil
}
