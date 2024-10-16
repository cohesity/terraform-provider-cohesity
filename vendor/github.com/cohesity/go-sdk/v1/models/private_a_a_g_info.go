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
)

// PrivateAAGInfo Contains information about an Always on Availability Group (aka AAG).
//
// See http://goo.gl/3NvMMW for details.
//
// swagger:model PrivateAAGInfo
type PrivateAAGInfo struct {

	// The backup preference for this AAG that was found on the SQL server.
	AutomatedBackupPreference *int32 `json:"automatedBackupPreference,omitempty"`

	// Common information for all the databases that belong to this AAG.
	DbInfoVec []*AAGDbInfo `json:"dbInfoVec"`

	// Error if any associated with this AAG. This is typically returned by the
	// Agent during AAG population. The AAGs in the entity hierarchy are not
	// expected to have this field set.
	Error *PrivateErrorProto `json:"error,omitempty"`

	// The unique identifier of the availability group.
	GroupID *string `json:"groupId,omitempty"`

	// Host type of the AAG replicas.
	HostType *int32 `json:"hostType,omitempty"`

	// The name of this AAG.
	Name *string `json:"name,omitempty"`

	// All the known replicas for the AAG.
	//
	// If returned by the agent, this will only contain all the replicas with
	// respect to the local replica that is querying for this information.
	//
	// In the entity hierarchy view, this should contain the union of the states
	// seen at all the replicas.
	ReplicaVec []*PrivateAAGReplicaInfo `json:"replicaVec"`
}

// Validate validates this private a a g info
func (m *PrivateAAGInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDbInfoVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateAAGInfo) validateDbInfoVec(formats strfmt.Registry) error {
	if swag.IsZero(m.DbInfoVec) { // not required
		return nil
	}

	for i := 0; i < len(m.DbInfoVec); i++ {
		if swag.IsZero(m.DbInfoVec[i]) { // not required
			continue
		}

		if m.DbInfoVec[i] != nil {
			if err := m.DbInfoVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dbInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dbInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrivateAAGInfo) validateError(formats strfmt.Registry) error {
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

func (m *PrivateAAGInfo) validateReplicaVec(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicaVec) { // not required
		return nil
	}

	for i := 0; i < len(m.ReplicaVec); i++ {
		if swag.IsZero(m.ReplicaVec[i]) { // not required
			continue
		}

		if m.ReplicaVec[i] != nil {
			if err := m.ReplicaVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicaVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicaVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this private a a g info based on the context it is used
func (m *PrivateAAGInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDbInfoVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplicaVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateAAGInfo) contextValidateDbInfoVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DbInfoVec); i++ {

		if m.DbInfoVec[i] != nil {

			if swag.IsZero(m.DbInfoVec[i]) { // not required
				return nil
			}

			if err := m.DbInfoVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dbInfoVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dbInfoVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrivateAAGInfo) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PrivateAAGInfo) contextValidateReplicaVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReplicaVec); i++ {

		if m.ReplicaVec[i] != nil {

			if swag.IsZero(m.ReplicaVec[i]) { // not required
				return nil
			}

			if err := m.ReplicaVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicaVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicaVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrivateAAGInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateAAGInfo) UnmarshalBinary(b []byte) error {
	var res PrivateAAGInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
