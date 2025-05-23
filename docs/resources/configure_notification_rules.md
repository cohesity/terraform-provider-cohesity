---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cohesity_configure_notification_rules Resource - terraform-provider-cohesity"
subcategory: ""
description: |-
  
---

# cohesity_configure_notification_rules (Resource)



## Example Usage

```terraform
resource "cohesity_configure_notification_rules" "rule1" {
  rule_name        = "rule1"

  email_delivery_targets {
    email_address  = "example@example.com"
    recipient_type = "kTo"
  }

  email_delivery_targets {
    email_address  = "example2@example.com"
    recipient_type = "kCc"
  }
  
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `rule_name` (String)

### Optional

- `alert_types` (List of Number)
- `categories` (List of String)
- `email_delivery_targets` (Block List) (see [below for nested schema](#nestedblock--email_delivery_targets))
- `severities` (List of String)
- `snmp_enabled` (Boolean)
- `syslog_enabled` (Boolean)
- `tenant_id` (String)
- `webhook_delivery_targets` (Block List) (see [below for nested schema](#nestedblock--webhook_delivery_targets))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--email_delivery_targets"></a>
### Nested Schema for `email_delivery_targets`

Required:

- `email_address` (String) Main recepient email addresses

Optional:

- `recipient_type` (String) Whether to add recipient as To or CC	// Enum: ['kTo','kCc']


<a id="nestedblock--webhook_delivery_targets"></a>
### Nested Schema for `webhook_delivery_targets`

Required:

- `webhook_url` (String) Destination webhook URL

Optional:

- `curl_options` (String) Options for webhook
