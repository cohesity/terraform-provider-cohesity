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

// RecoverExchangeDbsSnapshotParams Recover Exchange Databases Snapshot Params.
//
// Specifies the snapshot parameters to recover Exchange databases.
//
// swagger:model RecoverExchangeDbsSnapshotParams
type RecoverExchangeDbsSnapshotParams struct {
	CommonRecoverObjectSnapshotParams

	// Specifies the recovery target configuration if the recovery is set to a different location which is different from the original source.
	RecoveryTargetConfig *ExchangeRecoveryTargetConfig `json:"recoveryTargetConfig,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoverExchangeDbsSnapshotParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonRecoverObjectSnapshotParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonRecoverObjectSnapshotParams = aO0

	// AO1
	var dataAO1 struct {
		RecoveryTargetConfig *ExchangeRecoveryTargetConfig `json:"recoveryTargetConfig,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.RecoveryTargetConfig = dataAO1.RecoveryTargetConfig

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoverExchangeDbsSnapshotParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonRecoverObjectSnapshotParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		RecoveryTargetConfig *ExchangeRecoveryTargetConfig `json:"recoveryTargetConfig,omitempty"`
	}

	dataAO1.RecoveryTargetConfig = m.RecoveryTargetConfig

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recover exchange dbs snapshot params
func (m *RecoverExchangeDbsSnapshotParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonRecoverObjectSnapshotParams
	if err := m.CommonRecoverObjectSnapshotParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryTargetConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverExchangeDbsSnapshotParams) validateRecoveryTargetConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.RecoveryTargetConfig) { // not required
		return nil
	}

	if m.RecoveryTargetConfig != nil {
		if err := m.RecoveryTargetConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoveryTargetConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoveryTargetConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover exchange dbs snapshot params based on the context it is used
func (m *RecoverExchangeDbsSnapshotParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonRecoverObjectSnapshotParams
	if err := m.CommonRecoverObjectSnapshotParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoveryTargetConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverExchangeDbsSnapshotParams) contextValidateRecoveryTargetConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoveryTargetConfig != nil {

		if swag.IsZero(m.RecoveryTargetConfig) { // not required
			return nil
		}

		if err := m.RecoveryTargetConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recoveryTargetConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recoveryTargetConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverExchangeDbsSnapshotParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverExchangeDbsSnapshotParams) UnmarshalBinary(b []byte) error {
	var res RecoverExchangeDbsSnapshotParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
