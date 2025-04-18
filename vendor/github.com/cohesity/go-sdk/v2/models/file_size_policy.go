// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FileSizePolicy Specifies the file's selection rule by file size eg.
// 1. select files greather than 10 Bytes.
// 2. select files less than 20 TiB.
// 3. select files greather than 5 MiB.
// type: object
//
// swagger:model FileSizePolicy
type FileSizePolicy struct {

	// Specifies condition for the file selection.
	// Enum: ["GreaterThan","SmallerThan"]
	Condition *string `json:"condition,omitempty"`

	// Specifies the number of bytes.
	NBytes *int64 `json:"nBytes,omitempty"`
}

// Validate validates this file size policy
func (m *FileSizePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fileSizePolicyTypeConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GreaterThan","SmallerThan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileSizePolicyTypeConditionPropEnum = append(fileSizePolicyTypeConditionPropEnum, v)
	}
}

const (

	// FileSizePolicyConditionGreaterThan captures enum value "GreaterThan"
	FileSizePolicyConditionGreaterThan string = "GreaterThan"

	// FileSizePolicyConditionSmallerThan captures enum value "SmallerThan"
	FileSizePolicyConditionSmallerThan string = "SmallerThan"
)

// prop value enum
func (m *FileSizePolicy) validateConditionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileSizePolicyTypeConditionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileSizePolicy) validateCondition(formats strfmt.Registry) error {
	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	// value enum
	if err := m.validateConditionEnum("condition", "body", *m.Condition); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this file size policy based on context it is used
func (m *FileSizePolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileSizePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileSizePolicy) UnmarshalBinary(b []byte) error {
	var res FileSizePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
