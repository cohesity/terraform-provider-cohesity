// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAzureApplicationResponseParams Azure Application Creation parameters
//
// Specifies the response parameters containing the Azure apps created within the Microsoft365 domain.
//
// swagger:model CreateAzureApplicationResponseParams
type CreateAzureApplicationResponseParams struct {

	// Specifies a list of Microsoft365 azure application credentials needed to authenticate & authorize users for Office 365/Azure Workflows.
	// Min Items: 1
	Microsoft365AppCredentialsList []*Office365AppCredentials `json:"microsoft365AppCredentialsList"`
}

// Validate validates this create azure application response params
func (m *CreateAzureApplicationResponseParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMicrosoft365AppCredentialsList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureApplicationResponseParams) validateMicrosoft365AppCredentialsList(formats strfmt.Registry) error {
	if swag.IsZero(m.Microsoft365AppCredentialsList) { // not required
		return nil
	}

	iMicrosoft365AppCredentialsListSize := int64(len(m.Microsoft365AppCredentialsList))

	if err := validate.MinItems("microsoft365AppCredentialsList", "body", iMicrosoft365AppCredentialsListSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Microsoft365AppCredentialsList); i++ {
		if swag.IsZero(m.Microsoft365AppCredentialsList[i]) { // not required
			continue
		}

		if m.Microsoft365AppCredentialsList[i] != nil {
			if err := m.Microsoft365AppCredentialsList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("microsoft365AppCredentialsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("microsoft365AppCredentialsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create azure application response params based on the context it is used
func (m *CreateAzureApplicationResponseParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMicrosoft365AppCredentialsList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureApplicationResponseParams) contextValidateMicrosoft365AppCredentialsList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Microsoft365AppCredentialsList); i++ {

		if m.Microsoft365AppCredentialsList[i] != nil {

			if swag.IsZero(m.Microsoft365AppCredentialsList[i]) { // not required
				return nil
			}

			if err := m.Microsoft365AppCredentialsList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("microsoft365AppCredentialsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("microsoft365AppCredentialsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAzureApplicationResponseParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureApplicationResponseParams) UnmarshalBinary(b []byte) error {
	var res CreateAzureApplicationResponseParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
