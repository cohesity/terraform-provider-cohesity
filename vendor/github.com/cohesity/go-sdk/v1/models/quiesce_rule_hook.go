// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuiesceRuleHook Message to specify the container and the commands to run before or after
// taking volume snapshots.
//
// swagger:model QuiesceRule_Hook
type QuiesceRuleHook struct {

	// Command to execute specified as an array.
	Commands []string `json:"commands"`

	// Container within the pod where commands need to be run.
	Container *string `json:"container,omitempty"`

	// If there is an error executing the hook, fail backup.
	FailOnError *bool `json:"failOnError,omitempty"`

	// How long to wait for the command to finish executing. Defaults to 30s.
	Timeout *int64 `json:"timeout,omitempty"`
}

// Validate validates this quiesce rule hook
func (m *QuiesceRuleHook) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this quiesce rule hook based on context it is used
func (m *QuiesceRuleHook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuiesceRuleHook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuiesceRuleHook) UnmarshalBinary(b []byte) error {
	var res QuiesceRuleHook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
