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

// RecoveryOracleTaskInfo Recovery Task Info
//
// Specifies the info about a recovery task.
//
// swagger:model RecoveryOracleTaskInfo
type RecoveryOracleTaskInfo struct {

	// Specifies the progress monitor id.
	// Read Only: true
	ProgressTaskID *string `json:"progressTaskId,omitempty"`

	// Specifies the status of the recovery.
	// Read Only: true
	// Enum: ["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold"]
	Status *string `json:"status,omitempty"`

	// Specifies the start time in Unix timestamp epoch in microseconds.
	// Read Only: true
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`

	// Specifies the end time in Unix timestamp epoch in microseconds.
	// Read Only: true
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`
}

// Validate validates this recovery oracle task info
func (m *RecoveryOracleTaskInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var recoveryOracleTaskInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoveryOracleTaskInfoTypeStatusPropEnum = append(recoveryOracleTaskInfoTypeStatusPropEnum, v)
	}
}

const (

	// RecoveryOracleTaskInfoStatusAccepted captures enum value "Accepted"
	RecoveryOracleTaskInfoStatusAccepted string = "Accepted"

	// RecoveryOracleTaskInfoStatusRunning captures enum value "Running"
	RecoveryOracleTaskInfoStatusRunning string = "Running"

	// RecoveryOracleTaskInfoStatusCanceled captures enum value "Canceled"
	RecoveryOracleTaskInfoStatusCanceled string = "Canceled"

	// RecoveryOracleTaskInfoStatusCanceling captures enum value "Canceling"
	RecoveryOracleTaskInfoStatusCanceling string = "Canceling"

	// RecoveryOracleTaskInfoStatusFailed captures enum value "Failed"
	RecoveryOracleTaskInfoStatusFailed string = "Failed"

	// RecoveryOracleTaskInfoStatusMissed captures enum value "Missed"
	RecoveryOracleTaskInfoStatusMissed string = "Missed"

	// RecoveryOracleTaskInfoStatusSucceeded captures enum value "Succeeded"
	RecoveryOracleTaskInfoStatusSucceeded string = "Succeeded"

	// RecoveryOracleTaskInfoStatusSucceededWithWarning captures enum value "SucceededWithWarning"
	RecoveryOracleTaskInfoStatusSucceededWithWarning string = "SucceededWithWarning"

	// RecoveryOracleTaskInfoStatusOnHold captures enum value "OnHold"
	RecoveryOracleTaskInfoStatusOnHold string = "OnHold"

	// RecoveryOracleTaskInfoStatusFinalizing captures enum value "Finalizing"
	RecoveryOracleTaskInfoStatusFinalizing string = "Finalizing"

	// RecoveryOracleTaskInfoStatusSkipped captures enum value "Skipped"
	RecoveryOracleTaskInfoStatusSkipped string = "Skipped"

	// RecoveryOracleTaskInfoStatusLegalHold captures enum value "LegalHold"
	RecoveryOracleTaskInfoStatusLegalHold string = "LegalHold"
)

// prop value enum
func (m *RecoveryOracleTaskInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoveryOracleTaskInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoveryOracleTaskInfo) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this recovery oracle task info based on the context it is used
func (m *RecoveryOracleTaskInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProgressTaskID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTimeUsecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndTimeUsecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryOracleTaskInfo) contextValidateProgressTaskID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "progressTaskId", "body", m.ProgressTaskID); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryOracleTaskInfo) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryOracleTaskInfo) contextValidateStartTimeUsecs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "startTimeUsecs", "body", m.StartTimeUsecs); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryOracleTaskInfo) contextValidateEndTimeUsecs(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endTimeUsecs", "body", m.EndTimeUsecs); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryOracleTaskInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryOracleTaskInfo) UnmarshalBinary(b []byte) error {
	var res RecoveryOracleTaskInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
