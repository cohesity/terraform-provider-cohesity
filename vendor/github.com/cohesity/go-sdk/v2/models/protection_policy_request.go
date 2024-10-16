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

// ProtectionPolicyRequest Protection Policy Request.
//
// Specifies the request to create a Protection Policy.
//
// swagger:model ProtectionPolicyRequest
type ProtectionPolicyRequest struct {
	ProtectionPolicy

	// Specifies the parent policy template id to which the policy is linked to.
	TemplateID *string `json:"templateId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProtectionPolicyRequest) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ProtectionPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ProtectionPolicy = aO0

	// AO1
	var dataAO1 struct {
		TemplateID *string `json:"templateId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.TemplateID = dataAO1.TemplateID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProtectionPolicyRequest) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ProtectionPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		TemplateID *string `json:"templateId,omitempty"`
	}

	dataAO1.TemplateID = m.TemplateID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this protection policy request
func (m *ProtectionPolicyRequest) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this protection policy request based on the context it is used
func (m *ProtectionPolicyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (m *ProtectionPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionPolicyRequest) UnmarshalBinary(b []byte) error {
	var res ProtectionPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
