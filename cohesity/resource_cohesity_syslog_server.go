package cohesity

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/services"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func resourceCohesitySyslogServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCohesitySyslogServerCreate,
		ReadContext:   resourceCohesitySyslogServerRead,
		UpdateContext: resourceCohesitySyslogServerUpdate,
		DeleteContext: resourceCohesitySyslogServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			"port": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 65535),
			},
			"protocol": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"tcp", "udp"}, false),
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ca_certificate": {
				Type:     schema.TypeString,
				Required: true,
			},
			"program_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "List of program names.",
			},
		},
	}
}

func resourceCohesitySyslogServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	syslogServerParam := services.BuildSyslogServerParam(d, true)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer token with the given details, %s", err.Error())
	}
	url, err := services.NormalizeHTTPSAddress(config.ClusterVIP)
	if err != nil {
		return diag.Errorf("Failed to get a url for cluster with the given details, %s", err.Error())
	}

	syslogServer, err := services.CreateSyslogServer(url, token, syslogServerParam)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO] Successfully created syslog server: %+v, id: %d", syslogServer, *syslogServer.ID)
	idInt64 := int64(*syslogServer.ID)
	idString := strconv.FormatInt(idInt64, 10)
	d.SetId(idString)
	return resourceCohesitySyslogServerRead(ctx, d, m)
}

func resourceCohesitySyslogServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer token with the given details, %s", err.Error())
	}
	url, err := services.NormalizeHTTPSAddress(config.ClusterVIP)
	if err != nil {
		return diag.Errorf("Failed to get a url for cluster with the given details, %s", err.Error())
	}

	syslogServer, err := services.GetSyslogServer(url, token, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", syslogServer.Name)
	d.Set("ip_address", syslogServer.IP)
	d.Set("port", syslogServer.Port)
	d.Set("protocol", syslogServer.Protocol)
	d.Set("is_enabled", syslogServer.Enabled)
	d.Set("is_tls_enabled", syslogServer.IsTLSEnabled)
	d.Set("ca_certificate", syslogServer.CaCertificate)
	d.Set("program_names", syslogServer.ProgramNameList)
	idInt64 := int64(*syslogServer.ID)
	idString := strconv.FormatInt(idInt64, 10)
	d.SetId(idString)
	return nil
}

func resourceCohesitySyslogServerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer token with the given details, %s", err.Error())
	}
	url, err := services.NormalizeHTTPSAddress(config.ClusterVIP)
	if err != nil {
		return diag.Errorf("Failed to get a url for cluster with the given details, %s", err.Error())
	}

	syslogServerParam := services.BuildSyslogServerParam(d, false)

	_, err = services.UpdateSyslogServer(url, token, d.Id(), syslogServerParam)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceCohesitySyslogServerRead(ctx, d, m)
}

func resourceCohesitySyslogServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}

	err = client.DeleteSyslogServer(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
