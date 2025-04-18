// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServicePatchLevel Patch level of a service. It is the number of patches applied for the service on the cluster. If a service is never patched the patch level is 0. If two patches were applied, patch level is 2.
//
// swagger:model ServicePatchLevel
type ServicePatchLevel struct {

	// Specifies the name of the service.
	Service *string `json:"service,omitempty"`

	// Specifies patch level of the service after the patch operation.
	PatchLevel *int64 `json:"patchLevel,omitempty"`

	// Specifies the version of the service patch after the patch operation.
	PatchVersion *string `json:"patchVersion,omitempty"`

	// Specifies patch level of the service before the patch operation.
	StartLevel *int64 `json:"startLevel,omitempty"`

	// Specifies the version of the service running on the cluster before the patch operation.
	StartVersion *string `json:"startVersion,omitempty"`
}

// Validate validates this service patch level
func (m *ServicePatchLevel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service patch level based on context it is used
func (m *ServicePatchLevel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServicePatchLevel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServicePatchLevel) UnmarshalBinary(b []byte) error {
	var res ServicePatchLevel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
