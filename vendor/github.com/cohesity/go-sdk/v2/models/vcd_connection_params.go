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

// VcdConnectionParams Parameters to connect and query VMware config file.
//
// Specifies the parameters to connect to a seed node and fetch information from its config file.
//
// swagger:model VcdConnectionParams
type VcdConnectionParams struct {

	// IP or hostname of the source.
	// Required: true
	Host *string `json:"host"`

	Credentials
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VcdConnectionParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Host *string `json:"host"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Host = dataAO0.Host

	// AO1
	var aO1 Credentials
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Credentials = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VcdConnectionParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Host *string `json:"host"`
	}

	dataAO0.Host = m.Host

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.Credentials)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vcd connection params
func (m *VcdConnectionParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with Credentials
	if err := m.Credentials.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcdConnectionParams) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vcd connection params based on the context it is used
func (m *VcdConnectionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Credentials
	if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VcdConnectionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcdConnectionParams) UnmarshalBinary(b []byte) error {
	var res VcdConnectionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
