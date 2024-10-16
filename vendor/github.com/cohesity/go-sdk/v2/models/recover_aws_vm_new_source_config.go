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

// RecoverAwsVMNewSourceConfig Recover AWS VMs New Source Config.
//
// Specifies the new destination Source configuration where the VMs will be recovered.
//
// swagger:model RecoverAwsVmNewSourceConfig
type RecoverAwsVMNewSourceConfig struct {

	// Specifies the encryption configuration.
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// Specifies the pair of public and private key used to login into the VM
	KeyPair *RecoveryObjectIdentifier `json:"keyPair,omitempty"`

	// Specifies the networking configuration to be applied to the recovered VMs.
	// Required: true
	NetworkConfig *RecoverAwsVMNewSourceNetworkConfig `json:"networkConfig"`

	// Specifies the AWS region in which to deploy the VM.
	// Required: true
	Region *RecoveryObjectIdentifier `json:"region"`

	// Specifies the id of the parent source to recover the VMs.
	// Required: true
	Source *RecoveryObjectIdentifier `json:"source"`
}

// Validate validates this recover aws Vm new source config
func (m *RecoverAwsVMNewSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryptionConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyPair(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAwsVMNewSourceConfig) validateEncryptionConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionConfig) { // not required
		return nil
	}

	if m.EncryptionConfig != nil {
		if err := m.EncryptionConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAwsVMNewSourceConfig) validateKeyPair(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyPair) { // not required
		return nil
	}

	if m.KeyPair != nil {
		if err := m.KeyPair.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyPair")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyPair")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAwsVMNewSourceConfig) validateNetworkConfig(formats strfmt.Registry) error {

	if err := validate.Required("networkConfig", "body", m.NetworkConfig); err != nil {
		return err
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

func (m *RecoverAwsVMNewSourceConfig) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAwsVMNewSourceConfig) validateSource(formats strfmt.Registry) error {

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

// ContextValidate validate this recover aws Vm new source config based on the context it is used
func (m *RecoverAwsVMNewSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEncryptionConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyPair(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverAwsVMNewSourceConfig) contextValidateEncryptionConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.EncryptionConfig != nil {

		if swag.IsZero(m.EncryptionConfig) { // not required
			return nil
		}

		if err := m.EncryptionConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionConfig")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAwsVMNewSourceConfig) contextValidateKeyPair(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyPair != nil {

		if swag.IsZero(m.KeyPair) { // not required
			return nil
		}

		if err := m.KeyPair.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyPair")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyPair")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAwsVMNewSourceConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkConfig != nil {

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

func (m *RecoverAwsVMNewSourceConfig) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {

		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverAwsVMNewSourceConfig) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *RecoverAwsVMNewSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverAwsVMNewSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverAwsVMNewSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
