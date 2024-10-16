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

// ProtectionRunInstance Protection Job Run Instance.
//
// Specifies the status of one Job Run.
// A Job Run can have one Backup Run and zero or more Copy Runs.
//
// swagger:model ProtectionRunInstance
type ProtectionRunInstance struct {

	// Specifies details about the Backup task.
	// A Backup task captures the original backup snapshots.
	BackupRun *BackupRun `json:"backupRun,omitempty"`

	// Array of Copy Run Tasks.
	//
	// Specifies details about the Copy tasks of this Job Run.
	// A Copy task copies the captured snapshots to an external target
	// or a Remote Cohesity Cluster.
	CopyRun []*CopyRun `json:"copyRun"`

	// Specifies the id of the Protection Job that was run.
	JobID *int64 `json:"jobId,omitempty"`

	// Specifies the name of the Protection Job name that was run.
	JobName *string `json:"jobName,omitempty"`

	// Specifies the globally unique id of the Protection Job that was run.
	JobUID struct {
		UniversalID
	} `json:"jobUid,omitempty"`

	// Specifies the shell information about the protection run. This
	// will only be populated if OnlyReturnShellInfo is sent as true.
	ProtectionShellInfo *ProtectionShellInfo `json:"protectionShellInfo,omitempty"`

	// Specifies the Storage Domain (View Box) to store the backed up data.
	// Specify the id of the Storage Domain (View Box).
	ViewBoxID *int64 `json:"viewBoxId,omitempty"`
}

// Validate validates this protection run instance
func (m *ProtectionRunInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupRun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCopyRun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionShellInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionRunInstance) validateBackupRun(formats strfmt.Registry) error {
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

func (m *ProtectionRunInstance) validateCopyRun(formats strfmt.Registry) error {
	if swag.IsZero(m.CopyRun) { // not required
		return nil
	}

	for i := 0; i < len(m.CopyRun); i++ {
		if swag.IsZero(m.CopyRun[i]) { // not required
			continue
		}

		if m.CopyRun[i] != nil {
			if err := m.CopyRun[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copyRun" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copyRun" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionRunInstance) validateJobUID(formats strfmt.Registry) error {
	if swag.IsZero(m.JobUID) { // not required
		return nil
	}

	return nil
}

func (m *ProtectionRunInstance) validateProtectionShellInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionShellInfo) { // not required
		return nil
	}

	if m.ProtectionShellInfo != nil {
		if err := m.ProtectionShellInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protectionShellInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protectionShellInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this protection run instance based on the context it is used
func (m *ProtectionRunInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupRun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCopyRun(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJobUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtectionShellInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionRunInstance) contextValidateBackupRun(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProtectionRunInstance) contextValidateCopyRun(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CopyRun); i++ {

		if m.CopyRun[i] != nil {

			if swag.IsZero(m.CopyRun[i]) { // not required
				return nil
			}

			if err := m.CopyRun[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copyRun" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copyRun" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionRunInstance) contextValidateJobUID(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ProtectionRunInstance) contextValidateProtectionShellInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectionShellInfo != nil {

		if swag.IsZero(m.ProtectionShellInfo) { // not required
			return nil
		}

		if err := m.ProtectionShellInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protectionShellInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protectionShellInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionRunInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionRunInstance) UnmarshalBinary(b []byte) error {
	var res ProtectionRunInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
