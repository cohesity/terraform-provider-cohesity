// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtectionRunShellInfoProto A message that contains shell information about a backup run.
//
// NOTE: This information is derived from BackupJobRunMetadataProto which is
//
// kept in memory by Magneto. Magneto does not need to unpack the entire
// run metadata to furnish this information, making the RPC that serves
// this info extremely efficient.
//
// swagger:model ProtectionRunShellInfoProto
type ProtectionRunShellInfoProto struct {

	// The type of the scheduled backup run.
	BackupType *int32 `json:"backupType,omitempty"`

	// The end time of this run.
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`

	// The error encountered by this run (if any).
	Error *PrivateErrorProto `json:"error,omitempty"`

	// The instance id of this run.
	JobInstanceID *int64 `json:"jobInstanceId,omitempty"`

	// State of the this run.
	PublicStatus *int32 `json:"publicStatus,omitempty"`

	// The start time of this run.
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`
}

// Validate validates this protection run shell info proto
func (m *ProtectionRunShellInfoProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionRunShellInfoProto) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this protection run shell info proto based on the context it is used
func (m *ProtectionRunShellInfoProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionRunShellInfoProto) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionRunShellInfoProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionRunShellInfoProto) UnmarshalBinary(b []byte) error {
	var res ProtectionRunShellInfoProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
