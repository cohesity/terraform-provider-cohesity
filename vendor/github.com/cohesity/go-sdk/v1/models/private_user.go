// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PrivateUser Message encapsulating the details of a user.
//
// swagger:model PrivateUser
type PrivateUser struct {

	// Email address of the user.
	EmailAddress *string `json:"emailAddress,omitempty"`

	// Name of the user.
	Name *string `json:"name,omitempty"`
}

// Validate validates this private user
func (m *PrivateUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this private user based on context it is used
func (m *PrivateUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrivateUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateUser) UnmarshalBinary(b []byte) error {
	var res PrivateUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
