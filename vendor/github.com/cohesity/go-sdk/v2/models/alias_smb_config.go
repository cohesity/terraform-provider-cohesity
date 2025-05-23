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

// AliasSmbConfig Message defining SMB config for IRIS. SMB config contains SMB encryption flags, SMB discoverable flag and Share level permissions.
//
// swagger:model AliasSmbConfig
type AliasSmbConfig struct {

	// Whether the share is discoverable.
	DiscoveryEnabled *bool `json:"discoveryEnabled,omitempty"`

	// Whether SMB encryption is enabled for this share. Encryption is supported only by SMB 3.x dialects. Dialects that do not support would still access data in unencrypted format.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`

	// Whether to enforce encryption for all the sessions for this view. When enabled all unencrypted sessions are disallowed.
	EncryptionRequired *bool `json:"encryptionRequired,omitempty"`

	// Share level permissions. Note: Supported Access: FullControl, Modify, ReadOnly. Supported type: Allow, Deny.
	Permissions []*SmbPermission `json:"permissions"`

	// Specifies a list of super user sids. Duplicate SIDs are not allowed.
	// Unique: true
	SuperUserSids []string `json:"superUserSids"`

	// Indicate if offline file caching is supported.
	CachingEnabled *bool `json:"cachingEnabled,omitempty"`

	// Indicate if share level permission is cleared by user.
	IsShareLevelPermissionEmpty *bool `json:"isShareLevelPermissionEmpty,omitempty"`

	// Indicate the operation lock is enabled by this view.
	OplockEnabled *bool `json:"oplockEnabled,omitempty"`

	// Whether file open handles are persisted to scribe to survive bridge process crash. When set to false, open handles will be kept in memory until the current node has exclusive ticket for the entity handle. When the entity is opened from another node, the exclusive ticket would be revoked from the node. In revoke control flow, the current node would persist the state to scribe. On acquiring the exclusive ticket,another node would read the file open handles from scribe and resume the handling of operation.
	ContinuousAvailability *bool `json:"continuousAvailability,omitempty"`
}

// Validate validates this alias smb config
func (m *AliasSmbConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuperUserSids(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AliasSmbConfig) validatePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AliasSmbConfig) validateSuperUserSids(formats strfmt.Registry) error {
	if swag.IsZero(m.SuperUserSids) { // not required
		return nil
	}

	if err := validate.UniqueItems("superUserSids", "body", m.SuperUserSids); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this alias smb config based on the context it is used
func (m *AliasSmbConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AliasSmbConfig) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Permissions); i++ {

		if m.Permissions[i] != nil {

			if swag.IsZero(m.Permissions[i]) { // not required
				return nil
			}

			if err := m.Permissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AliasSmbConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AliasSmbConfig) UnmarshalBinary(b []byte) error {
	var res AliasSmbConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
