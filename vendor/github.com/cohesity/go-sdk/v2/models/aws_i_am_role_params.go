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

// AwsIAmRoleParams AWS IAmRole params.
//
// Specifies the parameters which are specific to IAmRole Authentication Method for AWS External Target.
//
// swagger:model AwsIAmRoleParams
type AwsIAmRoleParams struct {

	// Specifies the Account Id of the external target.
	AccountID *string `json:"accountId,omitempty"`

	// Specifies the I Am Role of the external target.
	// Required: true
	IAmRole *string `json:"iAmRole"`
}

// Validate validates this aws i am role params
func (m *AwsIAmRoleParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIAmRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsIAmRoleParams) validateIAmRole(formats strfmt.Registry) error {

	if err := validate.Required("iAmRole", "body", m.IAmRole); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this aws i am role params based on context it is used
func (m *AwsIAmRoleParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsIAmRoleParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsIAmRoleParams) UnmarshalBinary(b []byte) error {
	var res AwsIAmRoleParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
