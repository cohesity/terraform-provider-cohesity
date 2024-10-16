// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TenantGroupUpdateParameters Tenant Group Update Details.
//
// Specifies group update details about a tenant.
//
// swagger:model TenantGroupUpdateParameters
type TenantGroupUpdateParameters struct {

	// Specifies the array of Sid of the groups.
	Sids []string `json:"sids"`

	// Specifies the unique id of the tenant.
	TenantID *string `json:"tenantId,omitempty"`
}

// Validate validates this tenant group update parameters
func (m *TenantGroupUpdateParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tenant group update parameters based on context it is used
func (m *TenantGroupUpdateParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TenantGroupUpdateParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantGroupUpdateParameters) UnmarshalBinary(b []byte) error {
	var res TenantGroupUpdateParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
