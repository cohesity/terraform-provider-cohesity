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

// IsilonEnvParams Isilon Environment Parameters.
//
// Message to capture Ision-spcific environment parameters.
//
// swagger:model IsilonEnvParams
type IsilonEnvParams struct {

	// Mapping from access zone name to configuration.
	ZoneConfigMap map[string]IsilonEnvParamsZoneConfig `json:"zoneConfigMap,omitempty"`
}

// Validate validates this isilon env params
func (m *IsilonEnvParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateZoneConfigMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsilonEnvParams) validateZoneConfigMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ZoneConfigMap) { // not required
		return nil
	}

	for k := range m.ZoneConfigMap {

		if err := validate.Required("zoneConfigMap"+"."+k, "body", m.ZoneConfigMap[k]); err != nil {
			return err
		}
		if val, ok := m.ZoneConfigMap[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zoneConfigMap" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zoneConfigMap" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this isilon env params based on the context it is used
func (m *IsilonEnvParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateZoneConfigMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsilonEnvParams) contextValidateZoneConfigMap(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ZoneConfigMap {

		if val, ok := m.ZoneConfigMap[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IsilonEnvParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IsilonEnvParams) UnmarshalBinary(b []byte) error {
	var res IsilonEnvParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
