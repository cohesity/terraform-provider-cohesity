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

func resourceCohesityGCPExternalTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCohesityGCPExternalTargetCreate,
		ReadContext:   resourceCohesityGCPExternalTargetRead,
		UpdateContext: resourceCohesityGCPExternalTargetUpdate,
		DeleteContext: resourceCohesityGCPExternalTargetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
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

func resourceCohesityGCPExternalTargetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// clusterIp := d.Get("cluster_vip").(string)
	log.Printf("[INFO] starting create process")
	privateKeyFilePath := d.Get("client_private_key_file_path").(string)
	bucketName := d.Get("bucket_name").(string)
	projectId := d.Get("project_id").(string)
	clientEmail := d.Get("client_email").(string)
	tierType := d.Get("tier_type").(string)
	config := m.(utils.Config)
	clusterPassword := config.ClusterPassword
	supportPassword := config.SupportPassword
	clusterUsername := config.ClusterUsername
	clusterIp := config.ClusterVIP
	log.Printf("[INFO] pass: %s", clusterPassword)
	clusterPassword = utils.EscapeSpecialSymbols(clusterPassword)
	log.Printf("[INFO] pass: %s", clusterPassword)
	log.Printf("[INFO] cluster ip: %s", clusterIp)

	client, err := services.NewSSHClient(clusterIp, "support", supportPassword, time.Hour)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO] starting transfer process")

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
func resourceCohesityGCPExternalTargetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
func resourceCohesityGCPExternalTargetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
func resourceCohesityGCPExternalTargetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
