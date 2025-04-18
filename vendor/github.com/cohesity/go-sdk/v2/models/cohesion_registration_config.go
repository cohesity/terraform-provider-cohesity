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

// CohesionRegistrationConfig Cohesion Registration Config.
//
// Specifies the Cohesion Registration Config.
//
// swagger:model CohesionRegistrationConfig
type CohesionRegistrationConfig struct {

	// helios reg info
	HeliosRegInfo *CohesionHeliosRegistrationInfo `json:"heliosRegInfo,omitempty"`

	// helios connection info
	HeliosConnectionInfo *CohesionHeliosConnectionInfo `json:"heliosConnectionInfo,omitempty"`
}

// Validate validates this cohesion registration config
func (m *CohesionRegistrationConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeliosRegInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeliosConnectionInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CohesionRegistrationConfig) validateHeliosRegInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.HeliosRegInfo) { // not required
		return nil
	}

	if m.HeliosRegInfo != nil {
		if err := m.HeliosRegInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("heliosRegInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("heliosRegInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CohesionRegistrationConfig) validateHeliosConnectionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.HeliosConnectionInfo) { // not required
		return nil
	}

	if m.HeliosConnectionInfo != nil {
		if err := m.HeliosConnectionInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("heliosConnectionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("heliosConnectionInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cohesion registration config based on the context it is used
func (m *CohesionRegistrationConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeliosRegInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeliosConnectionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CohesionRegistrationConfig) contextValidateHeliosRegInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.HeliosRegInfo != nil {

		if swag.IsZero(m.HeliosRegInfo) { // not required
			return nil
		}

		if err := m.HeliosRegInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("heliosRegInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("heliosRegInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CohesionRegistrationConfig) contextValidateHeliosConnectionInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.HeliosConnectionInfo != nil {

		if swag.IsZero(m.HeliosConnectionInfo) { // not required
			return nil
		}

		if err := m.HeliosConnectionInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("heliosConnectionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("heliosConnectionInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CohesionRegistrationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CohesionRegistrationConfig) UnmarshalBinary(b []byte) error {
	var res CohesionRegistrationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
