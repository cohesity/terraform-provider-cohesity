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

// KubernetesLabel Kubernetes Label
//
// Represents a single Kubernetes label.
//
// swagger:model KubernetesLabel
type KubernetesLabel struct {

	// The key of the label, used to identify the label.
	// Read Only: true
	Key *string `json:"key,omitempty"`

	// The value associated with the label key.
	// Read Only: true
	Value *string `json:"value,omitempty"`
}

// Validate validates this kubernetes label
func (m *KubernetesLabel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this kubernetes label based on the context it is used
func (m *KubernetesLabel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesLabel) contextValidateKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesLabel) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesLabel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesLabel) UnmarshalBinary(b []byte) error {
	var res KubernetesLabel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
