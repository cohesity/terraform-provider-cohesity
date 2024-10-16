// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureApplication Message that represents an Azure application.
//
// swagger:model AzureApplication
type AzureApplication struct {

	// Id of the Azure application.
	ApplicationID *string `json:"applicationId,omitempty"`

	// Object Id of the Azure application.
	ApplicationObjectID *string `json:"applicationObjectId,omitempty"`

	// TODO(ENG-416039): Deprecate this and add application_key.
	// Encrypted application key.
	EncryptedApplicationKey *string `json:"encryptedApplicationKey,omitempty"`

	// Application key encrypted by magneto.
	InternalEncryptedApplicationKey []uint8 `json:"internalEncryptedApplicationKey"`
}

// Validate validates this azure application
func (m *AzureApplication) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure application based on context it is used
func (m *AzureApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureApplication) UnmarshalBinary(b []byte) error {
	var res AzureApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
