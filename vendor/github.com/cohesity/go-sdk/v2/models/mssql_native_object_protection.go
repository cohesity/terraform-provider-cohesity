// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MssqlNativeObjectProtection Specifies the object params to create Native based MSSQL Object Protection.
//
// swagger:model MssqlNativeObjectProtection
type MssqlNativeObjectProtection struct {

	// Specifies the ID of the object being protected. If this is a non leaf level object, then the object will be auto-protected unless leaf objects are specified for exclusion.
	// Required: true
	ID *int64 `json:"id"`
}

// Validate validates this mssql native object protection
func (m *MssqlNativeObjectProtection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MssqlNativeObjectProtection) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mssql native object protection based on context it is used
func (m *MssqlNativeObjectProtection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MssqlNativeObjectProtection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MssqlNativeObjectProtection) UnmarshalBinary(b []byte) error {
	var res MssqlNativeObjectProtection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
