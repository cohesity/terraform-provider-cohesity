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

// GlobalBandwidthSettings Global Target Bandwidth Settings
//
// Specifies the bandwidth setting of the External Target.
//
// swagger:model GlobalBandwidthSettings
type GlobalBandwidthSettings struct {

	// archival params
	ArchivalParams *ArchivalBandwidthSettings `json:"archivalParams,omitempty"`

	// tiering params
	TieringParams *TieringBandwidthSettings `json:"tieringParams,omitempty"`
}

// Validate validates this global bandwidth settings
func (m *GlobalBandwidthSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchivalParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTieringParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalBandwidthSettings) validateArchivalParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchivalParams) { // not required
		return nil
	}

	if m.ArchivalParams != nil {
		if err := m.ArchivalParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archivalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archivalParams")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalBandwidthSettings) validateTieringParams(formats strfmt.Registry) error {
	if swag.IsZero(m.TieringParams) { // not required
		return nil
	}

	if m.TieringParams != nil {
		if err := m.TieringParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tieringParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tieringParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this global bandwidth settings based on the context it is used
func (m *GlobalBandwidthSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchivalParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTieringParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalBandwidthSettings) contextValidateArchivalParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ArchivalParams != nil {

		if swag.IsZero(m.ArchivalParams) { // not required
			return nil
		}

		if err := m.ArchivalParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archivalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archivalParams")
			}
			return err
		}
	}

	return nil
}

func (m *GlobalBandwidthSettings) contextValidateTieringParams(ctx context.Context, formats strfmt.Registry) error {

	if m.TieringParams != nil {

		if swag.IsZero(m.TieringParams) { // not required
			return nil
		}

		if err := m.TieringParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tieringParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tieringParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalBandwidthSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalBandwidthSettings) UnmarshalBinary(b []byte) error {
	var res GlobalBandwidthSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
