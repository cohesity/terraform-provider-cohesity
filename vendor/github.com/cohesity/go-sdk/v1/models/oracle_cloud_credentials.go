// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OracleCloudCredentials Oracle Cloud Credentials.
//
// Specifies the Oracle Cloud Credentials to connect to an Oracle S3 Compatible
// vault account.
//
// Oracle Cloud Credentials Region, Access-Key-Id and Secret-Access-Key.
// Oracle Cloud properties Tenant and Tier Type.
//
// swagger:model OracleCloudCredentials
type OracleCloudCredentials struct {

	// Specifies access key to connect to Oracle S3 Compatible vault account.
	AccessKeyID *string `json:"accessKeyId,omitempty"`

	// Specifies the region for Oracle S3 Compatible vault account.
	Region *string `json:"region,omitempty"`

	// Specifies the secret access key for Oracle S3 Compatible vault account.
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`

	// Specifies the tenant which is part of the REST endpoints for Oracle S3
	// compatible vaults.
	Tenant *string `json:"tenant,omitempty"`

	// Specifies the storage class of Oracle vault.
	// OracleTierType specifies the storage class for Oracle.
	// 'kOracleTierStandard' indicates a tier type of Oracle properties that
	// requires fast, immediate and frequent access.
	// 'kOracleTierArchive' indicates a tier type of Oracle properties that is
	// rarely accesed and preserved for long times.
	// Enum: ["kOracleTierStandard","kOracleTierArchive"]
	TierType *string `json:"tierType,omitempty"`

	// Specifies the list of all tiers for Amazon account.
	Tiers []string `json:"tiers"`
}

// Validate validates this oracle cloud credentials
func (m *OracleCloudCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTierType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oracleCloudCredentialsTypeTierTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kOracleTierStandard","kOracleTierArchive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oracleCloudCredentialsTypeTierTypePropEnum = append(oracleCloudCredentialsTypeTierTypePropEnum, v)
	}
}

const (

	// OracleCloudCredentialsTierTypeKOracleTierStandard captures enum value "kOracleTierStandard"
	OracleCloudCredentialsTierTypeKOracleTierStandard string = "kOracleTierStandard"

	// OracleCloudCredentialsTierTypeKOracleTierArchive captures enum value "kOracleTierArchive"
	OracleCloudCredentialsTierTypeKOracleTierArchive string = "kOracleTierArchive"
)

// prop value enum
func (m *OracleCloudCredentials) validateTierTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oracleCloudCredentialsTypeTierTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OracleCloudCredentials) validateTierType(formats strfmt.Registry) error {
	if swag.IsZero(m.TierType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTierTypeEnum("tierType", "body", *m.TierType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oracle cloud credentials based on context it is used
func (m *OracleCloudCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleCloudCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleCloudCredentials) UnmarshalBinary(b []byte) error {
	var res OracleCloudCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
