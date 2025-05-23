---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cohesity_syslog_server Resource - terraform-provider-cohesity"
subcategory: ""
description: |-
  
---

# cohesity_syslog_server (Resource)



## Example Usage

```terraform
resource "cohesity_syslog_server" "syslog" {
  name        = "syslog"
  ip_address = "0.0.0.0"
  port       = 514
  protocol   = "udp"
  ca_certificate = "*****"
  program_names = ["sshd"]
  is_enabled = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ca_certificate` (String)
- `ip_address` (String)
- `name` (String)
- `port` (Number)
- `program_names` (List of String) List of program names.
- `protocol` (String)

### Optional

- `is_enabled` (Boolean)

### Read-Only

- `id` (String) The ID of this resource.
