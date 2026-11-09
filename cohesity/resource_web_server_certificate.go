package cohesity

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func resourceWebServerCertificate() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages the cluster HTTPS certificate. Destroy forgets the resource and retains the installed certificate.",
		CreateContext: saveWebServerCertificate,
		UpdateContext: saveWebServerCertificate,
		ReadContext:   readWebServerCertificate,
		DeleteContext: func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics { return nil },
		Importer: &schema.ResourceImporter{StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
			if d.Id() != "webServer" {
				return nil, fmt.Errorf("use webServer as the import ID")
			}
			if err := d.Set("private_key_version", 1); err != nil {
				return nil, err
			}
			if err := d.Set("restart_iris", false); err != nil {
				return nil, err
			}
			return []*schema.ResourceData{d}, nil
		}},
		Schema: map[string]*schema.Schema{
			"certificate_pem":     {Type: schema.TypeString, Required: true, Description: "PEM leaf certificate followed by its issuing chain."},
			"private_key_pem":     {Type: schema.TypeString, Optional: true, Sensitive: true, ConflictsWith: []string{"private_key_pem_wo"}, Description: "Matching PEM private key persisted in Terraform state. Use private_key_pem_wo for ephemeral keys."},
			"private_key_pem_wo":  {Type: schema.TypeString, Optional: true, Sensitive: true, WriteOnly: true, ConflictsWith: []string{"private_key_pem"}, Description: "Matching PEM private key, never persisted in plan or state. Increment private_key_version when rotating it."},
			"private_key_version": {Type: schema.TypeInt, Optional: true, Default: 1, ValidateFunc: validation.IntAtLeast(1), Description: "Increment when rotating a write-only private key."},
			"restart_iris":        {Type: schema.TypeBool, Optional: true, Default: false, Description: "Restart only Iris after changed uploads to activate HTTPS; briefly interrupts UI/API clients."},
			"sha256_fingerprint":  {Type: schema.TypeString, Computed: true, Description: "SHA256 fingerprint of the configured leaf."},
			"expires_at":          {Type: schema.TypeString, Computed: true, Description: "Configured leaf expiry in RFC3339 format."},
		},
	}
}

func certificateClient(meta interface{}) (*client, error) {
	config := meta.(utils.Config)
	if config.APIKey == "" {
		return nil, fmt.Errorf("api_key is required for web server certificate management")
	}
	return newClient(config.Endpoint, config.APIKey, config.CACertificatePEM, config.BootstrapCertificatePEM)
}

func saveWebServerCertificate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, err := certificateClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	key := d.Get("private_key_pem").(string)
	raw, diags := d.GetRawConfigAt(cty.GetAttrPath("private_key_pem_wo"))
	if diags.HasError() {
		return diags
	}
	if !raw.IsNull() {
		if !raw.IsKnown() || raw.Type() != cty.String {
			return diag.Errorf("private_key_pem_wo must be a known string at apply")
		}
		if key != "" {
			return diag.Errorf("set exactly one of private_key_pem and private_key_pem_wo")
		}
		key = raw.AsString()
	}
	if key == "" {
		return diag.Errorf("set exactly one of private_key_pem and private_key_pem_wo")
	}
	cert := d.Get("certificate_pem").(string)
	if err = c.install(ctx, cert, key, d.Get("restart_iris").(bool)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("webServer")
	return setWebServerCertificateMetadata(d, cert)
}

func readWebServerCertificate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, err := certificateClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	remote, err := c.request(ctx, http.MethodGet, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	configured := d.Get("certificate_pem").(string)
	// Keep formatting when the DER certificate chain is unchanged.
	if sameCertificates(configured, remote) {
		remote = configured
	}
	if err = d.Set("certificate_pem", remote); err != nil {
		return diag.FromErr(err)
	}
	return setWebServerCertificateMetadata(d, remote)
}

func setWebServerCertificateMetadata(d *schema.ResourceData, cert string) diag.Diagnostics {
	chain, err := certificates(cert)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("sha256_fingerprint", fingerprint(chain[0])); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("expires_at", chain[0].NotAfter.UTC().Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
