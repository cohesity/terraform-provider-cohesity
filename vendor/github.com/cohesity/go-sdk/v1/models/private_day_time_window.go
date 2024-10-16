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

// PrivateDayTimeWindow Message encapsulating day time window.
//
// swagger:model PrivateDayTimeWindow
type PrivateDayTimeWindow struct {

	// Indicates the end time of the window.
	EndTime *PrivateDayTime `json:"endTime,omitempty"`

	// Indicates the start time of the window.
	StartTime *PrivateDayTime `json:"startTime,omitempty"`
}

// Validate validates this private day time window
func (m *PrivateDayTimeWindow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateDayTimeWindow) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if m.EndTime != nil {
		if err := m.EndTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endTime")
			}
			return err
		}
	}

	return nil
}

func (m *PrivateDayTimeWindow) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if m.StartTime != nil {
		if err := m.StartTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startTime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this private day time window based on the context it is used
func (m *PrivateDayTimeWindow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateDayTimeWindow) contextValidateEndTime(ctx context.Context, formats strfmt.Registry) error {

	if m.EndTime != nil {

		if swag.IsZero(m.EndTime) { // not required
			return nil
		}

		if err := m.EndTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endTime")
			}
			return err
		}
	}

	return nil
}

func (m *PrivateDayTimeWindow) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if m.StartTime != nil {

		if swag.IsZero(m.StartTime) { // not required
			return nil
		}

		if err := m.StartTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateDayTimeWindow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateDayTimeWindow) UnmarshalBinary(b []byte) error {
	var res PrivateDayTimeWindow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
