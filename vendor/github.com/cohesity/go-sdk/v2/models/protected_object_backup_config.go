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

// ProtectedObjectBackupConfig Specifies the backup configuration for protected object.
//
// swagger:model ProtectedObjectBackupConfig
type ProtectedObjectBackupConfig struct {
	CommonObjectProtectParams

	EnvSpecificObjectProtectionResponseParams

	// Specifies whether or not this configuration is applied to an autoprotected object rather than this specific object.
	IsAutoProtectConfig *bool `json:"isAutoProtectConfig,omitempty"`

	// Specifies the parent ID of the object which the backup configuration is applied to if this is an auto protect config.
	AutoProtectParentID *int64 `json:"autoProtectParentId,omitempty"`

	// Specifies whether or not protection has been paused on this object.
	IsPaused *bool `json:"isPaused,omitempty"`

	// Specifies whether or not protection has been deactivated on this object.
	IsActive *bool `json:"isActive,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProtectedObjectBackupConfig) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonObjectProtectParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonObjectProtectParams = aO0

	// AO1
	var aO1 EnvSpecificObjectProtectionResponseParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.EnvSpecificObjectProtectionResponseParams = aO1

	// AO2
	var dataAO2 struct {
		IsAutoProtectConfig *bool `json:"isAutoProtectConfig,omitempty"`

		AutoProtectParentID *int64 `json:"autoProtectParentId,omitempty"`

		IsPaused *bool `json:"isPaused,omitempty"`

		IsActive *bool `json:"isActive,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.IsAutoProtectConfig = dataAO2.IsAutoProtectConfig

	m.AutoProtectParentID = dataAO2.AutoProtectParentID

	m.IsPaused = dataAO2.IsPaused

	m.IsActive = dataAO2.IsActive

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProtectedObjectBackupConfig) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.CommonObjectProtectParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.EnvSpecificObjectProtectionResponseParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	var dataAO2 struct {
		IsAutoProtectConfig *bool `json:"isAutoProtectConfig,omitempty"`

		AutoProtectParentID *int64 `json:"autoProtectParentId,omitempty"`

		IsPaused *bool `json:"isPaused,omitempty"`

		IsActive *bool `json:"isActive,omitempty"`
	}

	dataAO2.IsAutoProtectConfig = m.IsAutoProtectConfig

	dataAO2.AutoProtectParentID = m.AutoProtectParentID

	dataAO2.IsPaused = m.IsPaused

	dataAO2.IsActive = m.IsActive

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this protected object backup config
func (m *ProtectedObjectBackupConfig) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonObjectProtectParams
	if err := m.CommonObjectProtectParams.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with EnvSpecificObjectProtectionResponseParams
	if err := m.EnvSpecificObjectProtectionResponseParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this protected object backup config based on the context it is used
func (m *ProtectedObjectBackupConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonObjectProtectParams
	if err := m.CommonObjectProtectParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with EnvSpecificObjectProtectionResponseParams
	if err := m.EnvSpecificObjectProtectionResponseParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectedObjectBackupConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectedObjectBackupConfig) UnmarshalBinary(b []byte) error {
	var res ProtectedObjectBackupConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
