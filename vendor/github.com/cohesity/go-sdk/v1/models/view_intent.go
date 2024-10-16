// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ViewIntent ViewIntent
//
// Specifies the Intent of the View.
//
// swagger:model ViewIntent
type ViewIntent struct {

	// Specifies the template Id from which the View is created.
	TemplateID *int64 `json:"TemplateId,omitempty"`

	// Specifies the template name from which the View is created.
	TemplateName *string `json:"TemplateName,omitempty"`
}

// Validate validates this view intent
func (m *ViewIntent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this view intent based on context it is used
func (m *ViewIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ViewIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewIntent) UnmarshalBinary(b []byte) error {
	var res ViewIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
