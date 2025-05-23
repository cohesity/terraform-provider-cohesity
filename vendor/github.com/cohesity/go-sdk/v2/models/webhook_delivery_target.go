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

// WebhookDeliveryTarget Webhook config for the alerts matching this rule.
//
// swagger:model WebhookDeliveryTarget
type WebhookDeliveryTarget struct {

	// Options for webhook
	CurlOptions string `json:"curlOptions,omitempty"`

	// Destination webhook URL
	// Required: true
	WebhookURL *string `json:"webhookUrl"`
}

// Validate validates this webhook delivery target
func (m *WebhookDeliveryTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebhookURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookDeliveryTarget) validateWebhookURL(formats strfmt.Registry) error {

	if err := validate.Required("webhookUrl", "body", m.WebhookURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this webhook delivery target based on context it is used
func (m *WebhookDeliveryTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookDeliveryTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookDeliveryTarget) UnmarshalBinary(b []byte) error {
	var res WebhookDeliveryTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
