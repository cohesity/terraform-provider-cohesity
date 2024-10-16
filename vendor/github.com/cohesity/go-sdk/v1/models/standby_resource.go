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

// StandbyResource Message to encapsulate resources to be used to create a standby entity on
// the primary environment. Each environment should define the parameters it
// needs to create an entity on the primary environment.
//
// swagger:model StandbyResource
type StandbyResource struct {

	// User defined recovery point objective for the standby VM. Using this RPO,
	// Magneto will hydrate the VMs.
	RecoveryPointObjectiveSecs *int64 `json:"recoveryPointObjectiveSecs,omitempty"`

	// Standby resources needed in a VMware environment.
	VmwareStandbyResource *VMwareStandbyResource `json:"vmwareStandbyResource,omitempty"`
}

// Validate validates this standby resource
func (m *StandbyResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVmwareStandbyResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandbyResource) validateVmwareStandbyResource(formats strfmt.Registry) error {
	if swag.IsZero(m.VmwareStandbyResource) { // not required
		return nil
	}

	if m.VmwareStandbyResource != nil {
		if err := m.VmwareStandbyResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareStandbyResource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareStandbyResource")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this standby resource based on the context it is used
func (m *StandbyResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVmwareStandbyResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandbyResource) contextValidateVmwareStandbyResource(ctx context.Context, formats strfmt.Registry) error {

	if m.VmwareStandbyResource != nil {

		if swag.IsZero(m.VmwareStandbyResource) { // not required
			return nil
		}

		if err := m.VmwareStandbyResource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareStandbyResource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareStandbyResource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StandbyResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandbyResource) UnmarshalBinary(b []byte) error {
	var res StandbyResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
