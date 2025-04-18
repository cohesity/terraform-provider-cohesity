// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonObjectSummary Common Object Entity Summary.
//
// Specifies an Common summary properties for an object.
//
// swagger:model CommonObjectSummary
type CommonObjectSummary struct {
	ObjectSummary

	// Specifies the count and size of protected and unprotected objects for the size.
	ProtectionStats []*ObjectProtectionStatsSummary `json:"protectionStats"`

	// Specifies the list of users, groups and tenants that have permissions for this object.
	Permissions *PermissionInfo `json:"permissions,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CommonObjectSummary) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ObjectSummary
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ObjectSummary = aO0

	// AO1
	var dataAO1 struct {
		ProtectionStats []*ObjectProtectionStatsSummary `json:"protectionStats"`

		Permissions *PermissionInfo `json:"permissions,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ProtectionStats = dataAO1.ProtectionStats

	m.Permissions = dataAO1.Permissions

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CommonObjectSummary) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ObjectSummary)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ProtectionStats []*ObjectProtectionStatsSummary `json:"protectionStats"`

		Permissions *PermissionInfo `json:"permissions,omitempty"`
	}

	dataAO1.ProtectionStats = m.ProtectionStats

	dataAO1.Permissions = m.Permissions

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this common object summary
func (m *CommonObjectSummary) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ObjectSummary
	if err := m.ObjectSummary.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonObjectSummary) validateProtectionStats(formats strfmt.Registry) error {

	if swag.IsZero(m.ProtectionStats) { // not required
		return nil
	}

	for i := 0; i < len(m.ProtectionStats); i++ {
		if swag.IsZero(m.ProtectionStats[i]) { // not required
			continue
		}

		if m.ProtectionStats[i] != nil {
			if err := m.ProtectionStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionStats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonObjectSummary) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this common object summary based on the context it is used
func (m *CommonObjectSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ObjectSummary
	if err := m.ObjectSummary.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonObjectSummary) contextValidateProtectionStats(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProtectionStats); i++ {

		if m.ProtectionStats[i] != nil {

			if swag.IsZero(m.ProtectionStats[i]) { // not required
				return nil
			}

			if err := m.ProtectionStats[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionStats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonObjectSummary) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	if m.Permissions != nil {

		if swag.IsZero(m.Permissions) { // not required
			return nil
		}

		if err := m.Permissions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonObjectSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonObjectSummary) UnmarshalBinary(b []byte) error {
	var res CommonObjectSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
