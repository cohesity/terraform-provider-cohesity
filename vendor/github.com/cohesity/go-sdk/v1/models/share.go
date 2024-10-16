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
)

// Share Share.
//
// Specifies the share details when request is made for list of shares
// filtered by ShareName parameter.
//
// swagger:model Share
type Share struct {

	// Array of SMB Paths.
	//
	// Specifies the possible paths that can be used to mount this Share
	// as a SMB share. If Active Directory has multiple account names;
	// each machine account has its own path.
	AllSmbMountPaths []string `json:"allSmbMountPaths"`

	// Specifies whether to enable filer audit log on this view alias.
	EnableFilerAuditLog *bool `json:"enableFilerAuditLog,omitempty"`

	// Specifies the SMB encryption for the View Alias. If set, it enables the
	// SMB encryption for the View Alias. Encryption is supported only by SMB 3.x
	// dialects. Dialects that do not support would still access data in
	// unencrypted format.
	EnableSmbEncryption *bool `json:"enableSmbEncryption,omitempty"`

	// If set, it enables discovery of view alias for SMB.
	EnableSmbViewDiscovery *bool `json:"enableSmbViewDiscovery,omitempty"`

	// Specifies the SMB encryption for all the sessions for the View Alias.
	// If set, encryption is enforced for all the sessions for the View Alias.
	// When enabled all future and existing unencrypted sessions are disallowed.
	EnforceSmbEncryption *bool `json:"enforceSmbEncryption,omitempty"`

	// Specifies the path for mounting this Share as an NFS share.
	NfsMountPath *string `json:"nfsMountPath,omitempty"`

	// Specifies the path information for this share.
	Path *string `json:"path,omitempty"`

	// Specifies the path to access this View as an S3 share.
	S3AccessPath *string `json:"s3AccessPath,omitempty"`

	// The name of the share.
	ShareName *string `json:"shareName,omitempty"`

	// Specifies a list of share level permissions.
	SharePermissions []*SmbPermission `json:"sharePermissions"`

	// Specifies the main path for mounting this Share as an SMB share.
	SmbMountPath *string `json:"smbMountPath,omitempty"`

	// Specifies a list of Subnets with IP addresses that have permissions to
	// access the View Alias. (Overrides the Subnets specified at the global
	// Cohesity Cluster level and View level.)
	SubnetWhitelist []*Subnet `json:"subnetWhitelist"`

	// Specifies a list of user sids who have Superuser access to this share.
	SuperUserSids []string `json:"superUserSids"`

	// Specifies the unique id of the tenant.
	TenantID *string `json:"tenantId,omitempty"`

	// Specifies the view id
	ViewID *int64 `json:"viewId,omitempty"`

	// Specifies the view name this share belongs to.
	ViewName *string `json:"viewName,omitempty"`
}

// Validate validates this share
func (m *Share) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSharePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Share) validateSharePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.SharePermissions) { // not required
		return nil
	}

	for i := 0; i < len(m.SharePermissions); i++ {
		if swag.IsZero(m.SharePermissions[i]) { // not required
			continue
		}

		if m.SharePermissions[i] != nil {
			if err := m.SharePermissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sharePermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sharePermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Share) validateSubnetWhitelist(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetWhitelist) { // not required
		return nil
	}

	for i := 0; i < len(m.SubnetWhitelist); i++ {
		if swag.IsZero(m.SubnetWhitelist[i]) { // not required
			continue
		}

		if m.SubnetWhitelist[i] != nil {
			if err := m.SubnetWhitelist[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnetWhitelist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnetWhitelist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this share based on the context it is used
func (m *Share) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSharePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetWhitelist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Share) contextValidateSharePermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SharePermissions); i++ {

		if m.SharePermissions[i] != nil {

			if swag.IsZero(m.SharePermissions[i]) { // not required
				return nil
			}

			if err := m.SharePermissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sharePermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sharePermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Share) contextValidateSubnetWhitelist(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubnetWhitelist); i++ {

		if m.SubnetWhitelist[i] != nil {

			if swag.IsZero(m.SubnetWhitelist[i]) { // not required
				return nil
			}

			if err := m.SubnetWhitelist[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnetWhitelist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnetWhitelist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Share) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Share) UnmarshalBinary(b []byte) error {
	var res Share
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
