// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsKmsConfigurationUpdateParams AWS KMS configuration updatable parameters.
//
// swagger:model AwsKmsConfigurationUpdateParams
type AwsKmsConfigurationUpdateParams struct {

	// Specify the ca certificate.
	CaCertificate *string `json:"caCertificate,omitempty"`

	// AWS account secret access key. Required when 'iamRoleArn' is not given.
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`

	// AWS account access key id. Required when 'iamRoleArn' is not given.
	AccessKeyID *string `json:"accessKeyId,omitempty"`

	// The IAM role which will be used to authenticate with AWS KMS. Required when 'accessKeyId' and 'secretAccessKey' fields are not provided.
	IamRoleArn *string `json:"iamRoleArn,omitempty"`

	// Enable SSL verification or not.
	VerifySSL *bool `json:"verifySSL,omitempty"`
}

// Validate validates this aws kms configuration update params
func (m *AwsKmsConfigurationUpdateParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws kms configuration update params based on context it is used
func (m *AwsKmsConfigurationUpdateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsKmsConfigurationUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsKmsConfigurationUpdateParams) UnmarshalBinary(b []byte) error {
	var res AwsKmsConfigurationUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
