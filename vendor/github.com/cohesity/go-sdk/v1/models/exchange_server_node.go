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

// ExchangeServerNode Represents an Exchange server that is standalone or DAG node.
//
// This object can be used when discovering a Windows box without knowing
// whether its standalone or DAG Exchange server.
//
// swagger:model ExchangeServerNode
type ExchangeServerNode struct {

	// Exchange server info
	Serverinfo *ExchangeServerInfo `json:"serverinfo,omitempty"`

	// Exchange server type.
	Type *int32 `json:"type,omitempty"`
}

// Validate validates this exchange server node
func (m *ExchangeServerNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerinfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeServerNode) validateServerinfo(formats strfmt.Registry) error {
	if swag.IsZero(m.Serverinfo) { // not required
		return nil
	}

	if m.Serverinfo != nil {
		if err := m.Serverinfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serverinfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serverinfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this exchange server node based on the context it is used
func (m *ExchangeServerNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServerinfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeServerNode) contextValidateServerinfo(ctx context.Context, formats strfmt.Registry) error {

	if m.Serverinfo != nil {

		if swag.IsZero(m.Serverinfo) { // not required
			return nil
		}

		if err := m.Serverinfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serverinfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serverinfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeServerNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeServerNode) UnmarshalBinary(b []byte) error {
	var res ExchangeServerNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
