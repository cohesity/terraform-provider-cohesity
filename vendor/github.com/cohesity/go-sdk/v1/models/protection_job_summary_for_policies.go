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

// ProtectionJobSummaryForPolicies ProtectionJobSummaryForPolicies is the summary of a Protection
// Jobs associated with the Specified Protection Policy. This is only
// populated for a policy of type kRegular.
//
// swagger:model ProtectionJobSummaryForPolicies
type ProtectionJobSummaryForPolicies struct {

	// Specifies details about the last Backup task.
	// A Backup task captures the original backup snapshots.
	BackupRun *BackupRun `json:"backupRun,omitempty"`

	// Specifies details about the Copy tasks of the Job Run.
	// A Copy task copies the captured snapshots to an external target
	// or a Remote Cohesity Cluster.
	CopyRuns []*CopyRun `json:"copyRuns"`

	// Specifies the Protection job information.
	ProtectionJob *ProtectionJob `json:"protectionJob,omitempty"`
}

// Validate validates this protection job summary for policies
func (m *ProtectionJobSummaryForPolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupRun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCopyRuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionJobSummaryForPolicies) validateBackupRun(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupRun) { // not required
		return nil
	}

	if m.BackupRun != nil {
		if err := m.BackupRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupRun")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupRun")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionJobSummaryForPolicies) validateCopyRuns(formats strfmt.Registry) error {
	if swag.IsZero(m.CopyRuns) { // not required
		return nil
	}

	for i := 0; i < len(m.CopyRuns); i++ {
		if swag.IsZero(m.CopyRuns[i]) { // not required
			continue
		}

		if m.CopyRuns[i] != nil {
			if err := m.CopyRuns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copyRuns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copyRuns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionJobSummaryForPolicies) validateProtectionJob(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionJob) { // not required
		return nil
	}

	if m.ProtectionJob != nil {
		if err := m.ProtectionJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protectionJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protectionJob")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this protection job summary for policies based on the context it is used
func (m *ProtectionJobSummaryForPolicies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupRun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCopyRuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionJobSummaryForPolicies) contextValidateBackupRun(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupRun != nil {

		if swag.IsZero(m.BackupRun) { // not required
			return nil
		}

		if err := m.BackupRun.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupRun")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupRun")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionJobSummaryForPolicies) contextValidateCopyRuns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CopyRuns); i++ {

		if m.CopyRuns[i] != nil {

			if swag.IsZero(m.CopyRuns[i]) { // not required
				return nil
			}

			if err := m.CopyRuns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copyRuns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copyRuns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionJobSummaryForPolicies) contextValidateProtectionJob(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionJob != nil {

		if swag.IsZero(m.ProtectionJob) { // not required
			return nil
		}

		if err := m.ProtectionJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protectionJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protectionJob")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionJobSummaryForPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionJobSummaryForPolicies) UnmarshalBinary(b []byte) error {
	var res ProtectionJobSummaryForPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
