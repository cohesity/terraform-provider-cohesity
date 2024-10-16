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

// RestoreExchangeParams Params specific to restoring an Exchange application.
//
// swagger:model RestoreExchangeParams
type RestoreExchangeParams struct {

	// Only applicable when ExchangeRestoreType.Type = kDatabase.
	DatabaseOptions *RestoreExchangeParamsDatabaseOptions `json:"databaseOptions,omitempty"`

	// Restore type.
	Type *int32 `json:"type,omitempty"`

	// Only applicable when ExchangeRestoreType.Type=kView.
	ViewOptions *RestoreExchangeParamsViewOptions `json:"viewOptions,omitempty"`
}

// Validate validates this restore exchange params
func (m *RestoreExchangeParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreExchangeParams) validateDatabaseOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DatabaseOptions) { // not required
		return nil
	}

	if m.DatabaseOptions != nil {
		if err := m.DatabaseOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databaseOptions")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreExchangeParams) validateViewOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ViewOptions) { // not required
		return nil
	}

	if m.ViewOptions != nil {
		if err := m.ViewOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viewOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("viewOptions")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this restore exchange params based on the context it is used
func (m *RestoreExchangeParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatabaseOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreExchangeParams) contextValidateDatabaseOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.DatabaseOptions != nil {

		if swag.IsZero(m.DatabaseOptions) { // not required
			return nil
		}

		if err := m.DatabaseOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databaseOptions")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreExchangeParams) contextValidateViewOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.ViewOptions != nil {

		if swag.IsZero(m.ViewOptions) { // not required
			return nil
		}

		if err := m.ViewOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viewOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("viewOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreExchangeParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreExchangeParams) UnmarshalBinary(b []byte) error {
	var res RestoreExchangeParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
