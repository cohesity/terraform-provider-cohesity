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

// RecoverAcropolisVMNewSourceConfig Recover Acropolis VMs New Source Config.
//
// Specifies the new destination Source configuration where the VMs will be recovered.
//
// swagger:model RecoverAcropolisVmNewSourceConfig
type RecoverAcropolisVMNewSourceConfig struct {

	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig *RecoverAcropolisVMNewSourceNetworkConfig `json:"networkConfig,omitempty"`

	// Specifies the id of the parent source to recover the VMs.
	// Required: true
	Source *RecoveryObjectIdentifier `json:"source"`

	// A storage container where the VM's files should be restored to.
	StorageContainer *RecoveryObjectIdentifier `json:"storageContainer,omitempty"`
}

// Validate validates this recover acropolis Vm new source config
func (m *RecoverAcropolisVMNewSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageContainer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAcropolisVMNewSourceConfig) validateNetworkConfig(formats strfmt.Registry) error {
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

func (m *RecoverAcropolisVMNewSourceConfig) validateSource(formats strfmt.Registry) error {

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

func (m *RecoverAcropolisVMNewSourceConfig) validateStorageContainer(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageContainer) { // not required
		return nil
	}

	if m.StorageContainer != nil {
		if err := m.StorageContainer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageContainer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageContainer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover acropolis Vm new source config based on the context it is used
func (m *RecoverAcropolisVMNewSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageContainer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAcropolisVMNewSourceConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverAcropolisVMNewSourceConfig) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverAcropolisVMNewSourceConfig) contextValidateStorageContainer(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageContainer != nil {

		if swag.IsZero(m.StorageContainer) { // not required
			return nil
		}

		if err := m.StorageContainer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageContainer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageContainer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverAcropolisVMNewSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverAcropolisVMNewSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverAcropolisVMNewSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
