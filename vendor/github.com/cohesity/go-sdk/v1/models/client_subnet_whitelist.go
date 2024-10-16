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

// ClientSubnetWhitelist ClientSubnetWhitelist is the struct containing the external client subnets
// that can communicate with this cluster.
//
// swagger:model ClientSubnetWhitelist
type ClientSubnetWhitelist struct {

	// ClientSubnetWhitelist is the list of client subnets.
	ClientSubnetWhitelist []*PrivateSubnet `json:"clientSubnetWhitelist"`
}

// Validate validates this client subnet whitelist
func (m *ClientSubnetWhitelist) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientSubnetWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientSubnetWhitelist) validateClientSubnetWhitelist(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientSubnetWhitelist) { // not required
		return nil
	}

	for i := 0; i < len(m.ClientSubnetWhitelist); i++ {
		if swag.IsZero(m.ClientSubnetWhitelist[i]) { // not required
			continue
		}

		if m.ClientSubnetWhitelist[i] != nil {
			if err := m.ClientSubnetWhitelist[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientSubnetWhitelist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clientSubnetWhitelist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this client subnet whitelist based on the context it is used
func (m *ClientSubnetWhitelist) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientSubnetWhitelist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientSubnetWhitelist) contextValidateClientSubnetWhitelist(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClientSubnetWhitelist); i++ {

		if m.ClientSubnetWhitelist[i] != nil {

			if swag.IsZero(m.ClientSubnetWhitelist[i]) { // not required
				return nil
			}

			if err := m.ClientSubnetWhitelist[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientSubnetWhitelist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clientSubnetWhitelist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientSubnetWhitelist) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientSubnetWhitelist) UnmarshalBinary(b []byte) error {
	var res ClientSubnetWhitelist
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
