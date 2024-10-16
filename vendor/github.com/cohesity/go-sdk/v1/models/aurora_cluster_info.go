// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuroraClusterInfo Message that contains the information of the Aurora database cluster.
//
// swagger:model AuroraClusterInfo
type AuroraClusterInfo struct {

	// Aws region of the Aurora DB cluster and S3 bucket.
	AwsRegion *string `json:"awsRegion,omitempty"`

	// Contains the database port of the Aurora cluster.
	DatabasePort *string `json:"databasePort,omitempty"`

	// Contains the postgres db user IAM role Arn.
	DbAccessIamRoleArn *string `json:"dbAccessIamRoleArn,omitempty"`

	// Database user for managing the databases on the Aurora cluster.
	// This user will have exclusive access on all the databases created for the
	// protection group and recovery for a particular tenant.
	DbUserName *string `json:"dbUserName,omitempty"`

	// Contains the host name of the Aurora cluster. This is the writer
	// end point of the Aurora cluster.
	HostName *string `json:"hostName,omitempty"`

	// Contains the kms encryption key used for encryption of data on the
	// Aurora cluster.
	KmsKeyArn *string `json:"kmsKeyArn,omitempty"`
}

// Validate validates this aurora cluster info
func (m *AuroraClusterInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aurora cluster info based on context it is used
func (m *AuroraClusterInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuroraClusterInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuroraClusterInfo) UnmarshalBinary(b []byte) error {
	var res AuroraClusterInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
