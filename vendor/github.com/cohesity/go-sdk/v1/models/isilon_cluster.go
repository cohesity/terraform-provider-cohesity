// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IsilonCluster Isilon Cluster.
//
// Specifies information about an Isilon Cluster.
//
// swagger:model IsilonCluster
type IsilonCluster struct {

	// Specifies the API version of an Isilon Cluster as string.
	APIVersionStr *string `json:"apiVersionStr,omitempty"`

	// Specifies the description of an Isilon Cluster.
	Description *string `json:"description,omitempty"`

	// Specifies a globally unique id of an Isilon Cluster.
	GUID *string `json:"guid,omitempty"`

	// Specifies the version of an Isilon Cluster.
	Version *string `json:"version,omitempty"`
}

// Validate validates this isilon cluster
func (m *IsilonCluster) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this isilon cluster based on context it is used
func (m *IsilonCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IsilonCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IsilonCluster) UnmarshalBinary(b []byte) error {
	var res IsilonCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
