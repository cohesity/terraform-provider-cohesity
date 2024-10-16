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

// SupportedPattern Specifies details of the pattern available for search available in an
// application such as Pattern Finder within Analytics Work Bench.
//
// swagger:model SupportedPattern
type SupportedPattern struct {

	// Specifies whether the pattern has been defined by the system or the user.
	IsSystemDefined *bool `json:"isSystemDefined,omitempty"`

	// Specifies the name of the Pattern.
	Name *string `json:"name,omitempty"`

	// Specifies the value of the pattern(Regex).
	Pattern *string `json:"pattern,omitempty"`

	// Specifies the Pattern type.
	// 'REGULAR' indicates that the pattern contains a regular expression.
	// 'TEMPLATE' indicates that the pattern has a pre defined input pattern such as
	// date of the form 'DD-MM-YYYY'.
	// Enum: ["REGULAR","TEMPLATE"]
	PatternType *string `json:"patternType,omitempty"`
}

// Validate validates this supported pattern
func (m *SupportedPattern) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatternType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var supportedPatternTypePatternTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REGULAR","TEMPLATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		supportedPatternTypePatternTypePropEnum = append(supportedPatternTypePatternTypePropEnum, v)
	}
}

const (

	// SupportedPatternPatternTypeREGULAR captures enum value "REGULAR"
	SupportedPatternPatternTypeREGULAR string = "REGULAR"

	// SupportedPatternPatternTypeTEMPLATE captures enum value "TEMPLATE"
	SupportedPatternPatternTypeTEMPLATE string = "TEMPLATE"
)

// prop value enum
func (m *SupportedPattern) validatePatternTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, supportedPatternTypePatternTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SupportedPattern) validatePatternType(formats strfmt.Registry) error {
	if swag.IsZero(m.PatternType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePatternTypeEnum("patternType", "body", *m.PatternType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this supported pattern based on context it is used
func (m *SupportedPattern) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupportedPattern) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportedPattern) UnmarshalBinary(b []byte) error {
	var res SupportedPattern
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
