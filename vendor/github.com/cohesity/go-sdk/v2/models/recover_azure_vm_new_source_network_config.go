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

// RecoverAzureVMNewSourceNetworkConfig Recover Azure VMs New Source Network configuration.
//
// Specifies the network config parameters to be applied for Azure VMs if recovering to new Source.
//
// swagger:model RecoverAzureVmNewSourceNetworkConfig
type RecoverAzureVMNewSourceNetworkConfig struct {

	// Specifies id of the resource group for the selected virtual network.
	NetworkResourceGroup *RecoveryObjectIdentifier `json:"networkResourceGroup,omitempty"`

	// Specifies the subnet within the above virtual network.
	// Required: true
	Subnet *RecoveryObjectIdentifier `json:"subnet"`

	// Specifies the Virtual Network.
	// Required: true
	VirtualNetwork *RecoveryObjectIdentifier `json:"virtualNetwork"`
}

// Validate validates this recover azure Vm new source network config
func (m *RecoverAzureVMNewSourceNetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAzureVMNewSourceNetworkConfig) validateNetworkResourceGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkResourceGroup) { // not required
		return nil
	}

	if m.NetworkResourceGroup != nil {
		if err := m.NetworkResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkResourceGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkResourceGroup")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAzureVMNewSourceNetworkConfig) validateSubnet(formats strfmt.Registry) error {

	if err := validate.Required("subnet", "body", m.Subnet); err != nil {
		return err
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAzureVMNewSourceNetworkConfig) validateVirtualNetwork(formats strfmt.Registry) error {

	if err := validate.Required("virtualNetwork", "body", m.VirtualNetwork); err != nil {
		return err
	}

	if m.VirtualNetwork != nil {
		if err := m.VirtualNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualNetwork")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover azure Vm new source network config based on the context it is used
func (m *RecoverAzureVMNewSourceNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworkResourceGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAzureVMNewSourceNetworkConfig) contextValidateNetworkResourceGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkResourceGroup != nil {

		if swag.IsZero(m.NetworkResourceGroup) { // not required
			return nil
		}

		if err := m.NetworkResourceGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkResourceGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkResourceGroup")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAzureVMNewSourceNetworkConfig) contextValidateSubnet(ctx context.Context, formats strfmt.Registry) error {

	if m.Subnet != nil {

		if err := m.Subnet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAzureVMNewSourceNetworkConfig) contextValidateVirtualNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualNetwork != nil {

		if err := m.VirtualNetwork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualNetwork")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverAzureVMNewSourceNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverAzureVMNewSourceNetworkConfig) UnmarshalBinary(b []byte) error {
	var res RecoverAzureVMNewSourceNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
