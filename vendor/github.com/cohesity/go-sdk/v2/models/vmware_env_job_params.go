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

// VmwareEnvJobParams Specifies job parameters applicable for all 'kVMware' Environment type Protection Sources in a Protection Job.
//
// swagger:model VmwareEnvJobParams
type VmwareEnvJobParams struct {

	// If true, takes a crash-consistent snapshot when app-consistent snapshot fails. Otherwise, the snapshot attempt is marked failed.
	FallbackToCrashConsistent *bool `json:"fallbackToCrashConsistent,omitempty"`

	// If true, skip physical RDM disks when backing up VMs. Otherwise, backup of VMs having physical RDM will fail.
	SkipPhysicalRdmDisks *bool `json:"skipPhysicalRdmDisks,omitempty"`

	// Specifies the list of Disks to be excluded from backing up. These disks are excluded from all Protection Sources in the Protection Job.
	ExcludedDisks []*DiskInfo `json:"excludedDisks"`
}

// Validate validates this vmware env job params
func (m *VmwareEnvJobParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExcludedDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareEnvJobParams) validateExcludedDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludedDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.ExcludedDisks); i++ {
		if swag.IsZero(m.ExcludedDisks[i]) { // not required
			continue
		}

		if m.ExcludedDisks[i] != nil {
			if err := m.ExcludedDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excludedDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excludedDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vmware env job params based on the context it is used
func (m *VmwareEnvJobParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExcludedDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareEnvJobParams) contextValidateExcludedDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExcludedDisks); i++ {

		if m.ExcludedDisks[i] != nil {

			if swag.IsZero(m.ExcludedDisks[i]) { // not required
				return nil
			}

			if err := m.ExcludedDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excludedDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excludedDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmwareEnvJobParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareEnvJobParams) UnmarshalBinary(b []byte) error {
	var res VmwareEnvJobParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
