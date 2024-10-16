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

// ExchangeAppParams Exchange App Parameters.
//
// Specifies the Exchange special parameters for the Protection Group.
//
// swagger:model ExchangeAppParams
type ExchangeAppParams struct {

	// Specifies the application id of the Exchange database which has to be protected.
	AppID *int64 `json:"appId,omitempty"`

	// Specifies the application name of the Exchange database which has to be protected.
	// Read Only: true
	AppName *string `json:"appName,omitempty"`
}

// Validate validates this exchange app params
func (m *ExchangeAppParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this exchange app params based on the context it is used
func (m *ExchangeAppParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeAppParams) contextValidateAppName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "appName", "body", m.AppName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeAppParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeAppParams) UnmarshalBinary(b []byte) error {
	var res ExchangeAppParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
