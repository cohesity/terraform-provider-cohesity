// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenerateM365DeviceCodeRequestParams Microsoft365 Device Authorization User Code Generation parameters
//
// Specifies the request parameters to generate user and device codes for Microsoft365 Device Authorization Grant for Azure PowerShell.
//
// swagger:model GenerateM365DeviceCodeRequestParams
type GenerateM365DeviceCodeRequestParams struct {

	// Specifies the Microsoft365 domain.
	Domain *string `json:"domain,omitempty"`
}

// Validate validates this generate m365 device code request params
func (m *GenerateM365DeviceCodeRequestParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this generate m365 device code request params based on context it is used
func (m *GenerateM365DeviceCodeRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenerateM365DeviceCodeRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenerateM365DeviceCodeRequestParams) UnmarshalBinary(b []byte) error {
	var res GenerateM365DeviceCodeRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
