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

// UnregisterSwiftParams Specifies the parameters to unregister a Swift service from Keystone server.
//
// swagger:model UnregisterSwiftParams
type UnregisterSwiftParams struct {

	// Specifies the tenant Id who's Swift service will be unregistered.
	// Required: true
	TenantID *string `json:"tenantId"`

	// Specifies the credentials of the Keystone server. Keystone configuration associated to the tenant will be used if this field is not specified.
	KeystoneCredentials *KeystoneCredentials `json:"keystoneCredentials,omitempty"`
}

// Validate validates this unregister swift params
func (m *UnregisterSwiftParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeystoneCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnregisterSwiftParams) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *UnregisterSwiftParams) validateKeystoneCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.KeystoneCredentials) { // not required
		return nil
	}

	if m.KeystoneCredentials != nil {
		if err := m.KeystoneCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keystoneCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keystoneCredentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this unregister swift params based on the context it is used
func (m *UnregisterSwiftParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKeystoneCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnregisterSwiftParams) contextValidateKeystoneCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.KeystoneCredentials != nil {

		if swag.IsZero(m.KeystoneCredentials) { // not required
			return nil
		}

		if err := m.KeystoneCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keystoneCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keystoneCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnregisterSwiftParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnregisterSwiftParams) UnmarshalBinary(b []byte) error {
	var res UnregisterSwiftParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
