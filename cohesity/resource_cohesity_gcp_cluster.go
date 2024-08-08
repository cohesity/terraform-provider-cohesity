package cohesity

import (
    "context"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "time"
)

func resourceCohesityGcpCluster() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceCohesityGcpClusterCreate,
        ReadContext:   resourceCohesityGcpClusterRead,
        DeleteContext: resourceCohesityGcpClusterDelete,
        UpdateContext: resourceCohesityGcpClusterUpdate,

        Schema: map[string]*schema.Schema{
            "cluster_config_file": {
                Type:     schema.TypeString,
                Required: true,
            },
            "gcp_key_file": {
                Type:     schema.TypeString,
                Required: true,
            },
            "gcp_bucket_config_file": {
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceCohesityGcpClusterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics

    config := m.(CVMProviderConfig)
    if config.Type != "ssh" {
        return diag.Errorf("resourceCohesityGcpClusterCreate expects type 'ssh', but got '%s'", config.Type)
	}
    
    host := config.Host
    user := config.User
    password := config.Password
    timeout, _ := time.ParseDuration(config.Timeout)

    clusterConfigFile := d.Get("cluster_config_file").(string)
    GCPKeyFile := d.Get("gcp_key_file").(string)
    GCPBucketConfigFile := d.Get("gcp_bucket_config_file").(string)

    client, err := newSSHClient(host, user, password, timeout)
    if err != nil {
        return diag.FromErr(err)
    }
    defer client.Close()

    err = transferFile(client, clusterConfigFile, "/home/cohesity/data/gcp/cluster_config.json")
    if err != nil {
        return diag.FromErr(err)
    }

    err = transferFile(client, GCPKeyFile, "/home/cohesity/data/gcp/GCPKey.json")
    if err != nil {
        return diag.FromErr(err)
    }
	err = transferFile(client, GCPBucketConfigFile, "/home/cohesity/data/gcp/gcp_bucket_config")
    if err != nil {
        return diag.FromErr(err)
    }
    cmd := "export PATH=/usr/local/bin:/usr/bin:/usr/local/sbin:/usr/sbin:/home/cohesity/.local/bin:/home/cohesity/bin:/home/cohesity/software/bin:/home/cohesity/software/bin:/home/cohesity/software/bin && " +
           "/home/cohesity/software/bin/cohesity_gcp_setup create_cluster -cohesity_gcp_params_file /home/cohesity/data/gcp/cluster_config.json -cohesity_gcp_machine_type n1-standard-32 -cohesity_gcp_replication_factor 2"
    
    err = executeSSHCommand(client, cmd)
    if err != nil {
        return diag.FromErr(err)
    }
    d.SetId(fmt.Sprintf("%s-%s", host, time.Now().String()))

    return diags
}

func resourceCohesityGcpClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    // No read implementation necessary for this example
    return diags
}

func resourceCohesityGcpClusterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    // No delete implementation necessary for this example
    return diags
}
func resourceCohesityGcpClusterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    // No delete implementation necessary for this example
    return diags
}
