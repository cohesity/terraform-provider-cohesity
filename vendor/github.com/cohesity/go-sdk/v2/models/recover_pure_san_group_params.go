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

// RecoverPureSanGroupParams Recover Pure SAN group Params.
//
// Specifies the parameters to recover Pure SAN group.
//
// swagger:model RecoverPureSanGroupParams
type RecoverPureSanGroupParams struct {

	// Specifies the environment of the recovery target. The corresponding target params must be filled out.
	// Required: true
	// Enum: ["kPure","kIbmFlashSystem"]
	TargetEnvironment *string `json:"targetEnvironment"`

	// Specifies the parameters of the Pure SAN group to recover to.
	PureTargetParams *RecoverPureGroupTargetParams `json:"pureTargetParams,omitempty"`
}

// Validate validates this recover pure san group params
func (m *RecoverPureSanGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePureTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var recoverPureSanGroupParamsTypeTargetEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kPure","kIbmFlashSystem"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverPureSanGroupParamsTypeTargetEnvironmentPropEnum = append(recoverPureSanGroupParamsTypeTargetEnvironmentPropEnum, v)
	}
}

const (

	// RecoverPureSanGroupParamsTargetEnvironmentKPure captures enum value "kPure"
	RecoverPureSanGroupParamsTargetEnvironmentKPure string = "kPure"

	// RecoverPureSanGroupParamsTargetEnvironmentKIbmFlashSystem captures enum value "kIbmFlashSystem"
	RecoverPureSanGroupParamsTargetEnvironmentKIbmFlashSystem string = "kIbmFlashSystem"
)

// prop value enum
func (m *RecoverPureSanGroupParams) validateTargetEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverPureSanGroupParamsTypeTargetEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverPureSanGroupParams) validateTargetEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironment", "body", m.TargetEnvironment); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetEnvironmentEnum("targetEnvironment", "body", *m.TargetEnvironment); err != nil {
		return err
	}

	return nil
}

func (m *RecoverPureSanGroupParams) validatePureTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.PureTargetParams) { // not required
		return nil
	}

	if m.PureTargetParams != nil {
		if err := m.PureTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pureTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pureTargetParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover pure san group params based on the context it is used
func (m *RecoverPureSanGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePureTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverPureSanGroupParams) contextValidatePureTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.PureTargetParams != nil {

		if swag.IsZero(m.PureTargetParams) { // not required
			return nil
		}

		if err := m.PureTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pureTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pureTargetParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverPureSanGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverPureSanGroupParams) UnmarshalBinary(b []byte) error {
	var res RecoverPureSanGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
