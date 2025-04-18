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

// MssqlConnectionParams Specifies the parameters to connect to a SQL node/cluster using given IP or hostname FQDN.
//
// swagger:model MssqlConnectionParams
type MssqlConnectionParams struct {
	MsSQLCommonConnectionParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MssqlConnectionParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MsSQLCommonConnectionParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MsSQLCommonConnectionParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MssqlConnectionParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.MsSQLCommonConnectionParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this mssql connection params
func (m *MssqlConnectionParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MsSQLCommonConnectionParams
	if err := m.MsSQLCommonConnectionParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this mssql connection params based on the context it is used
func (m *MssqlConnectionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MsSQLCommonConnectionParams
	if err := m.MsSQLCommonConnectionParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MssqlConnectionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MssqlConnectionParams) UnmarshalBinary(b []byte) error {
	var res MssqlConnectionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
