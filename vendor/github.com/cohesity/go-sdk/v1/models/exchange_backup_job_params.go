// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExchangeBackupJobParams Exchange Backup Job Parameters.
//
// Message to capture additional backup job params specific to Exchange.
//
// swagger:model ExchangeBackupJobParams
type ExchangeBackupJobParams struct {

	// Whether the backups should be copy-only.
	// If this is set to true, then Exchange server will not truncate logs after
	// backup.
	IsCopyOnlyFull *bool `json:"isCopyOnlyFull,omitempty"`
}

// Validate validates this exchange backup job params
func (m *ExchangeBackupJobParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exchange backup job params based on context it is used
func (m *ExchangeBackupJobParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeBackupJobParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeBackupJobParams) UnmarshalBinary(b []byte) error {
	var res ExchangeBackupJobParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
