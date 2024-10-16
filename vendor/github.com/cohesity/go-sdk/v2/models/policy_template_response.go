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

// PolicyTemplateResponse Protection Policy
//
// Specifies the details about the Protection Policy.
//
// swagger:model PolicyTemplateResponse
type PolicyTemplateResponse struct {
	ProtectionPolicy

	// Specifies a unique Policy id assigned by the Cohesity Cluster.
	ID *string `json:"id,omitempty"`

	// Specifies the number of policies linked to this policy template. Only applicable in case of policy template.
	NumLinkedPolicies *int64 `json:"numLinkedPolicies,omitempty"`

	// This field is set to true if this policy template qualifies to create more policies. If the template is partially filled and can not create a working policy then this field will be set to false.
	IsUsable *bool `json:"isUsable,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PolicyTemplateResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ProtectionPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ProtectionPolicy = aO0

	// AO1
	var dataAO1 struct {
		ID *string `json:"id,omitempty"`

		NumLinkedPolicies *int64 `json:"numLinkedPolicies,omitempty"`

		IsUsable *bool `json:"isUsable,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	m.NumLinkedPolicies = dataAO1.NumLinkedPolicies

	m.IsUsable = dataAO1.IsUsable

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PolicyTemplateResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ProtectionPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ID *string `json:"id,omitempty"`

		NumLinkedPolicies *int64 `json:"numLinkedPolicies,omitempty"`

		IsUsable *bool `json:"isUsable,omitempty"`
	}

	dataAO1.ID = m.ID

	dataAO1.NumLinkedPolicies = m.NumLinkedPolicies

	dataAO1.IsUsable = m.IsUsable

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this policy template response
func (m *PolicyTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProtectionPolicy
	if err := m.ProtectionPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this policy template response based on the context it is used
func (m *PolicyTemplateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProtectionPolicy
	if err := m.ProtectionPolicy.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTemplateResponse) UnmarshalBinary(b []byte) error {
	var res PolicyTemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
