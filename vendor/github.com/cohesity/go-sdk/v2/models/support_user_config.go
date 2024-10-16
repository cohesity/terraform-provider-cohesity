// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SupportUserConfig Specifies the support user's configuration on the Cohesity cluster such as if its shell password has been set and/or sudo access is granted.
//
// swagger:model SupportUserConfig
type SupportUserConfig struct {

	// Specifies if the support user has sudo access.
	EnableSudoAccess *bool `json:"enableSudoAccess,omitempty"`

	// Specifies if the password for the support user has been set.
	PasswordSet *bool `json:"passwordSet,omitempty"`
}

// Validate validates this support user config
func (m *SupportUserConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this support user config based on context it is used
func (m *SupportUserConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupportUserConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportUserConfig) UnmarshalBinary(b []byte) error {
	var res SupportUserConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
