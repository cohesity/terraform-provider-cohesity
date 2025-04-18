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

// CommonAwsExternalTargetParams AWS External Target Request Common Params.
//
// Specifies the common parameters which are specific to AWS related External Targets.
//
// swagger:model CommonAwsExternalTargetParams
type CommonAwsExternalTargetParams struct {

	// Specifies bucket name of the External Target.
	// Required: true
	BucketName *string `json:"bucketName"`

	// Specifies region of the External Target.
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this common aws external target params
func (m *CommonAwsExternalTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucketName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonAwsExternalTargetParams) validateBucketName(formats strfmt.Registry) error {

	if err := validate.Required("bucketName", "body", m.BucketName); err != nil {
		return err
	}

	return nil
}

func (m *CommonAwsExternalTargetParams) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this common aws external target params based on context it is used
func (m *CommonAwsExternalTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonAwsExternalTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonAwsExternalTargetParams) UnmarshalBinary(b []byte) error {
	var res CommonAwsExternalTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
