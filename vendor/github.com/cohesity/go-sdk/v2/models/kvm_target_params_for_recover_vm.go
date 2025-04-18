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

// KvmTargetParamsForRecoverVM KVM Target Params.
//
// Specifies the parameters for a KVM recovery target.
//
// swagger:model KvmTargetParamsForRecoverVm
type KvmTargetParamsForRecoverVM struct {

	// Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state.
	PowerOnVms *bool `json:"powerOnVms,omitempty"`

	// Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Snource with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained.
	RecoveryTargetConfig *KvmVMRecoveryTargetConfig `json:"recoveryTargetConfig,omitempty"`

	// Specifies params to rename the VMs that are recovered. If not specified, the original names of the VMs are preserved.
	RenameRecoveredVmsParams *RecoveredOrClonedVmsRenameConfig `json:"renameRecoveredVmsParams,omitempty"`

	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// Validate validates this kvm target params for recover Vm
func (m *KvmTargetParamsForRecoverVM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoveryTargetConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenameRecoveredVmsParams(formats); err != nil {
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

func (m *KvmTargetParamsForRecoverVM) validateRecoveryTargetConfig(formats strfmt.Registry) error {
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

func (m *KvmTargetParamsForRecoverVM) validateRenameRecoveredVmsParams(formats strfmt.Registry) error {
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

func (m *KvmTargetParamsForRecoverVM) validateVlanConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this kvm target params for recover Vm based on the context it is used
func (m *KvmTargetParamsForRecoverVM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecoveryTargetConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRenameRecoveredVmsParams(ctx, formats); err != nil {
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

func (m *KvmTargetParamsForRecoverVM) contextValidateRecoveryTargetConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *KvmTargetParamsForRecoverVM) contextValidateRenameRecoveredVmsParams(ctx context.Context, formats strfmt.Registry) error {

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

func (m *KvmTargetParamsForRecoverVM) contextValidateVlanConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *KvmTargetParamsForRecoverVM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KvmTargetParamsForRecoverVM) UnmarshalBinary(b []byte) error {
	var res KvmTargetParamsForRecoverVM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
