// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateEulaConfig Update EULA Version.
//
// Specifies the update to the End User License Agreement information.
//
// swagger:model UpdateEulaConfig
type UpdateEulaConfig struct {

	// Specifies the version of the End User License Agreement that was accepted.
	SignedVersion *int64 `json:"signedVersion,omitempty"`
}

// Validate validates this update eula config
func (m *UpdateEulaConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update eula config based on context it is used
func (m *UpdateEulaConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateEulaConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateEulaConfig) UnmarshalBinary(b []byte) error {
	var res UpdateEulaConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
