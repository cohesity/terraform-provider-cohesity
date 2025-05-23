// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WormProperties WORM Archival data.
//
// Specifies the WORM related properties for this archive.
//
// swagger:model WormProperties
type WormProperties struct {

	// Specifies whether this archive run is WORM compliant
	IsArchiveWormCompliant *bool `json:"isArchiveWormCompliant,omitempty"`

	// Specifies reason of archive not being worm compliant.
	WormNonComplianceReason *string `json:"wormNonComplianceReason,omitempty"`

	// Specifies the time at which the WORM protection expires.
	WormExpiryTimeUsecs *int64 `json:"wormExpiryTimeUsecs,omitempty"`
}

// Validate validates this worm properties
func (m *WormProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this worm properties based on context it is used
func (m *WormProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WormProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WormProperties) UnmarshalBinary(b []byte) error {
	var res WormProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
