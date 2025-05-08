package cohesity

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/services"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func resourceCohesityConfigureNotificationRules() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceCohesityConfigureNotificationRulesCreate,
		ReadContext:   resourceCohesityConfigureNotificationRulesRead,
		UpdateContext: resourceCohesityConfigureNotificationRulesUpdate,
		DeleteContext: resourceCohesityConfigureNotificationRulesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"rule_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"syslog_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"snmp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"email_delivery_targets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Define the fields of EmailDeliveryTarget here
						"email_address": {

							Type:         schema.TypeString,
							Required:     true,
							Description:  "Main recepient email addresses",
							ValidateFunc: utils.IsValidEmail,
						},
						"recipient_type": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "Whether to add recipient as To or CC	// Enum: ['kTo','kCc']",
							Default:      "kTo",
							ValidateFunc: validation.StringInSlice([]string{"kTo", "kCc"}, false),
						},
					},
				},
			},
			"webhook_delivery_targets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Define the fields of WebhookDeliveryTarget here
						"webhook_url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Destination webhook URL",
						},
						"curl_options": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Options for webhook",
						},
					},
				},
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"severities": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					Default:      []string{"kCritical", "kWarning", "kInfo"},
					ValidateFunc: validation.StringInSlice([]string{"kCritical", "kWarning", "kInfo"}, false),
				},
			},
			"categories": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					Default:      categories,
					ValidateFunc: validation.StringInSlice(categories, false),
				},
			},
		},
	}
}

var categories = []string{
	"kDisk",
	"kNode",
	"kCluster",
	"kChassis",
	"kPowerSupply",
	"kCPU",
	"kMemory",
	"kTemperature",
	"kFan",
	"kNIC",
	"kFirmware",
	"kNodeHealth",
	"kOperatingSystem",
	"kDataPath",
	"kMetadata",
	"kIndexing",
	"kHelios",
	"kAppMarketPlace",
	"kSystemService",
	"kLicense",
	"kSecurity",
	"kUpgrade",
	"kClusterManagement",
	"kAuditLog",
	"kNetworking",
	"kConfiguration",
	"kStorageUsage",
	"kFaultTolerance",
	"kBackupRestore",
	"kArchivalRestore",
	"kRemoteReplication",
	"kQuota",
	"kCDP",
	"kViewFailover",
	"kDisasterRecovery",
	"kStorageDevice",
	"kStoragePool",
	"kGeneralSoftwareFailure",
	"kAgent",
}

func resourceCohesityConfigureNotificationRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	// Build the notification configuration parameters
	notificationRule, err := services.BuildNotificationRule(d, true)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO] notificationRule: %+v", notificationRule)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer token with the given details, %s", err.Error())
	}
	url, err := services.NormalizeHTTPSAddress(config.ClusterVIP)
	if err != nil {
		return diag.Errorf("Failed to get a url for cluster with the given details, %s", err.Error())
	}
	resp, err := services.CreateNotificationRule(url, token, *notificationRule)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[INFO] notificationRule: %+v, \n token:%s", *resp, token)

	// Set the ID of the created notification configuration in the resource data
	d.SetId(fmt.Sprintf("%d", *resp.RuleID))

	return nil
}

func resourceCohesityConfigureNotificationRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer token with the given details, %s", err.Error())
	}
	url, err := services.NormalizeHTTPSAddress(config.ClusterVIP)
	if err != nil {
		return diag.Errorf("Failed to get a url for cluster with the given details, %s", err.Error())
	}

	// Get the notification configuration by ID
	notificationConfig, err := services.GetNotificationRule(url, token, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Set the resource data with the retrieved notification configuration
	d.Set("rule_name", notificationConfig.RuleName)
	d.Set("syslog_enabled", notificationConfig.SyslogEnabled)
	d.Set("snmp_enabled", notificationConfig.SnmpEnabled)
	d.Set("email_delivery_targets", notificationConfig.EmailDeliveryTargets)
	d.Set("webhook_delivery_targets", notificationConfig.WebHookDeliveryTargets)
	d.Set("alert_names", notificationConfig.AlertTypeList)
	d.Set("severities", notificationConfig.Severities)
	d.Set("categories", notificationConfig.Categories)
	d.Set("tenant_id", notificationConfig.TenantID)

	return nil
}

func resourceCohesityConfigureNotificationRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer token with the given details, %s", err.Error())
	}
	url, err := services.NormalizeHTTPSAddress(config.ClusterVIP)
	if err != nil {
		return diag.Errorf("Failed to get a url for cluster with the given details, %s", err.Error())
	}
	// Build the updated notification configuration parameters
	notificationConfig, err := services.BuildNotificationRule(d, false)
	if err != nil {
		return diag.FromErr(err)
	}
	// Update the notification configuration
	_, err = services.UpdateNotificationRule(url, token, *notificationConfig)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceCohesityConfigureNotificationRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(utils.Config)
	token, err := services.GetAccessToken(config)
	if err != nil {
		return diag.Errorf("Failed to get a bearer token with the given details, %s", err.Error())
	}
	url, err := services.NormalizeHTTPSAddress(config.ClusterVIP)
	if err != nil {
		return diag.Errorf("Failed to get a url for cluster with the given details, %s", err.Error())
	}

	// Delete the notification configuration by ID
	err = services.DeleteNotificationRule(url, token, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
