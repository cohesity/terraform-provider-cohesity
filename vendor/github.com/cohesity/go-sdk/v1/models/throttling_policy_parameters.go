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

// ThrottlingPolicyParameters Throttling Policy.
//
// Specifies the throttling policy for a registered Protection Source.
//
// swagger:model ThrottlingPolicyParameters
type ThrottlingPolicyParameters struct {

	// Specifies whether datastore streams are configured for all datastores
	// that are part of the registered entity. If set to true, number of
	// streams from Cohesity cluster to the registered entity will be limited
	// to the value set for maxConcurrentStreams. If not set or set to false,
	// there is no max limit for the number of concurrent streams.
	EnforceMaxStreams *bool `json:"enforceMaxStreams,omitempty"`

	// Specifies whether no. of backups are configured for the registered entity. If
	// set to true, number of backups made by Cohesity cluster in the registered
	// entity will be limited to the value set for
	// RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no
	// max limit for the number of concurrent backups.
	EnforceRegisteredSourceMaxBackups *bool `json:"enforceRegisteredSourceMaxBackups,omitempty"`

	// Indicates whether read operations to the datastores, which are
	// part of the registered Protection Source, are throttled.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Specifies the thresholds that should be applied to all
	// datastores that are part of the registered Object.
	LatencyThresholds *LatencyThresholds `json:"latencyThresholds,omitempty"`

	// Specifies the limit on the number of streams Cohesity cluster will make
	// concurrently to the datastores of the registered entity. This limit is
	// enforced only when the flag enforceMaxStreams is set to true.
	MaxConcurrentStreams *int32 `json:"maxConcurrentStreams,omitempty"`

	// Specifies the NAS specific source throttling parameters during source
	// registration of the source.
	NasSourceParams *NasSourceThrottlingParams `json:"nasSourceParams,omitempty"`

	// Specifies the limit on the number of backups Cohesity cluster will make
	// concurrently to the registered entity. This limit is
	// enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.
	RegisteredSourceMaxConcurrentBackups *int32 `json:"registeredSourceMaxConcurrentBackups,omitempty"`

	// Specifies the storage array snapshot configuration parameters.
	// Valid only when IsStorageArraySnapshotEnabled is true.
	StorageArraySnapshotConfig *StorageArraySnapshotConfigParams `json:"storageArraySnapshotConfig,omitempty"`
}

// Validate validates this throttling policy parameters
func (m *ThrottlingPolicyParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLatencyThresholds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNasSourceParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArraySnapshotConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThrottlingPolicyParameters) validateLatencyThresholds(formats strfmt.Registry) error {
	if swag.IsZero(m.LatencyThresholds) { // not required
		return nil
	}

	if m.LatencyThresholds != nil {
		if err := m.LatencyThresholds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latencyThresholds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latencyThresholds")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicyParameters) validateNasSourceParams(formats strfmt.Registry) error {
	if swag.IsZero(m.NasSourceParams) { // not required
		return nil
	}

	if m.NasSourceParams != nil {
		if err := m.NasSourceParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nasSourceParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nasSourceParams")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicyParameters) validateStorageArraySnapshotConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageArraySnapshotConfig) { // not required
		return nil
	}

	if m.StorageArraySnapshotConfig != nil {
		if err := m.StorageArraySnapshotConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshotConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshotConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this throttling policy parameters based on the context it is used
func (m *ThrottlingPolicyParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLatencyThresholds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNasSourceParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageArraySnapshotConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThrottlingPolicyParameters) contextValidateLatencyThresholds(ctx context.Context, formats strfmt.Registry) error {

	if m.LatencyThresholds != nil {

		if swag.IsZero(m.LatencyThresholds) { // not required
			return nil
		}

		if err := m.LatencyThresholds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latencyThresholds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latencyThresholds")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicyParameters) contextValidateNasSourceParams(ctx context.Context, formats strfmt.Registry) error {

	if m.NasSourceParams != nil {

		if swag.IsZero(m.NasSourceParams) { // not required
			return nil
		}

		if err := m.NasSourceParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nasSourceParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nasSourceParams")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicyParameters) contextValidateStorageArraySnapshotConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageArraySnapshotConfig != nil {

		if swag.IsZero(m.StorageArraySnapshotConfig) { // not required
			return nil
		}

		if err := m.StorageArraySnapshotConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshotConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshotConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThrottlingPolicyParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThrottlingPolicyParameters) UnmarshalBinary(b []byte) error {
	var res ThrottlingPolicyParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
