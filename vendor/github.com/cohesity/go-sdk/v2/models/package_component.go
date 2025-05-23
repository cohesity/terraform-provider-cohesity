// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackageComponent "Software upgrade sub package. Aplicable for one helios package"
//
// swagger:model PackageComponent
type PackageComponent struct {

	// Name of sub package
	PackageName string `json:"packageName,omitempty"`

	// Version of sub package.
	Version string `json:"version,omitempty"`

	// Release Version of sub package.
	Release string `json:"release,omitempty"`
}

// Validate validates this package component
func (m *PackageComponent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this package component based on context it is used
func (m *PackageComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackageComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageComponent) UnmarshalBinary(b []byte) error {
	var res PackageComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
