// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SourceAttributeFilter Specifies a pair of source filter attribute and its possible values.
//
// swagger:model SourceAttributeFilter
type SourceAttributeFilter struct {

	// Specifies the filter attribute for the source.
	FilterAttribute *string `json:"filterAttribute,omitempty"`

	// Specifies the list of attribute values for above filter.
	AttributeValues []string `json:"attributeValues"`
}

// Validate validates this source attribute filter
func (m *SourceAttributeFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this source attribute filter based on context it is used
func (m *SourceAttributeFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SourceAttributeFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceAttributeFilter) UnmarshalBinary(b []byte) error {
	var res SourceAttributeFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
