// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OracleVlanInfo Oracle Protection Group vlan info
//
// Specifies details about capturing Cohesity cluster VLAN info for Oracle workflow.
//
// swagger:model OracleVlanInfo
type OracleVlanInfo struct {

	// Specifies the list of Ips in this VLAN.
	IPList []string `json:"ipList"`

	// Specifies the gateway of this VLAN.
	Gateway *string `json:"gateway,omitempty"`

	// Specifies the Id of this VLAN.
	ID int32 `json:"id,omitempty"`

	// Specifies the subnet Ip for this VLAN.
	SubnetIP *string `json:"subnetIp,omitempty"`
}

// Validate validates this oracle vlan info
func (m *OracleVlanInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this oracle vlan info based on context it is used
func (m *OracleVlanInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleVlanInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleVlanInfo) UnmarshalBinary(b []byte) error {
	var res OracleVlanInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
