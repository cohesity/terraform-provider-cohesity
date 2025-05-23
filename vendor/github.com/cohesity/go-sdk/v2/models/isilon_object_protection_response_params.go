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

// IsilonObjectProtectionResponseParams Specifies the parameters which are specific to Isilon object protection.
//
// swagger:model IsilonObjectProtectionResponseParams
type IsilonObjectProtectionResponseParams struct {
	CommonNasProtectionParams

	IsilonObjectProtectionParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IsilonObjectProtectionResponseParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonNasProtectionParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonNasProtectionParams = aO0

	// AO1
	var aO1 IsilonObjectProtectionParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.IsilonObjectProtectionParams = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IsilonObjectProtectionResponseParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonNasProtectionParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.IsilonObjectProtectionParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this isilon object protection response params
func (m *IsilonObjectProtectionResponseParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonNasProtectionParams
	if err := m.CommonNasProtectionParams.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with IsilonObjectProtectionParams
	if err := m.IsilonObjectProtectionParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this isilon object protection response params based on the context it is used
func (m *IsilonObjectProtectionResponseParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonNasProtectionParams
	if err := m.CommonNasProtectionParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with IsilonObjectProtectionParams
	if err := m.IsilonObjectProtectionParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IsilonObjectProtectionResponseParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IsilonObjectProtectionResponseParams) UnmarshalBinary(b []byte) error {
	var res IsilonObjectProtectionResponseParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
