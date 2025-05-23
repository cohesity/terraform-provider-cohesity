// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VersionInfo Version Info
//
// # Id containing the unique identifer and version
//
// swagger:model VersionInfo
type VersionInfo struct {

	// Unique identifier for the string entity. This field is used to uniquely distinguish different entities within the system.
	ID *string `json:"id,omitempty"`

	// Version number associated with the string id. This can be used to track different versions of the entity id over time. The string ID assigned to the an entity may change (infrequently) across software versions.
	Version *int64 `json:"version,omitempty"`
}

// Validate validates this version info
func (m *VersionInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this version info based on context it is used
func (m *VersionInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VersionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionInfo) UnmarshalBinary(b []byte) error {
	var res VersionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
