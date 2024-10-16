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

// RestoreKVMVMsParams This message defines the KVM specific VMs restore params.
//
// swagger:model RestoreKVMVMsParams
type RestoreKVMVMsParams struct {

	// The resource (KVMH host) to which the restored VM will be attached.
	//
	// This field is optional for a kRecoverVMs task if the VMs are being
	// restored to its original parent source. If not specified, restored VMs
	// will be attached to its original host. This field is mandatory if the
	// VMs are being restored to a different parent source.
	ClusterEntity *EntityProto `json:"clusterEntity,omitempty"`

	// A datacenter where the VM's files should be restored to. This
	// field is optional if VM is being restored to its original parent
	// source. If not specified, the VM's files will be restored to their
	// original datacenter locations. This field is mandatory if VM is being
	// restored to a different datacenter entity or to a different parent source.
	DatacenterEntity *EntityProto `json:"datacenterEntity,omitempty"`

	// The power state configuration to be applied to the restored VM.
	//
	// Semantics for kRecoverVMs task: By default, VMs are restored in their
	// original power state.
	//
	// This proto can be used to override the default behavior for the restore
	// task.
	PowerStateConfig *PowerStateConfigProto `json:"powerStateConfig,omitempty"`

	// By default, VMs are restored with their original name. This field can
	// be used to specify the transformation ( i.e prefix/suffix) to be applied
	// to the source VM name to derive the new name of the restored VM.
	RenameRestoredObjectParam *RenameObjectParamProto `json:"renameRestoredObjectParam,omitempty"`

	// The network configuration to be applied to the restored VMs.
	//
	// Semantics for kRecoverVMs task: By default, if the VMs are being restored
	// to their original location, then original network configuration will be
	// preserved. If objects are being restored to different location (i.e.,
	// different parent source), then network will be detached.
	//
	// Below field can be used to override the default network configuration of
	// the restored VMs.
	//
	// If user want to keep the original network setting for kRecoverVMs task,
	// then this proto should not be set.
	RestoredObjectsNetworkConfig *RestoredObjectNetworkConfigProto `json:"restoredObjectsNetworkConfig,omitempty"`

	// A storagedomain where the VM's disk should be restored to. This
	// field is optional if VM is being restored to its original parent
	// source. If not specified, the VM's disk will be restored to their
	// original storagedomain. This field is mandatory if VM is being
	// restored to a different resource entity or to a different parent source.
	StoragedomainEntity *EntityProto `json:"storagedomainEntity,omitempty"`
}

// Validate validates this restore k VM v ms params
func (m *RestoreKVMVMsParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatacenterEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerStateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenameRestoredObjectParam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoredObjectsNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragedomainEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreKVMVMsParams) validateClusterEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterEntity) { // not required
		return nil
	}

	if m.ClusterEntity != nil {
		if err := m.ClusterEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) validateDatacenterEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.DatacenterEntity) { // not required
		return nil
	}

	if m.DatacenterEntity != nil {
		if err := m.DatacenterEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datacenterEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datacenterEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) validatePowerStateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.PowerStateConfig) { // not required
		return nil
	}

	if m.PowerStateConfig != nil {
		if err := m.PowerStateConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("powerStateConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("powerStateConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) validateRenameRestoredObjectParam(formats strfmt.Registry) error {
	if swag.IsZero(m.RenameRestoredObjectParam) { // not required
		return nil
	}

	if m.RenameRestoredObjectParam != nil {
		if err := m.RenameRestoredObjectParam.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("renameRestoredObjectParam")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("renameRestoredObjectParam")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) validateRestoredObjectsNetworkConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoredObjectsNetworkConfig) { // not required
		return nil
	}

	if m.RestoredObjectsNetworkConfig != nil {
		if err := m.RestoredObjectsNetworkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoredObjectsNetworkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoredObjectsNetworkConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) validateStoragedomainEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.StoragedomainEntity) { // not required
		return nil
	}

	if m.StoragedomainEntity != nil {
		if err := m.StoragedomainEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storagedomainEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storagedomainEntity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this restore k VM v ms params based on the context it is used
func (m *RestoreKVMVMsParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatacenterEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowerStateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRenameRestoredObjectParam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoredObjectsNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoragedomainEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreKVMVMsParams) contextValidateClusterEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterEntity != nil {

		if swag.IsZero(m.ClusterEntity) { // not required
			return nil
		}

		if err := m.ClusterEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) contextValidateDatacenterEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.DatacenterEntity != nil {

		if swag.IsZero(m.DatacenterEntity) { // not required
			return nil
		}

		if err := m.DatacenterEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datacenterEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datacenterEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) contextValidatePowerStateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.PowerStateConfig != nil {

		if swag.IsZero(m.PowerStateConfig) { // not required
			return nil
		}

		if err := m.PowerStateConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("powerStateConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("powerStateConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) contextValidateRenameRestoredObjectParam(ctx context.Context, formats strfmt.Registry) error {

	if m.RenameRestoredObjectParam != nil {

		if swag.IsZero(m.RenameRestoredObjectParam) { // not required
			return nil
		}

		if err := m.RenameRestoredObjectParam.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("renameRestoredObjectParam")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("renameRestoredObjectParam")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) contextValidateRestoredObjectsNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.RestoredObjectsNetworkConfig != nil {

		if swag.IsZero(m.RestoredObjectsNetworkConfig) { // not required
			return nil
		}

		if err := m.RestoredObjectsNetworkConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoredObjectsNetworkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoredObjectsNetworkConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreKVMVMsParams) contextValidateStoragedomainEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.StoragedomainEntity != nil {

		if swag.IsZero(m.StoragedomainEntity) { // not required
			return nil
		}

		if err := m.StoragedomainEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storagedomainEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storagedomainEntity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreKVMVMsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreKVMVMsParams) UnmarshalBinary(b []byte) error {
	var res RestoreKVMVMsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
