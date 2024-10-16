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

// ExchangeDAG Represents top level Exchange DAG information.The DAG may or may not have an
// Admin Access Point (AAP). For IP-less DAGs, there won't be an AAP. We do not
// care about AAP, so it is not included in this DAG proto.
//
// swagger:model ExchangeDAG
type ExchangeDAG struct {

	// Exchange DAG identity. When queried, if id.guid is empty, then its not a
	// valid DAG.
	ID *ExchangeDAGIdentity `json:"id,omitempty"`
}

// Validate validates this exchange d a g
func (m *ExchangeDAG) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeDAG) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this exchange d a g based on the context it is used
func (m *ExchangeDAG) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeDAG) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if m.ID != nil {

		if swag.IsZero(m.ID) { // not required
			return nil
		}

		if err := m.ID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeDAG) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeDAG) UnmarshalBinary(b []byte) error {
	var res ExchangeDAG
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
