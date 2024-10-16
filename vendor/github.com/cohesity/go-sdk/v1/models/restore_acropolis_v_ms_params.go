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

// RestoreAcropolisVMsParams This message defines the Acropolis specific VMs restore params.
//
// swagger:model RestoreAcropolisVMsParams
type RestoreAcropolisVMsParams struct {

	// Whether to perform copy recovery instead of instant recovery.
	CopyRecovery *bool `json:"copyRecovery,omitempty"`

	// The power state configuration to be applied to the restored VM.
	//
	// Semantics for kRecoverVMs task: By default, VMs are restored in their
	// original power state.
	//
	// This proto can be used to override the default behavior for the restore
	// task.
	PowerStateConfig *PowerStateConfigProto `json:"powerStateConfig,omitempty"`

	// Whether disks that were excluded during backup should be recovered as
	// blank disks.
	RecoverExcludedDisksAsBlankDisks *bool `json:"recoverExcludedDisksAsBlankDisks,omitempty"`

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

	// A storage container where the VM's files should be restored to. This
	// field is optional if the VM is being restored to its original parent
	// source. If not specified, the VM's files will be restored to their
	// original storage container(s). This field is mandatory if the VMs are
	// being restored to a different parent source.
	StorageContainer *EntityProto `json:"storageContainer,omitempty"`

	// UUID config to use for restored VM.
	//
	// Semantics for kRecoverVMs task: By default, recovered VMs have new UUIDs
	// for them.
	//
	// This proto can be used to override the default behavior for the restore
	// task.
	UUIDConfig *UUIDConfigProto `json:"uuidConfig,omitempty"`
}

// Validate validates this restore acropolis v ms params
func (m *RestoreAcropolisVMsParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePowerStateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenameRestoredObjectParam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoredObjectsNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUIDConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreAcropolisVMsParams) validatePowerStateConfig(formats strfmt.Registry) error {
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

func (m *RestoreAcropolisVMsParams) validateRenameRestoredObjectParam(formats strfmt.Registry) error {
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

func (m *RestoreAcropolisVMsParams) validateRestoredObjectsNetworkConfig(formats strfmt.Registry) error {
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

func (m *RestoreAcropolisVMsParams) validateStorageContainer(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageContainer) { // not required
		return nil
	}

	if m.StorageContainer != nil {
		if err := m.StorageContainer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageContainer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageContainer")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAcropolisVMsParams) validateUUIDConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.UUIDConfig) { // not required
		return nil
	}

	if m.UUIDConfig != nil {
		if err := m.UUIDConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uuidConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uuidConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this restore acropolis v ms params based on the context it is used
func (m *RestoreAcropolisVMsParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePowerStateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRenameRestoredObjectParam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoredObjectsNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageContainer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUIDConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreAcropolisVMsParams) contextValidatePowerStateConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RestoreAcropolisVMsParams) contextValidateRenameRestoredObjectParam(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RestoreAcropolisVMsParams) contextValidateRestoredObjectsNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RestoreAcropolisVMsParams) contextValidateStorageContainer(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageContainer != nil {

		if swag.IsZero(m.StorageContainer) { // not required
			return nil
		}

		if err := m.StorageContainer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageContainer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageContainer")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAcropolisVMsParams) contextValidateUUIDConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.UUIDConfig != nil {

		if swag.IsZero(m.UUIDConfig) { // not required
			return nil
		}

		if err := m.UUIDConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uuidConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uuidConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreAcropolisVMsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreAcropolisVMsParams) UnmarshalBinary(b []byte) error {
	var res RestoreAcropolisVMsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
