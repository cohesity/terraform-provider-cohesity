// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Tenant Tenant
//
// Specifies a tenant object.
//
// swagger:model Tenant
type Tenant struct {
	TenantInfo
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Tenant) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TenantInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TenantInfo = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Tenant) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TenantInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tenant
func (m *Tenant) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TenantInfo
	if err := m.TenantInfo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this tenant based on the context it is used
func (m *Tenant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TenantInfo
	if err := m.TenantInfo.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Tenant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tenant) UnmarshalBinary(b []byte) error {
	var res Tenant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
