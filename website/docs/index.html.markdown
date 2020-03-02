---
layout: "Cohesity"
page_title: "Provider: Cohesity"
description: |-
  The Cohesity terraform provider is used to interact with cohesity and automate cluster workflows using Cohesity REST API's.
---

# Cohesity Terraform Provider

Cohesity terraform provider is used to interact with cohesity and automate cluster workflows using Cohesity REST API's. The provider needs to be configured with proper credentails prior to managing its resources. Currently it supports physical, cloud, virtual edition cluster creation and destruction workflows

Use the navigation to the left to read about the available resources.

## Example Usage

```
provider "cohesity" {
        cluster_vip = "10.2.45.143"
        cluster_username = "abc"
        cluster_password = "abc"
        cluster_domain = "LOCAL"

}
```

#### Argument Reference

- cluster_vip - (Required, string) IP or hostname of Cohesity cluster node. This can also be read from **COHESITY_IP** environment variable
- cluster_username - (Required, string) Cohesity cluster username. This can also be read from **COHESITY_USERNAME** environment variable
- cluster_password - (Required, string) Cohesity cluster password. This can also be read from **COHESITY_PASSWORD** environment variable
- cluster_domain - (Optional, string) The domain name of cohesity user. Defaults to **LOCAL**
