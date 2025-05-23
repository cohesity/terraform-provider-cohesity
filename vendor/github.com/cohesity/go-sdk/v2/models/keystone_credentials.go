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

// KeystoneCredentials Specifies user credentials of a Keystone server.
//
// swagger:model KeystoneCredentials
type KeystoneCredentials struct {

	// Specifies parameters related to Keystone administrator.
	// Required: true
	AdminCreds *KeystoneAdminParams `json:"adminCreds"`

	// Specifies parameters related to Keystone scope.
	// Required: true
	Scope *KeystoneScopeParams `json:"scope"`
}

// Validate validates this keystone credentials
func (m *KeystoneCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminCreds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeystoneCredentials) validateAdminCreds(formats strfmt.Registry) error {

	if err := validate.Required("adminCreds", "body", m.AdminCreds); err != nil {
		return err
	}

	if m.AdminCreds != nil {
		if err := m.AdminCreds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminCreds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminCreds")
			}
			return err
		}
	}

	return nil
}

func (m *KeystoneCredentials) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("scope", "body", m.Scope); err != nil {
		return err
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this keystone credentials based on the context it is used
func (m *KeystoneCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdminCreds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeystoneCredentials) contextValidateAdminCreds(ctx context.Context, formats strfmt.Registry) error {

	if m.AdminCreds != nil {

		if err := m.AdminCreds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminCreds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminCreds")
			}
			return err
		}
	}

	return nil
}

func (m *KeystoneCredentials) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {

		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeystoneCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeystoneCredentials) UnmarshalBinary(b []byte) error {
	var res KeystoneCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
