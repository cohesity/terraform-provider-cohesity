package cohesity

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-cohesity/cohesity/utils"
)

func TestProviderAuthenticationModes(t *testing.T) {
	for _, name := range []string{"COHESITY_API_KEY", "COHESITY_IP", "COHESITY_USERNAME", "COHESITY_PASSWORD"} {
		t.Setenv(name, "")
	}
	tests := []struct {
		name      string
		config    map[string]interface{}
		wantError bool
	}{
		{"existing username/password", map[string]interface{}{"cluster_vip": "backup.example.com", "cluster_username": "admin", "cluster_password": "password"}, false},
		{"API key only", map[string]interface{}{"endpoint": "https://backup.example.com", "api_key": "test-secret"}, false},
		{"missing authentication", map[string]interface{}{"cluster_vip": "backup.example.com"}, true},
		{"missing endpoint", map[string]interface{}{"api_key": "test-secret"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, Provider().Schema, tt.config)
			meta, diags := providerConfigure(context.Background(), d)
			if diags.HasError() != tt.wantError {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}
			if !tt.wantError && meta.(utils.Config).Endpoint != "https://backup.example.com" {
				t.Fatal("unexpected certificate endpoint")
			}
		})
	}
}
