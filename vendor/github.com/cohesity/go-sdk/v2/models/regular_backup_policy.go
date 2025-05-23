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

// RegularBackupPolicy Incremental, Full and Retention Policy.
//
// Specifies the Incremental and Full policy settings and also the common Retention policy settings."
//
// swagger:model RegularBackupPolicy
type RegularBackupPolicy struct {

	// Specifies the Incremental backup schedule and retention of a Protection Policy.
	Incremental *IncrementalBackupPolicy `json:"incremental,omitempty"`

	// Specifies the Full backup schedule of a Protection Policy.
	Full *FullBackupPolicy `json:"full,omitempty"`

	// Specifies multiple schedules and retentions for full backup. Specify either of the 'full' or 'fullBackups' values. Its recommended to use 'fullBaackups' value since 'full' will be deprecated after few releases.
	FullBackups []*FullScheduleAndRetention `json:"fullBackups"`

	// Specifies the Retention period for incremental and full backup in days, months or years.
	Retention *Retention `json:"retention,omitempty"`

	// Specifies the primary backup target settings for regular backups. Specifying this field shows that instead of local backups on Cohesity cluster, primary backup location is different such as Cloud Archives like s3 or azure.
	PrimaryBackupTarget *PrimaryBackupTarget `json:"primaryBackupTarget,omitempty"`
}

// Validate validates this regular backup policy
func (m *RegularBackupPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncremental(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFull(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullBackups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetention(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryBackupTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegularBackupPolicy) validateIncremental(formats strfmt.Registry) error {
	if swag.IsZero(m.Incremental) { // not required
		return nil
	}

	if m.Incremental != nil {
		if err := m.Incremental.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incremental")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incremental")
			}
			return err
		}
	}

	return nil
}

func (m *RegularBackupPolicy) validateFull(formats strfmt.Registry) error {
	if swag.IsZero(m.Full) { // not required
		return nil
	}

	if m.Full != nil {
		if err := m.Full.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("full")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("full")
			}
			return err
		}
	}

	return nil
}

func (m *RegularBackupPolicy) validateFullBackups(formats strfmt.Registry) error {
	if swag.IsZero(m.FullBackups) { // not required
		return nil
	}

	for i := 0; i < len(m.FullBackups); i++ {
		if swag.IsZero(m.FullBackups[i]) { // not required
			continue
		}

		if m.FullBackups[i] != nil {
			if err := m.FullBackups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fullBackups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fullBackups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RegularBackupPolicy) validateRetention(formats strfmt.Registry) error {
	if swag.IsZero(m.Retention) { // not required
		return nil
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

func (m *RegularBackupPolicy) validatePrimaryBackupTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.PrimaryBackupTarget) { // not required
		return nil
	}

	if m.PrimaryBackupTarget != nil {
		if err := m.PrimaryBackupTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryBackupTarget")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryBackupTarget")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this regular backup policy based on the context it is used
func (m *RegularBackupPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIncremental(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFull(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullBackups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRetention(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryBackupTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegularBackupPolicy) contextValidateIncremental(ctx context.Context, formats strfmt.Registry) error {

	if m.Incremental != nil {

		if swag.IsZero(m.Incremental) { // not required
			return nil
		}

		if err := m.Incremental.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incremental")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incremental")
			}
			return err
		}
	}

	return nil
}

func (m *RegularBackupPolicy) contextValidateFull(ctx context.Context, formats strfmt.Registry) error {

	if m.Full != nil {

		if swag.IsZero(m.Full) { // not required
			return nil
		}

		if err := m.Full.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("full")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("full")
			}
			return err
		}
	}

	return nil
}

func (m *RegularBackupPolicy) contextValidateFullBackups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FullBackups); i++ {

		if m.FullBackups[i] != nil {

			if swag.IsZero(m.FullBackups[i]) { // not required
				return nil
			}

			if err := m.FullBackups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fullBackups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fullBackups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RegularBackupPolicy) contextValidateRetention(ctx context.Context, formats strfmt.Registry) error {

	if m.Retention != nil {

		if swag.IsZero(m.Retention) { // not required
			return nil
		}

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

func (m *RegularBackupPolicy) contextValidatePrimaryBackupTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryBackupTarget != nil {

		if swag.IsZero(m.PrimaryBackupTarget) { // not required
			return nil
		}

		if err := m.PrimaryBackupTarget.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryBackupTarget")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryBackupTarget")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegularBackupPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegularBackupPolicy) UnmarshalBinary(b []byte) error {
	var res RegularBackupPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
