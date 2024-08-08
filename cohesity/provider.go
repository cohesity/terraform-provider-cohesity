package cohesity

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "context"
    // "log"
    "github.com/hashicorp/terraform-plugin-log/tflog"
)
type CVMProviderConfig struct {
	Type       string
    Host       string
    User       string
    Password   string
    PrivateKey string
    Timeout    string
}
// Provider represents a resource provider in terraform
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the schema set ('cluster', 'ssh')",
			},
			"cluster_vip": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_IP", ""),
				Description: "IP or hostname of Cohesity cluster node",
			},
			"cluster_username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_USERNAME", ""),
				Description: "Cohesity cluster username",
			},
			"cluster_password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("COHESITY_PASSWORD", ""),
				Description: "Cohesity cluster password",
			},
			"cluster_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "LOCAL",
				Description: "The domain name of cohesity user",
			},
			"ssh_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Host for SSH connection",
			},
			"ssh_user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User for SSH connection",
			},
			"ssh_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password for SSH connection",
			},
			"ssh_timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "60m",
				Description: "Timeout for SSH connection",
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
			"cohesity_gcp_cluster": 			 resourceCohesityGcpCluster(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	typeField, ok := d.GetOk("type")
	if !ok {
		return nil, diag.Errorf("Missing required type field")
	}
	switch typeField {
	case "cluster":
		vip, ok1 := d.GetOk("cluster_vip")
		username, ok2 := d.GetOk("cluster_username")
		password, ok3 := d.GetOk("cluster_password")
		domain, _ := d.GetOk("cluster_domain")

		if !ok1 || !ok2 || !ok3 {
			return nil, diag.Errorf("Missing required fields for cohesity type")
		}

		return Config{
			Type:			   typeField.(string),
			clusterVip:        vip.(string),
			clusterUsername:   username.(string),
			clusterPassword:   password.(string),
			clusterDomain:     domain.(string),
		}, nil
	case "ssh":
		host, _ := d.GetOk("ssh_host")
		user, ok2 := d.GetOk("ssh_user")
		password, _ := d.GetOk("ssh_password")
		timeout, _ := d.GetOk("ssh_timeout")

		if !ok2 {
			return nil, diag.Errorf("Missing user required fields for ssh type")
		}
		tflog.Info(ctx,"Successfully configured ssh provider "+host.(string))
		return CVMProviderConfig{
			Type:		typeField.(string),
			Host:       host.(string),
			User:       user.(string),
			Password: password.(string),
			Timeout:    timeout.(string),
		}, nil
	default:
		return nil, diag.Errorf("unknown type: %s", typeField)
	}
}