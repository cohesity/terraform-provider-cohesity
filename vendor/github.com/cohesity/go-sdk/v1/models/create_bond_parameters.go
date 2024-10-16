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

// CreateBondParameters Create Bond Parameters.
//
// Specifies the parameters needed to create a bond.
//
// swagger:model CreateBondParameters
type CreateBondParameters struct {

	// Specifies the bonding mode to use for this bond. If not specified,
	// this value will default to 'kActiveBackup'.
	// 'kActiveBackup' indicates active backup bonding mode.
	// 'k802_3ad' indicates 802.3ad bonding mode.
	// Enum: ["kActiveBackup","k802_3ad"]
	BondingMode *string `json:"bondingMode,omitempty"`

	// Specifies a unique name to identify the bond being created.
	// Required: true
	Name *string `json:"name"`

	// Specifies the names of the secondaries of this bond.
	// Required: true
	Slaves []string `json:"slaves"`
}

// Validate validates this create bond parameters
func (m *CreateBondParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBondingMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlaves(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createBondParametersTypeBondingModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kActiveBackup","k802_3ad"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createBondParametersTypeBondingModePropEnum = append(createBondParametersTypeBondingModePropEnum, v)
	}
}

const (

	// CreateBondParametersBondingModeKActiveBackup captures enum value "kActiveBackup"
	CreateBondParametersBondingModeKActiveBackup string = "kActiveBackup"

	// CreateBondParametersBondingModeK8023ad captures enum value "k802_3ad"
	CreateBondParametersBondingModeK8023ad string = "k802_3ad"
)

// prop value enum
func (m *CreateBondParameters) validateBondingModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createBondParametersTypeBondingModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateBondParameters) validateBondingMode(formats strfmt.Registry) error {
	if swag.IsZero(m.BondingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateBondingModeEnum("bondingMode", "body", *m.BondingMode); err != nil {
		return err
	}

	return nil
}

func (m *CreateBondParameters) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateBondParameters) validateSlaves(formats strfmt.Registry) error {

	if err := validate.Required("slaves", "body", m.Slaves); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create bond parameters based on context it is used
func (m *CreateBondParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateBondParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBondParameters) UnmarshalBinary(b []byte) error {
	var res CreateBondParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
