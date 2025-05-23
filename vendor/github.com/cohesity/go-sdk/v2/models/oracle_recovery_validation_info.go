// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OracleRecoveryValidationInfo OracleRecoveryValidationInfo
//
// Specifies information related to Oracle Recovery Validation.
//
// swagger:model OracleRecoveryValidationInfo
type OracleRecoveryValidationInfo struct {

	// Specifies whether we will be creating a dummy oracle instance to run the validations. Generally if source and target location are different we create a dummy oracle instance else we use the source db.
	CreateDummyInstance *bool `json:"createDummyInstance,omitempty"`
}

// Validate validates this oracle recovery validation info
func (m *OracleRecoveryValidationInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this oracle recovery validation info based on context it is used
func (m *OracleRecoveryValidationInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleRecoveryValidationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleRecoveryValidationInfo) UnmarshalBinary(b []byte) error {
	var res OracleRecoveryValidationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
