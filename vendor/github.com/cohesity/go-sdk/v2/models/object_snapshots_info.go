// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ObjectSnapshotsInfo Object Snapshots Information.
//
// Specifies the snapshots information for every Protection Group for a given object.
//
// swagger:model ObjectSnapshotsInfo
type ObjectSnapshotsInfo struct {

	// Specifies the archival snapshots information.
	ArchivalSnapshotsInfo []*ObjectArchivalSnapshotInfo `json:"archivalSnapshotsInfo"`

	// Specifies the indexing status of objects in this snapshot.<br> 'InProgress' indicates the indexing is in progress.<br> 'Done' indicates indexing is done.<br> 'NoIndex' indicates indexing is not applicable.<br> 'Error' indicates indexing failed with error.
	// Enum: ["InProgress","Done","NoIndex","Error"]
	IndexingStatus *string `json:"indexingStatus,omitempty"`

	// Specifies id of the Protection Group.
	ProtectionGroupID *string `json:"protectionGroupId,omitempty"`

	// Specifies name of the Protection Group.
	ProtectionGroupName *string `json:"protectionGroupName,omitempty"`

	// Specifies the instance id of the protection run which create the snapshot.
	RunInstanceID *int64 `json:"runInstanceId,omitempty"`

	// Specifies the source protection group id in case of replication.
	SourceGroupID *string `json:"sourceGroupId,omitempty"`

	// Specifies the Storage Domain id where the backup data of Object is present.
	StorageDomainID *int64 `json:"storageDomainId,omitempty"`

	// Specifies the name of Storage Domain id where the backup data of Object is present
	StorageDomainName *string `json:"storageDomainName,omitempty"`

	// Specifies the id of Protection Group Run.
	ProtectionRunID *string `json:"protectionRunId,omitempty"`

	// Specifies the type of protection run created this snapshot.
	// Enum: ["kRegular","kFull","kLog","kSystem","kHydrateCDP","kStorageArraySnapshot"]
	RunType *string `json:"runType,omitempty"`

	// Specifies the start time of Protection Group Run in Unix timestamp epoch in microseconds.
	ProtectionRunStartTimeUsecs *int64 `json:"protectionRunStartTimeUsecs,omitempty"`

	// Specifies the end time of Protection Group Run in Unix timestamp epoch in microseconds.
	ProtectionRunEndTimeUsecs *int64 `json:"protectionRunEndTimeUsecs,omitempty"`

	// Specifies the local snapshot information.
	LocalSnapshotInfo *ObjectLocalSnapshotInfo `json:"localSnapshotInfo,omitempty"`
}

// Validate validates this object snapshots info
func (m *ObjectSnapshotsInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchivalSnapshotsInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalSnapshotInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectSnapshotsInfo) validateArchivalSnapshotsInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchivalSnapshotsInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.ArchivalSnapshotsInfo); i++ {
		if swag.IsZero(m.ArchivalSnapshotsInfo[i]) { // not required
			continue
		}

		if m.ArchivalSnapshotsInfo[i] != nil {
			if err := m.ArchivalSnapshotsInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("archivalSnapshotsInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("archivalSnapshotsInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var objectSnapshotsInfoTypeIndexingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["InProgress","Done","NoIndex","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectSnapshotsInfoTypeIndexingStatusPropEnum = append(objectSnapshotsInfoTypeIndexingStatusPropEnum, v)
	}
}

const (

	// ObjectSnapshotsInfoIndexingStatusInProgress captures enum value "InProgress"
	ObjectSnapshotsInfoIndexingStatusInProgress string = "InProgress"

	// ObjectSnapshotsInfoIndexingStatusDone captures enum value "Done"
	ObjectSnapshotsInfoIndexingStatusDone string = "Done"

	// ObjectSnapshotsInfoIndexingStatusNoIndex captures enum value "NoIndex"
	ObjectSnapshotsInfoIndexingStatusNoIndex string = "NoIndex"

	// ObjectSnapshotsInfoIndexingStatusError captures enum value "Error"
	ObjectSnapshotsInfoIndexingStatusError string = "Error"
)

// prop value enum
func (m *ObjectSnapshotsInfo) validateIndexingStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectSnapshotsInfoTypeIndexingStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectSnapshotsInfo) validateIndexingStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexingStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateIndexingStatusEnum("indexingStatus", "body", *m.IndexingStatus); err != nil {
		return err
	}

	return nil
}

var objectSnapshotsInfoTypeRunTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kRegular","kFull","kLog","kSystem","kHydrateCDP","kStorageArraySnapshot"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectSnapshotsInfoTypeRunTypePropEnum = append(objectSnapshotsInfoTypeRunTypePropEnum, v)
	}
}

const (

	// ObjectSnapshotsInfoRunTypeKRegular captures enum value "kRegular"
	ObjectSnapshotsInfoRunTypeKRegular string = "kRegular"

	// ObjectSnapshotsInfoRunTypeKFull captures enum value "kFull"
	ObjectSnapshotsInfoRunTypeKFull string = "kFull"

	// ObjectSnapshotsInfoRunTypeKLog captures enum value "kLog"
	ObjectSnapshotsInfoRunTypeKLog string = "kLog"

	// ObjectSnapshotsInfoRunTypeKSystem captures enum value "kSystem"
	ObjectSnapshotsInfoRunTypeKSystem string = "kSystem"

	// ObjectSnapshotsInfoRunTypeKHydrateCDP captures enum value "kHydrateCDP"
	ObjectSnapshotsInfoRunTypeKHydrateCDP string = "kHydrateCDP"

	// ObjectSnapshotsInfoRunTypeKStorageArraySnapshot captures enum value "kStorageArraySnapshot"
	ObjectSnapshotsInfoRunTypeKStorageArraySnapshot string = "kStorageArraySnapshot"
)

// prop value enum
func (m *ObjectSnapshotsInfo) validateRunTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectSnapshotsInfoTypeRunTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectSnapshotsInfo) validateRunType(formats strfmt.Registry) error {
	if swag.IsZero(m.RunType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRunTypeEnum("runType", "body", *m.RunType); err != nil {
		return err
	}

	return nil
}

func (m *ObjectSnapshotsInfo) validateLocalSnapshotInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalSnapshotInfo) { // not required
		return nil
	}

	if m.LocalSnapshotInfo != nil {
		if err := m.LocalSnapshotInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localSnapshotInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localSnapshotInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this object snapshots info based on the context it is used
func (m *ObjectSnapshotsInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchivalSnapshotsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalSnapshotInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectSnapshotsInfo) contextValidateArchivalSnapshotsInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ArchivalSnapshotsInfo); i++ {

		if m.ArchivalSnapshotsInfo[i] != nil {

			if swag.IsZero(m.ArchivalSnapshotsInfo[i]) { // not required
				return nil
			}

			if err := m.ArchivalSnapshotsInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("archivalSnapshotsInfo" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("archivalSnapshotsInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectSnapshotsInfo) contextValidateLocalSnapshotInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalSnapshotInfo != nil {

		if swag.IsZero(m.LocalSnapshotInfo) { // not required
			return nil
		}

		if err := m.LocalSnapshotInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localSnapshotInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localSnapshotInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectSnapshotsInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectSnapshotsInfo) UnmarshalBinary(b []byte) error {
	var res ObjectSnapshotsInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
