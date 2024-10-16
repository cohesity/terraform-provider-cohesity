// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdDomain AD Domain.
//
// Specifies information about an AD Domain.
//
// swagger:model AdDomain
type AdDomain struct {

	// Specifies DNS root.
	DNSRoot *string `json:"dnsRoot,omitempty"`

	// Specifies AD forest name.
	Forest *string `json:"forest,omitempty"`

	// Specifies Identity information of the domain.
	Identity *AdDomainIdentity `json:"identity,omitempty"`

	// Specifies AD NetBIOS name.
	NetbiosName *string `json:"netbiosName,omitempty"`

	// Specifies parent domain name.
	ParentDomain *string `json:"parentDomain,omitempty"`

	// Specifies tombstone time in days.
	TombstoneDays *uint32 `json:"tombstoneDays,omitempty"`
}

// Validate validates this ad domain
func (m *AdDomain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdDomain) validateIdentity(formats strfmt.Registry) error {
	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {
		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ad domain based on the context it is used
func (m *AdDomain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdDomain) contextValidateIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.Identity != nil {

		if swag.IsZero(m.Identity) { // not required
			return nil
		}

		if err := m.Identity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdDomain) UnmarshalBinary(b []byte) error {
	var res AdDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
