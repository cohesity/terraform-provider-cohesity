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

// AzureSubscription Specifies the details of an Azure subscription
//
// swagger:model AzureSubscription
type AzureSubscription struct {

	// Specifies the id of Azure subscription.
	// Required: true
	SubscriptionID *string `json:"subscriptionId"`
}

// Validate validates this azure subscription
func (m *AzureSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureSubscription) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionId", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this azure subscription based on context it is used
func (m *AzureSubscription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureSubscription) UnmarshalBinary(b []byte) error {
	var res AzureSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
