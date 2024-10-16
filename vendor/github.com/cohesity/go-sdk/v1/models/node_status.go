// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeStatus Node Status.
//
// Specifies the status of each node in the cluster being created.
//
// swagger:model NodeStatus
type NodeStatus struct {

	// Specifies an optional message relating to the node status.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Specifies the IPMI IP of the node (if physical cluster).
	IpmiIP *string `json:"ipmiIp,omitempty"`

	// Specifies the ID of the node.
	NodeID *int64 `json:"nodeId,omitempty"`

	// For physical nodes this will specify the IP address of the node.
	NodeIP *string `json:"nodeIp,omitempty"`
}

// Validate validates this node status
func (m *NodeStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node status based on context it is used
func (m *NodeStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeStatus) UnmarshalBinary(b []byte) error {
	var res NodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
