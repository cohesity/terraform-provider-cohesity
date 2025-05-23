// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StoragePolicyOverride Storage Policy Override.
//
// Specifies if inline deduplication and compression settings inherited from
// Storage Domain (View Box) should be disabled for this View.
//
// swagger:model StoragePolicyOverride
type StoragePolicyOverride struct {

	// If false, the inline deduplication and compression settings inherited
	// from the Storage Domain (View Box) apply to this View.
	// If true, both inline deduplication and compression are disabled for this
	// View. This can only be set to true if inline deduplication is set for
	// the Storage Domain (View Box).
	DisableInlineDedupAndCompression *bool `json:"disableInlineDedupAndCompression,omitempty"`

	// If it is set to true, we would disable dedup for writes made in this
	// view irrespective of the view box's storage policy.
	DisableDedup *bool `json:"disableDedup,omitempty"`
}

// Validate validates this storage policy override
func (m *StoragePolicyOverride) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage policy override based on context it is used
func (m *StoragePolicyOverride) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StoragePolicyOverride) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePolicyOverride) UnmarshalBinary(b []byte) error {
	var res StoragePolicyOverride
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
