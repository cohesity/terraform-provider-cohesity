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

// IdentityProviderConfiguration Identity provider configuration
//
// swagger:model IdentityProviderConfiguration
type IdentityProviderConfiguration struct {
	CommonIdentityProviderConfiguration

	// Specifies id of idp configuration
	ID *int64 `json:"id,omitempty"`

	// Specifies the tenant id if the idp is configured for a tenant. If this is not set, this idp configuration is used for the cluster level users and for all users of tenants not having an idp configuration.
	TenantID *string `json:"tenantId,omitempty"`

	// Specifies name of the vendor providing idp service
	Name *string `json:"name,omitempty"`

	// Specifies domain of idp configuration
	Domain *string `json:"domain,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IdentityProviderConfiguration) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonIdentityProviderConfiguration
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonIdentityProviderConfiguration = aO0

	// AO1
	var dataAO1 struct {
		ID *int64 `json:"id,omitempty"`

		TenantID *string `json:"tenantId,omitempty"`

		Name *string `json:"name,omitempty"`

		Domain *string `json:"domain,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	m.TenantID = dataAO1.TenantID

	m.Name = dataAO1.Name

	m.Domain = dataAO1.Domain

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IdentityProviderConfiguration) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonIdentityProviderConfiguration)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ID *int64 `json:"id,omitempty"`

		TenantID *string `json:"tenantId,omitempty"`

		Name *string `json:"name,omitempty"`

		Domain *string `json:"domain,omitempty"`
	}

	dataAO1.ID = m.ID

	dataAO1.TenantID = m.TenantID

	dataAO1.Name = m.Name

	dataAO1.Domain = m.Domain

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this identity provider configuration
func (m *IdentityProviderConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonIdentityProviderConfiguration
	if err := m.CommonIdentityProviderConfiguration.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this identity provider configuration based on the context it is used
func (m *IdentityProviderConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonIdentityProviderConfiguration
	if err := m.CommonIdentityProviderConfiguration.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IdentityProviderConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityProviderConfiguration) UnmarshalBinary(b []byte) error {
	var res IdentityProviderConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
