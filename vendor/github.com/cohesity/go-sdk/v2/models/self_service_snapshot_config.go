// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SelfServiceSnapshotConfig Specifies the self service snapshot config of a view.
//
// swagger:model SelfServiceSnapshotConfig
type SelfServiceSnapshotConfig struct {

	// Specifies if self service snapshot feature is enabled. If this is set to true, the feature will also be enabled for NFS protocol. This field is deprecated.
	Enabled *bool `json:"enabled,omitempty"`

	// Specifies if previouse versions feature is enabled with SMB protocol.
	PreviousVersionsEnabled *bool `json:"previousVersionsEnabled,omitempty"`

	// Specifies if self service snapshot feature is enabled for NFS protocol.
	NfsAccessEnabled *bool `json:"nfsAccessEnabled,omitempty"`

	// Specifies if self service snapshot feature is enabled for SMB protocol.
	SmbAccessEnabled *bool `json:"smbAccessEnabled,omitempty"`

	// Specifies the directory name for the snapshots.
	SnapshotDirectoryName *string `json:"snapshotDirectoryName,omitempty"`

	// Specifies the alternate directory name for the snapshots. If it is not set, this feature for SMB protocol will not work.
	AlternateSnapshotDirectoryName *string `json:"alternateSnapshotDirectoryName,omitempty"`

	// Specifies a list of sids who has access to the snapshots.
	AllowAccessSids []string `json:"allowAccessSids,omitempty"`

	// Specifies a list of sids who does not have access to the snapshots. This field overrides 'allowAccessSids'.
	DenyAccessSids []string `json:"denyAccessSids,omitempty"`
}

// Validate validates this self service snapshot config
func (m *SelfServiceSnapshotConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this self service snapshot config based on context it is used
func (m *SelfServiceSnapshotConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SelfServiceSnapshotConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SelfServiceSnapshotConfig) UnmarshalBinary(b []byte) error {
	var res SelfServiceSnapshotConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
