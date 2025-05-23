// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClientStatsInLastHours Specifies the Client stats for last hours.
//
// swagger:model ClientStatsInLastHours
type ClientStatsInLastHours struct {

	// Specifies the time range.
	LastHours *int64 `json:"lastHours,omitempty"`

	// Specifies the stats value.
	Value *int64 `json:"value,omitempty"`
}

// Validate validates this client stats in last hours
func (m *ClientStatsInLastHours) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this client stats in last hours based on context it is used
func (m *ClientStatsInLastHours) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientStatsInLastHours) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientStatsInLastHours) UnmarshalBinary(b []byte) error {
	var res ClientStatsInLastHours
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
