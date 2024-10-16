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

// RestoreObjectState Restore Object State.
//
// Specifies the state of an individual object in a Restore Task.
//
// swagger:model RestoreObjectState
type RestoreObjectState struct {

	// Specifies if an error occurred during the restore operation.
	Error *RequestError `json:"error,omitempty"`

	// Specifies the status of an object during a Restore Task.
	// 'kFilesCloned' indicates that the cloning has completed.
	// 'kFetchedEntityInfo' indicates that information about the object was
	// fetched from the primary source.
	// 'kVMCreated' indicates that the new VM was created.
	// 'kRelocationStarted' indicates that restoring to a different
	// resource pool has started.
	// 'kFinished' indicates that the Restore Task has finished.
	// Whether it was successful or not is indicated by the error code that
	// that is stored with the Restore Task.
	// 'kAborted' indicates that the Restore Task was aborted before
	// trying to restore this object. This can happen if the
	// Restore Task encounters a global error.
	// For example during a 'kCloneVMs' Restore Task, the datastore
	// could not be mounted. The entire Restore Task is aborted
	// before trying to create VMs on the primary source.
	// 'kDataCopyStarted' indicates that the disk copy is started.
	// 'kInProgress' captures a generic in-progress state and can be used by restore
	// operations that don't track individual states.
	// 'kVMDeleted' captures a VM that has been deleted from the vCenter upon a failed
	// import to vCD.
	// 'kCancelled' captures the cancellation of the restore for the entity.
	// Enum: ["kFilesCloned","kFetchedEntityInfo","kVMCreated","kRelocationStarted","kFinished","kAborted","kDataCopyStarted","kInProgress","kVMDeleted","kCancelled"]
	ObjectStatus *string `json:"objectStatus,omitempty"`

	// Specifies the id of the Resource Pool that the restored
	// object is attached to.
	// For a 'kRecoverVMs' Restore Task, an object can be recovered
	// back to its original resource pool. This means while recovering a
	// set of objects, this field can reference different resource pools.
	// For a 'kCloneVMs' Restore Task, all objects are attached to the
	// same resource pool. However, this field will still be
	// populated.
	// NOTE: This field may not be populated if the restore of the object
	// fails.
	ResourcePoolID *int64 `json:"resourcePoolId,omitempty"`

	// Specifies the Id of the recovered or cloned object.
	// NOTE: For a Restore Task that is recovering or cloning an object
	// in the VMware environment, after the VM is created it is storage
	// vMotioned to its primary datastore. If storage
	// vMotion fails, the Cohesity Cluster marks the recovery task as failed.
	// However, this field is still populated with the id of the
	// recovered VM. This id can be used later to
	// clean up the VM from primary environment (i.e, the vCenter Server).
	RestoredObjectID *int64 `json:"restoredObjectId,omitempty"`

	// Specifies the Protection Source id of the object to be recovered or
	// cloned.
	SourceObjectID *int64 `json:"sourceObjectId,omitempty"`
}

// Validate validates this restore object state
func (m *RestoreObjectState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreObjectState) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

var restoreObjectStateTypeObjectStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kFilesCloned","kFetchedEntityInfo","kVMCreated","kRelocationStarted","kFinished","kAborted","kDataCopyStarted","kInProgress","kVMDeleted","kCancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		restoreObjectStateTypeObjectStatusPropEnum = append(restoreObjectStateTypeObjectStatusPropEnum, v)
	}
}

const (

	// RestoreObjectStateObjectStatusKFilesCloned captures enum value "kFilesCloned"
	RestoreObjectStateObjectStatusKFilesCloned string = "kFilesCloned"

	// RestoreObjectStateObjectStatusKFetchedEntityInfo captures enum value "kFetchedEntityInfo"
	RestoreObjectStateObjectStatusKFetchedEntityInfo string = "kFetchedEntityInfo"

	// RestoreObjectStateObjectStatusKVMCreated captures enum value "kVMCreated"
	RestoreObjectStateObjectStatusKVMCreated string = "kVMCreated"

	// RestoreObjectStateObjectStatusKRelocationStarted captures enum value "kRelocationStarted"
	RestoreObjectStateObjectStatusKRelocationStarted string = "kRelocationStarted"

	// RestoreObjectStateObjectStatusKFinished captures enum value "kFinished"
	RestoreObjectStateObjectStatusKFinished string = "kFinished"

	// RestoreObjectStateObjectStatusKAborted captures enum value "kAborted"
	RestoreObjectStateObjectStatusKAborted string = "kAborted"

	// RestoreObjectStateObjectStatusKDataCopyStarted captures enum value "kDataCopyStarted"
	RestoreObjectStateObjectStatusKDataCopyStarted string = "kDataCopyStarted"

	// RestoreObjectStateObjectStatusKInProgress captures enum value "kInProgress"
	RestoreObjectStateObjectStatusKInProgress string = "kInProgress"

	// RestoreObjectStateObjectStatusKVMDeleted captures enum value "kVMDeleted"
	RestoreObjectStateObjectStatusKVMDeleted string = "kVMDeleted"

	// RestoreObjectStateObjectStatusKCancelled captures enum value "kCancelled"
	RestoreObjectStateObjectStatusKCancelled string = "kCancelled"
)

// prop value enum
func (m *RestoreObjectState) validateObjectStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, restoreObjectStateTypeObjectStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RestoreObjectState) validateObjectStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateObjectStatusEnum("objectStatus", "body", *m.ObjectStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this restore object state based on the context it is used
func (m *RestoreObjectState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreObjectState) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreObjectState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreObjectState) UnmarshalBinary(b []byte) error {
	var res RestoreObjectState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
