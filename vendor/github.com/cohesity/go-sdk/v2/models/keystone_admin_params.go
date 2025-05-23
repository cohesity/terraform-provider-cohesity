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

// KeystoneAdminParams Specifies administrator credentials of a Keystone.
//
// swagger:model KeystoneAdminParams
type KeystoneAdminParams struct {

	// Specifies the administrator domain name.
	// Required: true
	Domain *string `json:"domain"`

	// Specifies the username of Keystone administrator.
	// Required: true
	Username *string `json:"username"`

	// Specifies the password of Keystone administrator.
	Password *string `json:"password,omitempty"`
}

// Validate validates this keystone admin params
func (m *KeystoneAdminParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeystoneAdminParams) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *KeystoneAdminParams) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this keystone admin params based on context it is used
func (m *KeystoneAdminParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeystoneAdminParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeystoneAdminParams) UnmarshalBinary(b []byte) error {
	var res KeystoneAdminParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
