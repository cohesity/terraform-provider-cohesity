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

// MountVolumesHyperVParams HyperV specific additional params.
//
// swagger:model MountVolumesHyperVParams
type MountVolumesHyperVParams struct {

	// Optional setting which will determine if the volumes need to be onlined
	// within the target environment after attaching the disks.
	// NOTE: If this is set to true, then target_entity_credentials must be
	// specified and HyperV Integration Services must be installed on the VM.
	BringDisksOnline *bool `json:"bringDisksOnline,omitempty"`

	// Credentials to access the target entity such as a VM.
	// This is not required if bring_disks_online is set to false.
	TargetEntityCredentials *Credentials `json:"targetEntityCredentials,omitempty"`
}

// Validate validates this mount volumes hyper v params
func (m *MountVolumesHyperVParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetEntityCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountVolumesHyperVParams) validateTargetEntityCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetEntityCredentials) { // not required
		return nil
	}

	if m.TargetEntityCredentials != nil {
		if err := m.TargetEntityCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetEntityCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetEntityCredentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mount volumes hyper v params based on the context it is used
func (m *MountVolumesHyperVParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTargetEntityCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountVolumesHyperVParams) contextValidateTargetEntityCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetEntityCredentials != nil {

		if swag.IsZero(m.TargetEntityCredentials) { // not required
			return nil
		}

		if err := m.TargetEntityCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetEntityCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetEntityCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MountVolumesHyperVParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountVolumesHyperVParams) UnmarshalBinary(b []byte) error {
	var res MountVolumesHyperVParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
