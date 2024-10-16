// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsKmsUpdateParams AwsKmsUpdateParams to define AWS KMS config.
//
// swagger:model AwsKmsUpdateParams
type AwsKmsUpdateParams struct {

	// Access key id needed to access the cloud account.
	// When update cluster config, should encrypte accessKeyId with cluster ID.
	AccessKeyID *string `json:"accessKeyId,omitempty"`

	// Specify the ca certificate path.
	CaCertificatePath *string `json:"caCertificatePath,omitempty"`

	// Specifies the IAM role which will be used to access the security
	// credentials required for API calls.
	IamRoleArn *string `json:"iamRoleArn,omitempty"`

	// Secret access key needed to access the cloud account. This is
	// encrypted with the cluster id.
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`

	// Specify whether to verify SSL when connect with AWS KMS.
	// Default is true.
	VerifySSL *bool `json:"verifySSL,omitempty"`
}

// Validate validates this aws kms update params
func (m *AwsKmsUpdateParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws kms update params based on context it is used
func (m *AwsKmsUpdateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsKmsUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsKmsUpdateParams) UnmarshalBinary(b []byte) error {
	var res AwsKmsUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
