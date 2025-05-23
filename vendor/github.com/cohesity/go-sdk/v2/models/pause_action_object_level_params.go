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

// PauseActionObjectLevelParams Specifies the request parameters for Pause action on a Protected object.
//
// swagger:model PauseActionObjectLevelParams
type PauseActionObjectLevelParams struct {
	ProtectionObjectInput
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PauseActionObjectLevelParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ProtectionObjectInput
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ProtectionObjectInput = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PauseActionObjectLevelParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ProtectionObjectInput)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pause action object level params
func (m *PauseActionObjectLevelParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProtectionObjectInput
	if err := m.ProtectionObjectInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this pause action object level params based on the context it is used
func (m *PauseActionObjectLevelParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProtectionObjectInput
	if err := m.ProtectionObjectInput.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PauseActionObjectLevelParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PauseActionObjectLevelParams) UnmarshalBinary(b []byte) error {
	var res PauseActionObjectLevelParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
