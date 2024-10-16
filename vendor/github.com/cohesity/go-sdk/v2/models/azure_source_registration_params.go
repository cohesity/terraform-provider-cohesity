// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AzureSourceRegistrationParams Specifies the paramaters to register an Azure source.
//
// swagger:model AzureSourceRegistrationParams
type AzureSourceRegistrationParams struct {

	// Specifies whether the registration is at tenant level or subscription level.
	// Required: true
	// Enum: ["kTenant","kSubscription"]
	RegistrationLevel *string `json:"registrationLevel"`

	// Specifies whether the type of registration is express or manual.
	// Required: true
	// Enum: ["kExpress","kManual"]
	RegistrationWorkflow *string `json:"registrationWorkflow"`

	// Specifies Tenant Id of the active directory of Azure account. Accpets both Azure tanant Id and tenant domain name.
	AzureTenantID *string `json:"azureTenantId,omitempty"`

	// Specifies the credentials for a list of applications from azure active directory.
	ApplicationCredentials []*AzureApplicationCredentials `json:"applicationCredentials"`

	// Specifies the list subscription ids to be registered.
	// Min Items: 1
	// Unique: true
	SubscriptionDetails []*AzureSubscription `json:"subscriptionDetails"`

	// The use cases for which the source is to be registered.
	// Min Items: 1
	// Unique: true
	UseCases []string `json:"useCases"`
}

// Validate validates this azure source registration params
func (m *AzureSourceRegistrationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegistrationLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrationWorkflow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseCases(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var azureSourceRegistrationParamsTypeRegistrationLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kTenant","kSubscription"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureSourceRegistrationParamsTypeRegistrationLevelPropEnum = append(azureSourceRegistrationParamsTypeRegistrationLevelPropEnum, v)
	}
}

const (

	// AzureSourceRegistrationParamsRegistrationLevelKTenant captures enum value "kTenant"
	AzureSourceRegistrationParamsRegistrationLevelKTenant string = "kTenant"

	// AzureSourceRegistrationParamsRegistrationLevelKSubscription captures enum value "kSubscription"
	AzureSourceRegistrationParamsRegistrationLevelKSubscription string = "kSubscription"
)

// prop value enum
func (m *AzureSourceRegistrationParams) validateRegistrationLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, azureSourceRegistrationParamsTypeRegistrationLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AzureSourceRegistrationParams) validateRegistrationLevel(formats strfmt.Registry) error {

	if err := validate.Required("registrationLevel", "body", m.RegistrationLevel); err != nil {
		return err
	}

	// value enum
	if err := m.validateRegistrationLevelEnum("registrationLevel", "body", *m.RegistrationLevel); err != nil {
		return err
	}

	return nil
}

var azureSourceRegistrationParamsTypeRegistrationWorkflowPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kExpress","kManual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureSourceRegistrationParamsTypeRegistrationWorkflowPropEnum = append(azureSourceRegistrationParamsTypeRegistrationWorkflowPropEnum, v)
	}
}

const (

	// AzureSourceRegistrationParamsRegistrationWorkflowKExpress captures enum value "kExpress"
	AzureSourceRegistrationParamsRegistrationWorkflowKExpress string = "kExpress"

	// AzureSourceRegistrationParamsRegistrationWorkflowKManual captures enum value "kManual"
	AzureSourceRegistrationParamsRegistrationWorkflowKManual string = "kManual"
)

// prop value enum
func (m *AzureSourceRegistrationParams) validateRegistrationWorkflowEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, azureSourceRegistrationParamsTypeRegistrationWorkflowPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AzureSourceRegistrationParams) validateRegistrationWorkflow(formats strfmt.Registry) error {

	if err := validate.Required("registrationWorkflow", "body", m.RegistrationWorkflow); err != nil {
		return err
	}

	// value enum
	if err := m.validateRegistrationWorkflowEnum("registrationWorkflow", "body", *m.RegistrationWorkflow); err != nil {
		return err
	}

	return nil
}

func (m *AzureSourceRegistrationParams) validateApplicationCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationCredentials) { // not required
		return nil
	}

	for i := 0; i < len(m.ApplicationCredentials); i++ {
		if swag.IsZero(m.ApplicationCredentials[i]) { // not required
			continue
		}

		if m.ApplicationCredentials[i] != nil {
			if err := m.ApplicationCredentials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applicationCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applicationCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AzureSourceRegistrationParams) validateSubscriptionDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.SubscriptionDetails) { // not required
		return nil
	}

	iSubscriptionDetailsSize := int64(len(m.SubscriptionDetails))

	if err := validate.MinItems("subscriptionDetails", "body", iSubscriptionDetailsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("subscriptionDetails", "body", m.SubscriptionDetails); err != nil {
		return err
	}

	for i := 0; i < len(m.SubscriptionDetails); i++ {
		if swag.IsZero(m.SubscriptionDetails[i]) { // not required
			continue
		}

		if m.SubscriptionDetails[i] != nil {
			if err := m.SubscriptionDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptionDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subscriptionDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var azureSourceRegistrationParamsUseCasesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVirtualMachine","kSQL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureSourceRegistrationParamsUseCasesItemsEnum = append(azureSourceRegistrationParamsUseCasesItemsEnum, v)
	}
}

func (m *AzureSourceRegistrationParams) validateUseCasesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, azureSourceRegistrationParamsUseCasesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AzureSourceRegistrationParams) validateUseCases(formats strfmt.Registry) error {
	if swag.IsZero(m.UseCases) { // not required
		return nil
	}

	iUseCasesSize := int64(len(m.UseCases))

	if err := validate.MinItems("useCases", "body", iUseCasesSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("useCases", "body", m.UseCases); err != nil {
		return err
	}

	for i := 0; i < len(m.UseCases); i++ {

		// value enum
		if err := m.validateUseCasesItemsEnum("useCases"+"."+strconv.Itoa(i), "body", m.UseCases[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this azure source registration params based on the context it is used
func (m *AzureSourceRegistrationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubscriptionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureSourceRegistrationParams) contextValidateApplicationCredentials(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ApplicationCredentials); i++ {

		if m.ApplicationCredentials[i] != nil {

			if swag.IsZero(m.ApplicationCredentials[i]) { // not required
				return nil
			}

			if err := m.ApplicationCredentials[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applicationCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applicationCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AzureSourceRegistrationParams) contextValidateSubscriptionDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubscriptionDetails); i++ {

		if m.SubscriptionDetails[i] != nil {

			if swag.IsZero(m.SubscriptionDetails[i]) { // not required
				return nil
			}

			if err := m.SubscriptionDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptionDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subscriptionDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureSourceRegistrationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureSourceRegistrationParams) UnmarshalBinary(b []byte) error {
	var res AzureSourceRegistrationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
