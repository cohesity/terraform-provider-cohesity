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

// RecoverHyperVVMStandaloneHostSourceConfig Recover HyperV VMs New Standalone Host Source Config.
//
// Specifies the new destination Source configuration where the VMs will be recovered.
//
// swagger:model RecoverHyperVVmStandaloneHostSourceConfig
type RecoverHyperVVMStandaloneHostSourceConfig struct {

	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig *RecoverHyperVVMNewSourceNetworkConfig `json:"networkConfig,omitempty"`

	// Specifies the id of the parent source to recover the VMs.
	// Required: true
	Source *RecoveryObjectIdentifier `json:"source"`

	// Specifies the datastore object where the VMs' files should be recovered to.
	// Required: true
	Volume *RecoveryObjectIdentifier `json:"volume"`
}

// Validate validates this recover hyper v Vm standalone host source config
func (m *RecoverHyperVVMStandaloneHostSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverHyperVVMStandaloneHostSourceConfig) validateNetworkConfig(formats strfmt.Registry) error {
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

func (m *RecoverHyperVVMStandaloneHostSourceConfig) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverHyperVVMStandaloneHostSourceConfig) validateVolume(formats strfmt.Registry) error {

	if err := validate.Required("volume", "body", m.Volume); err != nil {
		return err
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover hyper v Vm standalone host source config based on the context it is used
func (m *RecoverHyperVVMStandaloneHostSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverHyperVVMStandaloneHostSourceConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverHyperVVMStandaloneHostSourceConfig) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverHyperVVMStandaloneHostSourceConfig) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {

		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverHyperVVMStandaloneHostSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverHyperVVMStandaloneHostSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverHyperVVMStandaloneHostSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
