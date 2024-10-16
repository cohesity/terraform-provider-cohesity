// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubnetInfo SubnetInfo
//
// Subnet information.
//
// swagger:model SubnetInfo
type SubnetInfo struct {

	// Subnet IP.
	SubnetIP *string `json:"subnetIp,omitempty"`

	// Subnet netmask bits.
	NetmaskBits *int32 `json:"netmaskBits,omitempty"`

	// Gateway.
	Gateway *string `json:"gateway,omitempty"`
}

// Validate validates this subnet info
func (m *SubnetInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subnet info based on context it is used
func (m *SubnetInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubnetInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubnetInfo) UnmarshalBinary(b []byte) error {
	var res SubnetInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
