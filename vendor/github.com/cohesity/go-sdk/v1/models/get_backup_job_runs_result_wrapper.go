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

// GetBackupJobRunsResultWrapper get backup job runs result wrapper
//
// swagger:model GetBackupJobRunsResultWrapper
type GetBackupJobRunsResultWrapper struct {

	// BackupJobRuns is the struct for BackupJobRunsProto used by magneto.
	BackupJobRuns *BackupJobRunsProto `json:"backupJobRuns,omitempty"`

	// Extension field
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	// Tenant information of tenants having access to the backup job.
	Tenants []*TenantInfo `json:"tenants"`
}

// Validate validates this get backup job runs result wrapper
func (m *GetBackupJobRunsResultWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupJobRuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenants(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBackupJobRunsResultWrapper) validateBackupJobRuns(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupJobRuns) { // not required
		return nil
	}

	if m.BackupJobRuns != nil {
		if err := m.BackupJobRuns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupJobRuns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupJobRuns")
			}
			return err
		}
	}

	return nil
}

func (m *GetBackupJobRunsResultWrapper) validateTenants(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenants) { // not required
		return nil
	}

	for i := 0; i < len(m.Tenants); i++ {
		if swag.IsZero(m.Tenants[i]) { // not required
			continue
		}

		if m.Tenants[i] != nil {
			if err := m.Tenants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get backup job runs result wrapper based on the context it is used
func (m *GetBackupJobRunsResultWrapper) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupJobRuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenants(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBackupJobRunsResultWrapper) contextValidateBackupJobRuns(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupJobRuns != nil {

		if swag.IsZero(m.BackupJobRuns) { // not required
			return nil
		}

		if err := m.BackupJobRuns.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backupJobRuns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backupJobRuns")
			}
			return err
		}
	}

	return nil
}

func (m *GetBackupJobRunsResultWrapper) contextValidateTenants(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tenants); i++ {

		if m.Tenants[i] != nil {

			if swag.IsZero(m.Tenants[i]) { // not required
				return nil
			}

			if err := m.Tenants[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tenants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tenants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetBackupJobRunsResultWrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBackupJobRunsResultWrapper) UnmarshalBinary(b []byte) error {
	var res GetBackupJobRunsResultWrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
