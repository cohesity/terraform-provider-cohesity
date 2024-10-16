// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MigrateOracleCloneParams Refer MigrateOracleCloneParams message in /magneto/base/magneto.proto
//
// OracleUpdateRestoreTaskOptions encapsulates information needed for executing
// migration of a successful Oracle Clone task.
//
// swagger:model MigrateOracleCloneParams
type MigrateOracleCloneParams struct {

	// The delay in secs, after which the migration of the Clone should be
	// triggered.
	// "0" - Means the Migration should start as soon as the Clone successfully
	// finished. That is the default behaviour.
	//
	// 0 < - Means the migration will start when user triggers migration from
	// UI.
	//
	// > 0 - Means the delay in seconds after which the migration of the clone
	// would be triggered. [ For now this is not supported ]
	DelaySecs *int64 `json:"delaySecs,omitempty"`

	// Vector of paths where the contents (i.e. datafile/redo-logs) of the clone
	// DB will be migrated.
	TargetPathVec []string `json:"targetPathVec"`
}

// Validate validates this migrate oracle clone params
func (m *MigrateOracleCloneParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this migrate oracle clone params based on context it is used
func (m *MigrateOracleCloneParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MigrateOracleCloneParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateOracleCloneParams) UnmarshalBinary(b []byte) error {
	var res MigrateOracleCloneParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
