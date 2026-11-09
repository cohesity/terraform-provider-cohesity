package cohesity

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"testing"
)

func TestAccCertificateLifecycle(t *testing.T) {
	f := newFixture(t)
	bootstrap := f.cert
	t.Setenv("COHESITY_USERNAME", "")
	t.Setenv("COHESITY_PASSWORD", "")
	t.Setenv("COHESITY_IP", "")
	f.deferredReload = true
	cert, key := issue(t)
	config := func(version int) string {
		return fmt.Sprintf(`
variable "api_key" {
 type = string
 sensitive = true
 ephemeral = true
 default = "test-secret"
}
variable "private_key" {
 type = string
 sensitive = true
 ephemeral = true
 default = %q
}
provider "cohesity" {
 endpoint = %q
 api_key = var.api_key
 ca_certificate_pem = %q
 bootstrap_certificate_pem = %q
}
resource "cohesity_web_server_certificate" "test" {
 certificate_pem = %q
 private_key_pem_wo = var.private_key
 restart_iris = true
 private_key_version = %d
}
`, key, f.server.URL, cert, bootstrap, cert, version)
	}
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){"cohesity": func() (*schema.Provider, error) { return Provider(), nil }},
		CheckDestroy: func(*terraform.State) error {
			if f.deletes != 0 {
				return fmt.Errorf("destroy changed production certificate")
			}
			return nil
		},
		Steps: []resource.TestStep{
			{Config: config(1), Check: resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("cohesity_web_server_certificate.test", "id", "webServer"), resource.TestCheckNoResourceAttr("cohesity_web_server_certificate.test", "private_key_pem_wo"))},
			{Config: config(1), PlanOnly: true, ExpectNonEmptyPlan: false},
			{ResourceName: "cohesity_web_server_certificate.test", ImportState: true, ImportStateId: "webServer", ImportStateVerify: true, ImportStateVerifyIgnore: []string{"restart_iris"}},
			{Config: config(2), Check: resource.TestCheckResourceAttr("cohesity_web_server_certificate.test", "private_key_version", "2")},
			{PreConfig: func() { other, otherKey := issue(t); f.mu.Lock(); f.cert = other; f.key = otherKey; f.mu.Unlock() }, Config: config(2), PlanOnly: true, ExpectNonEmptyPlan: true},
			{Config: config(2)},
		},
	})
}
