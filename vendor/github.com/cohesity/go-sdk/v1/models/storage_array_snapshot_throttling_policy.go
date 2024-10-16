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

// StorageArraySnapshotThrottlingPolicy Specifies storage array snapshot throttling policy parameters.
//
// swagger:model StorageArraySnapshotThrottlingPolicy
type StorageArraySnapshotThrottlingPolicy struct {

	// Specifies the volume id of the storage array snapshot config.
	ID *int64 `json:"id,omitempty"`

	// Specifies if the storage array snapshot max snapshots config is enabled or
	// not.
	IsMaxSnapshotsConfigEnabled *bool `json:"isMaxSnapshotsConfigEnabled,omitempty"`

	// Specifies if the storage array snapshot max space config is enabled or
	// not.
	IsMaxSpaceConfigEnabled *bool `json:"isMaxSpaceConfigEnabled,omitempty"`

	// Specifies the storage array snapshot max snapshot config for this
	// volume/lun. Valid only when IsMaxSnapshotsConfigEnabled is true.
	MaxSnapshotConfig *StorageArraySnapshotMaxSnapshotConfigParams `json:"maxSnapshotConfig,omitempty"`

	// Specifies the storage array snapshot free space config for this
	// volume/lun. Valid only when IsMaxSpaceConfigEnabled is true.
	MaxSpaceConfig *StorageArraySnapshotMaxSpaceConfigParams `json:"maxSpaceConfig,omitempty"`
}

// Validate validates this storage array snapshot throttling policy
func (m *StorageArraySnapshotThrottlingPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxSnapshotConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxSpaceConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageArraySnapshotThrottlingPolicy) validateMaxSnapshotConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxSnapshotConfig) { // not required
		return nil
	}

	if m.MaxSnapshotConfig != nil {
		if err := m.MaxSnapshotConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maxSnapshotConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maxSnapshotConfig")
			}
			return err
		}
	}

	return nil
}

func (m *StorageArraySnapshotThrottlingPolicy) validateMaxSpaceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxSpaceConfig) { // not required
		return nil
	}

	if m.MaxSpaceConfig != nil {
		if err := m.MaxSpaceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maxSpaceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maxSpaceConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this storage array snapshot throttling policy based on the context it is used
func (m *StorageArraySnapshotThrottlingPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaxSnapshotConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxSpaceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageArraySnapshotThrottlingPolicy) contextValidateMaxSnapshotConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxSnapshotConfig != nil {

		if swag.IsZero(m.MaxSnapshotConfig) { // not required
			return nil
		}

		if err := m.MaxSnapshotConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maxSnapshotConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maxSnapshotConfig")
			}
			return err
		}
	}

	return nil
}

func (m *StorageArraySnapshotThrottlingPolicy) contextValidateMaxSpaceConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxSpaceConfig != nil {

		if swag.IsZero(m.MaxSpaceConfig) { // not required
			return nil
		}

		if err := m.MaxSpaceConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maxSpaceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("maxSpaceConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageArraySnapshotThrottlingPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageArraySnapshotThrottlingPolicy) UnmarshalBinary(b []byte) error {
	var res StorageArraySnapshotThrottlingPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
