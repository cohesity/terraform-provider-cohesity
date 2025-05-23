// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterHardwareInfo Specifies a hardware type for motherboard of the nodes that make dnsServerIps this Cohesity Cluster
//
// swagger:model ClusterHardwareInfo
type ClusterHardwareInfo struct {

	// hardware models
	HardwareModels []string `json:"hardwareModels,omitempty"`

	// hardware vendors
	HardwareVendors []string `json:"hardwareVendors,omitempty"`
}

// Validate validates this cluster hardware info
func (m *ClusterHardwareInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster hardware info based on context it is used
func (m *ClusterHardwareInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterHardwareInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterHardwareInfo) UnmarshalBinary(b []byte) error {
	var res ClusterHardwareInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
