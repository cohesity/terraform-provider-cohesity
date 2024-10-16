package cohesity

import (
	"context"
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
			},
		},
	}
}

func resourceCohesitySourceGcpCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV1(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}
	sourceId, err := client.AddGCPProtectionSource(d.Get("client_email_address").(string), d.Get("client_private_key").(string), d.Get("project_id").(string), d.Get("vpc_network").(string), d.Get("vpc_subnetwork").(string))
	if err != nil {
		return diag.Errorf("Failed to create a GCP source with the given details, %s", err.Error())
	}
	d.SetId(strconv.FormatInt(sourceId, 10))
	return nil
}
func resourceCohesitySourceGcpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV1(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}
	sourceInfo, err := client.GetGcpProtectionSource(d.Id())
	if err != nil {
		return diag.Errorf("Failed to get details of the Gcp source, %s", err.Error())
	}
	gcpInfo := sourceInfo.Entity.GcpEntity
	d.Set("client_email_address", *gcpInfo.OwnerID)
	d.Set("project_id", *gcpInfo.ProjectID)
	d.Set("vpc_network", *gcpInfo.VpcNetwork)
	d.Set("vpc_subnetwork", *gcpInfo.VpcSubnetwork)
	return nil
}
func resourceCohesitySourceGcpUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// config := m.(utils.Config)
	// client, err := services.NewCohesityClientV1(config)
	// if err != nil {
	// 	return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	// }
	// not implemented yet
	if d.HasChanges("client_email_address", "project_id", "vpc_network", "vpc_subnetwork", "client_private_key") {
		return diag.Errorf("changes are not yet allowed, to be implemented")
	}
	return nil
}
func resourceCohesitySourceGcpDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV1(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}
	err = client.DeleteProtectionSource(d.Id())
	if err != nil {
		return diag.Errorf("Failed to delete the GCP source, %s", err.Error())
	}
	return nil
}
