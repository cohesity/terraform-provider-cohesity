// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArchiveWORMProperties archive w o r m properties
//
// swagger:model ArchiveWORMProperties
type ArchiveWORMProperties struct {

	// Indicates the reason of archive not being worm compliant.
	ArchiveWormNonComplianceReason *string `json:"archiveWormNonComplianceReason,omitempty"`

	// Retention time for objects protected by data lock in external targets.
	ExternalTargetDataLockRetentionSecs *int64 `json:"externalTargetDataLockRetentionSecs,omitempty"`

	// Indicates whether this archive run is WORM compliant. For an archive to be
	// compliant, the following criteria must be met:
	// 1. The external target must be capable of supporting object-level
	// time-based locks.
	// 2. The protection policy must have DataLock enabled for the archival
	// section.
	// 3. 'WORM on external target' must be chosen in the archival section of
	// the policy
	// 4. The archival job should have succeeded in adding/extending locks
	// on all new and referred objects.
	IsArchiveWormCompliant *bool `json:"isArchiveWormCompliant,omitempty"`

	// Mode of data lock retention set on the target. Today we only support
	// kAdministrative and kCompliance. Legal hold on the objects in the external
	// target is not supported.
	PolicyType *int32 `json:"policyType,omitempty"`

	// Time at which the WORM protection expires. It will be populated with
	// the same value from ArchivalJobParams.DataLockConstraintsProto.
	// data_lock_expiry_usecs when is_archive_worm_compliant is true.
	WormExpiryTimeUsecs *int64 `json:"wormExpiryTimeUsecs,omitempty"`
}

// Validate validates this archive w o r m properties
func (m *ArchiveWORMProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this archive w o r m properties based on context it is used
func (m *ArchiveWORMProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArchiveWORMProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchiveWORMProperties) UnmarshalBinary(b []byte) error {
	var res ArchiveWORMProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
