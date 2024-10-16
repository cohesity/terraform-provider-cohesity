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

// AzureCoolBlobParams Azure Cool Blob Request Params.
//
// # Specifies the parameters which are specific to Azure related with tier type Cool Blob
//
// swagger:model AzureCoolBlobParams
type AzureCoolBlobParams struct {

	// Specifies the category of the external target.
	// Required: true
	// Enum: ["Azure","AzureStandard","AzureGovCloud"]
	Category *string `json:"category"`

	// Specifies the name of the Azure function app, which is the host of Azure functions.
	FunctionAppName *string `json:"functionAppName,omitempty"`

	// Specifies the access key to deploy Azure function to function app
	FunctionAppDeploymentKey *string `json:"functionAppDeploymentKey,omitempty"`
}

// Validate validates this azure cool blob params
func (m *AzureCoolBlobParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var azureCoolBlobParamsTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Azure","AzureStandard","AzureGovCloud"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureCoolBlobParamsTypeCategoryPropEnum = append(azureCoolBlobParamsTypeCategoryPropEnum, v)
	}
}

const (

	// AzureCoolBlobParamsCategoryAzure captures enum value "Azure"
	AzureCoolBlobParamsCategoryAzure string = "Azure"

	// AzureCoolBlobParamsCategoryAzureStandard captures enum value "AzureStandard"
	AzureCoolBlobParamsCategoryAzureStandard string = "AzureStandard"

	// AzureCoolBlobParamsCategoryAzureGovCloud captures enum value "AzureGovCloud"
	AzureCoolBlobParamsCategoryAzureGovCloud string = "AzureGovCloud"
)

// prop value enum
func (m *AzureCoolBlobParams) validateCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, azureCoolBlobParamsTypeCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AzureCoolBlobParams) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", *m.Category); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this azure cool blob params based on context it is used
func (m *AzureCoolBlobParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureCoolBlobParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCoolBlobParams) UnmarshalBinary(b []byte) error {
	var res AzureCoolBlobParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
