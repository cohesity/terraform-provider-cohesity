// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeInterfacePair Node Ids and interface tuple.
//
// Specifies a node ID and interface tuple.
//
// swagger:model NodeInterfacePair
type NodeInterfacePair struct {

	// Specifies the name of the interface.
	IfaceName *string `json:"ifaceName,omitempty"`

	// Specifies list of node IDs.
	NodeID *int64 `json:"nodeId,omitempty"`
}

// Validate validates this node interface pair
func (m *NodeInterfacePair) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node interface pair based on context it is used
func (m *NodeInterfacePair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeInterfacePair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeInterfacePair) UnmarshalBinary(b []byte) error {
	var res NodeInterfacePair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
