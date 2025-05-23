// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CassandraSourceRegistrationPatchParams Update cassandra source request parameters.
//
// Specifies parameters to update cassandra source.
//
// swagger:model CassandraSourceRegistrationPatchParams
type CassandraSourceRegistrationPatchParams struct {

	// Any one seed node of the Cassandra cluster.
	SeedNode *string `json:"seedNode,omitempty"`

	// Directory path containing Cassandra configuration YAML file.
	ConfigDirectory *string `json:"configDirectory,omitempty"`

	// Directory from where DSE specific configuration can be read. This should be set only when you are using the DSE distribution of Cassandra.
	DseConfigurationDirectory *string `json:"dseConfigurationDirectory,omitempty"`

	// Set to true if this cluster has DSE tiered storage.
	IsDseTieredStorage *bool `json:"isDseTieredStorage,omitempty"`

	// Set to true if this cluster has DSE Authenticator.
	IsDseAuthenticator *bool `json:"isDseAuthenticator,omitempty"`

	// SSH username + password that will be used for reading configuration file and for scp backup.Either 'sshPasswordCredentials' or 'sshPrivateKeyCredentials' are required.
	SSHPasswordCredentials *SSHPasswordCredentials `json:"sshPasswordCredentials,omitempty"`

	// SSH  userID + privateKey that will be used for reading configuration file and for scp backup.
	SSHPrivateKeyCredentials *SSHPrivateKeyCredentials `json:"sshPrivateKeyCredentials,omitempty"`

	// Data centers for this cluster.
	DataCenterNames []string `json:"dataCenterNames"`

	// Commit Logs backup location on cassandra nodes
	CommitLogBackupLocation *string `json:"commitLogBackupLocation,omitempty"`

	// Principal for the kerberos connection. (This is required only if your Cassandra has Kerberos authentication. Please refer to the user guide.)
	KerberosPrincipal *string `json:"kerberosPrincipal,omitempty"`

	// Contains information about the DSE Solr on this cluster.
	DseSolrInfo *DSESolrInfo `json:"dseSolrInfo,omitempty"`

	// cassandra credentials
	CassandraCredentials *CassandraSourceRegistrationPatchParamsCassandraCredentials `json:"cassandraCredentials,omitempty"`

	// jmx credentials
	JmxCredentials *CassandraSourceRegistrationPatchParamsJmxCredentials `json:"jmxCredentials,omitempty"`
}

// Validate validates this cassandra source registration patch params
func (m *CassandraSourceRegistrationPatchParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSSHPasswordCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHPrivateKeyCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDseSolrInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCassandraCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJmxCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CassandraSourceRegistrationPatchParams) validateSSHPasswordCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHPasswordCredentials) { // not required
		return nil
	}

	if m.SSHPasswordCredentials != nil {
		if err := m.SSHPasswordCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshPasswordCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sshPasswordCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParams) validateSSHPrivateKeyCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHPrivateKeyCredentials) { // not required
		return nil
	}

	if m.SSHPrivateKeyCredentials != nil {
		if err := m.SSHPrivateKeyCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshPrivateKeyCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sshPrivateKeyCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParams) validateDseSolrInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DseSolrInfo) { // not required
		return nil
	}

	if m.DseSolrInfo != nil {
		if err := m.DseSolrInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dseSolrInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dseSolrInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParams) validateCassandraCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.CassandraCredentials) { // not required
		return nil
	}

	if m.CassandraCredentials != nil {
		if err := m.CassandraCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParams) validateJmxCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.JmxCredentials) { // not required
		return nil
	}

	if m.JmxCredentials != nil {
		if err := m.JmxCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jmxCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jmxCredentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cassandra source registration patch params based on the context it is used
func (m *CassandraSourceRegistrationPatchParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSSHPasswordCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSSHPrivateKeyCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDseSolrInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCassandraCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJmxCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CassandraSourceRegistrationPatchParams) contextValidateSSHPasswordCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.SSHPasswordCredentials != nil {

		if swag.IsZero(m.SSHPasswordCredentials) { // not required
			return nil
		}

		if err := m.SSHPasswordCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshPasswordCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sshPasswordCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParams) contextValidateSSHPrivateKeyCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.SSHPrivateKeyCredentials != nil {

		if swag.IsZero(m.SSHPrivateKeyCredentials) { // not required
			return nil
		}

		if err := m.SSHPrivateKeyCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshPrivateKeyCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sshPrivateKeyCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParams) contextValidateDseSolrInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DseSolrInfo != nil {

		if swag.IsZero(m.DseSolrInfo) { // not required
			return nil
		}

		if err := m.DseSolrInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dseSolrInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dseSolrInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParams) contextValidateCassandraCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.CassandraCredentials != nil {

		if swag.IsZero(m.CassandraCredentials) { // not required
			return nil
		}

		if err := m.CassandraCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParams) contextValidateJmxCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.JmxCredentials != nil {

		if swag.IsZero(m.JmxCredentials) { // not required
			return nil
		}

		if err := m.JmxCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jmxCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jmxCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CassandraSourceRegistrationPatchParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CassandraSourceRegistrationPatchParams) UnmarshalBinary(b []byte) error {
	var res CassandraSourceRegistrationPatchParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CassandraSourceRegistrationPatchParamsCassandraCredentials Cassandra Credentials for this cluster.
//
// swagger:model CassandraSourceRegistrationPatchParamsCassandraCredentials
type CassandraSourceRegistrationPatchParamsCassandraCredentials struct {

	// Cassandra password.
	// Required: true
	Password *string `json:"password"`

	// Cassandra username.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this cassandra source registration patch params cassandra credentials
func (m *CassandraSourceRegistrationPatchParamsCassandraCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CassandraSourceRegistrationPatchParamsCassandraCredentials) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("cassandraCredentials"+"."+"password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParamsCassandraCredentials) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("cassandraCredentials"+"."+"username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cassandra source registration patch params cassandra credentials based on context it is used
func (m *CassandraSourceRegistrationPatchParamsCassandraCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CassandraSourceRegistrationPatchParamsCassandraCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CassandraSourceRegistrationPatchParamsCassandraCredentials) UnmarshalBinary(b []byte) error {
	var res CassandraSourceRegistrationPatchParamsCassandraCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CassandraSourceRegistrationPatchParamsJmxCredentials JMX Credentials for this cluster. These should be the same for all the nodes
//
// swagger:model CassandraSourceRegistrationPatchParamsJmxCredentials
type CassandraSourceRegistrationPatchParamsJmxCredentials struct {

	// JMX password.
	// Required: true
	Password *string `json:"password"`

	// JMX username.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this cassandra source registration patch params jmx credentials
func (m *CassandraSourceRegistrationPatchParamsJmxCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CassandraSourceRegistrationPatchParamsJmxCredentials) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("jmxCredentials"+"."+"password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CassandraSourceRegistrationPatchParamsJmxCredentials) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("jmxCredentials"+"."+"username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cassandra source registration patch params jmx credentials based on context it is used
func (m *CassandraSourceRegistrationPatchParamsJmxCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CassandraSourceRegistrationPatchParamsJmxCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CassandraSourceRegistrationPatchParamsJmxCredentials) UnmarshalBinary(b []byte) error {
	var res CassandraSourceRegistrationPatchParamsJmxCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
