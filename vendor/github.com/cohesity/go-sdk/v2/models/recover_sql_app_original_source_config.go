// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecoverSQLAppOriginalSourceConfig Recover Sql App New Source Config.
//
// Specifies the additional Source configuration parameters when databases will be recovered to original location.
//
// swagger:model RecoverSqlAppOriginalSourceConfig
type RecoverSQLAppOriginalSourceConfig struct {
	CommonSQLAppSourceConfig

	// Specifies the directory where to put the database data files. Missing directory will be automatically created. If you are overwriting the existing database then this field will be ignored.
	DataFileDirectoryLocation *string `json:"dataFileDirectoryLocation,omitempty"`

	// Specifies the directory where to put the database log files. Missing directory will be automatically created. If you are overwriting the existing database then this field will be ignored.
	LogFileDirectoryLocation *string `json:"logFileDirectoryLocation,omitempty"`

	// Set this to true if tail logs are to be captured before the recovery operation. This is only applicable if database is not being renamed.
	CaptureTailLogs *bool `json:"captureTailLogs,omitempty"`

	// Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery.
	NewDatabaseName *string `json:"newDatabaseName,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RecoverSQLAppOriginalSourceConfig) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonSQLAppSourceConfig
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonSQLAppSourceConfig = aO0

	// now for regular properties
	var propsRecoverSQLAppOriginalSourceConfig struct {
		DataFileDirectoryLocation *string `json:"dataFileDirectoryLocation,omitempty"`

		LogFileDirectoryLocation *string `json:"logFileDirectoryLocation,omitempty"`

		CaptureTailLogs *bool `json:"captureTailLogs,omitempty"`

		NewDatabaseName *string `json:"newDatabaseName,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsRecoverSQLAppOriginalSourceConfig); err != nil {
		return err
	}
	m.DataFileDirectoryLocation = propsRecoverSQLAppOriginalSourceConfig.DataFileDirectoryLocation

	m.LogFileDirectoryLocation = propsRecoverSQLAppOriginalSourceConfig.LogFileDirectoryLocation

	m.CaptureTailLogs = propsRecoverSQLAppOriginalSourceConfig.CaptureTailLogs

	m.NewDatabaseName = propsRecoverSQLAppOriginalSourceConfig.NewDatabaseName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RecoverSQLAppOriginalSourceConfig) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.CommonSQLAppSourceConfig)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsRecoverSQLAppOriginalSourceConfig struct {
		DataFileDirectoryLocation *string `json:"dataFileDirectoryLocation,omitempty"`

		LogFileDirectoryLocation *string `json:"logFileDirectoryLocation,omitempty"`

		CaptureTailLogs *bool `json:"captureTailLogs,omitempty"`

		NewDatabaseName *string `json:"newDatabaseName,omitempty"`
	}
	propsRecoverSQLAppOriginalSourceConfig.DataFileDirectoryLocation = m.DataFileDirectoryLocation

	propsRecoverSQLAppOriginalSourceConfig.LogFileDirectoryLocation = m.LogFileDirectoryLocation

	propsRecoverSQLAppOriginalSourceConfig.CaptureTailLogs = m.CaptureTailLogs

	propsRecoverSQLAppOriginalSourceConfig.NewDatabaseName = m.NewDatabaseName

	jsonDataPropsRecoverSQLAppOriginalSourceConfig, errRecoverSQLAppOriginalSourceConfig := swag.WriteJSON(propsRecoverSQLAppOriginalSourceConfig)
	if errRecoverSQLAppOriginalSourceConfig != nil {
		return nil, errRecoverSQLAppOriginalSourceConfig
	}
	_parts = append(_parts, jsonDataPropsRecoverSQLAppOriginalSourceConfig)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this recover Sql app original source config
func (m *RecoverSQLAppOriginalSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSQLAppSourceConfig
	if err := m.CommonSQLAppSourceConfig.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this recover Sql app original source config based on the context it is used
func (m *RecoverSQLAppOriginalSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSQLAppSourceConfig
	if err := m.CommonSQLAppSourceConfig.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RecoverSQLAppOriginalSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverSQLAppOriginalSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverSQLAppOriginalSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
