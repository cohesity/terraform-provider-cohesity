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

// CassandraSourceConfigParams Parameters fetched by reading cassandra config file.
//
// Specifies the parameters fetched by reading cassandra configuration on the seed node.
//
// swagger:model CassandraSourceConfigParams
type CassandraSourceConfigParams struct {

	// Seed nodes of this cluster.
	Seeds []string `json:"seeds"`

	// Is JMX Authentication enabled in this cluster ?
	IsJmxAuthEnable *bool `json:"isJmxAuthEnable,omitempty"`

	// Info about specific cassandra ports.
	CassandraPortInfo *CassandraPortInfo `json:"cassandraPortInfo,omitempty"`

	// Cassandra security related info.
	CassandraSecurityInfo *CassandraSecurityInfo `json:"cassandraSecurityInfo,omitempty"`

	// Data centers for this cluster.
	DataCenterNames []string `json:"dataCenterNames"`

	// Commit Logs backup location on cassandra nodes
	CommitLogBackupLocation *string `json:"commitLogBackupLocation,omitempty"`

	// Endpoint snitch used for this cluster.
	EndpointSnitch *string `json:"endpointSnitch,omitempty"`

	// Cassandra partitioner required in compaction.
	CassandraPartitioner *string `json:"cassandraPartitioner,omitempty"`

	// Populated if cassandraAuthType is Kerberos.
	KerberosSaslProtocol *string `json:"kerberosSaslProtocol,omitempty"`

	// Cassandra Version.
	CassandraVersion *string `json:"cassandraVersion,omitempty"`

	// DSE Version
	DseVersion *string `json:"dseVersion,omitempty"`
}

// Validate validates this cassandra source config params
func (m *CassandraSourceConfigParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCassandraPortInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCassandraSecurityInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CassandraSourceConfigParams) validateCassandraPortInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CassandraPortInfo) { // not required
		return nil
	}

	if m.CassandraPortInfo != nil {
		if err := m.CassandraPortInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraPortInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraPortInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceConfigParams) validateCassandraSecurityInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CassandraSecurityInfo) { // not required
		return nil
	}

	if m.CassandraSecurityInfo != nil {
		if err := m.CassandraSecurityInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraSecurityInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraSecurityInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cassandra source config params based on the context it is used
func (m *CassandraSourceConfigParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCassandraPortInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCassandraSecurityInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CassandraSourceConfigParams) contextValidateCassandraPortInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CassandraPortInfo != nil {

		if swag.IsZero(m.CassandraPortInfo) { // not required
			return nil
		}

		if err := m.CassandraPortInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraPortInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraPortInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceConfigParams) contextValidateCassandraSecurityInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CassandraSecurityInfo != nil {

		if swag.IsZero(m.CassandraSecurityInfo) { // not required
			return nil
		}

		if err := m.CassandraSecurityInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraSecurityInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraSecurityInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CassandraSourceConfigParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CassandraSourceConfigParams) UnmarshalBinary(b []byte) error {
	var res CassandraSourceConfigParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
