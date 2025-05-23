// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VerifyIpmiUser Specifies the params for verifying ipmi user.
//
// swagger:model VerifyIpmiUser
type VerifyIpmiUser struct {

	// Specifies the node id of the node for which ipmi user info needs to be verified. This parameter is incompatible with 'nodeIp'.
	NodeID *string `json:"nodeId,omitempty"`

	// Specifies the IP Address of the node for which ipmi user needs to be verified. This parameter is incompatible with 'nodeId'.
	NodeIP *string `json:"nodeIp,omitempty"`

	// Specifies the ipmi username to be verified for given node.
	Username *string `json:"username,omitempty"`

	// Specifies the password of the ipmi user provided that needs to be verified.
	Password *string `json:"password,omitempty"`
}

// Validate validates this verify ipmi user
func (m *VerifyIpmiUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this verify ipmi user based on context it is used
func (m *VerifyIpmiUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VerifyIpmiUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VerifyIpmiUser) UnmarshalBinary(b []byte) error {
	var res VerifyIpmiUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
