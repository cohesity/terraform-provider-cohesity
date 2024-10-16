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

// AcropolisTargetParamsForRecoverVM Acropolis Target Params.
//
// Specifies the parameters for an Acropolis recovery target.
//
// swagger:model AcropolisTargetParamsForRecoverVm
type AcropolisTargetParamsForRecoverVM struct {

	// Specifies type of Recovery Process to be used. InstantRecovery/CopyRecovery etc... Default value is InstantRecovery.
	// Enum: ["InstantRecovery","CopyRecovery"]
	RecoveryProcessType string `json:"recoveryProcessType,omitempty"`

	// Specifies whether to power on vms after recovery. If not specified, or false, recovered vms will be in powered off state.
	PowerOnVms *bool `json:"powerOnVms,omitempty"`

	// Specifies whether to continue recovering other vms if one of vms failed to recover. Default value is false.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// Specifies whether to recover excluded disk while performing recovery of a VM by creating empty disks for them. Default value is false.
	RecoverExcludedDisk *bool `json:"recoverExcludedDisk,omitempty"`

	// Specifies the recovery target configuration if recovery has to be done to a different location which is different from original source or to original Source with different configuration. If not specified, then the recovery of the vms will be performed to original location with all configuration parameters retained.
	RecoveryTargetConfig *RecoverAcropolisVMTargetConfig `json:"recoveryTargetConfig,omitempty"`

	// Specifies params to rename the VMs that are recovered. If not specified, the original names of the VMs are preserved.
	RenameRecoveredVmsParams *RecoveredOrClonedVmsRenameConfig `json:"renameRecoveredVmsParams,omitempty"`

	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// Validate validates this acropolis target params for recover Vm
func (m *AcropolisTargetParamsForRecoverVM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoveryProcessType(formats); err != nil {
		res = append(res, err)
	}

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

var acropolisTargetParamsForRecoverVmTypeRecoveryProcessTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InstantRecovery","CopyRecovery"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		acropolisTargetParamsForRecoverVmTypeRecoveryProcessTypePropEnum = append(acropolisTargetParamsForRecoverVmTypeRecoveryProcessTypePropEnum, v)
	}
}

const (

	// AcropolisTargetParamsForRecoverVMRecoveryProcessTypeInstantRecovery captures enum value "InstantRecovery"
	AcropolisTargetParamsForRecoverVMRecoveryProcessTypeInstantRecovery string = "InstantRecovery"

	// AcropolisTargetParamsForRecoverVMRecoveryProcessTypeCopyRecovery captures enum value "CopyRecovery"
	AcropolisTargetParamsForRecoverVMRecoveryProcessTypeCopyRecovery string = "CopyRecovery"
)

// prop value enum
func (m *AcropolisTargetParamsForRecoverVM) validateRecoveryProcessTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, acropolisTargetParamsForRecoverVmTypeRecoveryProcessTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AcropolisTargetParamsForRecoverVM) validateRecoveryProcessType(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoveryProcessType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecoveryProcessTypeEnum("recoveryProcessType", "body", m.RecoveryProcessType); err != nil {
		return err
	}

	return nil
}

func (m *AcropolisTargetParamsForRecoverVM) validateRecoveryTargetConfig(formats strfmt.Registry) error {
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

func (m *AcropolisTargetParamsForRecoverVM) validateRenameRecoveredVmsParams(formats strfmt.Registry) error {
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

func (m *AcropolisTargetParamsForRecoverVM) validateVlanConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this acropolis target params for recover Vm based on the context it is used
func (m *AcropolisTargetParamsForRecoverVM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *AcropolisTargetParamsForRecoverVM) contextValidateRecoveryTargetConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AcropolisTargetParamsForRecoverVM) contextValidateRenameRecoveredVmsParams(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AcropolisTargetParamsForRecoverVM) contextValidateVlanConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AcropolisTargetParamsForRecoverVM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcropolisTargetParamsForRecoverVM) UnmarshalBinary(b []byte) error {
	var res AcropolisTargetParamsForRecoverVM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
