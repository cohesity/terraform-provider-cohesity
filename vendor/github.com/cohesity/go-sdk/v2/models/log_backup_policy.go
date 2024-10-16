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

// LogBackupPolicy Log Backup (Databases) Schedule and Retention.
//
// Specifies log backup settings for a Protection Group.
//
// swagger:model LogBackupPolicy
type LogBackupPolicy struct {

	// Specifies the schedule settings for log backup.
	// Required: true
	Schedule *LogSchedule `json:"schedule"`

	// Specifies the Retention period of a log backup in days, months or years.
	// Required: true
	Retention *Retention `json:"retention"`
}

// Validate validates this log backup policy
func (m *LogBackupPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetention(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogBackupPolicy) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *LogBackupPolicy) validateRetention(formats strfmt.Registry) error {

	if err := validate.Required("retention", "body", m.Retention); err != nil {
		return err
	}

	if m.Retention != nil {
		if err := m.Retention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retention")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log backup policy based on the context it is used
func (m *LogBackupPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRetention(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogBackupPolicy) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *LogBackupPolicy) contextValidateRetention(ctx context.Context, formats strfmt.Registry) error {

	if m.Retention != nil {

		if err := m.Retention.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retention")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogBackupPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogBackupPolicy) UnmarshalBinary(b []byte) error {
	var res LogBackupPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
