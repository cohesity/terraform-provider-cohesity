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

// KeystoneACLProto Protobuf that describes the Keystone access control list (ACL) permissions
// for a swift container.
// Note: Keystone ACL is applicable for only keystone authenticated users.
//
// swagger:model KeystoneACLProto
type KeystoneACLProto struct {

	// Grantees with read permission.
	ReadGrantees *KeystoneACLProtoGrantees `json:"readGrantees,omitempty"`

	// Grantees with write permission.
	WriteGrantees *KeystoneACLProtoGrantees `json:"writeGrantees,omitempty"`
}

// Validate validates this keystone ACL proto
func (m *KeystoneACLProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReadGrantees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWriteGrantees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeystoneACLProto) validateReadGrantees(formats strfmt.Registry) error {
	if swag.IsZero(m.ReadGrantees) { // not required
		return nil
	}

	if m.ReadGrantees != nil {
		if err := m.ReadGrantees.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readGrantees")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("readGrantees")
			}
			return err
		}
	}

	return nil
}

func (m *KeystoneACLProto) validateWriteGrantees(formats strfmt.Registry) error {
	if swag.IsZero(m.WriteGrantees) { // not required
		return nil
	}

	if m.WriteGrantees != nil {
		if err := m.WriteGrantees.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writeGrantees")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("writeGrantees")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this keystone ACL proto based on the context it is used
func (m *KeystoneACLProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReadGrantees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWriteGrantees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeystoneACLProto) contextValidateReadGrantees(ctx context.Context, formats strfmt.Registry) error {

	if m.ReadGrantees != nil {

		if swag.IsZero(m.ReadGrantees) { // not required
			return nil
		}

		if err := m.ReadGrantees.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readGrantees")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("readGrantees")
			}
			return err
		}
	}

	return nil
}

func (m *KeystoneACLProto) contextValidateWriteGrantees(ctx context.Context, formats strfmt.Registry) error {

	if m.WriteGrantees != nil {

		if swag.IsZero(m.WriteGrantees) { // not required
			return nil
		}

		if err := m.WriteGrantees.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writeGrantees")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("writeGrantees")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeystoneACLProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeystoneACLProto) UnmarshalBinary(b []byte) error {
	var res KeystoneACLProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
