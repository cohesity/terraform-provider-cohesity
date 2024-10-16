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

// CreateOrUpdateAPIKeyRequest Specified the request to create a new user API key.
//
// swagger:model CreateOrUpdateAPIKeyRequest
type CreateOrUpdateAPIKeyRequest struct {

	// Specifies the API key name.
	// Required: true
	Name *string `json:"name"`

	// Specifies if the API key is active.
	IsActive *bool `json:"isActive,omitempty"`

	// Specifies the time in milliseconds when the API key will expire. Set null for no expiry.
	ExpiryTimeMsecs *int64 `json:"expiryTimeMsecs,omitempty"`
}

// Validate validates this create or update API key request
func (m *CreateOrUpdateAPIKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateOrUpdateAPIKeyRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create or update API key request based on context it is used
func (m *CreateOrUpdateAPIKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateOrUpdateAPIKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOrUpdateAPIKeyRequest) UnmarshalBinary(b []byte) error {
	var res CreateOrUpdateAPIKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
