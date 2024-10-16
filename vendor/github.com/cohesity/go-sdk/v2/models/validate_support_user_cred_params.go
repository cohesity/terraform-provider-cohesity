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

// ValidateSupportUserCredParams Specifies the support user credentials to validate.
//
// swagger:model ValidateSupportUserCredParams
type ValidateSupportUserCredParams struct {

	// Specifies the password of the user to validate.
	// Required: true
	// Min Length: 1
	Password *string `json:"password"`
}

// Validate validates this validate support user cred params
func (m *ValidateSupportUserCredParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValidateSupportUserCredParams) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", *m.Password, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this validate support user cred params based on context it is used
func (m *ValidateSupportUserCredParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ValidateSupportUserCredParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValidateSupportUserCredParams) UnmarshalBinary(b []byte) error {
	var res ValidateSupportUserCredParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
