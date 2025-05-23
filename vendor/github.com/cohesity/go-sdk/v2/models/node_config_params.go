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

// NodeConfigParams Specifies the configuration of the nodes.
//
// swagger:model NodeConfigParams
type NodeConfigParams struct {

	// Specifies the node ID for this node.
	// Required: true
	ID *int64 `json:"id"`

	// Specifies the IP address for the node.
	// Required: true
	IP *string `json:"ip"`

	// Specifies whether to use the node for compute only.
	IsComputeNode bool `json:"isComputeNode,omitempty"`

	// Specifies IPMI IP for the node.
	IpmiIP string `json:"ipmiIp,omitempty"`
}

// Validate validates this node config params
func (m *NodeConfigParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeConfigParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NodeConfigParams) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this node config params based on context it is used
func (m *NodeConfigParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeConfigParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeConfigParams) UnmarshalBinary(b []byte) error {
	var res NodeConfigParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
