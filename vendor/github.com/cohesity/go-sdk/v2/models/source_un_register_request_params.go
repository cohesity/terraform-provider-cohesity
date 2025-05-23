// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SourceUnRegisterRequestParams Specifies the parameters to unregister the Protection Source.
//
// swagger:model SourceUnRegisterRequestParams
type SourceUnRegisterRequestParams struct {

	// Specifies the parameters to unregister the azure Protection Source which was registered via Express Mode.
	AzureParams *AzureSourceUnRegisterParams `json:"azureParams,omitempty"`
}

// Validate validates this source un register request params
func (m *SourceUnRegisterRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceUnRegisterRequestParams) validateAzureParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureParams) { // not required
		return nil
	}

	if m.AzureParams != nil {
		if err := m.AzureParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this source un register request params based on the context it is used
func (m *SourceUnRegisterRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAzureParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceUnRegisterRequestParams) contextValidateAzureParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureParams != nil {

		if swag.IsZero(m.AzureParams) { // not required
			return nil
		}

		if err := m.AzureParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourceUnRegisterRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceUnRegisterRequestParams) UnmarshalBinary(b []byte) error {
	var res SourceUnRegisterRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
