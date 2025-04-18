// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecoverKvmVMOriginalSourceNetworkConfig Recover KVM VMs Original Source Network configuration.
//
// Specifies the network config parameters to be applied for KVM VMs if recovering to original Source.
//
// swagger:model RecoverKvmVmOriginalSourceNetworkConfig
type RecoverKvmVMOriginalSourceNetworkConfig struct {

	// If this is set to true, then the network will be detached from the recovered VMs. All the other networking parameters set will be ignored if set to true. Default value is false.
	DetachNetwork *bool `json:"detachNetwork,omitempty"`
}

// Validate validates this recover kvm Vm original source network config
func (m *RecoverKvmVMOriginalSourceNetworkConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this recover kvm Vm original source network config based on context it is used
func (m *RecoverKvmVMOriginalSourceNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecoverKvmVMOriginalSourceNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverKvmVMOriginalSourceNetworkConfig) UnmarshalBinary(b []byte) error {
	var res RecoverKvmVMOriginalSourceNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
