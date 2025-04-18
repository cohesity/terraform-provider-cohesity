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

// NfsConfig Specifies the NFS config settings for this View.
//
// swagger:model NfsConfig
type NfsConfig struct {

	// If set, it enables discovery of view for NFS.
	EnableNfsViewDiscovery *bool `json:"enableNfsViewDiscovery,omitempty"`

	// If set, it enables NFS UNIX Authentication
	EnableNfsUnixAuthentication *bool `json:"enableNfsUnixAuthentication,omitempty"`

	// If set, it enables NFS Kerberos Authentication
	EnableNfsKerberosAuthentication *bool `json:"enableNfsKerberosAuthentication,omitempty"`

	// If set, it enables NFS Kerberos Integrity
	EnableNfsKerberosIntegrity *bool `json:"enableNfsKerberosIntegrity,omitempty"`

	// If set, it enables NFS Kerberos Privacy
	EnableNfsKerberosPrivacy *bool `json:"enableNfsKerberosPrivacy,omitempty"`

	// If set, it enables NFS weak cache consistency.
	EnableNfsWcc *bool `json:"enableNfsWcc,omitempty"`

	// Specifies the NFS all squash config.
	NfsAllSquash *NfsSquash `json:"nfsAllSquash,omitempty"`

	// Specifies the NFS root permission config of the view file system.
	NfsRootPermissions *NfsRootPermissions `json:"nfsRootPermissions,omitempty"`

	// Specifies the NFS root squash config.
	NfsRootSquash *NfsSquash `json:"nfsRootSquash,omitempty"`
}

// Validate validates this nfs config
func (m *NfsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNfsAllSquash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsRootPermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsRootSquash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsConfig) validateNfsAllSquash(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsAllSquash) { // not required
		return nil
	}

	if m.NfsAllSquash != nil {
		if err := m.NfsAllSquash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsAllSquash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsAllSquash")
			}
			return err
		}
	}

	return nil
}

func (m *NfsConfig) validateNfsRootPermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsRootPermissions) { // not required
		return nil
	}

	if m.NfsRootPermissions != nil {
		if err := m.NfsRootPermissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsRootPermissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsRootPermissions")
			}
			return err
		}
	}

	return nil
}

func (m *NfsConfig) validateNfsRootSquash(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsRootSquash) { // not required
		return nil
	}

	if m.NfsRootSquash != nil {
		if err := m.NfsRootSquash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsRootSquash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsRootSquash")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nfs config based on the context it is used
func (m *NfsConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNfsAllSquash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfsRootPermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfsRootSquash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NfsConfig) contextValidateNfsAllSquash(ctx context.Context, formats strfmt.Registry) error {

	if m.NfsAllSquash != nil {

		if swag.IsZero(m.NfsAllSquash) { // not required
			return nil
		}

		if err := m.NfsAllSquash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsAllSquash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsAllSquash")
			}
			return err
		}
	}

	return nil
}

func (m *NfsConfig) contextValidateNfsRootPermissions(ctx context.Context, formats strfmt.Registry) error {

	if m.NfsRootPermissions != nil {

		if swag.IsZero(m.NfsRootPermissions) { // not required
			return nil
		}

		if err := m.NfsRootPermissions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsRootPermissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsRootPermissions")
			}
			return err
		}
	}

	return nil
}

func (m *NfsConfig) contextValidateNfsRootSquash(ctx context.Context, formats strfmt.Registry) error {

	if m.NfsRootSquash != nil {

		if swag.IsZero(m.NfsRootSquash) { // not required
			return nil
		}

		if err := m.NfsRootSquash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsRootSquash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsRootSquash")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NfsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NfsConfig) UnmarshalBinary(b []byte) error {
	var res NfsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
