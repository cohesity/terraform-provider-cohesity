// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PhysicalNodeConfiguration Physical Node Configuration.
//
// Specifies the configuration for a node in the Cluster.
//
// swagger:model PhysicalNodeConfiguration
type PhysicalNodeConfiguration struct {

	// Specifies the Node ID for this node.
	NodeID *int64 `json:"nodeId,omitempty"`

	// Specifies the Node IP address for this node.
	NodeIP *string `json:"nodeIp,omitempty"`

	// Specifies IPMI IP for this node.
	NodeIpmiIP *string `json:"nodeIpmiIp,omitempty"`

	// Specifies whether to use the Node for compute only.
	UseAsComputeNode *bool `json:"useAsComputeNode,omitempty"`
}

// Validate validates this physical node configuration
func (m *PhysicalNodeConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this physical node configuration based on context it is used
func (m *PhysicalNodeConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PhysicalNodeConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhysicalNodeConfiguration) UnmarshalBinary(b []byte) error {
	var res PhysicalNodeConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
