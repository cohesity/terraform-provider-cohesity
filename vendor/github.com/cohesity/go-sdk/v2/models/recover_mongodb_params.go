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

// RecoverMongodbParams Recover MongoDB params.
//
// Specifies the parameters to recover MongoDB objects.
//
// swagger:model RecoverMongodbParams
type RecoverMongodbParams struct {
	CommonNoSQLRecoveryOptions

	// Specifies the local snapshot ids of the Objects to be recovered.
	// Required: true
	Snapshots []*RecoverMongodbSnapshotParams `json:"snapshots"`

	// A suffix that is to be applied to all recovered objects.
	Suffix *string `json:"suffix,omitempty"`

	// Specifies whether to recover User and roles at the time of recovery.
	RecoverUserRoles *bool `json:"recoverUserRoles,omitempty"`

	// Specifies whether to recover Zones/shard tags at the time of recovery.
	RecoverZonesTags *bool `json:"recoverZonesTags,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoverMongodbParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonNoSQLRecoveryOptions
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonNoSQLRecoveryOptions = aO0

	// AO1
	var dataAO1 struct {
		Snapshots []*RecoverMongodbSnapshotParams `json:"snapshots"`

		Suffix *string `json:"suffix,omitempty"`

		RecoverUserRoles *bool `json:"recoverUserRoles,omitempty"`

		RecoverZonesTags *bool `json:"recoverZonesTags,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Snapshots = dataAO1.Snapshots

	m.Suffix = dataAO1.Suffix

	m.RecoverUserRoles = dataAO1.RecoverUserRoles

	m.RecoverZonesTags = dataAO1.RecoverZonesTags

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoverMongodbParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonNoSQLRecoveryOptions)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Snapshots []*RecoverMongodbSnapshotParams `json:"snapshots"`

		Suffix *string `json:"suffix,omitempty"`

		RecoverUserRoles *bool `json:"recoverUserRoles,omitempty"`

		RecoverZonesTags *bool `json:"recoverZonesTags,omitempty"`
	}

	dataAO1.Snapshots = m.Snapshots

	dataAO1.Suffix = m.Suffix

	dataAO1.RecoverUserRoles = m.RecoverUserRoles

	dataAO1.RecoverZonesTags = m.RecoverZonesTags

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recover mongodb params
func (m *RecoverMongodbParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonNoSQLRecoveryOptions
	if err := m.CommonNoSQLRecoveryOptions.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverMongodbParams) validateSnapshots(formats strfmt.Registry) error {

	if err := validate.Required("snapshots", "body", m.Snapshots); err != nil {
		return err
	}

	for i := 0; i < len(m.Snapshots); i++ {
		if swag.IsZero(m.Snapshots[i]) { // not required
			continue
		}

		if m.Snapshots[i] != nil {
			if err := m.Snapshots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this recover mongodb params based on the context it is used
func (m *RecoverMongodbParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonNoSQLRecoveryOptions
	if err := m.CommonNoSQLRecoveryOptions.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverMongodbParams) contextValidateSnapshots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Snapshots); i++ {

		if m.Snapshots[i] != nil {

			if swag.IsZero(m.Snapshots[i]) { // not required
				return nil
			}

			if err := m.Snapshots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverMongodbParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverMongodbParams) UnmarshalBinary(b []byte) error {
	var res RecoverMongodbParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
