// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VmwareTargetParamsForRecoverFileAndFolder VMware Target Params for Recover File And Folder
//
// Specifies the parameters for a VMware recovery target.
//
// swagger:model VmwareTargetParamsForRecoverFileAndFolder
type VmwareTargetParamsForRecoverFileAndFolder struct {

	// Specifies whether to recover to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
	// Required: true
	RecoverToOriginalTarget *bool `json:"recoverToOriginalTarget"`

	// Specifies whether to overwrite the existing files. Default is true.
	OverwriteExisting *bool `json:"overwriteExisting,omitempty"`

	// Specifies whether to preserve original attributes. Default is true.
	PreserveAttributes *bool `json:"preserveAttributes,omitempty"`

	// Specifies whether to continue recovering other files if one of files or folders failed to recover. Default value is false.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// Specifies the configuration for recovering to a new target.
	NewTargetConfig *VmwareRecoverFilesNewTargetConfig `json:"newTargetConfig,omitempty"`

	// Specifies the configuration for recovering to the original target.
	OriginalTargetConfig *VmwareRecoverFilesOriginalTargetConfig `json:"originalTargetConfig,omitempty"`

	// Specifies VLAN Params associated with the recovered files and folders. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// Validate validates this vmware target params for recover file and folder
func (m *VmwareTargetParamsForRecoverFileAndFolder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoverToOriginalTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewTargetConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalTargetConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareTargetParamsForRecoverFileAndFolder) validateRecoverToOriginalTarget(formats strfmt.Registry) error {

	if err := validate.Required("recoverToOriginalTarget", "body", m.RecoverToOriginalTarget); err != nil {
		return err
	}

	return nil
}

func (m *VmwareTargetParamsForRecoverFileAndFolder) validateNewTargetConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NewTargetConfig) { // not required
		return nil
	}

	if m.NewTargetConfig != nil {
		if err := m.NewTargetConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newTargetConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newTargetConfig")
			}
			return err
		}
	}

	return nil
}

func (m *VmwareTargetParamsForRecoverFileAndFolder) validateOriginalTargetConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalTargetConfig) { // not required
		return nil
	}

	if m.OriginalTargetConfig != nil {
		if err := m.OriginalTargetConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originalTargetConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originalTargetConfig")
			}
			return err
		}
	}

	return nil
}

func (m *VmwareTargetParamsForRecoverFileAndFolder) validateVlanConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanConfig) { // not required
		return nil
	}

	if m.VlanConfig != nil {
		if err := m.VlanConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vmware target params for recover file and folder based on the context it is used
func (m *VmwareTargetParamsForRecoverFileAndFolder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNewTargetConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginalTargetConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmwareTargetParamsForRecoverFileAndFolder) contextValidateNewTargetConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NewTargetConfig != nil {

		if swag.IsZero(m.NewTargetConfig) { // not required
			return nil
		}

		if err := m.NewTargetConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newTargetConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newTargetConfig")
			}
			return err
		}
	}

	return nil
}

func (m *VmwareTargetParamsForRecoverFileAndFolder) contextValidateOriginalTargetConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginalTargetConfig != nil {

		if swag.IsZero(m.OriginalTargetConfig) { // not required
			return nil
		}

		if err := m.OriginalTargetConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originalTargetConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originalTargetConfig")
			}
			return err
		}
	}

	return nil
}

func (m *VmwareTargetParamsForRecoverFileAndFolder) contextValidateVlanConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.VlanConfig != nil {

		if swag.IsZero(m.VlanConfig) { // not required
			return nil
		}

		if err := m.VlanConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmwareTargetParamsForRecoverFileAndFolder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTargetParamsForRecoverFileAndFolder) UnmarshalBinary(b []byte) error {
	var res VmwareTargetParamsForRecoverFileAndFolder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
