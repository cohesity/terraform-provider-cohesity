package cohesity

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"ssh": {
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
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
				}},
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
			"cohesity_gcp_cluster":              resourceCohesityGcpCluster(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	vip, ok := d.GetOk("cluster_vip")

	if ok {
		username, ok2 := d.GetOk("cluster_username")
		password, ok3 := d.GetOk("cluster_password")
		domain, _ := d.GetOk("cluster_domain")

		if !ok2 || !ok3 {
			return nil, diag.Errorf("Missing cluster_username (required fields)")
		}
		if !ok3 {
			return nil, diag.Errorf("Missing cluster_password (required fields)")
		}

		return Config{
			clusterVip:      vip.(string),
			clusterUsername: username.(string),
			clusterPassword: password.(string),
			clusterDomain:   domain.(string),
		}, nil
	}
	if v, ok := d.GetOk("ssh"); ok{
		sshConfig := v.([]interface{})[0].(map[string]interface{})
		host, _ := sshConfig["ssh_host"].(string)
		user, ok1 := sshConfig["ssh_user"].(string)
		password, ok2 := sshConfig["ssh_password"].(string)
		timeout, _ := sshConfig["ssh_timeout"].(string)
		if !ok1 || user == "" {
			return nil, diag.Errorf("Missing ssh_username (required field) for ssh type")
		}
		if !ok2 || password == "" {
			return nil, diag.Errorf("Missing ssh_password (required field) for ssh type")
		}
		tflog.Info(ctx, "Successfully configured ssh provider "+host)
		return CVMProviderConfig{
			Host:     host,
			User:     user,
			Password: password,
			Timeout:  timeout,
		}, nil
	}

	return nil, diag.Errorf("Fields incorrect or missing")
}
