package cohesity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

// Provider represents a resource provider in terraform
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint":                  {Type: schema.TypeString, Optional: true, Description: "HTTPS origin for certificate management. Defaults to https://cluster_vip."},
			"api_key":                   {Type: schema.TypeString, Optional: true, Sensitive: true, DefaultFunc: schema.EnvDefaultFunc("COHESITY_API_KEY", ""), Description: "API key with CLUSTER_VIEW and CLUSTER_MODIFY for web server certificate management. Other resources require username/password authentication."},
			"ca_certificate_pem":        {Type: schema.TypeString, Optional: true, Description: "Additional trusted CA certificates for certificate management."},
			"bootstrap_certificate_pem": {Type: schema.TypeString, Optional: true, Description: "PEM bundle of explicitly verified existing server leaves, accepted only when normal CA validation fails."},
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
				Sensitive:   true,
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
			"cohesity_web_server_certificate":       resourceWebServerCertificate(),
			"cohesity_source_vmware":                resourceCohesitySourceVMware(),
			"cohesity_job_vmware":                   resourceCohesityJobVMware(),
			"cohesity_job_run":                      resourceCohesityJobRun(),
			"cohesity_ngce_cluster":                 resourceCohesityNGCECluster(),
			"cohesity_gcp_external_target":          resourceCohesityGCPExternalTarget(),
			"cohesity_source_netapp":                resourceCohesitySourceNetapp(),
			"cohesity_source_gcp":                   resourceCohesitySourceGcp(),
			"cohesity_configure_snmp":               resourceCohesityConfigureSnmp(),
			"cohesity_configure_smtp":               resourceCohesityConfigureSmtp(),
			"cohesity_syslog_server":                resourceCohesitySyslogServer(),
			"cohesity_configure_notification_rules": resourceCohesityConfigureNotificationRules(),
			"cohesity_backup_policy":                resourceCohesityBackupPolicy(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// providerConfigure reads the configuration and returns it as a Config struct
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Retrieve the configuration values
	config := utils.Config{
		Endpoint:                d.Get("endpoint").(string),
		APIKey:                  d.Get("api_key").(string),
		CACertificatePEM:        d.Get("ca_certificate_pem").(string),
		BootstrapCertificatePEM: d.Get("bootstrap_certificate_pem").(string),
		ClusterVIP:              d.Get("cluster_vip").(string),
		ClusterUsername:         d.Get("cluster_username").(string),
		ClusterPassword:         d.Get("cluster_password").(string),
		SupportPassword:         d.Get("support_password").(string),
		ClusterDomain:           d.Get("cluster_domain").(string),
	}

	if config.ClusterVIP == "" && config.Endpoint == "" {
		return nil, diag.Errorf("set cluster_vip or endpoint")
	}
	if config.Endpoint == "" {
		config.Endpoint = "https://" + config.ClusterVIP
	}
	if config.APIKey == "" && (config.ClusterUsername == "" || config.ClusterPassword == "") {
		return nil, diag.Errorf("cluster_username and cluster_password are required without an api_key")
	}
	return config, diags
}
