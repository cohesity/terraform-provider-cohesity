package cohesity

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider represents a resource provider in terraform
func Provider() *schema.Provider {
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
			"support_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "support password for the Cohesity cluster",
			},
			"cluster_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "LOCAL",
				Description: "The domain name of cohesity user",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cohesity_source_vmware":       resourceCohesitySourceVMware(),
			"cohesity_job_vmware":          resourceCohesityJobVMware(),
			"cohesity_job_run":             resourceCohesityJobRun(),
			"cohesity_ngce_cluster":        resourceCohesityNGCECluster(),
			"cohesity_gcp_external_target": resourceGCPExternalTarget(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// providerConfigure reads the configuration and returns it as a Config struct
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Retrieve the configuration values
	config := Config{
		ClusterVIP:      d.Get("cluster_vip").(string),
		ClusterUsername: d.Get("cluster_username").(string),
		ClusterPassword: d.Get("cluster_password").(string),
		SupportPassword: d.Get("support_password").(string),
		ClusterDomain:   d.Get("cluster_domain").(string),
	}
	log.Printf("[INFO] vip: %s", config.ClusterVIP)

	// Validate mandatory fields
	if config.ClusterUsername == "" {
		return nil, diag.Errorf("cluster_username is required")
	}

	if config.ClusterPassword == "" {
		return nil, diag.Errorf("cluster_password is required")
	}

	return config, diags
}
