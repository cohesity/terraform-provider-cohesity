// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IbmTenantCOSResourceConfig IbmTenantCOSResourceConfig
//
// Specifies the details of COS resource configuration required for posting metrics and trackinb billing information for IBM tenants.
//
// swagger:model IbmTenantCOSResourceConfig
type IbmTenantCOSResourceConfig struct {

	// Specifies the resource COS resource configuration endpoint that will be used for fetching bucket usage for a given tenant.
	ResourceURL *string `json:"resourceURL,omitempty"`
}

// Validate validates this ibm tenant c o s resource config
func (m *IbmTenantCOSResourceConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibm tenant c o s resource config based on context it is used
func (m *IbmTenantCOSResourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbmTenantCOSResourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbmTenantCOSResourceConfig) UnmarshalBinary(b []byte) error {
	var res IbmTenantCOSResourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
