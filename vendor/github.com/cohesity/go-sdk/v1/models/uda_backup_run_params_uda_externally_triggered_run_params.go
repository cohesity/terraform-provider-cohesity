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

// UdaBackupRunParamsUdaExternallyTriggeredRunParams uda backup run params uda externally triggered run params
//
// swagger:model UdaBackupRunParams_UdaExternallyTriggeredRunParams
type UdaBackupRunParamsUdaExternallyTriggeredRunParams struct {

	// Specifies a map of custom arguments to be supplied to the plugin.
	ArgsMap map[string]UdaCustomArgument `json:"argsMap,omitempty"`

	// Specifies the IP or FQDN of the source host where this
	// backup will run.
	ControlNode *string `json:"controlNode,omitempty"`

	// A unique identifier for this externally triggerd run,
	// to be sent to the plugin.
	EtRunID *string `json:"etRunId,omitempty"`
}

// Validate validates this uda backup run params uda externally triggered run params
func (m *UdaBackupRunParamsUdaExternallyTriggeredRunParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgsMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaBackupRunParamsUdaExternallyTriggeredRunParams) validateArgsMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ArgsMap) { // not required
		return nil
	}

	for k := range m.ArgsMap {

		if err := validate.Required("argsMap"+"."+k, "body", m.ArgsMap[k]); err != nil {
			return err
		}
		if val, ok := m.ArgsMap[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("argsMap" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("argsMap" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this uda backup run params uda externally triggered run params based on the context it is used
func (m *UdaBackupRunParamsUdaExternallyTriggeredRunParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArgsMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaBackupRunParamsUdaExternallyTriggeredRunParams) contextValidateArgsMap(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ArgsMap {

		if val, ok := m.ArgsMap[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UdaBackupRunParamsUdaExternallyTriggeredRunParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaBackupRunParamsUdaExternallyTriggeredRunParams) UnmarshalBinary(b []byte) error {
	var res UdaBackupRunParamsUdaExternallyTriggeredRunParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
