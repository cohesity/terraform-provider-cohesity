// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CheckRestoreAppTaskResultFullSnapshotInfo Encapsulates info about a full or incremental snapshot to be displayed to
// the user.
//
// swagger:model CheckRestoreAppTaskResult_FullSnapshotInfo
type CheckRestoreAppTaskResultFullSnapshotInfo struct {

	// The type of backup, whether it is full or incremental.
	BackupType *int32 `json:"backupType,omitempty"`

	// Number of bytes to transfer in order to restore to this snapshot. Should
	// be the same as the logical size of the entity that was backed up.
	EntitySize *int64 `json:"entitySize,omitempty"`

	// The start time of the backup job run containing this snapshot.
	RunStartTimeUsecs *int64 `json:"runStartTimeUsecs,omitempty"`
}

// Validate validates this check restore app task result full snapshot info
func (m *CheckRestoreAppTaskResultFullSnapshotInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this check restore app task result full snapshot info based on context it is used
func (m *CheckRestoreAppTaskResultFullSnapshotInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CheckRestoreAppTaskResultFullSnapshotInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckRestoreAppTaskResultFullSnapshotInfo) UnmarshalBinary(b []byte) error {
	var res CheckRestoreAppTaskResultFullSnapshotInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
