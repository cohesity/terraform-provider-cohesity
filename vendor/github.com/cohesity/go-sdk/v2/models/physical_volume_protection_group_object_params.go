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

// PhysicalVolumeProtectionGroupObjectParams Specifies object parameters for creating physical volume Protection Groups.
//
// swagger:model PhysicalVolumeProtectionGroupObjectParams
type PhysicalVolumeProtectionGroupObjectParams struct {

	// Specifies the ID of the object protected.
	// Required: true
	ID *int64 `json:"id"`

	// Specifies the name of the object protected.
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Specifies the list of GUIDs of volumes protected. If empty, then all volumes will be protected by default.
	VolumeGuids []string `json:"volumeGuids"`

	// Specifies whether or not to take a system backup. Applicable only for windows sources.
	EnableSystemBackup *bool `json:"enableSystemBackup,omitempty"`

	// Specifies writer names which should be excluded from physical volume based backups for a given source.
	ExcludedVssWriters []string `json:"excludedVssWriters"`
}

// Validate validates this physical volume protection group object params
func (m *PhysicalVolumeProtectionGroupObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhysicalVolumeProtectionGroupObjectParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this physical volume protection group object params based on the context it is used
func (m *PhysicalVolumeProtectionGroupObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhysicalVolumeProtectionGroupObjectParams) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PhysicalVolumeProtectionGroupObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhysicalVolumeProtectionGroupObjectParams) UnmarshalBinary(b []byte) error {
	var res PhysicalVolumeProtectionGroupObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
