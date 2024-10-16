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

// RecoverAcropolisVMOriginalSourceConfig Recover Acropolis VMs Original Source Config.
//
// Specifies the Source configuration if VM's are being recovered to Original Source.
//
// swagger:model RecoverAcropolisVmOriginalSourceConfig
type RecoverAcropolisVMOriginalSourceConfig struct {

	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig *RecoverAcropolisVMOriginalSourceNetworkConfig `json:"networkConfig,omitempty"`
}

// Validate validates this recover acropolis Vm original source config
func (m *RecoverAcropolisVMOriginalSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAcropolisVMOriginalSourceConfig) validateNetworkConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkConfig) { // not required
		return nil
	}

	if m.NetworkConfig != nil {
		if err := m.NetworkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover acropolis Vm original source config based on the context it is used
func (m *RecoverAcropolisVMOriginalSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAcropolisVMOriginalSourceConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkConfig != nil {

		if swag.IsZero(m.NetworkConfig) { // not required
			return nil
		}

		if err := m.NetworkConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverAcropolisVMOriginalSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverAcropolisVMOriginalSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverAcropolisVMOriginalSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
