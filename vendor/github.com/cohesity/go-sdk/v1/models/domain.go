// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Domain Domain.
//
// Specifies a domain and its trusted domains.
//
// swagger:model Domain
type Domain struct {

	// Specifies the domain name.
	DomainName *string `json:"domainName,omitempty"`

	// Specifies a list of trusted domains of this domain.
	TrustedDomains []string `json:"trustedDomains"`
}

// Validate validates this domain
func (m *Domain) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain based on context it is used
func (m *Domain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Domain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Domain) UnmarshalBinary(b []byte) error {
	var res Domain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
