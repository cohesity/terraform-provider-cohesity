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

// ProtectionSummary Specifies a summary of an object protection.
//
// swagger:model ProtectionSummary
type ProtectionSummary struct {

	// Specifies the policy name for this group.
	PolicyName *string `json:"policyName,omitempty"`

	// Specifies the policy id for this protection.
	PolicyID *string `json:"policyId,omitempty"`

	// Specifies the storage domain id of this protection. Format is clusterId:clusterIncarnationId:storageDomainId.
	StorageDomainID *string `json:"storageDomainId,omitempty"`

	// Specifies the status of last local back up run.
	// Enum: ["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold","Paused"]
	LastBackupRunStatus *string `json:"lastBackupRunStatus,omitempty"`

	// Specifies the status of last archival run.
	// Enum: ["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold","Paused"]
	LastArchivalRunStatus *string `json:"lastArchivalRunStatus,omitempty"`

	// Specifies the status of last replication run.
	// Enum: ["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold","Paused"]
	LastReplicationRunStatus *string `json:"lastReplicationRunStatus,omitempty"`

	// Specifies if the sla is violated in last run.
	LastRunSLAViolated *bool `json:"lastRunSlaViolated,omitempty"`
}

// Validate validates this protection summary
func (m *ProtectionSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastBackupRunStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastArchivalRunStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastReplicationRunStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var protectionSummaryTypeLastBackupRunStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold","Paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectionSummaryTypeLastBackupRunStatusPropEnum = append(protectionSummaryTypeLastBackupRunStatusPropEnum, v)
	}
}

const (

	// ProtectionSummaryLastBackupRunStatusAccepted captures enum value "Accepted"
	ProtectionSummaryLastBackupRunStatusAccepted string = "Accepted"

	// ProtectionSummaryLastBackupRunStatusRunning captures enum value "Running"
	ProtectionSummaryLastBackupRunStatusRunning string = "Running"

	// ProtectionSummaryLastBackupRunStatusCanceled captures enum value "Canceled"
	ProtectionSummaryLastBackupRunStatusCanceled string = "Canceled"

	// ProtectionSummaryLastBackupRunStatusCanceling captures enum value "Canceling"
	ProtectionSummaryLastBackupRunStatusCanceling string = "Canceling"

	// ProtectionSummaryLastBackupRunStatusFailed captures enum value "Failed"
	ProtectionSummaryLastBackupRunStatusFailed string = "Failed"

	// ProtectionSummaryLastBackupRunStatusMissed captures enum value "Missed"
	ProtectionSummaryLastBackupRunStatusMissed string = "Missed"

	// ProtectionSummaryLastBackupRunStatusSucceeded captures enum value "Succeeded"
	ProtectionSummaryLastBackupRunStatusSucceeded string = "Succeeded"

	// ProtectionSummaryLastBackupRunStatusSucceededWithWarning captures enum value "SucceededWithWarning"
	ProtectionSummaryLastBackupRunStatusSucceededWithWarning string = "SucceededWithWarning"

	// ProtectionSummaryLastBackupRunStatusOnHold captures enum value "OnHold"
	ProtectionSummaryLastBackupRunStatusOnHold string = "OnHold"

	// ProtectionSummaryLastBackupRunStatusFinalizing captures enum value "Finalizing"
	ProtectionSummaryLastBackupRunStatusFinalizing string = "Finalizing"

	// ProtectionSummaryLastBackupRunStatusSkipped captures enum value "Skipped"
	ProtectionSummaryLastBackupRunStatusSkipped string = "Skipped"

	// ProtectionSummaryLastBackupRunStatusLegalHold captures enum value "LegalHold"
	ProtectionSummaryLastBackupRunStatusLegalHold string = "LegalHold"

	// ProtectionSummaryLastBackupRunStatusPaused captures enum value "Paused"
	ProtectionSummaryLastBackupRunStatusPaused string = "Paused"
)

// prop value enum
func (m *ProtectionSummary) validateLastBackupRunStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectionSummaryTypeLastBackupRunStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectionSummary) validateLastBackupRunStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.LastBackupRunStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateLastBackupRunStatusEnum("lastBackupRunStatus", "body", *m.LastBackupRunStatus); err != nil {
		return err
	}

	return nil
}

var protectionSummaryTypeLastArchivalRunStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold","Paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectionSummaryTypeLastArchivalRunStatusPropEnum = append(protectionSummaryTypeLastArchivalRunStatusPropEnum, v)
	}
}

