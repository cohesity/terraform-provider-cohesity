// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SanBackupJobParams SAN Backup Job Parameters.
//
// Message to capture any additional backup params for SAN environment.
//
// swagger:model SanBackupJobParams
type SanBackupJobParams struct {

	// Whether backup should continue use secured snapshot. For example IBM
	// FlashSystem SAN env uses this param to create safeguarded snapshot.
	UseSecuredSnapshot *bool `json:"useSecuredSnapshot,omitempty"`
}

// Validate validates this san backup job params
func (m *SanBackupJobParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this san backup job params based on context it is used
func (m *SanBackupJobParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SanBackupJobParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SanBackupJobParams) UnmarshalBinary(b []byte) error {
	var res SanBackupJobParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
