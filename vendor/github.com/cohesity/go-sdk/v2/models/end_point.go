// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndPoint Interface Information
//
// Specifies information about a node interface.
//
// swagger:model EndPoint
type EndPoint struct {

	// Index of the interface as given by 'ip a' command.
	Index int32 `json:"index,omitempty"`

	// Name of the interface like bond0.
	Name string `json:"name,omitempty"`

	// IP addresses on the interface
	IPAddresses []string `json:"ipAddresses"`
}

// Validate validates this end point
func (m *EndPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this end point based on context it is used
func (m *EndPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndPoint) UnmarshalBinary(b []byte) error {
	var res EndPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
