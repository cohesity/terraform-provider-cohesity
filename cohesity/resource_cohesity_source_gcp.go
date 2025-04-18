package cohesity

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/services"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func resourceCohesitySourceGcp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCohesitySourceGcpCreate,
		ReadContext:   resourceCohesitySourceGcpRead,
		UpdateContext: resourceCohesitySourceGcpUpdate,
		DeleteContext: resourceCohesitySourceGcpDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"client_email_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_private_key": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vpc_network": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: utils.SuppressNetworkNameDiff,
			},
			"vpc_subnetwork": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: utils.SuppressNetworkNameDiff,
			},
		},
	}
}

func resourceCohesitySourceGcpCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer token with the given details, %s", err.Error())
	}
	log.Printf("[INFO]token:%v",token)
	sourceId, err := services.AddGCPProtectionSource(config.ClusterVIP, token, d.Get("client_email_address").(string), d.Get("client_private_key").(string), d.Get("project_id").(string), d.Get("vpc_network").(string), d.Get("vpc_subnetwork").(string))
	if err != nil {
		return diag.Errorf("Failed to create a GCP source with the given details, %s", err.Error())
	}
	d.SetId(strconv.FormatInt(sourceId, 10))
	return nil
}
func resourceCohesitySourceGcpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer toke with the given details, %s", err.Error())
	}
	sourceInfo, err := services.GetGcpProtectionSource(config.ClusterVIP, token, d.Id())
	if err != nil {
		return diag.Errorf("Failed to get details of the Gcp source, %s", err.Error())
	}
	gcpInfo := sourceInfo.Entity.GcpEntity
	d.Set("client_email_address", gcpInfo.OwnerID)
	d.Set("project_id", gcpInfo.ProjectID)
	d.Set("vpc_network", gcpInfo.VpcNetwork)
	d.Set("vpc_subnetwork", gcpInfo.VpcSubnetwork)
	return nil
}
func resourceCohesitySourceGcpUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChange( "project_id"){
		return diag.Errorf("changes to project Id are not allowed")
	}
	if d.HasChanges("client_email_address", "vpc_network", "vpc_subnetwork", "client_private_key") {
		config := m.(utils.Config)
		token, err := services.GetAccessToken(config)
		if err != nil {
			return diag.Errorf("Failed to get a bearer toke with the given details, %s", err.Error())
		}
		err = services.UpdateGcpProtectionSource(config.ClusterVIP, token, d.Id(), d.Get("client_email_address").(string), d.Get("client_private_key").(string),d.Get("project_id").(string), d.Get("vpc_network").(string), d.Get("vpc_subnetwork").(string))
		if err != nil {
			return diag.Errorf("Failed to update the GCP source with the given details, %s", err.Error())
		}
		return nil
	}
	return nil
}
func resourceCohesitySourceGcpDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer token with the given details, %s", err.Error())
	}
	err = services.DeleteProtectionSource(config.ClusterVIP, token, d.Id())
	if err != nil {
		return diag.Errorf("Failed to delete the GCP source, %s", err.Error())
	}
	return nil
}
