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

// VMFilter Specifies the VM filter details.
//
// swagger:model VMFilter
type VMFilter struct {
	Filter

	// Specifies whether the provided filter string is case sensitive or not. This needs to be explicitly set to true if user is trying to filter by case sensitive expressions. The default value is assumed to be false.
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VMFilter) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Filter
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Filter = aO0

	// AO1
	var dataAO1 struct {
		CaseSensitive *bool `json:"caseSensitive,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CaseSensitive = dataAO1.CaseSensitive

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VMFilter) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Filter)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CaseSensitive *bool `json:"caseSensitive,omitempty"`
	}

	dataAO1.CaseSensitive = m.CaseSensitive

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this VM filter
func (m *VMFilter) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Filter
	if err := m.Filter.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this VM filter based on the context it is used
func (m *VMFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Filter
	if err := m.Filter.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VMFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMFilter) UnmarshalBinary(b []byte) error {
	var res VMFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
