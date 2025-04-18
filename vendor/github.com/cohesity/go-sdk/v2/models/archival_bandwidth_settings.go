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

// ArchivalBandwidthSettings Archival Bandwidth Settings
//
// Specifies the global bandwidth setting of the Archival External Target.
//
// swagger:model ArchivalBandwidthSettings
type ArchivalBandwidthSettings struct {
	CommonBandwidthSettings
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ArchivalBandwidthSettings) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonBandwidthSettings
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonBandwidthSettings = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ArchivalBandwidthSettings) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.CommonBandwidthSettings)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this archival bandwidth settings
func (m *ArchivalBandwidthSettings) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonBandwidthSettings
	if err := m.CommonBandwidthSettings.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this archival bandwidth settings based on the context it is used
func (m *ArchivalBandwidthSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonBandwidthSettings
	if err := m.CommonBandwidthSettings.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ArchivalBandwidthSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchivalBandwidthSettings) UnmarshalBinary(b []byte) error {
	var res ArchivalBandwidthSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
