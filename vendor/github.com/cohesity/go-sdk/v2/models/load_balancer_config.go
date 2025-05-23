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

// LoadBalancerConfig Load balancer VIP config for OneHelios cluster.
//
// swagger:model LoadBalancerConfig
type LoadBalancerConfig struct {

	// Specifies host name of the Helios endpoint.
	HostName *string `json:"hostName,omitempty"`

	// Specifies list of Virtual IP Addresses.
	VirtualIPVec []string `json:"virtualIpVec"`

	// Specifies subnet.
	Subnet *SubnetDefinition `json:"subnet,omitempty"`

	// Specifies gateway.
	Gateway *string `json:"gateway,omitempty"`
}

// Validate validates this load balancer config
func (m *LoadBalancerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancerConfig) validateSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this load balancer config based on the context it is used
func (m *LoadBalancerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubnet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancerConfig) contextValidateSubnet(ctx context.Context, formats strfmt.Registry) error {

	if m.Subnet != nil {

		if swag.IsZero(m.Subnet) { // not required
			return nil
		}

		if err := m.Subnet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadBalancerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadBalancerConfig) UnmarshalBinary(b []byte) error {
	var res LoadBalancerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
