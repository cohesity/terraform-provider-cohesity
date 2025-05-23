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

// MSSQLVolumeProtectionGroupHostParams MSSQL Volume Protection Group Container Params
//
// Specifies the host specific parameters for a host container in this protection group. Objects specified here should only be MSSQL root containers and will not be protected unless they are also specified in the 'objects' list. This list is just for specifying source level settings.
//
// swagger:model MSSQLVolumeProtectionGroupHostParams
type MSSQLVolumeProtectionGroupHostParams struct {

	// Specifies the id of the host container on which databases are hosted.
	// Required: true
	HostID *int64 `json:"hostId"`

	// Specifies the name of the host container on which databases are hosted.
	// Read Only: true
	HostName *string `json:"hostName,omitempty"`

	// Specifies the list of volume GUIDs to be protected. If not specified, all the volumes of the host will be protected. Note that volumes of host on which databases are hosted are protected even if its not mentioned in this list.
	VolumeGuids []string `json:"volumeGuids"`

	// Specifies whether to enable system/bmr backup using 3rd party tools installed on agent host.
	EnableSystemBackup *bool `json:"enableSystemBackup,omitempty"`
}

// Validate validates this m s SQL volume protection group host params
func (m *MSSQLVolumeProtectionGroupHostParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MSSQLVolumeProtectionGroupHostParams) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("hostId", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this m s SQL volume protection group host params based on the context it is used
func (m *MSSQLVolumeProtectionGroupHostParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHostName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MSSQLVolumeProtectionGroupHostParams) contextValidateHostName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "hostName", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MSSQLVolumeProtectionGroupHostParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MSSQLVolumeProtectionGroupHostParams) UnmarshalBinary(b []byte) error {
	var res MSSQLVolumeProtectionGroupHostParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
