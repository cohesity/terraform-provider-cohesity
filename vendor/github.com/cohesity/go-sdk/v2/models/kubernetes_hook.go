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

// KubernetesHook Kubernetes Hook
//
// Specifies the set of parameters for applying scripts.
//
// swagger:model KubernetesHook
type KubernetesHook struct {

	// Specifies the name of the container.
	Container *string `json:"container,omitempty"`

	// Specifies the commands.
	// Required: true
	// Unique: true
	Commands []string `json:"commands"`

	// Specifies whether to fail on error or not.
	FailOnError *bool `json:"failOnError,omitempty"`

	// Specifies timeout for the operation.
	Timeout *int64 `json:"timeout,omitempty"`
}

// Validate validates this kubernetes hook
func (m *KubernetesHook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesHook) validateCommands(formats strfmt.Registry) error {

	if err := validate.Required("commands", "body", m.Commands); err != nil {
		return err
	}

	if err := validate.UniqueItems("commands", "body", m.Commands); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kubernetes hook based on context it is used
func (m *KubernetesHook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesHook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesHook) UnmarshalBinary(b []byte) error {
	var res KubernetesHook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
