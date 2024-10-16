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

// RestoreObjectParams The parameters that control the various options to restore a RestoreObject.
//
// swagger:model RestoreObjectParams
type RestoreObjectParams struct {

	// The action to perform.
	Action *int32 `json:"action,omitempty"`

	// A datastore entity where the object's files should be restored to. This
	// field is optional if object is being restored to its original parent
	// source. If not specified, the object's files will be restored to their
	// original datastore locations. This field is mandatory if object is being
	// restored to a different resource pool or to a different parent source.
	DatastoreEntity *EntityProto `json:"datastoreEntity,omitempty"`

	// The power state configuration to be applied to the restored object.
	//
	// Semantics for kCloneVMs task: By default, objects are restored in the
	// powered off state.
	//
	// Semantics for kRecoverVMs task: By default, objects are restored in the
	// powered on state.
	//
	// This proto can be used to override the default behavior for kCloneVMs or
	// kRecoverVMs task.
	PowerStateConfig *PowerStateConfigProto `json:"powerStateConfig,omitempty"`

	// By default, objects are restored with their original name. This field can
	// be used to specify the transformation ( i.e prefix/suffix) to be applied
	// to the source object name to derive the new name of the restored object.
	RenameRestoredObjectParam *RenameObjectParamProto `json:"renameRestoredObjectParam,omitempty"`

	// The resource pool entity where the restored objects will be attached. This
	// field is mandatory for a kCloneVMs task.
	//
	// This field is optional for a kRecoverVMs task if the objects are being
	// restored to its original parent source. If not specified, restored objects
	// will be attached to its original resource entity. This field is mandatory
	// if objects are being restored to a different parent source.
	ResourcePoolEntity *EntityProto `json:"resourcePoolEntity,omitempty"`

	// An optional registered parent source to which objects are to be restored.
	// If not specified, objects are restored back to the original source that
	// was managing the objects.
	RestoreParentSource *EntityProto `json:"restoreParentSource,omitempty"`

	// The network configuration to be applied to the restored object.
	//
	// Semantics for kCloneVMs task: By default, if objects are being restored to
	// their original location, then network will be disabled. If objects are
	// being restored to a different location (i.e., either to different resource
	// pool or to different parent source), then network will be detached.
	//
	// Semantics for kRecoverVMs task: By default, if objects are being restored
	// to their original location, then original network configuration will be
	// preserved. If objects are being restored to different location, (i.e.,
	// either to different resource pool or to different parent source), then
	// network will be detached.
	//
	// Below field can be used to override the default network configuration of
	// the restored objects.
	//
	// If user want to keep the original network setting for kRecoverVMs task,
	// then this proto should not be set.
	RestoredObjectsNetworkConfig *RestoredObjectNetworkConfigProto `json:"restoredObjectsNetworkConfig,omitempty"`

	// Target view into which the objects are to be cloned.
	ViewName *string `json:"viewName,omitempty"`
}

// Validate validates this restore object params
func (m *RestoreObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastoreEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerStateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenameRestoredObjectParam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcePoolEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreParentSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoredObjectsNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreObjectParams) validateDatastoreEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.DatastoreEntity) { // not required
		return nil
	}

	if m.DatastoreEntity != nil {
		if err := m.DatastoreEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastoreEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastoreEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreObjectParams) validatePowerStateConfig(formats strfmt.Registry) error {
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

func (m *RestoreObjectParams) validateRenameRestoredObjectParam(formats strfmt.Registry) error {
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

func (m *RestoreObjectParams) validateResourcePoolEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourcePoolEntity) { // not required
		return nil
	}

	if m.ResourcePoolEntity != nil {
		if err := m.ResourcePoolEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourcePoolEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourcePoolEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreObjectParams) validateRestoreParentSource(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoreParentSource) { // not required
		return nil
	}

	if m.RestoreParentSource != nil {
		if err := m.RestoreParentSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreParentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreParentSource")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreObjectParams) validateRestoredObjectsNetworkConfig(formats strfmt.Registry) error {
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

// ContextValidate validate this restore object params based on the context it is used
func (m *RestoreObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatastoreEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePowerStateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRenameRestoredObjectParam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourcePoolEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoreParentSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoredObjectsNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreObjectParams) contextValidateDatastoreEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.DatastoreEntity != nil {

		if swag.IsZero(m.DatastoreEntity) { // not required
			return nil
		}

		if err := m.DatastoreEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastoreEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastoreEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreObjectParams) contextValidatePowerStateConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RestoreObjectParams) contextValidateRenameRestoredObjectParam(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RestoreObjectParams) contextValidateResourcePoolEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourcePoolEntity != nil {

		if swag.IsZero(m.ResourcePoolEntity) { // not required
			return nil
		}

		if err := m.ResourcePoolEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourcePoolEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourcePoolEntity")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreObjectParams) contextValidateRestoreParentSource(ctx context.Context, formats strfmt.Registry) error {

	if m.RestoreParentSource != nil {

		if swag.IsZero(m.RestoreParentSource) { // not required
			return nil
		}

		if err := m.RestoreParentSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restoreParentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restoreParentSource")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreObjectParams) contextValidateRestoredObjectsNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *RestoreObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreObjectParams) UnmarshalBinary(b []byte) error {
	var res RestoreObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
