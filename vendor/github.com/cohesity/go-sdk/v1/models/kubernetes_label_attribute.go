// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesLabelAttribute Kubernetes Label Attributes.
//
// Specifies a Kubernetes label.
//
// swagger:model KubernetesLabelAttribute
type KubernetesLabelAttribute struct {

	// Specifies the Cohesity id of the K8s label.
	ID *int64 `json:"id,omitempty"`

	// Specifies the appended key and value of the K8s label.
	Name *string `json:"name,omitempty"`

	// Specifies Kubernetes Unique Identifier (UUID) of the
	// K8s label.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this kubernetes label attribute
func (m *KubernetesLabelAttribute) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes label attribute based on context it is used
func (m *KubernetesLabelAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesLabelAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesLabelAttribute) UnmarshalBinary(b []byte) error {
	var res KubernetesLabelAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
