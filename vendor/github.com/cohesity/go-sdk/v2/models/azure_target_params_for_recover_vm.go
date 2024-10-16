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

// AzureTargetParamsForRecoverVM Azure Recovery Target Params
//
// Specifies the parameters for an Azure recovery target.
//
// swagger:model AzureTargetParamsForRecoverVm
type AzureTargetParamsForRecoverVM struct {

	// Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state.
	PowerOnVms *bool `json:"powerOnVms,omitempty"`

	// Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained.
	RecoveryTargetConfig *AzureVMRecoveryTargetConfig `json:"recoveryTargetConfig,omitempty"`

	// Specifies params to rename the VMs that are recovered. If not specified, the original names of the VMs are preserved.
	RenameRecoveredVmsParams *RecoveredOrClonedVmsRenameConfig `json:"renameRecoveredVmsParams,omitempty"`
}

// Validate validates this azure target params for recover Vm
func (m *AzureTargetParamsForRecoverVM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoveryTargetConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenameRecoveredVmsParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureTargetParamsForRecoverVM) validateRecoveryTargetConfig(formats strfmt.Registry) error {
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

func (m *AzureTargetParamsForRecoverVM) validateRenameRecoveredVmsParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RenameRecoveredVmsParams) { // not required
		return nil
	}

	if m.RenameRecoveredVmsParams != nil {
		if err := m.RenameRecoveredVmsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("renameRecoveredVmsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("renameRecoveredVmsParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this azure target params for recover Vm based on the context it is used
func (m *AzureTargetParamsForRecoverVM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecoveryTargetConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRenameRecoveredVmsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureTargetParamsForRecoverVM) contextValidateRecoveryTargetConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AzureTargetParamsForRecoverVM) contextValidateRenameRecoveredVmsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RenameRecoveredVmsParams != nil {

		if swag.IsZero(m.RenameRecoveredVmsParams) { // not required
			return nil
		}

		if err := m.RenameRecoveredVmsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("renameRecoveredVmsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("renameRecoveredVmsParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureTargetParamsForRecoverVM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureTargetParamsForRecoverVM) UnmarshalBinary(b []byte) error {
	var res AzureTargetParamsForRecoverVM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
