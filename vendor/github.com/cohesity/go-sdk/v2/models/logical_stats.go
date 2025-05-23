// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogicalStats Specifies the total logical data size of all the local and Cloud Tier data stored by the Cohesity Cluster
//
// swagger:model LogicalStats
type LogicalStats struct {

	// Specifies the size of data before reduction by deduplication and compression.
	TotalLogicalUsageBytes *int64 `json:"totalLogicalUsageBytes,omitempty"`
}

// Validate validates this logical stats
func (m *LogicalStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this logical stats based on context it is used
func (m *LogicalStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogicalStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogicalStats) UnmarshalBinary(b []byte) error {
	var res LogicalStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
