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

// AirgapConfig Airgap configuration
//
// Specifies the Airgap configuration.
//
// swagger:model AirgapConfig
type AirgapConfig struct {

	// Specifies Airgap should be enabled or disabled.
	// Enum: ["Enable","Disable"]
	AirgapStatus *string `json:"airgapStatus,omitempty"`

	// Specifies the firewall profiles allowed when Airgap is enabled.
	ExceptionProfiles []string `json:"exceptionProfiles,omitempty"`
}

// Validate validates this airgap config
func (m *AirgapConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirgapStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var airgapConfigTypeAirgapStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Enable","Disable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		airgapConfigTypeAirgapStatusPropEnum = append(airgapConfigTypeAirgapStatusPropEnum, v)
	}
}

const (

	// AirgapConfigAirgapStatusEnable captures enum value "Enable"
	AirgapConfigAirgapStatusEnable string = "Enable"

	// AirgapConfigAirgapStatusDisable captures enum value "Disable"
	AirgapConfigAirgapStatusDisable string = "Disable"
)

// prop value enum
func (m *AirgapConfig) validateAirgapStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, airgapConfigTypeAirgapStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AirgapConfig) validateAirgapStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AirgapStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateAirgapStatusEnum("airgapStatus", "body", *m.AirgapStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this airgap config based on context it is used
func (m *AirgapConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AirgapConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AirgapConfig) UnmarshalBinary(b []byte) error {
	var res AirgapConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
