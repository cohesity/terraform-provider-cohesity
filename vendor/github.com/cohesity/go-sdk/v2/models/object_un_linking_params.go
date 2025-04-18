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

// ObjectUnLinkingParams Specifies the parameters required for unlinking objects. This is currently used as a part of vm migration workflow.
//
// swagger:model ObjectUnLinkingParams
type ObjectUnLinkingParams struct {
	CommonObjectsActionParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ObjectUnLinkingParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonObjectsActionParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonObjectsActionParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ObjectUnLinkingParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.CommonObjectsActionParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this object un linking params
func (m *ObjectUnLinkingParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonObjectsActionParams
	if err := m.CommonObjectsActionParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this object un linking params based on the context it is used
func (m *ObjectUnLinkingParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonObjectsActionParams
	if err := m.CommonObjectsActionParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectUnLinkingParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectUnLinkingParams) UnmarshalBinary(b []byte) error {
	var res ObjectUnLinkingParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
