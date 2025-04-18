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

// FleetNetworkParams Fleet Network Params.
//
// Specifies various network params for the fleet.
//
// swagger:model FleetNetworkParams
type FleetNetworkParams struct {

	// Specifies vpc for the fleet.
	// Required: true
	Vpc *string `json:"vpc"`

	// Specifies subnet for the fleet.
	// Required: true
	Subnet *string `json:"subnet"`
}

// Validate validates this fleet network params
func (m *FleetNetworkParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FleetNetworkParams) validateVpc(formats strfmt.Registry) error {

	if err := validate.Required("vpc", "body", m.Vpc); err != nil {
		return err
	}

	return nil
}

func (m *FleetNetworkParams) validateSubnet(formats strfmt.Registry) error {

	if err := validate.Required("subnet", "body", m.Subnet); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fleet network params based on context it is used
func (m *FleetNetworkParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FleetNetworkParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FleetNetworkParams) UnmarshalBinary(b []byte) error {
	var res FleetNetworkParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
