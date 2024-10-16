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

// RecoverExchangeAppParams Recover Exchange App params.
//
// Specifies the parameters to recover Sql databases.
//
// swagger:model RecoverExchangeAppParams
type RecoverExchangeAppParams struct {

	// Specifies the environment of the recovery target. The corresponding params below must be filled out.
	// Required: true
	// Enum: ["kExchange"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// Specifies the params for recovering to an Exchange host.
	ExchangeTargetParams *ExchangeTargetParamsForRecoverExchangeApp `json:"exchangeTargetParams,omitempty"`
}

// Validate validates this recover exchange app params
func (m *RecoverExchangeAppParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExchangeTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var recoverExchangeAppParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kExchange"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverExchangeAppParamsTypeTargetEnvironmentPropEnum = append(recoverExchangeAppParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// RecoverExchangeAppParamsTargetEnvironmentKExchange captures enum value "kExchange"
	RecoverExchangeAppParamsTargetEnvironmentKExchange string = "kExchange"
)

// prop value enum
func (m *RecoverExchangeAppParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverExchangeAppParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverExchangeAppParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *RecoverExchangeAppParams) validateExchangeTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ExchangeTargetParams) { // not required
		return nil
	}

	if m.ExchangeTargetParams != nil {
		if err := m.ExchangeTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exchangeTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exchangeTargetParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover exchange app params based on the context it is used
func (m *RecoverExchangeAppParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExchangeTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverExchangeAppParams) contextValidateExchangeTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ExchangeTargetParams != nil {

		if swag.IsZero(m.ExchangeTargetParams) { // not required
			return nil
		}

		if err := m.ExchangeTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exchangeTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exchangeTargetParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverExchangeAppParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverExchangeAppParams) UnmarshalBinary(b []byte) error {
	var res RecoverExchangeAppParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
