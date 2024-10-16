// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetappVersionTuple Specifies the Netapp version Tuple.
//
// swagger:model NetappVersionTuple
type NetappVersionTuple struct {

	// Netapp generation.
	Generation *int32 `json:"generation,omitempty"`

	// Major version number.
	MajorVersion *int32 `json:"majorVersion,omitempty"`

	// Minor version number.
	MinorVersion *int32 `json:"minorVersion,omitempty"`
}

// Validate validates this netapp version tuple
func (m *NetappVersionTuple) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this netapp version tuple based on context it is used
func (m *NetappVersionTuple) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetappVersionTuple) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetappVersionTuple) UnmarshalBinary(b []byte) error {
	var res NetappVersionTuple
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
