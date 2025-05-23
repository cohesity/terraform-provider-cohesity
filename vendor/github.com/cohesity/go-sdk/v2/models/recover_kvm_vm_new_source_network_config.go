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

// RecoverKvmVMNewSourceNetworkConfig Recover KVM VMs New Source Network configuration.
//
// Specifies the network config parameters to be applied for KVM VMs if recovering to new Source.
//
// swagger:model RecoverKvmVmNewSourceNetworkConfig
type RecoverKvmVMNewSourceNetworkConfig struct {

	// If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false.
	DetachNetwork *bool `json:"detachNetwork,omitempty"`

	// Specifies the new network configuration of the Kvm recovery.
	NewNetworkConfig *RecoverKvmVMNewNetworkConfig `json:"newNetworkConfig,omitempty"`
}

// Validate validates this recover kvm Vm new source network config
func (m *RecoverKvmVMNewSourceNetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverKvmVMNewSourceNetworkConfig) validateNewNetworkConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NewNetworkConfig) { // not required
		return nil
	}

	if m.NewNetworkConfig != nil {
		if err := m.NewNetworkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newNetworkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newNetworkConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover kvm Vm new source network config based on the context it is used
func (m *RecoverKvmVMNewSourceNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNewNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverKvmVMNewSourceNetworkConfig) contextValidateNewNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NewNetworkConfig != nil {

		if swag.IsZero(m.NewNetworkConfig) { // not required
			return nil
		}

		if err := m.NewNetworkConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newNetworkConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newNetworkConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverKvmVMNewSourceNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverKvmVMNewSourceNetworkConfig) UnmarshalBinary(b []byte) error {
	var res RecoverKvmVMNewSourceNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
