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

// RestoreRDSPostgresParams Aws RDS environment specific restore Parameters.
//
// swagger:model RestoreRDSPostgresParams
type RestoreRDSPostgresParams struct {

	// Target Parameters to be filled if Restore target is AWS.
	AwsTargetParams *AwsTargetParams `json:"awsTargetParams,omitempty"`

	// If false, recovery will fail if the database (with same name as this
	// request) exists on the target server.
	// If true, recovery will delete/overwrite the existing database as part of
	// recovery.
	OverwriteDatabase *bool `json:"overwriteDatabase,omitempty"`

	// Specifies the prefix to be prepended to the object name
	// after the recovery.
	PrefixToDatabaseName *string `json:"prefixToDatabaseName,omitempty"`

	// Specifies the suffix to be appended to the object name
	// after the recovery.
	SuffixToDatabaseName *string `json:"suffixToDatabaseName,omitempty"`
}

// Validate validates this restore r d s postgres params
func (m *RestoreRDSPostgresParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsTargetParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreRDSPostgresParams) validateAwsTargetParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsTargetParams) { // not required
		return nil
	}

	if m.AwsTargetParams != nil {
		if err := m.AwsTargetParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsTargetParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this restore r d s postgres params based on the context it is used
func (m *RestoreRDSPostgresParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsTargetParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreRDSPostgresParams) contextValidateAwsTargetParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsTargetParams != nil {

		if swag.IsZero(m.AwsTargetParams) { // not required
			return nil
		}

		if err := m.AwsTargetParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsTargetParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsTargetParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreRDSPostgresParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreRDSPostgresParams) UnmarshalBinary(b []byte) error {
	var res RestoreRDSPostgresParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
