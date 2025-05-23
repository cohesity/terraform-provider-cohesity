// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceState Service state struct
//
// # Structure to hold service status
//
// swagger:model ServiceState
type ServiceState struct {

	// Name of service
	Name *string `json:"name,omitempty"`

	// State of service
	State *string `json:"state,omitempty"`
}

// Validate validates this service state
func (m *ServiceState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service state based on context it is used
func (m *ServiceState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceState) UnmarshalBinary(b []byte) error {
	var res ServiceState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
