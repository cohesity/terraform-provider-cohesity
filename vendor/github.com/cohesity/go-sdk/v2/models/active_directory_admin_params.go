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

// ActiveDirectoryAdminParams Specifies the params of a user with administrative privilege of an Active Directory.
//
// swagger:model ActiveDirectoryAdminParams
type ActiveDirectoryAdminParams struct {

	// Specifies the user name.
	// Required: true
	Username *string `json:"username"`

	// Specifies the password of the user.
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this active directory admin params
func (m *ActiveDirectoryAdminParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectoryAdminParams) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectoryAdminParams) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this active directory admin params based on context it is used
func (m *ActiveDirectoryAdminParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectoryAdminParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryAdminParams) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryAdminParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
