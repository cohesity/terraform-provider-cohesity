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

// ErrorClass Specifies a class of error with name and count of that class.
//
// swagger:model ErrorClass
type ErrorClass struct {

	// class name
	// Enum: ["File","Folder","Persistent","Intermittent","Discovery","Ingestion"]
	ClassName string `json:"className,omitempty"`

	// count
	Count *int64 `json:"count,omitempty"`
}

// Validate validates this error class
func (m *ErrorClass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var errorClassTypeClassNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["File","Folder","Persistent","Intermittent","Discovery","Ingestion"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorClassTypeClassNamePropEnum = append(errorClassTypeClassNamePropEnum, v)
	}
}

const (

	// ErrorClassClassNameFile captures enum value "File"
	ErrorClassClassNameFile string = "File"

	// ErrorClassClassNameFolder captures enum value "Folder"
	ErrorClassClassNameFolder string = "Folder"

	// ErrorClassClassNamePersistent captures enum value "Persistent"
	ErrorClassClassNamePersistent string = "Persistent"

	// ErrorClassClassNameIntermittent captures enum value "Intermittent"
	ErrorClassClassNameIntermittent string = "Intermittent"

	// ErrorClassClassNameDiscovery captures enum value "Discovery"
	ErrorClassClassNameDiscovery string = "Discovery"

	// ErrorClassClassNameIngestion captures enum value "Ingestion"
	ErrorClassClassNameIngestion string = "Ingestion"
)

// prop value enum
func (m *ErrorClass) validateClassNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, errorClassTypeClassNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ErrorClass) validateClassName(formats strfmt.Registry) error {
	if swag.IsZero(m.ClassName) { // not required
		return nil
	}

	// value enum
	if err := m.validateClassNameEnum("className", "body", m.ClassName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error class based on context it is used
func (m *ErrorClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorClass) UnmarshalBinary(b []byte) error {
	var res ErrorClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
