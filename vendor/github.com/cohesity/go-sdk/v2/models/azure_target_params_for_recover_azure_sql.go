// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AzureTargetParamsForRecoverAzureSQL Specifies the recovery target params for Azure SQL target config.
//
// swagger:model AzureTargetParamsForRecoverAzureSql
type AzureTargetParamsForRecoverAzureSQL struct {

	// Specifies the parameter whether the recovery should be performed to a new or an existing target.
	// Required: true
	RecoverToNewSource *bool `json:"recoverToNewSource"`

	// Specifies the new destination Source configuration parameters where the Azure SQL instances will be recovered. This is mandatory if recoverToNewSource is set to true.
	NewSourceConfig *RecoverAzureSQLNewSourceConfig `json:"newSourceConfig,omitempty"`
}

// Validate validates this azure target params for recover azure Sql
func (m *AzureTargetParamsForRecoverAzureSQL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoverToNewSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewSourceConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureTargetParamsForRecoverAzureSQL) validateRecoverToNewSource(formats strfmt.Registry) error {

	if err := validate.Required("recoverToNewSource", "body", m.RecoverToNewSource); err != nil {
		return err
	}

	return nil
}

func (m *AzureTargetParamsForRecoverAzureSQL) validateNewSourceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NewSourceConfig) { // not required
		return nil
	}

	if m.NewSourceConfig != nil {
		if err := m.NewSourceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newSourceConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this azure target params for recover azure Sql based on the context it is used
func (m *AzureTargetParamsForRecoverAzureSQL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNewSourceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureTargetParamsForRecoverAzureSQL) contextValidateNewSourceConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NewSourceConfig != nil {

		if swag.IsZero(m.NewSourceConfig) { // not required
			return nil
		}

		if err := m.NewSourceConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newSourceConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureTargetParamsForRecoverAzureSQL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureTargetParamsForRecoverAzureSQL) UnmarshalBinary(b []byte) error {
	var res AzureTargetParamsForRecoverAzureSQL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
