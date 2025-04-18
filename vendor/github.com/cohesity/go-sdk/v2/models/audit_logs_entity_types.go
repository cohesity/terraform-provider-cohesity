// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuditLogsEntityTypes Specifies entity types of audit logs.
//
// swagger:model AuditLogsEntityTypes
type AuditLogsEntityTypes struct {

	// Specifies a list of audit logs entity types.
	EntityTypes []string `json:"entityTypes"`
}

// Validate validates this audit logs entity types
func (m *AuditLogsEntityTypes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this audit logs entity types based on context it is used
func (m *AuditLogsEntityTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditLogsEntityTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLogsEntityTypes) UnmarshalBinary(b []byte) error {
	var res AuditLogsEntityTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
