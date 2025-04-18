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

// TieringExternalTargetParams Specifies the parameters which are specific to Tiering purpose type External Targets.
//
// swagger:model TieringExternalTargetParams
type TieringExternalTargetParams struct {
	CommonTieringExternalTargetParams

	// azure params
	AzureParams *TieringAzureExternalTargetParams `json:"azureParams,omitempty"`

	// gcp params
	GcpParams *TieringGcpExternalTargetParams `json:"gcpParams,omitempty"`

	// aws params
	AwsParams *TieringAwsExternalTargetParams `json:"awsParams,omitempty"`

	// oracle params
	OracleParams *TieringOracleExternalTargetParams `json:"oracleParams,omitempty"`

	// s3 comp params
	S3CompParams *TieringS3CompExternalTargetParams `json:"s3CompParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TieringExternalTargetParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonTieringExternalTargetParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonTieringExternalTargetParams = aO0

	// AO1
	var dataAO1 struct {
		AzureParams *TieringAzureExternalTargetParams `json:"azureParams,omitempty"`

		GcpParams *TieringGcpExternalTargetParams `json:"gcpParams,omitempty"`

		AwsParams *TieringAwsExternalTargetParams `json:"awsParams,omitempty"`

		OracleParams *TieringOracleExternalTargetParams `json:"oracleParams,omitempty"`

		S3CompParams *TieringS3CompExternalTargetParams `json:"s3CompParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AzureParams = dataAO1.AzureParams

	m.GcpParams = dataAO1.GcpParams

	m.AwsParams = dataAO1.AwsParams

	m.OracleParams = dataAO1.OracleParams

	m.S3CompParams = dataAO1.S3CompParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TieringExternalTargetParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonTieringExternalTargetParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AzureParams *TieringAzureExternalTargetParams `json:"azureParams,omitempty"`

		GcpParams *TieringGcpExternalTargetParams `json:"gcpParams,omitempty"`

		AwsParams *TieringAwsExternalTargetParams `json:"awsParams,omitempty"`

		OracleParams *TieringOracleExternalTargetParams `json:"oracleParams,omitempty"`

		S3CompParams *TieringS3CompExternalTargetParams `json:"s3CompParams,omitempty"`
	}

	dataAO1.AzureParams = m.AzureParams

	dataAO1.GcpParams = m.GcpParams

	dataAO1.AwsParams = m.AwsParams

	dataAO1.OracleParams = m.OracleParams

	dataAO1.S3CompParams = m.S3CompParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tiering external target params
func (m *TieringExternalTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonTieringExternalTargetParams
	if err := m.CommonTieringExternalTargetParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcpParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3CompParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TieringExternalTargetParams) validateAzureParams(formats strfmt.Registry) error {

	if swag.IsZero(m.AzureParams) { // not required
		return nil
	}

	if m.AzureParams != nil {
		if err := m.AzureParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureParams")
			}
			return err
		}
	}

	return nil
}

func (m *TieringExternalTargetParams) validateGcpParams(formats strfmt.Registry) error {

	if swag.IsZero(m.GcpParams) { // not required
		return nil
	}

	if m.GcpParams != nil {
		if err := m.GcpParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcpParams")
			}
			return err
		}
	}

	return nil
}

func (m *TieringExternalTargetParams) validateAwsParams(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsParams) { // not required
		return nil
	}

	if m.AwsParams != nil {
		if err := m.AwsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsParams")
			}
			return err
		}
	}

	return nil
}

func (m *TieringExternalTargetParams) validateOracleParams(formats strfmt.Registry) error {

	if swag.IsZero(m.OracleParams) { // not required
		return nil
	}

	if m.OracleParams != nil {
		if err := m.OracleParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracleParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oracleParams")
			}
			return err
		}
	}

	return nil
}

func (m *TieringExternalTargetParams) validateS3CompParams(formats strfmt.Registry) error {

	if swag.IsZero(m.S3CompParams) { // not required
		return nil
	}

	if m.S3CompParams != nil {
		if err := m.S3CompParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3CompParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3CompParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tiering external target params based on the context it is used
func (m *TieringExternalTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonTieringExternalTargetParams
	if err := m.CommonTieringExternalTargetParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcpParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOracleParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3CompParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TieringExternalTargetParams) contextValidateAzureParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureParams != nil {

		if swag.IsZero(m.AzureParams) { // not required
			return nil
		}

		if err := m.AzureParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureParams")
			}
			return err
		}
	}

	return nil
}

func (m *TieringExternalTargetParams) contextValidateGcpParams(ctx context.Context, formats strfmt.Registry) error {

	if m.GcpParams != nil {

		if swag.IsZero(m.GcpParams) { // not required
			return nil
		}

		if err := m.GcpParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcpParams")
			}
			return err
		}
	}

	return nil
}

func (m *TieringExternalTargetParams) contextValidateAwsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsParams != nil {

		if swag.IsZero(m.AwsParams) { // not required
			return nil
		}

		if err := m.AwsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsParams")
			}
			return err
		}
	}

	return nil
}

func (m *TieringExternalTargetParams) contextValidateOracleParams(ctx context.Context, formats strfmt.Registry) error {

	if m.OracleParams != nil {

		if swag.IsZero(m.OracleParams) { // not required
			return nil
		}

		if err := m.OracleParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracleParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oracleParams")
			}
			return err
		}
	}

	return nil
}

func (m *TieringExternalTargetParams) contextValidateS3CompParams(ctx context.Context, formats strfmt.Registry) error {

	if m.S3CompParams != nil {

		if swag.IsZero(m.S3CompParams) { // not required
			return nil
		}

		if err := m.S3CompParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3CompParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3CompParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TieringExternalTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TieringExternalTargetParams) UnmarshalBinary(b []byte) error {
	var res TieringExternalTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
