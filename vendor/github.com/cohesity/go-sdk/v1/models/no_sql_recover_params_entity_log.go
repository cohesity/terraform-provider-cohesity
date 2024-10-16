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

// NoSQLRecoverParamsEntityLog no Sql recover params entity log
//
// swagger:model NoSqlRecoverParams_EntityLog
type NoSQLRecoverParamsEntityLog struct {

	// Entity for a leaf level entity.
	Entity *EntityProto `json:"entity,omitempty"`

	// List of log file and time range to applied for hydrated
	// backup or for recovery. Each data event has a path of log file and the
	// valid sequencer range within that log file.
	LogDataVec []*NoSQLLogData `json:"logDataVec"`
}

// Validate validates this no Sql recover params entity log
func (m *NoSQLRecoverParamsEntityLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogDataVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NoSQLRecoverParamsEntityLog) validateEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *NoSQLRecoverParamsEntityLog) validateLogDataVec(formats strfmt.Registry) error {
	if swag.IsZero(m.LogDataVec) { // not required
		return nil
	}

	for i := 0; i < len(m.LogDataVec); i++ {
		if swag.IsZero(m.LogDataVec[i]) { // not required
			continue
		}

		if m.LogDataVec[i] != nil {
			if err := m.LogDataVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logDataVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logDataVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this no Sql recover params entity log based on the context it is used
func (m *NoSQLRecoverParamsEntityLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogDataVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NoSQLRecoverParamsEntityLog) contextValidateEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.Entity != nil {

		if swag.IsZero(m.Entity) { // not required
			return nil
		}

		if err := m.Entity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *NoSQLRecoverParamsEntityLog) contextValidateLogDataVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LogDataVec); i++ {

		if m.LogDataVec[i] != nil {

			if swag.IsZero(m.LogDataVec[i]) { // not required
				return nil
			}

			if err := m.LogDataVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logDataVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logDataVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NoSQLRecoverParamsEntityLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NoSQLRecoverParamsEntityLog) UnmarshalBinary(b []byte) error {
	var res NoSQLRecoverParamsEntityLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
