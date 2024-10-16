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

// UpdateUserParameters Specifies User properties to update.
//
// swagger:model UpdateUserParameters
type UpdateUserParameters struct {
	CommonUpdatableUserParams

	// Specifies the LOCAL user properties. This field is required when updating LOCAL Cohesity User params.
	LocalUserParams *LocalUserUpdateParams `json:"localUserParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UpdateUserParameters) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonUpdatableUserParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonUpdatableUserParams = aO0

	// AO1
	var dataAO1 struct {
		LocalUserParams *LocalUserUpdateParams `json:"localUserParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.LocalUserParams = dataAO1.LocalUserParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UpdateUserParameters) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonUpdatableUserParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		LocalUserParams *LocalUserUpdateParams `json:"localUserParams,omitempty"`
	}

	dataAO1.LocalUserParams = m.LocalUserParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this update user parameters
func (m *UpdateUserParameters) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonUpdatableUserParams
	if err := m.CommonUpdatableUserParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalUserParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUserParameters) validateLocalUserParams(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalUserParams) { // not required
		return nil
	}

	if m.LocalUserParams != nil {
		if err := m.LocalUserParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localUserParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localUserParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update user parameters based on the context it is used
func (m *UpdateUserParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonUpdatableUserParams
	if err := m.CommonUpdatableUserParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalUserParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateUserParameters) contextValidateLocalUserParams(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalUserParams != nil {

		if swag.IsZero(m.LocalUserParams) { // not required
			return nil
		}

		if err := m.LocalUserParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localUserParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localUserParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUserParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserParameters) UnmarshalBinary(b []byte) error {
	var res UpdateUserParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
