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