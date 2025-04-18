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

// SwiftParams Specifies the parameters to update a Swift configuration.
//
// swagger:model SwiftParams
type SwiftParams struct {

	// Specifies the tenant Id who will use this Swift configuration.
	// Required: true
	TenantID *string `json:"tenantId"`

	// Specifies the associated Keystone configuration Id.
	KeystoneID *int64 `json:"keystoneId,omitempty"`

	// Specifies a list of roles that can operate on Cohesity Swift service.
	// Min Items: 1
	// Unique: true
	OperatorRoles []string `json:"operatorRoles"`
}

// Validate validates this swift params
func (m *SwiftParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwiftParams) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *SwiftParams) validateOperatorRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.OperatorRoles) { // not required
		return nil
	}

	iOperatorRolesSize := int64(len(m.OperatorRoles))

	if err := validate.MinItems("operatorRoles", "body", iOperatorRolesSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("operatorRoles", "body", m.OperatorRoles); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this swift params based on context it is used
func (m *SwiftParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SwiftParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwiftParams) UnmarshalBinary(b []byte) error {
	var res SwiftParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
