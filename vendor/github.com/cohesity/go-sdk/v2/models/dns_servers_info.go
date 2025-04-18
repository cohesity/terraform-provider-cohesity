// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DNSServersInfo DnsServersInfo
//
// List of DNS servers in cluster.
//
// swagger:model DnsServersInfo
type DNSServersInfo struct {

	// List of DNS servers in cluster.
	DNSServers []string `json:"dnsServers"`
}

// Validate validates this Dns servers info
func (m *DNSServersInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Dns servers info based on context it is used
func (m *DNSServersInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSServersInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSServersInfo) UnmarshalBinary(b []byte) error {
	var res DNSServersInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
