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

// RecoverKvmVMNewSourceConfig Recover KVM VMs New Source Config.
//
// Specifies the new destination Source configuration where the VMs will be recovered.
//
// swagger:model RecoverKvmVmNewSourceConfig
type RecoverKvmVMNewSourceConfig struct {

	// Specifies the resource (KVMH host) to which the restored VM will be attached
	// Required: true
	Cluster *RecoveryObjectIdentifier `json:"cluster"`

	// Specifies the datacenter where the VM's files should be restored to.
	// Required: true
	DataCenter *RecoveryObjectIdentifier `json:"dataCenter"`

	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig *RecoverKvmVMNewSourceNetworkConfig `json:"networkConfig,omitempty"`

	// Specifies the id of the parent source to recover the VMs.
	// Required: true
	Source *RecoveryObjectIdentifier `json:"source"`

	// Specifies the Storage Domain where the VM's disk should be restored to.
	// Required: true
	StorageDomain *RecoveryObjectIdentifier `json:"storageDomain"`
}

// Validate validates this recover kvm Vm new source config
func (m *RecoverKvmVMNewSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataCenter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageDomain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverKvmVMNewSourceConfig) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverKvmVMNewSourceConfig) validateDataCenter(formats strfmt.Registry) error {

	if err := validate.Required("dataCenter", "body", m.DataCenter); err != nil {
		return err
	}

	if m.DataCenter != nil {
		if err := m.DataCenter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataCenter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataCenter")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverKvmVMNewSourceConfig) validateNetworkConfig(formats strfmt.Registry) error {
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

func (m *RecoverKvmVMNewSourceConfig) validateSource(formats strfmt.Registry) error {

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

func (m *RecoverKvmVMNewSourceConfig) validateStorageDomain(formats strfmt.Registry) error {

	if err := validate.Required("storageDomain", "body", m.StorageDomain); err != nil {
		return err
	}

	if m.StorageDomain != nil {
		if err := m.StorageDomain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageDomain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageDomain")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover kvm Vm new source config based on the context it is used
func (m *RecoverKvmVMNewSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataCenter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverKvmVMNewSourceConfig) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {

		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverKvmVMNewSourceConfig) contextValidateDataCenter(ctx context.Context, formats strfmt.Registry) error {

	if m.DataCenter != nil {

		if err := m.DataCenter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataCenter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataCenter")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverKvmVMNewSourceConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverKvmVMNewSourceConfig) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverKvmVMNewSourceConfig) contextValidateStorageDomain(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageDomain != nil {

		if err := m.StorageDomain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageDomain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageDomain")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverKvmVMNewSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverKvmVMNewSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverKvmVMNewSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
