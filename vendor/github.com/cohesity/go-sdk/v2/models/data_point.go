// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataPoint Specifies a data point.
//
// swagger:model DataPoint
type DataPoint struct {

	// Specifies the timestamp of the data point.
	TimestampMsecs *int64 `json:"timestampMsecs,omitempty"`

	// Specifies the data point value in string format.
	StringValue *string `json:"stringValue,omitempty"`

	// Specifies the data point value in double format.
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	// Specifies the data point value in int64 format.
	Int64Value *int64 `json:"int64Value,omitempty"`
}

// Validate validates this data point
func (m *DataPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this data point based on context it is used
func (m *DataPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataPoint) UnmarshalBinary(b []byte) error {
	var res DataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
