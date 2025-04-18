// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegisterTrustedCas Specifies the parameters to register a Certificate.
//
// swagger:model RegisterTrustedCas
type RegisterTrustedCas struct {

	// Specifies the certificates to be imported.
	// Required: true
	// Min Items: 1
	Certificates []*TrustedCaRequest `json:"certificates"`

	// Specifies if the certificates are only to be validated.
	OnlyValidate *bool `json:"onlyValidate,omitempty"`
}

// Validate validates this register trusted cas
func (m *RegisterTrustedCas) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterTrustedCas) validateCertificates(formats strfmt.Registry) error {

	if err := validate.Required("certificates", "body", m.Certificates); err != nil {
		return err
	}

	iCertificatesSize := int64(len(m.Certificates))

	if err := validate.MinItems("certificates", "body", iCertificatesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Certificates); i++ {
		if swag.IsZero(m.Certificates[i]) { // not required
			continue
		}

		if m.Certificates[i] != nil {
			if err := m.Certificates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this register trusted cas based on the context it is used
func (m *RegisterTrustedCas) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterTrustedCas) contextValidateCertificates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Certificates); i++ {

		if m.Certificates[i] != nil {

			if swag.IsZero(m.Certificates[i]) { // not required
				return nil
			}

			if err := m.Certificates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegisterTrustedCas) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterTrustedCas) UnmarshalBinary(b []byte) error {
	var res RegisterTrustedCas
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
