package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

type NotificationRule struct {

	// Specifies alert types this rule is applicable to.
	AlertTypeList []int32 `json:"alertTypeList"`

	// Specifies alert categories this rule is applicable to.
	// Specifies the category of an Alert.
	// kDisk - Alert associated with the disk.
	// kNode - Alert associated with general hardware on a specific node.
	// kCluster - Alert associated with general hardware in cluster level.
	// kChassis - Alert associated with the Chassis.
	// kPowerSupply - Alert associated with the power supply.
	// kCPU - Alert associated with the CPU usage.
	// kMemory - Alert associated with the RAM/Memory.
	// kTemperature - Alert associated with the temperature.
	// kFan - Alert associated with the fan.
	// kNIC - Alert associated with network chips and interfaces.
	// kFirmware - Alert associated with the firmware.
	// kNodeHealth - Alert associated with node health status.
	// kOperatingSystem - Alert associated with operating systems.
	// kDataPath - Alert associated with data management in the cluster.
	// kMetadata - Alert associated with metadata management.
	// kIndexing - Alert associated with indexing services.
	// kHelios - Alert associated with Helios.
	// kAppMarketPlace - Alert associated with App MarketPlace.
	// kSystemService -Alert associated with System service apps.
	// kLicense - Alert associated with licensing.
	// kSecurity - Alert associated with security.
	// kUpgrade - Alert associated with upgrade activities.
	// kClusterManagement - Alert associated with cluster management activities.
	// kAuditLog - Alert associated with audit log events.
	// kNetworking - Alert associated with networking issue.
	// kConfiguration - Alert associated with cluster or system configurations.
	// kStorageUsage - Alert associated with the disk/domain/cluster storage usage.
	// kFaultTolerance - Alert associated with the fault tolerance in different levels.
	// kBackupRestore - Alert associated with Backup-Restore job.
	// kArchivalRestore - Alert associated with Archival-Restore job.
	// kRemoteReplication - Alert associated with Replication job.
	// kQuota - Alert associated with Quotas.
	// kCDP - Alert associated with Continuous Data Protection.
	// kViewFailover - Alert associated with view Failover.
	// kDisasterRecovery - Alert associated with Disaster Recovery.
	// kStorageDevice - Alert associated with storage hardware(tape drives & libraries, Fiber HBAs used to attach devices, etc).
	// kStoragePool - Alert associated with storage pools -- logical groupings of similar kinds of storage hardware (disk, tape, etc) into which client data is stored.
	// kGeneralSoftwareFailure - Alert associated with general software failures - that don't fall into any known categories.
	// kAgent - Alert associated with agent based protection workloads.
	Categories []string `json:"categories"`

	// Specifies email addresses to be notified when an alert matching this
	// rule is generated.
	EmailDeliveryTargets []*EmailDeliveryTarget `json:"emailDeliveryTargets"`

	// Specifies id of the alert delivery rule.
	RuleID *int64 `json:"ruleId,omitempty"`

	// Specifies name of the alert delivery rule.
	RuleName *string `json:"ruleName,omitempty"`

	// Specifies alert severity types this rule is applicable to.
	// Specifies the severity level of an Alert.
	// kCritical - Alerts whose severity type is Critical.
	// kWarning - Alerts whose severity type is Warning.
	// kInfo - Alerts whose severity type is Info.
	Severities []string `json:"severities"`

	// Specifies whether SNMP notification to be invoked when an alert matching
	// this rule is generated.
	SnmpEnabled *bool `json:"snmpEnabled,omitempty"`

	// Specifies whether syslog notification to be invoked when an alert matching
	// this rule is generated.
	SyslogEnabled *bool `json:"syslogEnabled,omitempty"`

	// Specifies tenant id this rule is applicable to.
	TenantID *string `json:"tenantId,omitempty"`

	// Specifies external api urls to be invoked when an alert matching this
	// rule is generated.
	WebHookDeliveryTargets []*WebHookDeliveryTarget `json:"webHookDeliveryTargets"`
}
type WebHookDeliveryTarget struct {

	// Specifies curl options used to invoke external api url defined above.
	CurlOptions *string `json:"curlOptions,omitempty"`

	// external Api Url
	ExternalAPIURL *string `json:"externalApiUrl,omitempty"`
}

type EmailDeliveryTarget struct {

	// email address
	EmailAddress *string `json:"emailAddress,omitempty"`

	// Specifies the language in which the emails sent to the above defined
	// mail address should be in.
	Locale *string `json:"locale,omitempty"`

	// Specifies the recipient type on how the emails are to received.
	// The email recipient type.
	// 'kTo' indicates the primary receiver type
	// 'kCc' indicates the carbon copy receiver type
	// Enum: ["kTo","kCc"]
	RecipientType *string `json:"recipientType,omitempty"`
}

func expandEmailDeliveryTargets(raw []interface{}) ([]*EmailDeliveryTarget, error) {
	var targets []*EmailDeliveryTarget

	for i, item := range raw {
		m, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("element at index %d is not a map[string]interface{}", i)
		}

		// Required field
		emailRaw, ok := m["email_address"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'emailAddress' at index %d", i)
		}

		recipientType, _ := m["recipient_type"].(string) // optional

		target := &EmailDeliveryTarget{
			EmailAddress:  &emailRaw,
			RecipientType: &recipientType,
		}

		targets = append(targets, target)
	}

	return targets, nil
}

