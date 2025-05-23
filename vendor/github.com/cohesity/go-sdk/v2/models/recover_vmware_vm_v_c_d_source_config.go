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

// RecoverVmwareVMVCDSourceConfig Recover VMware VMs vCloudDirector Source Config.
//
// Specifies the new destination Source configuration where the VMs will be recovered for vCloudDirector sources.
//
// swagger:model RecoverVmwareVmVCDSourceConfig
type RecoverVmwareVMVCDSourceConfig struct {

	// Specifies the datastore objects where the object's files should be recovered to. This should only be specified if storageProfile is not specified.
	Datastores []*RecoveryObjectIdentifier `json:"datastores"`

	// Specifies the VDC organization network which will be attached with recoverd VM.
	OrgVdcNetwork *OrgVDCNetwork `json:"orgVdcNetwork,omitempty"`

	// Specifies the networking configuration to be applied to the recovered VMs.
	NetworkConfig *RecoverVmwareVMNewSourceNetworkConfig `json:"networkConfig,omitempty"`

	// Specifies the id of the parent source to recover the VMs.
	// Required: true
	Source *RecoveryObjectIdentifier `json:"source"`

	// Specifies the storage profile to which the objects should be recovered. This should only be specified if datastores are not specified.
	StorageProfile *VcdStorageProfileParams `json:"storageProfile,omitempty"`

	// Specifies the vApp object where the recovered objects will be attached.
	VApp *RecoveryObjectIdentifier `json:"vApp,omitempty"`

	// Specifies the VDC object where the recovered objects will be attached.
	// Required: true
	Vdc *RecoveryObjectIdentifier `json:"vdc"`
}

// Validate validates this recover vmware Vm v c d source config
func (m *RecoverVmwareVMVCDSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrgVdcNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVApp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVdc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverVmwareVMVCDSourceConfig) validateDatastores(formats strfmt.Registry) error {
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

func (m *RecoverVmwareVMVCDSourceConfig) validateOrgVdcNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.OrgVdcNetwork) { // not required
		return nil
	}

	if m.OrgVdcNetwork != nil {
		if err := m.OrgVdcNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orgVdcNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orgVdcNetwork")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverVmwareVMVCDSourceConfig) validateNetworkConfig(formats strfmt.Registry) error {
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

func (m *RecoverVmwareVMVCDSourceConfig) validateSource(formats strfmt.Registry) error {

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

func (m *RecoverVmwareVMVCDSourceConfig) validateStorageProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageProfile) { // not required
		return nil
	}

	if m.StorageProfile != nil {
		if err := m.StorageProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageProfile")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverVmwareVMVCDSourceConfig) validateVApp(formats strfmt.Registry) error {
	if swag.IsZero(m.VApp) { // not required
		return nil
	}

	if m.VApp != nil {
		if err := m.VApp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vApp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vApp")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverVmwareVMVCDSourceConfig) validateVdc(formats strfmt.Registry) error {

	if err := validate.Required("vdc", "body", m.Vdc); err != nil {
		return err
	}

	if m.Vdc != nil {
		if err := m.Vdc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vdc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vdc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover vmware Vm v c d source config based on the context it is used
func (m *RecoverVmwareVMVCDSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatastores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrgVdcNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVdc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverVmwareVMVCDSourceConfig) contextValidateDatastores(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverVmwareVMVCDSourceConfig) contextValidateOrgVdcNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.OrgVdcNetwork != nil {

		if swag.IsZero(m.OrgVdcNetwork) { // not required
			return nil
		}

		if err := m.OrgVdcNetwork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orgVdcNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orgVdcNetwork")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverVmwareVMVCDSourceConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverVmwareVMVCDSourceConfig) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RecoverVmwareVMVCDSourceConfig) contextValidateStorageProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageProfile != nil {

		if swag.IsZero(m.StorageProfile) { // not required
			return nil
		}

		if err := m.StorageProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageProfile")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverVmwareVMVCDSourceConfig) contextValidateVApp(ctx context.Context, formats strfmt.Registry) error {

	if m.VApp != nil {

		if swag.IsZero(m.VApp) { // not required
			return nil
		}

		if err := m.VApp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vApp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vApp")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverVmwareVMVCDSourceConfig) contextValidateVdc(ctx context.Context, formats strfmt.Registry) error {

	if m.Vdc != nil {

		if err := m.Vdc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vdc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vdc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverVmwareVMVCDSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverVmwareVMVCDSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverVmwareVMVCDSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
