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

// ViewUserQuotaSettings Specifies the user quota config on the View.
//
// swagger:model ViewUserQuotaSettings
type ViewUserQuotaSettings struct {

	// Specifies whether user quota is enabled for the View.
	// Required: true
	Enabled *bool `json:"enabled"`

	// Specifies the default user quota policy of the View.
	DefaultQuotaPolicy *QuotaPolicy `json:"defaultQuotaPolicy,omitempty"`
}

// Validate validates this view user quota settings
func (m *ViewUserQuotaSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultQuotaPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewUserQuotaSettings) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ViewUserQuotaSettings) validateDefaultQuotaPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultQuotaPolicy) { // not required
		return nil
	}

	if m.DefaultQuotaPolicy != nil {
		if err := m.DefaultQuotaPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultQuotaPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultQuotaPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this view user quota settings based on the context it is used
func (m *ViewUserQuotaSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultQuotaPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewUserQuotaSettings) contextValidateDefaultQuotaPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultQuotaPolicy != nil {

		if swag.IsZero(m.DefaultQuotaPolicy) { // not required
			return nil
		}

		if err := m.DefaultQuotaPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultQuotaPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("defaultQuotaPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ViewUserQuotaSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewUserQuotaSettings) UnmarshalBinary(b []byte) error {
	var res ViewUserQuotaSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
