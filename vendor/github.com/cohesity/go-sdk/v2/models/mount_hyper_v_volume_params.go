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

// MountHyperVVolumeParams Mount Volumes Params.
//
// Specifies the parameters to mount volumes.
//
// swagger:model MountHyperVVolumeParams
type MountHyperVVolumeParams struct {

	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	// Required: true
	// Enum: ["kHyperV"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// Specifies the params for recovering to a HyperV target.
	HypervTargetParams *HyperVTargetParamsForMountVolume `json:"hypervTargetParams,omitempty"`
}

// Validate validates this mount hyper v volume params
func (m *MountHyperVVolumeParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHypervTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mountHyperVVolumeParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kHyperV"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mountHyperVVolumeParamsTypeTargetEnvironmentPropEnum = append(mountHyperVVolumeParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// MountHyperVVolumeParamsTargetEnvironmentKHyperV captures enum value "kHyperV"
	MountHyperVVolumeParamsTargetEnvironmentKHyperV string = "kHyperV"
)

// prop value enum
func (m *MountHyperVVolumeParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mountHyperVVolumeParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MountHyperVVolumeParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *MountHyperVVolumeParams) validateHypervTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.HypervTargetParams) { // not required
		return nil
	}

	if m.HypervTargetParams != nil {
		if err := m.HypervTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hypervTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hypervTargetParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mount hyper v volume params based on the context it is used
func (m *MountHyperVVolumeParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHypervTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountHyperVVolumeParams) contextValidateHypervTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.HypervTargetParams != nil {

		if swag.IsZero(m.HypervTargetParams) { // not required
			return nil
		}

		if err := m.HypervTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hypervTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hypervTargetParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MountHyperVVolumeParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountHyperVVolumeParams) UnmarshalBinary(b []byte) error {
	var res MountHyperVVolumeParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
