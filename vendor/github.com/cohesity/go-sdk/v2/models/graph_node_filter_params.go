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

// GraphNodeFilterParams Determines filter that can be applied to query node.
//
// swagger:model GraphNodeFilterParams
type GraphNodeFilterParams struct {
	CommonGraphNodeFilterParams

	// Specifies the aad adapter specific node filter params.
	AadParams struct {
		AadGraphNodeFilterParams
	} `json:"aadParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GraphNodeFilterParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonGraphNodeFilterParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonGraphNodeFilterParams = aO0

	// now for regular properties
	var propsGraphNodeFilterParams struct {
		AadParams struct {
			AadGraphNodeFilterParams
		} `json:"aadParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsGraphNodeFilterParams); err != nil {
		return err
	}
	m.AadParams = propsGraphNodeFilterParams.AadParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GraphNodeFilterParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.CommonGraphNodeFilterParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsGraphNodeFilterParams struct {
		AadParams struct {
			AadGraphNodeFilterParams
		} `json:"aadParams,omitempty"`
	}
	propsGraphNodeFilterParams.AadParams = m.AadParams

	jsonDataPropsGraphNodeFilterParams, errGraphNodeFilterParams := swag.WriteJSON(propsGraphNodeFilterParams)
	if errGraphNodeFilterParams != nil {
		return nil, errGraphNodeFilterParams
	}
	_parts = append(_parts, jsonDataPropsGraphNodeFilterParams)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this graph node filter params
func (m *GraphNodeFilterParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonGraphNodeFilterParams
	if err := m.CommonGraphNodeFilterParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAadParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphNodeFilterParams) validateAadParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AadParams) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this graph node filter params based on the context it is used
func (m *GraphNodeFilterParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonGraphNodeFilterParams
	if err := m.CommonGraphNodeFilterParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAadParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphNodeFilterParams) contextValidateAadParams(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GraphNodeFilterParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphNodeFilterParams) UnmarshalBinary(b []byte) error {
	var res GraphNodeFilterParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
