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

// MSSQLFileProtectionGroupHostParams MSSQL File Protection Group Container Params
//
// Specifies the host specific parameters for a host container in this protection group. Objects specified here should only be MSSQL root containers and will not be protected unless they are also specified in the 'objects' list. This list is just for specifying source level settings.
//
// swagger:model MSSQLFileProtectionGroupHostParams
type MSSQLFileProtectionGroupHostParams struct {

	// Specifies the id of the host container on which databases are hosted.
	// Required: true
	HostID *int64 `json:"hostId"`

	// Specifies the name of the host container on which databases are hosted.
	// Read Only: true
	HostName *string `json:"hostName,omitempty"`

	// Specifies whether or not to disable source side deduplication on this source. The default behavior is false unless the user has set 'performSourceSideDeduplication' to true.
	DisableSourceSideDeduplication *bool `json:"disableSourceSideDeduplication,omitempty"`
}

// Validate validates this m s SQL file protection group host params
func (m *MSSQLFileProtectionGroupHostParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MSSQLFileProtectionGroupHostParams) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("hostId", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this m s SQL file protection group host params based on the context it is used
func (m *MSSQLFileProtectionGroupHostParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHostName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MSSQLFileProtectionGroupHostParams) contextValidateHostName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "hostName", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MSSQLFileProtectionGroupHostParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MSSQLFileProtectionGroupHostParams) UnmarshalBinary(b []byte) error {
	var res MSSQLFileProtectionGroupHostParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
