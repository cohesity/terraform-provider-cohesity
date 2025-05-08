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

func resourceCohesityConfigureSmtp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCohesityConfigureSmtpCreate,
		ReadContext:   resourceCohesityConfigureSmtpRead,
		UpdateContext: resourceCohesityConfigureSmtpUpdate,
		DeleteContext: resourceCohesityConfigureSmtpDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"smtp_server": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the IP address or the FQDN of the SMTP server.",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Specifies the SMTP port. Usually 465 or 587. For authenticated connection, it is generally 587.",
			},
			"sender_email_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is used for setting Sender field in SMTP header. This has to be in valid email format, and could be different from username, and it's a required field if the username is not of valid email address.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the username which will be used to connect to the SMTP server. If username is not specified, then it would imply that SMTP server is set up for unauthenticated access.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the password of the SMTP user. This is required if username is specified in the request.",
			},
			"use_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "This is set to true when the SMTP server uses SSL/TLS without supporting STARTTLS. Typically, this is used for port 465.",
			},
		},
	}
}

func resourceCohesityConfigureSmtpCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}

	smtpConfigParam := services.BuildSmtpConfigParam(d, true)

	_, err = client.EnableSmtpConfig(smtpConfigParam)
	if err != nil {
		return diag.Errorf("Failed to configure SMTP on the cluster: %s", err.Error())
	}
	d.SetId(fmt.Sprintf("smtp-config-%d", time.Now().Unix()))
	d.Set("password", smtpConfigParam.Password)

	return resourceCohesityConfigureSmtpRead(ctx, d, m)
}

func resourceCohesityConfigureSmtpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}
	smtpConfig, err := client.GetSmtpConfig()
	if err != nil {
		return diag.Errorf("Failed to get Smtp Config for the cluster, %s", err.Error())
	}

	d.Set("smtp_server", smtpConfig.Hostname)
	d.Set("port", smtpConfig.Port)
	// d.Set("sender_email_address", smtpConfig.SenderEmailAddress)
	d.Set("username", smtpConfig.Username)
	d.Set("use_ssl", smtpConfig.UseSSL)

	return nil
}

func resourceCohesityConfigureSmtpUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChanges("smtp_server", "port", "sender_email_address", "username", "password", "use_ssl") {
		config := m.(utils.Config)
		client, err := services.NewCohesityClientV2(config)
		if err != nil {
			return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
		}

		smtpConfigParam := services.BuildSmtpConfigParam(d, true)

		_, err = client.EnableSmtpConfig(smtpConfigParam)
		if err != nil {
			return diag.Errorf("Failed to update Smtp Config for the cluster, %s", err.Error())
		}
		d.Set("password", smtpConfigParam.Password)
		return resourceCohesityConfigureSmtpRead(ctx, d, m)
	}
	return nil
}

func resourceCohesityConfigureSmtpDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	client, err := services.NewCohesityClientV2(config)
	if err != nil {
		return diag.Errorf("Failed to create a client with the given details, %s", err.Error())
	}

	smtpConfigParam := services.BuildSmtpConfigParam(d, false)

	err = client.DisableSmtpConfig(smtpConfigParam)
	if err != nil {
		return diag.Errorf("Failed to disable Smtp Config for the cluster, %s", err.Error())
	}
	return nil
}
