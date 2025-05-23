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

// RecoverExchangeAppSnapshotParams Recover Exchange App Snapshot Params.
//
// Specifies the snapshot parameters to recover Exchange databases.
//
// swagger:model RecoverExchangeAppSnapshotParams
type RecoverExchangeAppSnapshotParams struct {
	CommonRecoverObjectSnapshotParams

	// Specifies the list of params for all the databases which have to be recovered.
	AppObjects []*ExchangeRecoverDatabaseParams `json:"appObjects"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoverExchangeAppSnapshotParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonRecoverObjectSnapshotParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonRecoverObjectSnapshotParams = aO0

	// AO1
	var dataAO1 struct {
		AppObjects []*ExchangeRecoverDatabaseParams `json:"appObjects"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AppObjects = dataAO1.AppObjects

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoverExchangeAppSnapshotParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonRecoverObjectSnapshotParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AppObjects []*ExchangeRecoverDatabaseParams `json:"appObjects"`
	}

	dataAO1.AppObjects = m.AppObjects

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recover exchange app snapshot params
func (m *RecoverExchangeAppSnapshotParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonRecoverObjectSnapshotParams
	if err := m.CommonRecoverObjectSnapshotParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverExchangeAppSnapshotParams) validateAppObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.AppObjects) { // not required
		return nil
	}

	for i := 0; i < len(m.AppObjects); i++ {
		if swag.IsZero(m.AppObjects[i]) { // not required
			continue
		}

		if m.AppObjects[i] != nil {
			if err := m.AppObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this recover exchange app snapshot params based on the context it is used
func (m *RecoverExchangeAppSnapshotParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonRecoverObjectSnapshotParams
	if err := m.CommonRecoverObjectSnapshotParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverExchangeAppSnapshotParams) contextValidateAppObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppObjects); i++ {

		if m.AppObjects[i] != nil {

			if swag.IsZero(m.AppObjects[i]) { // not required
				return nil
			}

			if err := m.AppObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appObjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverExchangeAppSnapshotParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverExchangeAppSnapshotParams) UnmarshalBinary(b []byte) error {
	var res RecoverExchangeAppSnapshotParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
