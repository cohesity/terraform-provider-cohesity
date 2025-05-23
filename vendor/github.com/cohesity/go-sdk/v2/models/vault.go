// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Vault Specifies the fields of vault.
//
// swagger:model Vault
type Vault struct {

	// Specifies Global vault id.
	GlobalVaultID *string `json:"globalVaultId,omitempty"`

	// Specifies name of vault.
	VaultName *string `json:"vaultName,omitempty"`

	// Specifies id of region where vault resides.
	RegionID *string `json:"regionId,omitempty"`

	// Specifies name of region where vault resides.
	RegionName *string `json:"regionName,omitempty"`
}

// Validate validates this vault
func (m *Vault) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vault based on context it is used
func (m *Vault) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Vault) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vault) UnmarshalBinary(b []byte) error {
	var res Vault
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
