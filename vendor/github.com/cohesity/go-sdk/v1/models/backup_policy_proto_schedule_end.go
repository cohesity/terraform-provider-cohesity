// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupPolicyProtoScheduleEnd Only one of the following fields should be set.
//
// swagger:model BackupPolicyProto_ScheduleEnd
type BackupPolicyProtoScheduleEnd struct {

	// If specified, the backup job will no longer be run after it has been run
	// these many times.
	EndAfterNumBackups *int64 `json:"endAfterNumBackups,omitempty"`

	// If specified, the backup job will no longer be run after this time.
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`
}

// Validate validates this backup policy proto schedule end
func (m *BackupPolicyProtoScheduleEnd) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this backup policy proto schedule end based on context it is used
func (m *BackupPolicyProtoScheduleEnd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupPolicyProtoScheduleEnd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupPolicyProtoScheduleEnd) UnmarshalBinary(b []byte) error {
	var res BackupPolicyProtoScheduleEnd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
