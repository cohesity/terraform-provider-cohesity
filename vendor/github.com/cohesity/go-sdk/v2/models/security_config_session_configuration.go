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

// SecurityConfigSessionConfiguration Specifies configuration for user sessions.
//
// swagger:model SecurityConfigSessionConfiguration
type SecurityConfigSessionConfiguration struct {

	// Specifies absolute session expiration time in seconds.
	// Minimum: 1
	AbsoluteTimeout *int64 `json:"absoluteTimeout,omitempty"`

	// Specifies inactivity session expiration time in seconds.
	// Minimum: 1
	InactivityTimeout *int64 `json:"inactivityTimeout,omitempty"`

	// Specifies if limitations on number of active sessions is enabled or not.
	LimitSessions *bool `json:"limitSessions,omitempty"`

	// Specifies maximum number of active sessions allowed per user. This applies only when limitSessions is enabled.
	// Minimum: 1
	SessionLimitPerUser *int64 `json:"sessionLimitPerUser,omitempty"`

	// Specifies maximum number of active sessions allowed system wide. This applies only when limitSessions is enabled.
	// Minimum: 1
	SessionLimitSystemWide *int64 `json:"sessionLimitSystemWide,omitempty"`
}

// Validate validates this security config session configuration
func (m *SecurityConfigSessionConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbsoluteTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInactivityTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionLimitPerUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionLimitSystemWide(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityConfigSessionConfiguration) validateAbsoluteTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.AbsoluteTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("absoluteTimeout", "body", *m.AbsoluteTimeout, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *SecurityConfigSessionConfiguration) validateInactivityTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.InactivityTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("inactivityTimeout", "body", *m.InactivityTimeout, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *SecurityConfigSessionConfiguration) validateSessionLimitPerUser(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionLimitPerUser) { // not required
		return nil
	}

	if err := validate.MinimumInt("sessionLimitPerUser", "body", *m.SessionLimitPerUser, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *SecurityConfigSessionConfiguration) validateSessionLimitSystemWide(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionLimitSystemWide) { // not required
		return nil
	}

	if err := validate.MinimumInt("sessionLimitSystemWide", "body", *m.SessionLimitSystemWide, 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this security config session configuration based on context it is used
func (m *SecurityConfigSessionConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityConfigSessionConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityConfigSessionConfiguration) UnmarshalBinary(b []byte) error {
	var res SecurityConfigSessionConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
