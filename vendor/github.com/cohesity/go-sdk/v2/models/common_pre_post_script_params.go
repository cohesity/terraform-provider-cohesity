// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CommonPrePostScriptParams Common PrePost Script Params
//
// Specifies the path to the remote script and any common parameters expected by the remote script.
//
// swagger:model CommonPrePostScriptParams
type CommonPrePostScriptParams struct {

	// Specifies the absolute path to the script on the remote host.
	// Required: true
	Path *string `json:"path"`

	// Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
	Params *string `json:"params,omitempty"`

	// Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
	// Minimum: 1
	TimeoutSecs *int32 `json:"timeoutSecs,omitempty"`

	// Specifies whether the script should be enabled, default value set to true.
	IsActive *bool `json:"isActive,omitempty"`
}

// Validate validates this common pre post script params
func (m *CommonPrePostScriptParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeoutSecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonPrePostScriptParams) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *CommonPrePostScriptParams) validateTimeoutSecs(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeoutSecs) { // not required
		return nil
	}

	if err := validate.MinimumInt("timeoutSecs", "body", int64(*m.TimeoutSecs), 1, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this common pre post script params based on context it is used
func (m *CommonPrePostScriptParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonPrePostScriptParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonPrePostScriptParams) UnmarshalBinary(b []byte) error {
	var res CommonPrePostScriptParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
