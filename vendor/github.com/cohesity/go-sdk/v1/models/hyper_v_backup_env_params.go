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

// HyperVBackupEnvParams HyperV Backup Environment Parameters.
//
// Message to capture any additional backup params for a HyperV environment.
//
// swagger:model HyperVBackupEnvParams
type HyperVBackupEnvParams struct {

	// Whether to fallback to take a crash-consistent snapshot incase taking
	// an app-consistent snapshot fails.
	AllowCrashConsistentSnapshot *bool `json:"allowCrashConsistentSnapshot,omitempty"`

	// The type of backup job to use. Default is to auto-detect the best type to
	// use based on the VMs to backup. End user may select RCT or VSS also.
	BackupJobType *int32 `json:"backupJobType,omitempty"`

	// List of Virtual Disk(s) to be excluded from the backup job. These disks
	// will be excluded for all VMs in this environment unless overriden by the
	// disk exclusion list from BackupSourceParams.HyperVBackupSourceParams.
	HypervDiskExclusionInfo []*HyperVDiskFilterProto `json:"hypervDiskExclusionInfo"`

	// List of Virtual Disk(s) to be included in the backup job for the source.
	// These disks will be included for all VMs in this environment and all other
	// disks will be excluded.
	// It can be overriden by the disk exclusion/inclusion list from
	// BackupSourceParams.HyperVBackupSourceParams.
	HypervDiskInclusionInfo []*HyperVDiskFilterProto `json:"hypervDiskInclusionInfo"`
}

// Validate validates this hyper v backup env params
func (m *HyperVBackupEnvParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHypervDiskExclusionInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHypervDiskInclusionInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperVBackupEnvParams) validateHypervDiskExclusionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.HypervDiskExclusionInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.HypervDiskExclusionInfo); i++ {
		if swag.IsZero(m.HypervDiskExclusionInfo[i]) { // not required
			continue
		}

		if m.HypervDiskExclusionInfo[i] != nil {
			if err := m.HypervDiskExclusionInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hypervDiskExclusionInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hypervDiskExclusionInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperVBackupEnvParams) validateHypervDiskInclusionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.HypervDiskInclusionInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.HypervDiskInclusionInfo); i++ {
		if swag.IsZero(m.HypervDiskInclusionInfo[i]) { // not required
			continue
		}

		if m.HypervDiskInclusionInfo[i] != nil {
			if err := m.HypervDiskInclusionInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hypervDiskInclusionInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hypervDiskInclusionInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hyper v backup env params based on the context it is used
func (m *HyperVBackupEnvParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHypervDiskExclusionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHypervDiskInclusionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HyperVBackupEnvParams) contextValidateHypervDiskExclusionInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HypervDiskExclusionInfo); i++ {

		if m.HypervDiskExclusionInfo[i] != nil {

			if swag.IsZero(m.HypervDiskExclusionInfo[i]) { // not required
				return nil
			}

			if err := m.HypervDiskExclusionInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hypervDiskExclusionInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hypervDiskExclusionInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HyperVBackupEnvParams) contextValidateHypervDiskInclusionInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HypervDiskInclusionInfo); i++ {

		if m.HypervDiskInclusionInfo[i] != nil {

			if swag.IsZero(m.HypervDiskInclusionInfo[i]) { // not required
				return nil
			}

			if err := m.HypervDiskInclusionInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hypervDiskInclusionInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hypervDiskInclusionInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperVBackupEnvParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperVBackupEnvParams) UnmarshalBinary(b []byte) error {
	var res HyperVBackupEnvParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
