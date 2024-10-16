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

// SchedulerProto Scheduler Proto.
//
// Specifies the scheduler structure which holds the various schedule jobs.
//
// swagger:model SchedulerProto
type SchedulerProto struct {

	// The array of the various scheduler jobs.
	SchedulerJobs []*SchedulerProtoSchedulerJob `json:"schedulerJobs"`
}

// Validate validates this scheduler proto
func (m *SchedulerProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedulerJobs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerProto) validateSchedulerJobs(formats strfmt.Registry) error {
	if swag.IsZero(m.SchedulerJobs) { // not required
		return nil
	}

	for i := 0; i < len(m.SchedulerJobs); i++ {
		if swag.IsZero(m.SchedulerJobs[i]) { // not required
			continue
		}

		if m.SchedulerJobs[i] != nil {
			if err := m.SchedulerJobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schedulerJobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schedulerJobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this scheduler proto based on the context it is used
func (m *SchedulerProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedulerJobs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchedulerProto) contextValidateSchedulerJobs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SchedulerJobs); i++ {

		if m.SchedulerJobs[i] != nil {

			if swag.IsZero(m.SchedulerJobs[i]) { // not required
				return nil
			}

			if err := m.SchedulerJobs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schedulerJobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schedulerJobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchedulerProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulerProto) UnmarshalBinary(b []byte) error {
	var res SchedulerProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
