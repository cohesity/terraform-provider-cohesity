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

// ObjectProtectionInfo Specifies the object info on cluster.
//
// swagger:model ObjectProtectionInfo
type ObjectProtectionInfo struct {

	// Specifies the object id.
	ObjectID *int64 `json:"objectId,omitempty"`

	// Specifies the source id.
	SourceID *int64 `json:"sourceId,omitempty"`

	// Specifies the view id for the object.
	ViewID *int64 `json:"viewId,omitempty"`

	// Specifies the region id where this object belongs to.
	RegionID *string `json:"regionId,omitempty"`

	// Specifies the cluster id where this object belongs to.
	ClusterID *int64 `json:"clusterId,omitempty"`

	// Specifies the cluster incarnation id where this object belongs to.
	ClusterIncarnationID *int64 `json:"clusterIncarnationId,omitempty"`

	// List of Tenants the object belongs to.
	TenantIds []string `json:"tenantIds"`

	// Specifies whether the object is deleted. Deleted objects can't be protected but can be recovered or unprotected.
	IsDeleted *bool `json:"isDeleted,omitempty"`

	// Specifies a list of protection groups protecting this object.
	ProtectionGroups []*ObjectProtectionGroupSummary `json:"protectionGroups"`

	// Specifies a list of object protections.
	ObjectBackupConfiguration []*ProtectionSummary `json:"objectBackupConfiguration"`

	// Specifies the status of the object's last protection run.
	// Enum: ["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold","Paused"]
	LastRunStatus *string `json:"lastRunStatus,omitempty"`
}

// Validate validates this object protection info
func (m *ObjectProtectionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtectionGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectBackupConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastRunStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectProtectionInfo) validateProtectionGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ProtectionGroups); i++ {
		if swag.IsZero(m.ProtectionGroups[i]) { // not required
			continue
		}

		if m.ProtectionGroups[i] != nil {
			if err := m.ProtectionGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectProtectionInfo) validateObjectBackupConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectBackupConfiguration) { // not required
		return nil
	}

	for i := 0; i < len(m.ObjectBackupConfiguration); i++ {
		if swag.IsZero(m.ObjectBackupConfiguration[i]) { // not required
			continue
		}

		if m.ObjectBackupConfiguration[i] != nil {
			if err := m.ObjectBackupConfiguration[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectBackupConfiguration" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectBackupConfiguration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var objectProtectionInfoTypeLastRunStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","LegalHold","Paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectProtectionInfoTypeLastRunStatusPropEnum = append(objectProtectionInfoTypeLastRunStatusPropEnum, v)
	}
}

const (

	// ObjectProtectionInfoLastRunStatusAccepted captures enum value "Accepted"
	ObjectProtectionInfoLastRunStatusAccepted string = "Accepted"

	// ObjectProtectionInfoLastRunStatusRunning captures enum value "Running"
	ObjectProtectionInfoLastRunStatusRunning string = "Running"

	// ObjectProtectionInfoLastRunStatusCanceled captures enum value "Canceled"
	ObjectProtectionInfoLastRunStatusCanceled string = "Canceled"

	// ObjectProtectionInfoLastRunStatusCanceling captures enum value "Canceling"
	ObjectProtectionInfoLastRunStatusCanceling string = "Canceling"

	// ObjectProtectionInfoLastRunStatusFailed captures enum value "Failed"
	ObjectProtectionInfoLastRunStatusFailed string = "Failed"

	// ObjectProtectionInfoLastRunStatusMissed captures enum value "Missed"
	ObjectProtectionInfoLastRunStatusMissed string = "Missed"

	// ObjectProtectionInfoLastRunStatusSucceeded captures enum value "Succeeded"
	ObjectProtectionInfoLastRunStatusSucceeded string = "Succeeded"

	// ObjectProtectionInfoLastRunStatusSucceededWithWarning captures enum value "SucceededWithWarning"
	ObjectProtectionInfoLastRunStatusSucceededWithWarning string = "SucceededWithWarning"

	// ObjectProtectionInfoLastRunStatusOnHold captures enum value "OnHold"
	ObjectProtectionInfoLastRunStatusOnHold string = "OnHold"

	// ObjectProtectionInfoLastRunStatusFinalizing captures enum value "Finalizing"
	ObjectProtectionInfoLastRunStatusFinalizing string = "Finalizing"

	// ObjectProtectionInfoLastRunStatusSkipped captures enum value "Skipped"
	ObjectProtectionInfoLastRunStatusSkipped string = "Skipped"

	// ObjectProtectionInfoLastRunStatusLegalHold captures enum value "LegalHold"
	ObjectProtectionInfoLastRunStatusLegalHold string = "LegalHold"

	// ObjectProtectionInfoLastRunStatusPaused captures enum value "Paused"
	ObjectProtectionInfoLastRunStatusPaused string = "Paused"
)

// prop value enum
func (m *ObjectProtectionInfo) validateLastRunStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectProtectionInfoTypeLastRunStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectProtectionInfo) validateLastRunStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.LastRunStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateLastRunStatusEnum("lastRunStatus", "body", *m.LastRunStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this object protection info based on the context it is used
func (m *ObjectProtectionInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtectionGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjectBackupConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectProtectionInfo) contextValidateProtectionGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProtectionGroups); i++ {

		if m.ProtectionGroups[i] != nil {

			if swag.IsZero(m.ProtectionGroups[i]) { // not required
				return nil
			}

			if err := m.ProtectionGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectProtectionInfo) contextValidateObjectBackupConfiguration(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ObjectBackupConfiguration); i++ {

		if m.ObjectBackupConfiguration[i] != nil {

			if swag.IsZero(m.ObjectBackupConfiguration[i]) { // not required
				return nil
			}

			if err := m.ObjectBackupConfiguration[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectBackupConfiguration" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectBackupConfiguration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectProtectionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectProtectionInfo) UnmarshalBinary(b []byte) error {
	var res ObjectProtectionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
