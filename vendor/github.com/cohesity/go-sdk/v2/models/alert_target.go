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

// AlertTarget Specifies an alert target to receive an alert.
//
// swagger:model AlertTarget
type AlertTarget struct {

	// Specifies an email address to receive an alert.
	// Required: true
	EmailAddress *string `json:"emailAddress"`

	// Specifies the language of the delivery target. Default value is 'en-us'.
	// Enum: ["en-us","ja-jp","zh-cn"]
	Language *string `json:"language,omitempty"`

	// Specifies the recipient type of email recipient. Default value is 'kTo'.
	// Enum: ["kTo","kCc"]
	RecipientType *string `json:"recipientType,omitempty"`
}

// Validate validates this alert target
func (m *AlertTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertTarget) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("emailAddress", "body", m.EmailAddress); err != nil {
		return err
	}

	return nil
}

var alertTargetTypeLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en-us","ja-jp","zh-cn"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertTargetTypeLanguagePropEnum = append(alertTargetTypeLanguagePropEnum, v)
	}
}

const (

	// AlertTargetLanguageEnDashUs captures enum value "en-us"
	AlertTargetLanguageEnDashUs string = "en-us"

	// AlertTargetLanguageJaDashJp captures enum value "ja-jp"
	AlertTargetLanguageJaDashJp string = "ja-jp"

	// AlertTargetLanguageZhDashCn captures enum value "zh-cn"
	AlertTargetLanguageZhDashCn string = "zh-cn"
)

// prop value enum
func (m *AlertTarget) validateLanguageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertTargetTypeLanguagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AlertTarget) validateLanguage(formats strfmt.Registry) error {
	if swag.IsZero(m.Language) { // not required
		return nil
	}

	// value enum
	if err := m.validateLanguageEnum("language", "body", *m.Language); err != nil {
		return err
	}

	return nil
}

var alertTargetTypeRecipientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kTo","kCc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertTargetTypeRecipientTypePropEnum = append(alertTargetTypeRecipientTypePropEnum, v)
	}
}

const (

	// AlertTargetRecipientTypeKTo captures enum value "kTo"
	AlertTargetRecipientTypeKTo string = "kTo"

	// AlertTargetRecipientTypeKCc captures enum value "kCc"
	AlertTargetRecipientTypeKCc string = "kCc"
)

// prop value enum
func (m *AlertTarget) validateRecipientTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertTargetTypeRecipientTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AlertTarget) validateRecipientType(formats strfmt.Registry) error {
	if swag.IsZero(m.RecipientType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecipientTypeEnum("recipientType", "body", *m.RecipientType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this alert target based on context it is used
func (m *AlertTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertTarget) UnmarshalBinary(b []byte) error {
	var res AlertTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
