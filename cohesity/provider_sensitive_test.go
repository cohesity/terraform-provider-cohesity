package cohesity

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var credentialAttrPattern = regexp.MustCompile(`(?i)(^|_)(password|passwd|secret|token|private_key|api_key|apikey|access_key|credentials?|passphrase)($|_)`)

func checkSensitive(t *testing.T, path string, s map[string]*schema.Schema) {
	for name, attr := range s {
		attrPath := path + "." + name
		if credentialAttrPattern.MatchString(name) && !attr.Sensitive {
			t.Errorf("%s looks like a credential but is not marked Sensitive", attrPath)
		}
		if elem, ok := attr.Elem.(*schema.Resource); ok {
			checkSensitive(t, attrPath, elem.Schema)
		}
	}
}

func TestProvider_credentialAttributesSensitive(t *testing.T) {
	p := Provider()
	checkSensitive(t, "provider", p.Schema)
	for name, r := range p.ResourcesMap {
		checkSensitive(t, name, r.Schema)
	}
	for name, ds := range p.DataSourcesMap {
		checkSensitive(t, "data."+name, ds.Schema)
	}
}
