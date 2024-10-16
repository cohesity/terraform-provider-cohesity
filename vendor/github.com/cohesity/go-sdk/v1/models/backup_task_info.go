// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupTaskInfo Parameters for a backup op.
//
// swagger:model BackupTaskInfo
type BackupTaskInfo struct {

	// Id of that particular instance
	InstanceID *string `json:"instanceId,omitempty"`

	// Name of the recovery task.
	Name *string `json:"name,omitempty"`

	// Denotes the start time of the backuptask, needed for deeplinking.
	StartTimeUsecs *string `json:"startTimeUsecs,omitempty"`

	// Id of the recovery task.
	TaskID *string `json:"taskId,omitempty"`
}

// Validate validates this backup task info
func (m *BackupTaskInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this backup task info based on context it is used
func (m *BackupTaskInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupTaskInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupTaskInfo) UnmarshalBinary(b []byte) error {
	var res BackupTaskInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
