// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FieldInfo Proto that contains specific information about an Sfdc Field.
//
// swagger:model FieldInfo
type FieldInfo struct {

	// Data Type of the field. Salesforce supports various data types.
	// More on this here -
	// https://developer.salesforce.com/docs/atlas.en-us.object_reference.meta/
	// object_reference/field_types.htm
	// https://developer.salesforce.com/docs/
	// atlas.en-us.220.0.object_reference.meta/object_reference/
	// primitive_data_types.htm
	DataType *string `json:"dataType,omitempty"`

	// field type
	FieldType *int32 `json:"fieldType,omitempty"`

	// Max length of field.
	Length *uint32 `json:"length,omitempty"`

	// Name of the field.
	Name *string `json:"name,omitempty"`

	// The maximum number of digits in a numeric value,
	// or the length of a text value.
	Precision *uint32 `json:"precision,omitempty"`

	// The number of digits to the right of the decimal point in a numeric value.
	// Must be less than the precision value.
	Scale *uint32 `json:"scale,omitempty"`
}

// Validate validates this field info
func (m *FieldInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this field info based on context it is used
func (m *FieldInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FieldInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FieldInfo) UnmarshalBinary(b []byte) error {
	var res FieldInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
