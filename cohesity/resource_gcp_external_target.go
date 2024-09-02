package cohesity

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/services"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func resourceGCPExternalTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGCPExternalTargetCreate,
		ReadContext:   resourceGCPExternalTargetRead,
		UpdateContext: resourceGCPExternalTargetUpdate,
		DeleteContext: resourceGCPExternalTargetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"cluster_vip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_private_key_file_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tier_type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGCPExternalTargetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clusterIp := d.Get("cluster_vip").(string)
	privateKeyFilePath := d.Get("client_private_key_file_path").(string)
	bucketName := d.Get("bucket_name").(string)
	projectId := d.Get("project_id").(string)
	clientEmail := d.Get("client_email").(string)
	tierType := d.Get("tier_type").(string)
	config := m.(Config)
	clusterPassword := config.ClusterPassword
	supportPassword := config.SupportPassword
	clusterUsername := config.ClusterUsername
	log.Printf("[INFO] pass: %s", clusterPassword)
	clusterPassword = utils.EscapeSpecialSymbols(clusterPassword)
	log.Printf("[INFO] pass: %s", clusterPassword)

	client, err := services.NewSSHClient(clusterIp, "support", supportPassword, time.Hour)
	if err != nil {
		return diag.FromErr(err)
	}
	err = services.TransferFile(client, privateKeyFilePath, "/tmp/gcp_bucket_config")
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO] transferred file")
	cmd := fmt.Sprintf("/home/support/bin/iris_cli --output=prettyjson --username=%s --password=%s --skip_password_prompt=true --skip_force_password_change=true vault create name=DefaultGCPExternalTarget usage-type=archival compression-enabled=true encryption-enabled=true vault-type=Google bucket-name=%s goog-project-id=%s goog-client-email=%s goog-client-priv-key-file-loc=/tmp/gcp_bucket_config tier-type=%s forever-incremental-enabled=true dedup-enabled=true view-box-name=DefaultStorageDomain; status=$?; sleep 30; exit $status", clusterUsername, clusterPassword, bucketName, projectId, clientEmail, tierType)
	log.Printf("[INFO] cmd: %s", cmd)
	output, err := services.ExecuteSSHCommand(client, cmd)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO] output: %s", output)
	d.SetId(bucketName)
	return nil
}
func resourceGCPExternalTargetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
func resourceGCPExternalTargetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
func resourceGCPExternalTargetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
