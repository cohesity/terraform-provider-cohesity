package cohesity

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/services"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func resourceCohesityConfigureSnmp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCohesityConfigureSnmpCreate,
		ReadContext:   resourceCohesityConfigureSnmpRead,
		UpdateContext: resourceCohesityConfigureSnmpUpdate,
		DeleteContext: resourceCohesityConfigureSnmpDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"agent_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     161,
				Description: "AgentPort is the TCP port SNMP agent is using.",
			},
			"read_user": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem:        userSchema(),
				Description: "SNMP Read User Info for this cluster",
			},
			"server": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Server is the IP address of Network Management System.",
			},
			"trap_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     162,
				Description: "TrapPort is the TCP port SNMP agent is using.",
			},
			"trap_user": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem:        userSchema(),
				Description: "SNMP Trap User Info for this cluster",
			},
			"snmp_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "V2C",
				Description: "SnmpVersion is the SNMP version to talk with SNMP agent. Options: V2C, V3",
			},
			"write_user": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem:        userSchema(),
				Description: "SNMP Write User Info for this cluster",
			},
		},
	}
}

func userSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AuthPassword is the authentication password for SNMP V3 users.",
			},
			"auth_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AuthPrototol is the authentication protocol for SNMP V3 users. Options are None, MD5, or SHA. Defaults to None",
			},
			"priv_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "PrivPassword is the privacy password for SNMP V3 users.",
			},
			"priv_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "PrivPrototol is the privacy protocol for SNMP V3 users. Options are None, DES, or AES. Defaults to None",
			},
			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "UserName is the user name to access SNMP V2 or SNMP V3 agent. Default is cohesityV2Public for SNMP Version V2C and cohesityV3Public for SNMP Version V3",
			},
		},
	}
}

func resourceCohesityConfigureSnmpCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}

	snmpConfigParam := services.BuildSnmpConfigParam(d, true)

	_, err = client.EnableSnmpConfig(snmpConfigParam)
	if err != nil {
		return diag.Errorf("Failed to configure SNMP on the cluster: %s", err.Error())
	}
	d.SetId(fmt.Sprintf("snmp-config-%d", time.Now().Unix()))

	return resourceCohesityConfigureSnmpRead(ctx, d, m)
}

func resourceCohesityConfigureSnmpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}
	snmpConfig, err := client.GetSnmpConfig()
	if err != nil {
		return diag.Errorf("Failed to get Snmp Config for the cluster, %s", err.Error())
	}

	d.Set("agent_port", snmpConfig.AgentPort)
	d.Set("server", snmpConfig.Server)
	d.Set("trap_port", snmpConfig.TrapPort)

	if *snmpConfig.Version == "kSnmpV2" {
		d.Set("snmp_version", "V2C")
	} else {
		d.Set("snmp_version", "V3")
	}

	if snmpConfig.ReadUser != nil {
		d.Set("read_user", services.BuildSnmpUserMap(snmpConfig.ReadUser))
	}
	if snmpConfig.WriteUser != nil {
		d.Set("write_user", services.BuildSnmpUserMap(snmpConfig.WriteUser))
	}
	if snmpConfig.TrapUser != nil {
		d.Set("trap_user", services.BuildSnmpUserMap(snmpConfig.TrapUser))
	}
	return nil
}

func resourceCohesityConfigureSnmpUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChanges("agent_port", "trap_port", "server", "snmp_version", "read_user", "write_user", "trap_user") {
		config := m.(utils.Config)
		client, err := services.NewCohesityClientV2(config)
		if err != nil {
			return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
		}

		snmpConfigParam := services.BuildSnmpConfigParam(d, true)

		_, err = client.DisableSnmpConfig(snmpConfigParam)
		if err != nil {
			return diag.Errorf("Failed to update Snmp Config for the cluster, %s", err.Error())
		}
		return resourceCohesityConfigureSnmpRead(ctx, d, m)
	}
	return nil
}

func resourceCohesityConfigureSnmpDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}

	snmpConfigParam := services.BuildSnmpConfigParam(d, false)

	_, err = client.DisableSnmpConfig(snmpConfigParam)
	if err != nil {
		return diag.Errorf("Failed to disable Snmp Config for the cluster, %s", err.Error())
	}
	return nil
}
