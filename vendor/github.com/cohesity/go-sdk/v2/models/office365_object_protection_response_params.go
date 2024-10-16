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

// Office365ObjectProtectionResponseParams Specifies the response parameters specific to Microsoft 365 User Mailbox object protection.
//
// swagger:model Office365ObjectProtectionResponseParams
type Office365ObjectProtectionResponseParams struct {
	Office365ObjectProtectionParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Office365ObjectProtectionResponseParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Office365ObjectProtectionParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Office365ObjectProtectionParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Office365ObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Office365ObjectProtectionParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this office365 object protection response params
func (m *Office365ObjectProtectionResponseParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Office365ObjectProtectionParams
	if err := m.Office365ObjectProtectionParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this office365 object protection response params based on the context it is used
func (m *Office365ObjectProtectionResponseParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Office365ObjectProtectionParams
	if err := m.Office365ObjectProtectionParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Office365ObjectProtectionResponseParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Office365ObjectProtectionResponseParams) UnmarshalBinary(b []byte) error {
	var res Office365ObjectProtectionResponseParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
