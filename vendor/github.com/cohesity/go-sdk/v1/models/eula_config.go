// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EulaConfig EULA Configuration.
//
// Specifies the End User License Agreement acceptance information.
//
// swagger:model EulaConfig
type EulaConfig struct {

	// Specifies the license key.
	// Required: true
	LicenseKey *string `json:"licenseKey"`

	// Specifies the login account name for the Cohesity user who accepted
	// the End User License Agreement.
	// Read Only: true
	SignedByUser *string `json:"signedByUser,omitempty"`

	// Specifies the time that the End User License Agreement was accepted.
	// Read Only: true
	SignedTime *int64 `json:"signedTime,omitempty"`

	// Specifies the version of the End User License Agreement that was accepted.
	// Required: true
	SignedVersion *int64 `json:"signedVersion"`
}

// Validate validates this eula config
func (m *EulaConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicenseKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignedVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EulaConfig) validateLicenseKey(formats strfmt.Registry) error {

	if err := validate.Required("licenseKey", "body", m.LicenseKey); err != nil {
		return err
	}

	return nil
}

func (m *EulaConfig) validateSignedVersion(formats strfmt.Registry) error {

	if err := validate.Required("signedVersion", "body", m.SignedVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this eula config based on the context it is used
func (m *EulaConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSignedByUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EulaConfig) contextValidateSignedByUser(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "signedByUser", "body", m.SignedByUser); err != nil {
		return err
	}

	return nil
}

func (m *EulaConfig) contextValidateSignedTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "signedTime", "body", m.SignedTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EulaConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EulaConfig) UnmarshalBinary(b []byte) error {
	var res EulaConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
