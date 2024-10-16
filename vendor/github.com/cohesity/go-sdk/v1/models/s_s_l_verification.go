// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SSLVerification Contains information regarding SSL verification
//
// swagger:model SSLVerification
type SSLVerification struct {

	// Contains the contents of CA cert/cert chain.
	CaCertificate *string `json:"caCertificate,omitempty"`

	// Whether SSL verification is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Optionally contains the server sha256 key.
	ServerSha256Key []uint8 `json:"serverSha256Key"`
}

// Validate validates this s s l verification
func (m *SSLVerification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s s l verification based on context it is used
func (m *SSLVerification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SSLVerification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSLVerification) UnmarshalBinary(b []byte) error {
	var res SSLVerification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
