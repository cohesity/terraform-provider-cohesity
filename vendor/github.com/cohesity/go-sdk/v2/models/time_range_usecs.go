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

// TimeRangeUsecs Specifies a valid time range in micro seconds.
//
// swagger:model TimeRangeUsecs
type TimeRangeUsecs struct {

	// Specifies the start time of this time range.
	// Required: true
	StartTimeUsecs *int64 `json:"startTimeUsecs"`

	// Specifies the end time of this time range.
	// Required: true
	EndTimeUsecs *int64 `json:"endTimeUsecs"`
}

// Validate validates this time range usecs
func (m *TimeRangeUsecs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStartTimeUsecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTimeUsecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeRangeUsecs) validateStartTimeUsecs(formats strfmt.Registry) error {

	if err := validate.Required("startTimeUsecs", "body", m.StartTimeUsecs); err != nil {
		return err
	}

	return nil
}

func (m *TimeRangeUsecs) validateEndTimeUsecs(formats strfmt.Registry) error {

	if err := validate.Required("endTimeUsecs", "body", m.EndTimeUsecs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this time range usecs based on context it is used
func (m *TimeRangeUsecs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeRangeUsecs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeRangeUsecs) UnmarshalBinary(b []byte) error {
	var res TimeRangeUsecs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
