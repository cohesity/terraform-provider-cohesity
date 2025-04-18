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

// IPPool IP pools from the vlan ip addresses, the IPs in a pool goes together.
//
// swagger:model IpPool
type IPPool struct {

	// Name of the IP pool.
	// Required: true
	Name *string `json:"name"`

	// IP addresses.
	// Required: true
	Ips []string `json:"ips"`
}

// Validate validates this Ip pool
func (m *IPPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPPool) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IPPool) validateIps(formats strfmt.Registry) error {

	if err := validate.Required("ips", "body", m.Ips); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Ip pool based on context it is used
func (m *IPPool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPPool) UnmarshalBinary(b []byte) error {
	var res IPPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
