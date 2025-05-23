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

// CreateBondParams Specifies the parameters needed to create a bond.
//
// swagger:model CreateBondParams
type CreateBondParams struct {

	// Specifies a unique name to identify the bond being created.
	// Required: true
	Name *string `json:"name"`

	// Specifies the names of the secondaries of this bond.
	// Required: true
	Slaves []*string `json:"slaves"`
}

// Validate validates this create bond params
func (m *CreateBondParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlaves(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBondParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateBondParams) validateSlaves(formats strfmt.Registry) error {

	if err := validate.Required("slaves", "body", m.Slaves); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create bond params based on context it is used
func (m *CreateBondParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateBondParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBondParams) UnmarshalBinary(b []byte) error {
	var res CreateBondParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
