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
	"github.com/go-openapi/validate"
)

// VmwareTargetParamsForMountVolume VMware Target Params.
//
// Specifies the parameters for a VMware mount volume target.
//
// swagger:model VmwareTargetParamsForMountVolume
type VmwareTargetParamsForMountVolume struct {

	// Specifies whether to mount to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
	// Required: true
	MountToOriginalTarget *bool `json:"mountToOriginalTarget"`

	// Specifies whether to perform a read-only mount. Default is false.
	ReadOnlyMount *bool `json:"readOnlyMount,omitempty"`

	// Specifies the names of volumes that need to be mounted. If this is not specified then all volumes that are part of the source VM will be mounted on the target VM.
	VolumeNames []string `json:"volumeNames"`

	// Specifies the mapping of original volumes and mounted volumes
	// Read Only: true
	MountedVolumeMapping []*MountedVolumeMapping `json:"mountedVolumeMapping"`

	// Specifies the configuration for mounting to a new target.
	NewTargetConfig *VMwareMountVolumesNewTargetConfig `json:"newTargetConfig,omitempty"`

	// Specifies the configuration for mounting to the original target.
	OriginalTargetConfig *VMwareMountVolumesOriginalTargetConfig `json:"originalTargetConfig,omitempty"`

	// Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// Validate validates this vmware target params for mount volume
func (m *VmwareTargetParamsForMountVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMountToOriginalTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountedVolumeMapping(formats); err != nil {
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

func (m *VmwareTargetParamsForMountVolume) validateMountToOriginalTarget(formats strfmt.Registry) error {

	if err := validate.Required("mountToOriginalTarget", "body", m.MountToOriginalTarget); err != nil {
		return err
	}

	return nil
}

func (m *VmwareTargetParamsForMountVolume) validateMountedVolumeMapping(formats strfmt.Registry) error {
	if swag.IsZero(m.MountedVolumeMapping) { // not required
		return nil
	}

	for i := 0; i < len(m.MountedVolumeMapping); i++ {
		if swag.IsZero(m.MountedVolumeMapping[i]) { // not required
			continue
		}

		if m.MountedVolumeMapping[i] != nil {
			if err := m.MountedVolumeMapping[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mountedVolumeMapping" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mountedVolumeMapping" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VmwareTargetParamsForMountVolume) validateNewTargetConfig(formats strfmt.Registry) error {
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

func (m *VmwareTargetParamsForMountVolume) validateOriginalTargetConfig(formats strfmt.Registry) error {
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

func (m *VmwareTargetParamsForMountVolume) validateVlanConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this vmware target params for mount volume based on the context it is used
func (m *VmwareTargetParamsForMountVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMountedVolumeMapping(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *VmwareTargetParamsForMountVolume) contextValidateMountedVolumeMapping(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mountedVolumeMapping", "body", []*MountedVolumeMapping(m.MountedVolumeMapping)); err != nil {
		return err
	}

	for i := 0; i < len(m.MountedVolumeMapping); i++ {

		if m.MountedVolumeMapping[i] != nil {

			if swag.IsZero(m.MountedVolumeMapping[i]) { // not required
				return nil
			}

			if err := m.MountedVolumeMapping[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mountedVolumeMapping" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mountedVolumeMapping" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VmwareTargetParamsForMountVolume) contextValidateNewTargetConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VmwareTargetParamsForMountVolume) contextValidateOriginalTargetConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VmwareTargetParamsForMountVolume) contextValidateVlanConfig(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VmwareTargetParamsForMountVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTargetParamsForMountVolume) UnmarshalBinary(b []byte) error {
	var res VmwareTargetParamsForMountVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
