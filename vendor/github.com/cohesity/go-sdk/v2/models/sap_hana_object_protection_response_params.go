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

// SapHanaObjectProtectionResponseParams Specifies the response parameters specific to SAP HANA object protection.
//
// swagger:model SapHanaObjectProtectionResponseParams
type SapHanaObjectProtectionResponseParams struct {
	SapHanaObjectProtectionParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SapHanaObjectProtectionResponseParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SapHanaObjectProtectionParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SapHanaObjectProtectionParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SapHanaObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SapHanaObjectProtectionParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this sap hana object protection response params
func (m *SapHanaObjectProtectionResponseParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SapHanaObjectProtectionParams
	if err := m.SapHanaObjectProtectionParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this sap hana object protection response params based on the context it is used
func (m *SapHanaObjectProtectionResponseParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SapHanaObjectProtectionParams
	if err := m.SapHanaObjectProtectionParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SapHanaObjectProtectionResponseParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SapHanaObjectProtectionResponseParams) UnmarshalBinary(b []byte) error {
	var res SapHanaObjectProtectionResponseParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