const (

	// ProtectionSummaryLastArchivalRunStatusAccepted captures enum value "Accepted"
	ProtectionSummaryLastArchivalRunStatusAccepted string = "Accepted"

	// ProtectionSummaryLastArchivalRunStatusRunning captures enum value "Running"
	ProtectionSummaryLastArchivalRunStatusRunning string = "Running"

	// ProtectionSummaryLastArchivalRunStatusCanceled captures enum value "Canceled"
	ProtectionSummaryLastArchivalRunStatusCanceled string = "Canceled"

	// ProtectionSummaryLastArchivalRunStatusCanceling captures enum value "Canceling"
	ProtectionSummaryLastArchivalRunStatusCanceling string = "Canceling"

	// ProtectionSummaryLastArchivalRunStatusFailed captures enum value "Failed"
	ProtectionSummaryLastArchivalRunStatusFailed string = "Failed"

	// ProtectionSummaryLastArchivalRunStatusMissed captures enum value "Missed"
	ProtectionSummaryLastArchivalRunStatusMissed string = "Missed"

	// ProtectionSummaryLastArchivalRunStatusSucceeded captures enum value "Succeeded"
	ProtectionSummaryLastArchivalRunStatusSucceeded string = "Succeeded"

	// ProtectionSummaryLastArchivalRunStatusSucceededWithWarning captures enum value "SucceededWithWarning"
	ProtectionSummaryLastArchivalRunStatusSucceededWithWarning string = "SucceededWithWarning"

	// ProtectionSummaryLastArchivalRunStatusOnHold captures enum value "OnHold"
	ProtectionSummaryLastArchivalRunStatusOnHold string = "OnHold"

	// ProtectionSummaryLastArchivalRunStatusFinalizing captures enum value "Finalizing"
	ProtectionSummaryLastArchivalRunStatusFinalizing string = "Finalizing"

	// ProtectionSummaryLastArchivalRunStatusSkipped captures enum value "Skipped"
	ProtectionSummaryLastArchivalRunStatusSkipped string = "Skipped"

	// ProtectionSummaryLastArchivalRunStatusLegalHold captures enum value "LegalHold"
	ProtectionSummaryLastArchivalRunStatusLegalHold string = "LegalHold"

	// ProtectionSummaryLastArchivalRunStatusPaused captures enum value "Paused"
	ProtectionSummaryLastArchivalRunStatusPaused string = "Paused"
)

// prop value enum
func (m *ProtectionSummary) validateLastArchivalRunStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectionSummaryTypeLastArchivalRunStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectionSummary) validateLastArchivalRunStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.LastArchivalRunStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateLastArchivalRunStatusEnum("lastArchivalRunStatus", "body", *m.LastArchivalRunStatus); err != nil {
		return err
	}

	return nil
}

var protectionSummaryTypeLastReplicationRunStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold","Paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectionSummaryTypeLastReplicationRunStatusPropEnum = append(protectionSummaryTypeLastReplicationRunStatusPropEnum, v)
	}
}

const (

	// ProtectionSummaryLastReplicationRunStatusAccepted captures enum value "Accepted"
	ProtectionSummaryLastReplicationRunStatusAccepted string = "Accepted"

	// ProtectionSummaryLastReplicationRunStatusRunning captures enum value "Running"
	ProtectionSummaryLastReplicationRunStatusRunning string = "Running"

	// ProtectionSummaryLastReplicationRunStatusCanceled captures enum value "Canceled"
	ProtectionSummaryLastReplicationRunStatusCanceled string = "Canceled"

	// ProtectionSummaryLastReplicationRunStatusCanceling captures enum value "Canceling"
	ProtectionSummaryLastReplicationRunStatusCanceling string = "Canceling"

	// ProtectionSummaryLastReplicationRunStatusFailed captures enum value "Failed"
	ProtectionSummaryLastReplicationRunStatusFailed string = "Failed"

	// ProtectionSummaryLastReplicationRunStatusMissed captures enum value "Missed"
	ProtectionSummaryLastReplicationRunStatusMissed string = "Missed"

	// ProtectionSummaryLastReplicationRunStatusSucceeded captures enum value "Succeeded"
	ProtectionSummaryLastReplicationRunStatusSucceeded string = "Succeeded"

	// ProtectionSummaryLastReplicationRunStatusSucceededWithWarning captures enum value "SucceededWithWarning"
	ProtectionSummaryLastReplicationRunStatusSucceededWithWarning string = "SucceededWithWarning"

	// ProtectionSummaryLastReplicationRunStatusOnHold captures enum value "OnHold"
	ProtectionSummaryLastReplicationRunStatusOnHold string = "OnHold"

	// ProtectionSummaryLastReplicationRunStatusFinalizing captures enum value "Finalizing"
	ProtectionSummaryLastReplicationRunStatusFinalizing string = "Finalizing"

	// ProtectionSummaryLastReplicationRunStatusSkipped captures enum value "Skipped"
	ProtectionSummaryLastReplicationRunStatusSkipped string = "Skipped"

	// ProtectionSummaryLastReplicationRunStatusLegalHold captures enum value "LegalHold"
	ProtectionSummaryLastReplicationRunStatusLegalHold string = "LegalHold"

	// ProtectionSummaryLastReplicationRunStatusPaused captures enum value "Paused"
	ProtectionSummaryLastReplicationRunStatusPaused string = "Paused"
)

// prop value enum
func (m *ProtectionSummary) validateLastReplicationRunStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectionSummaryTypeLastReplicationRunStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectionSummary) validateLastReplicationRunStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.LastReplicationRunStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateLastReplicationRunStatusEnum("lastReplicationRunStatus", "body", *m.LastReplicationRunStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this protection summary based on context it is used
func (m *ProtectionSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionSummary) UnmarshalBinary(b []byte) error {
	var res ProtectionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
