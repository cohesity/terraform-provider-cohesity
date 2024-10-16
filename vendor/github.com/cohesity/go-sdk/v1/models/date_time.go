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

// DateTime Message encapsulating date and time information.
//
// swagger:model DateTime
type DateTime struct {

	// Indicates day of the month.
	DayOfTheMonth *int32 `json:"dayOfTheMonth,omitempty"`

	// Indicates month for specific date.
	Month *int32 `json:"month,omitempty"`

	// Indicates the time in 24 hrs format.
	Time *Time `json:"time,omitempty"`

	// Indicates year for specific date.
	Year *int32 `json:"year,omitempty"`
}

// Validate validates this date time
func (m *DateTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DateTime) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if m.Time != nil {
		if err := m.Time.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this date time based on the context it is used
func (m *DateTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DateTime) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if m.Time != nil {

		if swag.IsZero(m.Time) { // not required
			return nil
		}

		if err := m.Time.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DateTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DateTime) UnmarshalBinary(b []byte) error {
	var res DateTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
