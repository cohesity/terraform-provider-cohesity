// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TenantCreateParameters Tenant Create Request.
//
// Specifies the settings used to create/add a new tenant.
//
// swagger:model TenantCreateParameters
type TenantCreateParameters struct {

	// Specifies whether bifrost (Ambassador proxy) is enabled for tenant.
	BifrostEnabled *bool `json:"bifrostEnabled,omitempty"`

	// The hostname for Cohesity cluster as seen by tenants and as is routable
	// from the tenant's network. Tenant's VLAN's hostname, if available can be
	// used instead but it is mandatory to provide this value if there's no VLAN
	// hostname to use. Also, when set, this field would take precedence over
	// VLAN hostname.
	ClusterHostname *string `json:"clusterHostname,omitempty"`

	// Set of IPs as seen from the tenant's network for the Cohesity cluster.
	// Only one from 'ClusterHostname' and 'ClusterIps' is needed.
	ClusterIps []string `json:"clusterIps"`

	// Specifies the description of this tenant.
	Description *string `json:"description,omitempty"`

	// Specifies whether this tenant is manged on helios
	IsManagedOnHelios *bool `json:"isManagedOnHelios,omitempty"`

	// Specifies the name of the tenant.
	Name *string `json:"name,omitempty"`

	// Specifies the organization suffix needed to construct tenant id. Tenant id
	// is not completely auto generated rather chosen by tenant/SP admin as we
	// needed same tenant id on multiple clusters to support replication across
	// clusters, etc.
	OrgSuffix *string `json:"orgSuffix,omitempty"`

	// Specifies the parent tenant of this tenant if available.
	ParentTenantID *string `json:"parentTenantId,omitempty"`

	// Service provider can optionally unsubscribe from the Tenant Alert Emails.
	SubscribeToAlertEmails *bool `json:"subscribeToAlertEmails,omitempty"`
}

// Validate validates this tenant create parameters
func (m *TenantCreateParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tenant create parameters based on context it is used
func (m *TenantCreateParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TenantCreateParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantCreateParameters) UnmarshalBinary(b []byte) error {
	var res TenantCreateParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
