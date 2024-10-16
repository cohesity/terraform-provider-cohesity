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

// ObjectProtectionRunSummary Object Protection Run Summary.
//
// Specifies the response body of the get object runs request.
//
// swagger:model ObjectProtectionRunSummary
type ObjectProtectionRunSummary struct {
	ObjectIdentifier

	ObjectProtectionRunInfo
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ObjectProtectionRunSummary) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ObjectIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ObjectIdentifier = aO0

	// AO1
	var aO1 ObjectProtectionRunInfo
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ObjectProtectionRunInfo = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ObjectProtectionRunSummary) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ObjectIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ObjectProtectionRunInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this object protection run summary
func (m *ObjectProtectionRunSummary) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ObjectIdentifier
	if err := m.ObjectIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ObjectProtectionRunInfo
	if err := m.ObjectProtectionRunInfo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this object protection run summary based on the context it is used
func (m *ObjectProtectionRunSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ObjectIdentifier
	if err := m.ObjectIdentifier.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ObjectProtectionRunInfo
	if err := m.ObjectProtectionRunInfo.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectProtectionRunSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectProtectionRunSummary) UnmarshalBinary(b []byte) error {
	var res ObjectProtectionRunSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
