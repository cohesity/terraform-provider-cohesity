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

// UpdateMFAResult update m f a result
//
// swagger:model UpdateMFAResult
type UpdateMFAResult struct {
	SupportTotpKeyInfo

	SupportMfaConfigInfo
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UpdateMFAResult) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SupportTotpKeyInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SupportTotpKeyInfo = aO0

	// AO1
	var aO1 SupportMfaConfigInfo
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SupportMfaConfigInfo = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UpdateMFAResult) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SupportTotpKeyInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SupportMfaConfigInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this update m f a result
func (m *UpdateMFAResult) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SupportTotpKeyInfo
	if err := m.SupportTotpKeyInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SupportMfaConfigInfo
	if err := m.SupportMfaConfigInfo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this update m f a result based on the context it is used
func (m *UpdateMFAResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SupportTotpKeyInfo
	if err := m.SupportTotpKeyInfo.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SupportMfaConfigInfo
	if err := m.SupportMfaConfigInfo.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateMFAResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateMFAResult) UnmarshalBinary(b []byte) error {
	var res UpdateMFAResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
