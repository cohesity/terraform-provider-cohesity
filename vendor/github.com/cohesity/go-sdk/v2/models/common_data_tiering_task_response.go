// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CommonDataTieringTaskResponse Specifies the data tiering task details.
//
// swagger:model CommonDataTieringTaskResponse
type CommonDataTieringTaskResponse struct {

	// Specifies the id of the data tiering task.
	ID *string `json:"id,omitempty"`

	// Specifies the name of the data tiering task.
	// Required: true
	Name *string `json:"name"`

	// Specifies a description of the data tiering task.
	Description *string `json:"description,omitempty"`

	// Specifies the alerting policy for the data tiering task.
	AlertPolicy *ProtectionGroupAlertingPolicy `json:"alertPolicy,omitempty"`

	// Specifies source for data tiering.
	Source *DataTieringSource `json:"source,omitempty"`

	// Specifies target for data tiering.
	Target *DataTieringTarget `json:"target,omitempty"`

	// Specifies schedule for data tiering
	Schedule *DataTieringSchedule `json:"schedule,omitempty"`

	// Type of data tiering task.
	// 'Downtier' indicates downtiering task.
	// 'Uptier' indicates uptiering task.
	// Required: true
	// Enum: ["Downtier","Uptier"]
	Type *string `json:"type"`

	// Specifies last run details for data tiering task.
	// Read Only: true
	LastRun *DataTieringTaskRun `json:"lastRun,omitempty"`

	// Whether the data tiering task is active or not.
	IsActive *bool `json:"isActive,omitempty"`

	// Whether the data tiering task is paused. New runs are not scheduled
	// for the paused tasks. Active run of the task (if any) is not
	// impacted.
	IsPaused *bool `json:"isPaused,omitempty"`

	// Tracks whether the backup job has actually been deleted.
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// Validate validates this common data tiering task response
func (m *CommonDataTieringTaskResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastRun(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonDataTieringTaskResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CommonDataTieringTaskResponse) validateAlertPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AlertPolicy) { // not required
		return nil
	}

	if m.AlertPolicy != nil {
		if err := m.AlertPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alertPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *CommonDataTieringTaskResponse) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *CommonDataTieringTaskResponse) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *CommonDataTieringTaskResponse) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
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

var commonDataTieringTaskResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Downtier","Uptier"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonDataTieringTaskResponseTypeTypePropEnum = append(commonDataTieringTaskResponseTypeTypePropEnum, v)
	}
}

const (

	// CommonDataTieringTaskResponseTypeDowntier captures enum value "Downtier"
	CommonDataTieringTaskResponseTypeDowntier string = "Downtier"

	// CommonDataTieringTaskResponseTypeUptier captures enum value "Uptier"
	CommonDataTieringTaskResponseTypeUptier string = "Uptier"
)

// prop value enum
func (m *CommonDataTieringTaskResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonDataTieringTaskResponseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonDataTieringTaskResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CommonDataTieringTaskResponse) validateLastRun(formats strfmt.Registry) error {
	if swag.IsZero(m.LastRun) { // not required
		return nil
	}

	if m.LastRun != nil {
		if err := m.LastRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastRun")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastRun")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this common data tiering task response based on the context it is used
func (m *CommonDataTieringTaskResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlertPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastRun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonDataTieringTaskResponse) contextValidateAlertPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AlertPolicy != nil {

		if swag.IsZero(m.AlertPolicy) { // not required
			return nil
		}

		if err := m.AlertPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alertPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *CommonDataTieringTaskResponse) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *CommonDataTieringTaskResponse) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *CommonDataTieringTaskResponse) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if swag.IsZero(m.Schedule) { // not required
			return nil
		}

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

func (m *CommonDataTieringTaskResponse) contextValidateLastRun(ctx context.Context, formats strfmt.Registry) error {

	if m.LastRun != nil {

		if swag.IsZero(m.LastRun) { // not required
			return nil
		}

		if err := m.LastRun.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastRun")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastRun")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonDataTieringTaskResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonDataTieringTaskResponse) UnmarshalBinary(b []byte) error {
	var res CommonDataTieringTaskResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
