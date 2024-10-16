// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UUIDConfigProto Proto to specify UUID config of an object.
//
// swagger:model UuidConfigProto
type UUIDConfigProto struct {

	// preserve Uuid
	PreserveUUID *bool `json:"preserveUuid,omitempty"`
}

// Validate validates this Uuid config proto
func (m *UUIDConfigProto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Uuid config proto based on context it is used
func (m *UUIDConfigProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UUIDConfigProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UUIDConfigProto) UnmarshalBinary(b []byte) error {
	var res UUIDConfigProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
