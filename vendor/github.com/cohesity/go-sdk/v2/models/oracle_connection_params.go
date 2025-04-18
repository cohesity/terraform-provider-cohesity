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

// OracleConnectionParams Specifies the parameters to connect to a Oracle node/cluster using given IP or hostname FQDN.
//
// swagger:model OracleConnectionParams
type OracleConnectionParams struct {

	// Specifies the unique identifier to locate the Oracle node or cluster. The host identifier can be IP address or FQDN.
	// Required: true
	Hostname *string `json:"hostname"`
}

// Validate validates this oracle connection params
func (m *OracleConnectionParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OracleConnectionParams) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this oracle connection params based on context it is used
func (m *OracleConnectionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OracleConnectionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OracleConnectionParams) UnmarshalBinary(b []byte) error {
	var res OracleConnectionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
