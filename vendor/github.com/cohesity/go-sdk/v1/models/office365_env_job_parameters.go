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

// Office365EnvJobParameters Office365 Environment Job Parameters.
//
// Specifies Office365 parameters applicable for all Office365 Environment
// type Protection Sources in a Protection Job. This encapsulates both OneDrive
// & Mailbox parameters.
//
// swagger:model Office365EnvJobParameters
type Office365EnvJobParameters struct {

	// Specifies OneDrive backup parameters.
	OnedriveParameters *OneDriveEnvJobParameters `json:"onedriveParameters,omitempty"`

	// Specifies Outlook backup parameters.
	OutlookParameters *OutlookEnvJobParameters `json:"outlookParameters,omitempty"`
}

// Validate validates this office365 env job parameters
func (m *Office365EnvJobParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOnedriveParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutlookParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Office365EnvJobParameters) validateOnedriveParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.OnedriveParameters) { // not required
		return nil
	}

	if m.OnedriveParameters != nil {
		if err := m.OnedriveParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onedriveParameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onedriveParameters")
			}
			return err
		}
	}

	return nil
}

func (m *Office365EnvJobParameters) validateOutlookParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.OutlookParameters) { // not required
		return nil
	}

	if m.OutlookParameters != nil {
		if err := m.OutlookParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outlookParameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outlookParameters")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this office365 env job parameters based on the context it is used
func (m *Office365EnvJobParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOnedriveParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutlookParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Office365EnvJobParameters) contextValidateOnedriveParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.OnedriveParameters != nil {

		if swag.IsZero(m.OnedriveParameters) { // not required
			return nil
		}

		if err := m.OnedriveParameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onedriveParameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onedriveParameters")
			}
			return err
		}
	}

	return nil
}

func (m *Office365EnvJobParameters) contextValidateOutlookParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.OutlookParameters != nil {

		if swag.IsZero(m.OutlookParameters) { // not required
			return nil
		}

		if err := m.OutlookParameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outlookParameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outlookParameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Office365EnvJobParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Office365EnvJobParameters) UnmarshalBinary(b []byte) error {
	var res Office365EnvJobParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
