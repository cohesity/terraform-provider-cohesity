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

// MssqlObjectProtectionRequestParams Specifies the request parameters specific to MSSQL object protection.
//
// swagger:model MssqlObjectProtectionRequestParams
type MssqlObjectProtectionRequestParams struct {
	MssqlObjectProtectionParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MssqlObjectProtectionRequestParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MssqlObjectProtectionParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MssqlObjectProtectionParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MssqlObjectProtectionRequestParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.MssqlObjectProtectionParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this mssql object protection request params
func (m *MssqlObjectProtectionRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MssqlObjectProtectionParams
	if err := m.MssqlObjectProtectionParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this mssql object protection request params based on the context it is used
func (m *MssqlObjectProtectionRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MssqlObjectProtectionParams
	if err := m.MssqlObjectProtectionParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MssqlObjectProtectionRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MssqlObjectProtectionRequestParams) UnmarshalBinary(b []byte) error {
	var res MssqlObjectProtectionRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
