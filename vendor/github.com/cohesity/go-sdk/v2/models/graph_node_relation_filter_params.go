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

// GraphNodeRelationFilterParams Determines filter that can be applied to query edge of a graph node.
//
// swagger:model GraphNodeRelationFilterParams
type GraphNodeRelationFilterParams struct {

	// Specifies the aad adapter specific relation filter params.
	AadParams struct {
		AadRelationFilterParams
	} `json:"aadParams,omitempty"`
}

// Validate validates this graph node relation filter params
func (m *GraphNodeRelationFilterParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAadParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphNodeRelationFilterParams) validateAadParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AadParams) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this graph node relation filter params based on the context it is used
func (m *GraphNodeRelationFilterParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAadParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphNodeRelationFilterParams) contextValidateAadParams(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GraphNodeRelationFilterParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphNodeRelationFilterParams) UnmarshalBinary(b []byte) error {
	var res GraphNodeRelationFilterParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
