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

// ProtectionPolicyResponse Protection Policy
//
// Specifies the details about the Protection Policy.
//
// swagger:model ProtectionPolicyResponse
type ProtectionPolicyResponse struct {
	ProtectionPolicy

	// Specifies a unique Policy id assigned by the Cohesity Cluster.
	ID *string `json:"id,omitempty"`

	// Specifies the parent policy template id to which the policy is linked to. This field is set only when policy is created from template.
	TemplateID *string `json:"templateId,omitempty"`

	// This field is set to true if the linked policy which is internally created from a policy templates qualifies as usable to create more policies on the cluster. If the linked policy is partially filled and can not create a working policy then this field will be set to false. In case of normal policy created on the cluster, this field wont be populated.
	IsUsable *bool `json:"isUsable,omitempty"`

	// This field is set to true when policy is the replicated policy.
	IsReplicated *bool `json:"isReplicated,omitempty"`

	// Specifies the number of protection groups using the protection policy.
	NumProtectionGroups *int64 `json:"numProtectionGroups,omitempty"`

	// Specifies the number of protected objects using the protection policy.
	NumProtectedObjects *int64 `json:"numProtectedObjects,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProtectionPolicyResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ProtectionPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ProtectionPolicy = aO0

	// AO1
	var dataAO1 struct {
		ID *string `json:"id,omitempty"`

		TemplateID *string `json:"templateId,omitempty"`

		IsUsable *bool `json:"isUsable,omitempty"`

		IsReplicated *bool `json:"isReplicated,omitempty"`

		NumProtectionGroups *int64 `json:"numProtectionGroups,omitempty"`

		NumProtectedObjects *int64 `json:"numProtectedObjects,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	m.TemplateID = dataAO1.TemplateID

	m.IsUsable = dataAO1.IsUsable

	m.IsReplicated = dataAO1.IsReplicated

	m.NumProtectionGroups = dataAO1.NumProtectionGroups

	m.NumProtectedObjects = dataAO1.NumProtectedObjects

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProtectionPolicyResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ProtectionPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ID *string `json:"id,omitempty"`

		TemplateID *string `json:"templateId,omitempty"`

		IsUsable *bool `json:"isUsable,omitempty"`

		IsReplicated *bool `json:"isReplicated,omitempty"`

		NumProtectionGroups *int64 `json:"numProtectionGroups,omitempty"`

		NumProtectedObjects *int64 `json:"numProtectedObjects,omitempty"`
	}

	dataAO1.ID = m.ID

	dataAO1.TemplateID = m.TemplateID

	dataAO1.IsUsable = m.IsUsable

	dataAO1.IsReplicated = m.IsReplicated

	dataAO1.NumProtectionGroups = m.NumProtectionGroups

	dataAO1.NumProtectedObjects = m.NumProtectedObjects

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this protection policy response
func (m *ProtectionPolicyResponse) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this protection policy response based on the context it is used
func (m *ProtectionPolicyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (m *ProtectionPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionPolicyResponse) UnmarshalBinary(b []byte) error {
	var res ProtectionPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
