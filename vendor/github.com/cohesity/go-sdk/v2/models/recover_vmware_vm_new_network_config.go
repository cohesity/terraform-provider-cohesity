// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecoverVmwareVMNewNetworkConfig Recover VMware VMs New Network configuration.
//
// Specifies the new network config parameters to be applied to VMware VMs.
//
// swagger:model RecoverVmwareVmNewNetworkConfig
type RecoverVmwareVMNewNetworkConfig struct {

	// Specifies whether the attached network should be left in disabled state. Default is false
	DisableNetwork *bool `json:"disableNetwork,omitempty"`

	// If this is true and we are attaching to a new network entity, then the VM's MAC address will be preserved on the new network. Default value is false.
	PreserveMacAddress *bool `json:"preserveMacAddress,omitempty"`

	// Specifies the target network mapping for each VM's network adapter.
	// Min Items: 0
	Mappings []*RecoverVmwareVMNewNetworkConfigMapping `json:"mappings"`

	// Specifies the network port group (i.e, either a standard switch port group or a distributed port group) that will attached to the recovered Object. This parameter is mandatory if detach network is specified as false.
	NetworkPortGroup *RecoveryObjectIdentifier `json:"networkPortGroup,omitempty"`
}

// Validate validates this recover vmware Vm new network config
func (m *RecoverVmwareVMNewNetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkPortGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverVmwareVMNewNetworkConfig) validateMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.Mappings) { // not required
		return nil
	}

	iMappingsSize := int64(len(m.Mappings))

	if err := validate.MinItems("mappings", "body", iMappingsSize, 0); err != nil {
		return err
	}

	for i := 0; i < len(m.Mappings); i++ {
		if swag.IsZero(m.Mappings[i]) { // not required
			continue
		}

		if m.Mappings[i] != nil {
			if err := m.Mappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoverVmwareVMNewNetworkConfig) validateNetworkPortGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkPortGroup) { // not required
		return nil
	}

	if m.NetworkPortGroup != nil {
		if err := m.NetworkPortGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPortGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkPortGroup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover vmware Vm new network config based on the context it is used
func (m *RecoverVmwareVMNewNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkPortGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverVmwareVMNewNetworkConfig) contextValidateMappings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Mappings); i++ {

		if m.Mappings[i] != nil {

			if swag.IsZero(m.Mappings[i]) { // not required
				return nil
			}

			if err := m.Mappings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoverVmwareVMNewNetworkConfig) contextValidateNetworkPortGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkPortGroup != nil {

		if swag.IsZero(m.NetworkPortGroup) { // not required
			return nil
		}

		if err := m.NetworkPortGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPortGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkPortGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverVmwareVMNewNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverVmwareVMNewNetworkConfig) UnmarshalBinary(b []byte) error {
	var res RecoverVmwareVMNewNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
