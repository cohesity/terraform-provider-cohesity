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

// RegisteredEntityCassandraParams Contains all params specified by the user while registering a cassandra
// entity.
//
// swagger:model RegisteredEntityCassandraParams
type RegisteredEntityCassandraParams struct {

	// cassandra additional params
	CassandraAdditionalParams *CassandraAdditionalParams `json:"cassandraAdditionalParams,omitempty"`

	// cassandra connect params
	CassandraConnectParams *PrivateCassandraConnectParams `json:"cassandraConnectParams,omitempty"`

	// Whether the request is discovery request or the final registration.
	IsDiscoveryRequest *bool `json:"isDiscoveryRequest,omitempty"`
}

// Validate validates this registered entity cassandra params
func (m *RegisteredEntityCassandraParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCassandraAdditionalParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCassandraConnectParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisteredEntityCassandraParams) validateCassandraAdditionalParams(formats strfmt.Registry) error {
	if swag.IsZero(m.CassandraAdditionalParams) { // not required
		return nil
	}

	if m.CassandraAdditionalParams != nil {
		if err := m.CassandraAdditionalParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraAdditionalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraAdditionalParams")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredEntityCassandraParams) validateCassandraConnectParams(formats strfmt.Registry) error {
	if swag.IsZero(m.CassandraConnectParams) { // not required
		return nil
	}

	if m.CassandraConnectParams != nil {
		if err := m.CassandraConnectParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraConnectParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraConnectParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this registered entity cassandra params based on the context it is used
func (m *RegisteredEntityCassandraParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCassandraAdditionalParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCassandraConnectParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisteredEntityCassandraParams) contextValidateCassandraAdditionalParams(ctx context.Context, formats strfmt.Registry) error {

	if m.CassandraAdditionalParams != nil {

		if swag.IsZero(m.CassandraAdditionalParams) { // not required
			return nil
		}

		if err := m.CassandraAdditionalParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraAdditionalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraAdditionalParams")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredEntityCassandraParams) contextValidateCassandraConnectParams(ctx context.Context, formats strfmt.Registry) error {

	if m.CassandraConnectParams != nil {

		if swag.IsZero(m.CassandraConnectParams) { // not required
			return nil
		}

		if err := m.CassandraConnectParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraConnectParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraConnectParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegisteredEntityCassandraParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisteredEntityCassandraParams) UnmarshalBinary(b []byte) error {
	var res RegisteredEntityCassandraParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
