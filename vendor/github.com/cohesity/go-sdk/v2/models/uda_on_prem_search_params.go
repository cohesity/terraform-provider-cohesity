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

// UdaOnPremSearchParams UdaOnPremSearchParams
//
// Parameters required to search Universal Data Adapter objects.
//
// swagger:model UdaOnPremSearchParams
type UdaOnPremSearchParams struct {
	UdaSearchParams

	OnPremSourceIds
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UdaOnPremSearchParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 UdaSearchParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.UdaSearchParams = aO0

	// AO1
	var aO1 OnPremSourceIds
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.OnPremSourceIds = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UdaOnPremSearchParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.UdaSearchParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.OnPremSourceIds)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this uda on prem search params
func (m *UdaOnPremSearchParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UdaSearchParams
	if err := m.UdaSearchParams.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with OnPremSourceIds
	if err := m.OnPremSourceIds.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this uda on prem search params based on the context it is used
func (m *UdaOnPremSearchParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UdaSearchParams
	if err := m.UdaSearchParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with OnPremSourceIds
	if err := m.OnPremSourceIds.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UdaOnPremSearchParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaOnPremSearchParams) UnmarshalBinary(b []byte) error {
	var res UdaOnPremSearchParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
