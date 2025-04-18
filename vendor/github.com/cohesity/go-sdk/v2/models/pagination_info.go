// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaginationInfo Specifies information needed to support pagination.
//
// swagger:model PaginationInfo
type PaginationInfo struct {

	// Specifies a cookie which can be passed in by the user in order to retrieve the next page of results.
	Cookie *string `json:"cookie,omitempty"`
}

// Validate validates this pagination info
func (m *PaginationInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pagination info based on context it is used
func (m *PaginationInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaginationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationInfo) UnmarshalBinary(b []byte) error {
	var res PaginationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
