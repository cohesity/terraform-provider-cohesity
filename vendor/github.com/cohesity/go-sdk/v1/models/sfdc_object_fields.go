// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SfdcObjectFields Sfdc Object Fields Info.
//
// Specifies information about a Sfdc Object Field.
//
// swagger:model SfdcObjectFields
type SfdcObjectFields struct {

	// Specifies the name of a Salesforce Object Field.
	Name *string `json:"name,omitempty"`
}

// Validate validates this sfdc object fields
func (m *SfdcObjectFields) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sfdc object fields based on context it is used
func (m *SfdcObjectFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SfdcObjectFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SfdcObjectFields) UnmarshalBinary(b []byte) error {
	var res SfdcObjectFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
