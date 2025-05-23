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

// CommonOracleAppSourceConfig Common Oracle App Source Config.
//
// Specifies a common parameters used when restoring back to original or new source.
//
// swagger:model CommonOracleAppSourceConfig
type CommonOracleAppSourceConfig struct {

	// Specifies the time in the past to which the Oracle db needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle db will be restored from the full/incremental snapshot.
	RestoreTimeUsecs *int64 `json:"restoreTimeUsecs,omitempty"`

	// Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases.
	DbChannels []*OracleDbChannel `json:"dbChannels"`

	// Specifies if database should be left in recovery mode.
	RecoveryMode *bool `json:"recoveryMode,omitempty"`

	// Specifies key value pairs of shell variables which defines the restore shell environment.
	ShellEvironmentVars []*ShellKeyValuePair `json:"shellEvironmentVars"`

	// Specifies whether database recovery performed should use scn value or not.
	UseScnForRestore *bool `json:"useScnForRestore,omitempty"`

	// Specifies information about list of objects (PDBs) to restore.
	GranularRestoreInfo *RecoverOracleGranularRestoreInfo `json:"granularRestoreInfo,omitempty"`

	// Specifies Range in Time, Scn or Sequence to restore archive logs of a DB.
	OracleArchiveLogInfo *OracleArchiveLogInfo `json:"oracleArchiveLogInfo,omitempty"`

	// Specifies parameters related to Oracle Recovery Validation.
	OracleRecoveryValidationInfo *OracleRecoveryValidationInfo `json:"oracleRecoveryValidationInfo,omitempty"`

	// Specifies parameters related to spfile/pfile restore.
	RestoreSpfileOrPfileInfo *RestoreSpfileOrPfileInfo `json:"restoreSpfileOrPfileInfo,omitempty"`
}

// Validate validates this common oracle app source config
func (m *CommonOracleAppSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDbChannels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShellEvironmentVars(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGranularRestoreInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleArchiveLogInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleRecoveryValidationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreSpfileOrPfileInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonOracleAppSourceConfig) validateDbChannels(formats strfmt.Registry) error {
	if swag.IsZero(m.DbChannels) { // not required
		return nil
	}

	for i := 0; i < len(m.DbChannels); i++ {
		if swag.IsZero(m.DbChannels[i]) { // not required
			continue
		}

		if m.DbChannels[i] != nil {
			if err := m.DbChannels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dbChannels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dbChannels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonOracleAppSourceConfig) validateShellEvironmentVars(formats strfmt.Registry) error {
	if swag.IsZero(m.ShellEvironmentVars) { // not required
		return nil
	}

	for i := 0; i < len(m.ShellEvironmentVars); i++ {
		if swag.IsZero(m.ShellEvironmentVars[i]) { // not required
			continue
		}

		if m.ShellEvironmentVars[i] != nil {
			if err := m.ShellEvironmentVars[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shellEvironmentVars" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shellEvironmentVars" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonOracleAppSourceConfig) validateGranularRestoreInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.GranularRestoreInfo) { // not required
		return nil
	}

	if m.GranularRestoreInfo != nil {
		if err := m.GranularRestoreInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("granularRestoreInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("granularRestoreInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CommonOracleAppSourceConfig) validateOracleArchiveLogInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.OracleArchiveLogInfo) { // not required
		return nil
	}

	if m.OracleArchiveLogInfo != nil {
		if err := m.OracleArchiveLogInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracleArchiveLogInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oracleArchiveLogInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CommonOracleAppSourceConfig) validateOracleRecoveryValidationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.OracleRecoveryValidationInfo) { // not required
		return nil
	}

	if m.OracleRecoveryValidationInfo != nil {
		if err := m.OracleRecoveryValidationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracleRecoveryValidationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oracleRecoveryValidationInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CommonOracleAppSourceConfig) validateRestoreSpfileOrPfileInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoreSpfileOrPfileInfo) { // not required
		return nil
	}

	if m.RestoreSpfileOrPfileInfo != nil {
		if err := m.RestoreSpfileOrPfileInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreSpfileOrPfileInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreSpfileOrPfileInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this common oracle app source config based on the context it is used
func (m *CommonOracleAppSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDbChannels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShellEvironmentVars(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGranularRestoreInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOracleArchiveLogInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOracleRecoveryValidationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoreSpfileOrPfileInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonOracleAppSourceConfig) contextValidateDbChannels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DbChannels); i++ {

		if m.DbChannels[i] != nil {

			if swag.IsZero(m.DbChannels[i]) { // not required
				return nil
			}

			if err := m.DbChannels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dbChannels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dbChannels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonOracleAppSourceConfig) contextValidateShellEvironmentVars(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ShellEvironmentVars); i++ {

		if m.ShellEvironmentVars[i] != nil {

			if swag.IsZero(m.ShellEvironmentVars[i]) { // not required
				return nil
			}

			if err := m.ShellEvironmentVars[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shellEvironmentVars" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shellEvironmentVars" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonOracleAppSourceConfig) contextValidateGranularRestoreInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.GranularRestoreInfo != nil {

		if swag.IsZero(m.GranularRestoreInfo) { // not required
			return nil
		}

		if err := m.GranularRestoreInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("granularRestoreInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("granularRestoreInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CommonOracleAppSourceConfig) contextValidateOracleArchiveLogInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.OracleArchiveLogInfo != nil {

		if swag.IsZero(m.OracleArchiveLogInfo) { // not required
			return nil
		}

		if err := m.OracleArchiveLogInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracleArchiveLogInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oracleArchiveLogInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CommonOracleAppSourceConfig) contextValidateOracleRecoveryValidationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.OracleRecoveryValidationInfo != nil {

		if swag.IsZero(m.OracleRecoveryValidationInfo) { // not required
			return nil
		}

		if err := m.OracleRecoveryValidationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracleRecoveryValidationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oracleRecoveryValidationInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CommonOracleAppSourceConfig) contextValidateRestoreSpfileOrPfileInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.RestoreSpfileOrPfileInfo != nil {

		if swag.IsZero(m.RestoreSpfileOrPfileInfo) { // not required
			return nil
		}

		if err := m.RestoreSpfileOrPfileInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreSpfileOrPfileInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreSpfileOrPfileInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonOracleAppSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonOracleAppSourceConfig) UnmarshalBinary(b []byte) error {
	var res CommonOracleAppSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
