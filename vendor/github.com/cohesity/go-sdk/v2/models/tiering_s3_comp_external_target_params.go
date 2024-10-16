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

// TieringS3CompExternalTargetParams S3 Compatible Tiering External Target Request Common Params for tiering purpose type.
//
// Specifies the parameters which are specific to S3 Compatible related External Targets of tiering purpose type.
//
// swagger:model TieringS3CompExternalTargetParams
type TieringS3CompExternalTargetParams struct {
	CommonS3CompExternalTargetParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TieringS3CompExternalTargetParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonS3CompExternalTargetParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonS3CompExternalTargetParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TieringS3CompExternalTargetParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.CommonS3CompExternalTargetParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tiering s3 comp external target params
func (m *TieringS3CompExternalTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonS3CompExternalTargetParams
	if err := m.CommonS3CompExternalTargetParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this tiering s3 comp external target params based on the context it is used
func (m *TieringS3CompExternalTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonS3CompExternalTargetParams
	if err := m.CommonS3CompExternalTargetParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TieringS3CompExternalTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TieringS3CompExternalTargetParams) UnmarshalBinary(b []byte) error {
	var res TieringS3CompExternalTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
