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
)

// ScheduleProto schedule proto
//
// swagger:model ScheduleProto
type ScheduleProto struct {

	// Specifies the time range within the days of the week. This field is
	// non-empty iff type == kPeriodicTimeWindows.
	PeriodicTimeWindows []*TimeWindow `json:"periodicTimeWindows"`

	// Specifies the time ranges in usecs. This field is non-empty iff
	// type == kCustomIntervals.
	TimeRanges []*TimeRangeUsecs `json:"timeRanges"`

	// Timezone of the user of this ScheduleProto. The timezones have unique
	// names of the form "Area/Location".
	Timezone *string `json:"timezone,omitempty"`

	// Specifies the type of schedule for this ScheduleProto.
	Type *int32 `json:"type,omitempty"`
}

// Validate validates this schedule proto
func (m *ScheduleProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeriodicTimeWindows(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleProto) validatePeriodicTimeWindows(formats strfmt.Registry) error {
	if swag.IsZero(m.PeriodicTimeWindows) { // not required
		return nil
	}

	for i := 0; i < len(m.PeriodicTimeWindows); i++ {
		if swag.IsZero(m.PeriodicTimeWindows[i]) { // not required
			continue
		}

		if m.PeriodicTimeWindows[i] != nil {
			if err := m.PeriodicTimeWindows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("periodicTimeWindows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("periodicTimeWindows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScheduleProto) validateTimeRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.TimeRanges); i++ {
		if swag.IsZero(m.TimeRanges[i]) { // not required
			continue
		}

		if m.TimeRanges[i] != nil {
			if err := m.TimeRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("timeRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("timeRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this schedule proto based on the context it is used
func (m *ScheduleProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePeriodicTimeWindows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleProto) contextValidatePeriodicTimeWindows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PeriodicTimeWindows); i++ {

		if m.PeriodicTimeWindows[i] != nil {

			if swag.IsZero(m.PeriodicTimeWindows[i]) { // not required
				return nil
			}

			if err := m.PeriodicTimeWindows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("periodicTimeWindows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("periodicTimeWindows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScheduleProto) contextValidateTimeRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TimeRanges); i++ {

		if m.TimeRanges[i] != nil {

			if swag.IsZero(m.TimeRanges[i]) { // not required
				return nil
			}

			if err := m.TimeRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("timeRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("timeRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleProto) UnmarshalBinary(b []byte) error {
	var res ScheduleProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
