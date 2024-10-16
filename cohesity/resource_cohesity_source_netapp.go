package cohesity

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/services"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func resourceCohesitySourceNetapp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCohesitySourceNetappCreate,
		ReadContext:   resourceCohesitySourceNetappRead,
		UpdateContext: resourceCohesitySourceNetappUpdate,
		DeleteContext: resourceCohesitySourceNetappDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"endpoint": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			"source_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.IsValidNetappSourceType,
			},
			"backup_smb_volumes": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"smb_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"smb_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"allowed_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsIPAddress,
				},
				DiffSuppressFunc: utils.SuppressNodeIPsDiff,
			},
			"disallowed_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsIPAddress,
				},
				DiffSuppressFunc: utils.SuppressNodeIPsDiff,
			},
		},
	}
}

func resourceCohesitySourceNetappCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}
	sourceNetapp := services.NewSourceNetapp(d.Get("username").(string), d.Get("password").(string), d.Get("endpoint").(string), d.Get("source_type").(string), d.Get("backup_smb_volumes").(bool))
	if d.Get("backup_smb_volumes").(bool) {
		// Set SMB credentials if backupSMBVolumes is true and credentials are not already set
		if d.Get("smb_username").(string) == "" || d.Get("smb_password").(string) == "" {
			return diag.Errorf("SMB credentials are required when backupSMBVolumes is true")
		}
		sourceNetapp = sourceNetapp.WithSMBCredentials(d.Get("smb_username").(string), d.Get("smb_password").(string))
	}
	if len(d.Get("allowed_ips").([]string)) != 0 {
		sourceNetapp = sourceNetapp.WithAllowedIPAddresses(d.Get("allowed_ips").([]string))
	}
	if len(d.Get("disallowed_ips").([]string)) != 0 {
		sourceNetapp = sourceNetapp.WithDeniedIPAddresses(d.Get("disallowed_ips").([]string))
	}
	sourceId, err := client.AddNetappProtectionSource(sourceNetapp)
	if err != nil {
		return diag.Errorf("Failed to create a Netapp source with the given details, %s", err.Error())
	}
	d.SetId(strconv.FormatInt(sourceId, 10))
	return nil
}
func resourceCohesitySourceNetappRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}
	sourceInfo, err := client.GetNetAppSource(d.Id())
	if err != nil {
		return diag.Errorf("Failed to get details of the Netapp source, %s", err.Error())
	}
	netappInfo := sourceInfo.NetappParams
	d.Set("username", *netappInfo.Credentials.Username)
	d.Set("source_type", *netappInfo.SourceType)
	d.Set("endpoint", *netappInfo.Endpoint)
	d.Set("backup_smb_volumes", *netappInfo.BackUpSMBVolumes)
	if *netappInfo.BackUpSMBVolumes {
		d.Set("smb_username", *netappInfo.SmbCredentials.Username)
	}
	d.Set("allowed_ips", netappInfo.FilterIPConfig.AllowedIPAddresses)
	d.Set("disallowed_ips", netappInfo.FilterIPConfig.DeniedIPAddresses)
	return nil
}
func resourceCohesitySourceNetappUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChanges("username", "source_type", "endpoint", "backup_smb_volumes", "smb_username", "allowed_ips", "disallowed_ips") {
		config := m.(utils.Config)
		client, err := services.NewCohesityClientV2(config)
		if err != nil {
			return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
		}
		sourceNetapp := services.NewSourceNetapp(d.Get("username").(string), d.Get("password").(string), d.Get("endpoint").(string), d.Get("source_type").(string), d.Get("backup_smb_volumes").(bool))
		if d.Get("backup_smb_volumes").(bool) {
			// Set SMB credentials if backupSMBVolumes is true and credentials are not already set
			if d.Get("smb_username").(string) == "" || d.Get("smb_password").(string) == "" {
				return diag.Errorf("SMB credentials are required when backupSMBVolumes is true")
			}
			sourceNetapp = sourceNetapp.WithSMBCredentials(d.Get("smb_username").(string), d.Get("smb_password").(string))
		}
		if len(d.Get("allowed_ips").([]string)) != 0 {
			sourceNetapp = sourceNetapp.WithAllowedIPAddresses(d.Get("allowed_ips").([]string))
		}
		if len(d.Get("disallowed_ips").([]string)) != 0 {
			sourceNetapp = sourceNetapp.WithDeniedIPAddresses(d.Get("disallowed_ips").([]string))
		}
		_, err = client.UpdateNetAppSource(sourceNetapp)
		if err != nil {
			return diag.Errorf("Failed to update the Netapp source with the given details, %s", err.Error())
		}
	}
	return nil
}
func resourceCohesitySourceNetappDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV1(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}
	err = client.DeleteProtectionSource(d.Id())
	if err != nil {
		return diag.Errorf("Failed to delete the Netapp source, %s", err.Error())
	}
	return nil
}
