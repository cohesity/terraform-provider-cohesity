// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorProto error proto
//
// swagger:model ErrorProto
type ErrorProto struct {

	// An optional detail.
	ErrorMsg *string `json:"errorMsg,omitempty"`

	// Error.
	Type *int32 `json:"type,omitempty"`
}

// Validate validates this error proto
func (m *ErrorProto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error proto based on context it is used
func (m *ErrorProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorProto) UnmarshalBinary(b []byte) error {
	var res ErrorProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
