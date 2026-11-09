---
page_title: "cohesity_web_server_certificate Resource - cohesity"
subcategory: ""
description: |-
  Manages the HTTPS server certificate on a Cohesity cluster.
---

# cohesity_web_server_certificate (Resource)

Installs a hostname-valid certificate chain and matching private key. The provider
requires `api_key` with `CLUSTER_VIEW` and `CLUSTER_MODIFY`, an HTTPS `endpoint`
(or `cluster_vip`), and system trust or `ca_certificate_pem`. For the initial
self-signed certificate, `bootstrap_certificate_pem` can contain explicitly
verified existing node leaves. Remove bootstrap pins after installing trusted TLS.
Username/password authentication remains required for the upstream resources.

```terraform
resource "cohesity_web_server_certificate" "main" {
  certificate_pem     = file("server-chain.pem")
  private_key_pem_wo  = ephemeral.aws_ssm_parameter.private_key.value
  private_key_version = 1
  restart_iris        = true
}
```

- `certificate_pem` (required): PEM leaf followed by its issuing chain.
- `private_key_pem_wo` (optional, write-only, sensitive): matching key; accepts
  ephemeral values with Terraform 1.11+. Set exactly one of this and
  `private_key_pem` (stateful, sensitive).
- `private_key_version` (optional, default `1`): increment when changing the
  write-only key. Certificate renewal is managed by the certificate issuer.
- `restart_iris` (optional, default `false`): restart only Iris after upload,
  briefly interrupting UI/API clients. The served certificate is checked after
  installation. Already-correct certificates skip upload and restart.
- `sha256_fingerprint`, `expires_at`, `id` (computed): certificate metadata.

Import with `terraform import cohesity_web_server_certificate.main webServer`.
Destroy only removes Terraform management; it retains the installed certificate.
