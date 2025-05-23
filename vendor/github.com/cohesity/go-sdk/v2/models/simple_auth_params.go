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

// SimpleAuthParams Specifies the parameters for LDAP with 'Simple' authentication type.
//
// swagger:model SimpleAuthParams
type SimpleAuthParams struct {

	// Specifies the user distinguished name that is used for LDAP authentication.
	// Required: true
	UserDistinguishedName *string `json:"userDistinguishedName"`

	// Specifies the user password that is used for LDAP authentication.
	UserPassword *string `json:"userPassword,omitempty"`

	// Specifies whether to use SSL for LDAP connections.
	UseSsl *bool `json:"useSsl,omitempty"`
}

// Validate validates this simple auth params
func (m *SimpleAuthParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserDistinguishedName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimpleAuthParams) validateUserDistinguishedName(formats strfmt.Registry) error {

	if err := validate.Required("userDistinguishedName", "body", m.UserDistinguishedName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this simple auth params based on context it is used
func (m *SimpleAuthParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SimpleAuthParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimpleAuthParams) UnmarshalBinary(b []byte) error {
	var res SimpleAuthParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
