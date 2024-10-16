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

// ClusterSnapshotPolicy Snapshot Policy struct.
//
// Describes the snapshot policy struct.
//
// swagger:model ClusterSnapshotPolicy
type ClusterSnapshotPolicy struct {

	// Snapshot schedule policy.
	SnapshotSchedulePolicy *SnapshotSchedulePolicy `json:"snapshotSchedulePolicy,omitempty"`

	// Snapshot retention policy.
	SnapshotRetentionPolicy *SnapshotRetentionPolicy `json:"snapshotRetentionPolicy,omitempty"`
}

// Validate validates this cluster snapshot policy
func (m *ClusterSnapshotPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshotSchedulePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotRetentionPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSnapshotPolicy) validateSnapshotSchedulePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotSchedulePolicy) { // not required
		return nil
	}

	if m.SnapshotSchedulePolicy != nil {
		if err := m.SnapshotSchedulePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotSchedulePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotSchedulePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSnapshotPolicy) validateSnapshotRetentionPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotRetentionPolicy) { // not required
		return nil
	}

	if m.SnapshotRetentionPolicy != nil {
		if err := m.SnapshotRetentionPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotRetentionPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotRetentionPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster snapshot policy based on the context it is used
func (m *ClusterSnapshotPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSnapshotSchedulePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotRetentionPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSnapshotPolicy) contextValidateSnapshotSchedulePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotSchedulePolicy != nil {

		if swag.IsZero(m.SnapshotSchedulePolicy) { // not required
			return nil
		}

		if err := m.SnapshotSchedulePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotSchedulePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotSchedulePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSnapshotPolicy) contextValidateSnapshotRetentionPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotRetentionPolicy != nil {

		if swag.IsZero(m.SnapshotRetentionPolicy) { // not required
			return nil
		}

		if err := m.SnapshotRetentionPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotRetentionPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotRetentionPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSnapshotPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSnapshotPolicy) UnmarshalBinary(b []byte) error {
	var res ClusterSnapshotPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
