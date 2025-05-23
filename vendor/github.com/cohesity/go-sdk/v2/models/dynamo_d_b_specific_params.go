// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DynamoDBSpecificParams AWS Dynamo DB source register parameters
//
// # Specifies the Dynamo DB specific parameters for source registration
//
// swagger:model DynamoDBSpecificParams
type DynamoDBSpecificParams struct {

	// Specifies the s3 bucket URI which is used for import and export during backup and recovery.
	S3URI *string `json:"s3URI,omitempty"`
}

// Validate validates this dynamo d b specific params
func (m *DynamoDBSpecificParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dynamo d b specific params based on context it is used
func (m *DynamoDBSpecificParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DynamoDBSpecificParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DynamoDBSpecificParams) UnmarshalBinary(b []byte) error {
	var res DynamoDBSpecificParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
