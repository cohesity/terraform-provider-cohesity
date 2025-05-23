// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthHeader Specifies structure of request header
//
// swagger:model AuthHeader
type AuthHeader struct {

	// Specifies key for the header.
	Key string `json:"key,omitempty"`

	// Specifies value for the header.
	Value string `json:"value,omitempty"`
}

// Validate validates this auth header
func (m *AuthHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auth header based on context it is used
func (m *AuthHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthHeader) UnmarshalBinary(b []byte) error {
	var res AuthHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