func expandWebhookDeliveryTargets(raw []interface{}) ([]*WebHookDeliveryTarget, error) {
	var targets []*WebHookDeliveryTarget

	for i, item := range raw {
		m, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("element at index %d is not a map[string]interface{}", i)
		}

		webhookURLRaw, ok := m["webhook_url"].(string)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'webhookUrl' at index %d", i)
		}

		curlOptions, _ := m["curl_options"].(string)

		target := &WebHookDeliveryTarget{
			ExternalAPIURL: &webhookURLRaw,
			CurlOptions:    &curlOptions,
		}

		targets = append(targets, target)
	}

	return targets, nil
}

func BuildNotificationRule(d *schema.ResourceData, isNew bool) (*NotificationRule, error) {
	notificationRule := NotificationRule{
		RuleName:      utils.StringPtr(d.Get("rule_name").(string)),
		SyslogEnabled: utils.BoolPtr(d.Get("syslog_enabled").(bool)),
		SnmpEnabled:   utils.BoolPtr(d.Get("snmp_enabled").(bool)),
	}
	if d.Get("tenant_id") != "" {
		notificationRule.TenantID = utils.StringPtr(d.Get("tenant_id").(string))
	}
	if !isNew {
		ruleId64, _ := strconv.ParseInt(d.Id(), 10, 64)
		notificationRule.RuleID = utils.Int64Ptr(ruleId64)
	}
	_, ok := d.GetOk("categories")
	if ok {
		categories, err := utils.InterfaceSliceToStringSlice(d.Get("categories").([]interface{}))
		if err != nil {
			return nil, err
		}
		notificationRule.Categories = categories
	}
	_, ok = d.GetOk("alert_names")
	if ok {
		alertNames, err := utils.InterfaceSliceToInt32Slice(d.Get("alert_types").([]interface{}))
		if err != nil {
			return nil, err
		}
		notificationRule.AlertTypeList = alertNames
	}
	_, ok = d.GetOk("severities")
	if ok {
		severities, err := utils.InterfaceSliceToStringSlice(d.Get("severities").([]interface{}))
		if err != nil {
			return nil, err
		}
		notificationRule.Severities = severities
	}
	_, ok = d.GetOk("email_delivery_targets")
	if ok {
		emailDeliveryTargets, err := expandEmailDeliveryTargets(d.Get("email_delivery_targets").([]interface{}))
		if err != nil {
			return nil, err
		}
		notificationRule.EmailDeliveryTargets = emailDeliveryTargets
	}
	if _, ok = d.GetOk("webhook_delivery_targets"); ok {
		webhookDeliveryTargets, err := expandWebhookDeliveryTargets(d.Get("webhook_delivery_targets").([]interface{}))
		if err != nil {
			return nil, err
		}
		notificationRule.WebHookDeliveryTargets = webhookDeliveryTargets
	}
	return &notificationRule, nil
}

func CreateNotificationRule(url, token string, notificationRule NotificationRule) (*NotificationRule, error) {
	resp, code, err := PostWithModel(fmt.Sprintf("%s/irisservices/api/v1/public/alertNotificationRules", url), token, notificationRule)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("the request failed with code %d, resp:%+v", code, string(resp))
	}
	var response NotificationRule
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Printf("failed to unmarshal response: %d ,: %v", resp, err)
		return nil, err
	}
	return &response, err
}

func UpdateNotificationRule(url, token string, notificationRule NotificationRule) (*NotificationRule, error) {
	resp, code, err := PutWithModel(fmt.Sprintf("%s/irisservices/api/v1/public/alertNotificationRules", url), token, notificationRule)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("the request failed with code %d", code)
	}
	var response NotificationRule
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Printf("failed to unmarshal response: %d ,: %v", resp, err)
		return nil, err
	}
	return &response, err
}
func DeleteNotificationRule(url, token, id string) error {
	resp, code, err := DeleteWithModel(fmt.Sprintf("%s/irisservices/api/v1/public/alertNotificationRules/%s", url, id), token)
	if code != 204 {
		return fmt.Errorf("the request failed with code %d, resp: %s", code, string(resp))
	}
	if err != nil {
		return err
	}
	return nil
}

func GetNotificationRule(url, token, id string) (*NotificationRule, error) {
	resp, code, err := GetWithModel(fmt.Sprintf("%s/v2/alerts/config/notification-rules?ids=%s", url, id), token)
	if code != 200 {
		return nil, fmt.Errorf("the request failed with code %d, resp: %s", code, string(resp))
	}
	if err != nil {
		return nil, err
	}
	var response []NotificationRule
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Printf("failed to unmarshal response: %d ,: %v", resp, err)
		return nil, err
	}
	if len(response) == 0 {
		return nil, fmt.Errorf("no info returned, may have been deleted")
	}
	return &response[0], nil
}
