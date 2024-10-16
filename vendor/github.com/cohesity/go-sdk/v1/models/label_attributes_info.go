// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LabelAttributesInfo This message encapsulates information about a label entity.
//
// swagger:model LabelAttributesInfo
type LabelAttributesInfo struct {

	// Entity ID of the label entity in EH.
	EntityID *int64 `json:"entityId,omitempty"`

	// Name of the label entity.
	Name *string `json:"name,omitempty"`

	// UUID of the label entity.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this label attributes info
func (m *LabelAttributesInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this label attributes info based on context it is used
func (m *LabelAttributesInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LabelAttributesInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelAttributesInfo) UnmarshalBinary(b []byte) error {
	var res LabelAttributesInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
