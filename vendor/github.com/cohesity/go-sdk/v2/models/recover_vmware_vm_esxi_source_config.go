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

// RecoverVmwareVMEsxiSourceConfig Recover VMware VMs ESXi Source Config.
//
// Specifies the new destination Source configuration where the VMs will be recovered for ESXi sources.
//
// swagger:model RecoverVmwareVmEsxiSourceConfig
type RecoverVmwareVMEsxiSourceConfig struct {

	// Specifies the datastore objects where the object's files should be recovered to.
	Datastores []*RecoveryObjectIdentifier `json:"datastores"`

	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig *RecoverVmwareVMNewSourceNetworkConfig `json:"networkConfig,omitempty"`

	// Specifies the resource pool object where the recovered objects will be attached.
	// Required: true
	ResourcePool *RecoveryObjectIdentifier `json:"resourcePool"`

	// Specifies the id of the parent source to recover the VMs.
	// Required: true
	Source *RecoveryObjectIdentifier `json:"source"`

	// Folder where the VMs should be created.
	VMFolder *RecoveryObjectIdentifier `json:"vmFolder,omitempty"`
}

// Validate validates this recover vmware Vm esxi source config
func (m *RecoverVmwareVMEsxiSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcePool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMFolder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverVmwareVMEsxiSourceConfig) validateDatastores(formats strfmt.Registry) error {
	if swag.IsZero(m.Datastores) { // not required
		return nil
	}

	for i := 0; i < len(m.Datastores); i++ {
		if swag.IsZero(m.Datastores[i]) { // not required
			continue
		}

		if m.Datastores[i] != nil {
			if err := m.Datastores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datastores" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datastores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoverVmwareVMEsxiSourceConfig) validateNetworkConfig(formats strfmt.Registry) error {
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

func (m *RecoverVmwareVMEsxiSourceConfig) validateResourcePool(formats strfmt.Registry) error {

	if err := validate.Required("resourcePool", "body", m.ResourcePool); err != nil {
		return err
	}

	if m.ResourcePool != nil {
		if err := m.ResourcePool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourcePool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourcePool")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverVmwareVMEsxiSourceConfig) validateSource(formats strfmt.Registry) error {

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

func (m *RecoverVmwareVMEsxiSourceConfig) validateVMFolder(formats strfmt.Registry) error {
	if swag.IsZero(m.VMFolder) { // not required
		return nil
	}

	if m.VMFolder != nil {
		if err := m.VMFolder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmFolder")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmFolder")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover vmware Vm esxi source config based on the context it is used
func (m *RecoverVmwareVMEsxiSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatastores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourcePool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMFolder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverVmwareVMEsxiSourceConfig) contextValidateDatastores(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Datastores); i++ {

		if m.Datastores[i] != nil {

			if swag.IsZero(m.Datastores[i]) { // not required
				return nil
			}

			if err := m.Datastores[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datastores" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datastores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoverVmwareVMEsxiSourceConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverVmwareVMEsxiSourceConfig) contextValidateResourcePool(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourcePool != nil {

		if err := m.ResourcePool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourcePool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourcePool")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverVmwareVMEsxiSourceConfig) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverVmwareVMEsxiSourceConfig) contextValidateVMFolder(ctx context.Context, formats strfmt.Registry) error {

	if m.VMFolder != nil {

		if swag.IsZero(m.VMFolder) { // not required
			return nil
		}

		if err := m.VMFolder.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmFolder")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmFolder")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverVmwareVMEsxiSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverVmwareVMEsxiSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverVmwareVMEsxiSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
