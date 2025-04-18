// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConstructRestoreMetaInfoOracleParams ConstructRestoreMetaInfoOracleParams
//
// # Params to fetch oracle restore meta info
//
// swagger:model ConstructRestoreMetaInfoOracleParams
type ConstructRestoreMetaInfoOracleParams struct {

	// Specifies the name of the Oracle database that we restore to.
	DbName *string `json:"dbName,omitempty"`

	// Specifies the base directory of Oracle at destination.
	BaseDir *string `json:"baseDir,omitempty"`

	// Specifies the home directory of Oracle at destination.
	HomeDir *string `json:"homeDir,omitempty"`

	// Specifies the location to put the database files(datafiles, logfiles etc.)
	DbFileDestination *string `json:"dbFileDestination,omitempty"`

	// Specifies whether operation is clone or not
	IsClone *bool `json:"isClone,omitempty"`

	// Specifies whether the operation is granular restore or not.
	IsGranularRestore *bool `json:"isGranularRestore,omitempty"`

	// Specifies whether the operation is recovery validation or not.
	IsRecoveryValidation *bool `json:"isRecoveryValidation,omitempty"`

	// Specifies whether the recovery is of type Disaster Recovery.
	IsDisasterRecovery *bool `json:"isDisasterRecovery,omitempty"`
}

// Validate validates this construct restore meta info oracle params
func (m *ConstructRestoreMetaInfoOracleParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this construct restore meta info oracle params based on context it is used
func (m *ConstructRestoreMetaInfoOracleParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConstructRestoreMetaInfoOracleParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConstructRestoreMetaInfoOracleParams) UnmarshalBinary(b []byte) error {
	var res ConstructRestoreMetaInfoOracleParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
