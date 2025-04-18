// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InterfaceGroupParams Parameters to update an interface group on the cluster.
//
// swagger:model InterfaceGroupParams
type InterfaceGroupParams struct {

	// Name of the interface group.
	// Required: true
	Name *string `json:"name"`

	// Type of the interface group.
	// Required: true
	// Enum: ["Bond","Loopback"]
	Type *string `json:"type"`

	// Node and interface parameters.
	// Required: true
	NodeInterfaceParams []*NodeInterfaceParams `json:"nodeInterfaceParams"`

	// Interface group network parameters.
	NetworkParams *InterfaceGroupNetworkParams `json:"networkParams,omitempty"`
}

// Validate validates this interface group params
func (m *InterfaceGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeInterfaceParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceGroupParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var interfaceGroupParamsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Bond","Loopback"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interfaceGroupParamsTypeTypePropEnum = append(interfaceGroupParamsTypeTypePropEnum, v)
	}
}

const (

	// InterfaceGroupParamsTypeBond captures enum value "Bond"
	InterfaceGroupParamsTypeBond string = "Bond"

	// InterfaceGroupParamsTypeLoopback captures enum value "Loopback"
	InterfaceGroupParamsTypeLoopback string = "Loopback"
)

// prop value enum
func (m *InterfaceGroupParams) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interfaceGroupParamsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InterfaceGroupParams) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceGroupParams) validateNodeInterfaceParams(formats strfmt.Registry) error {

	if err := validate.Required("nodeInterfaceParams", "body", m.NodeInterfaceParams); err != nil {
		return err
	}

	for i := 0; i < len(m.NodeInterfaceParams); i++ {
		if swag.IsZero(m.NodeInterfaceParams[i]) { // not required
			continue
		}

		if m.NodeInterfaceParams[i] != nil {
			if err := m.NodeInterfaceParams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeInterfaceParams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeInterfaceParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InterfaceGroupParams) validateNetworkParams(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkParams) { // not required
		return nil
	}

	if m.NetworkParams != nil {
		if err := m.NetworkParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this interface group params based on the context it is used
func (m *InterfaceGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodeInterfaceParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceGroupParams) contextValidateNodeInterfaceParams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeInterfaceParams); i++ {

		if m.NodeInterfaceParams[i] != nil {

			if swag.IsZero(m.NodeInterfaceParams[i]) { // not required
				return nil
			}

			if err := m.NodeInterfaceParams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeInterfaceParams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeInterfaceParams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InterfaceGroupParams) contextValidateNetworkParams(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkParams != nil {

		if swag.IsZero(m.NetworkParams) { // not required
			return nil
		}

		if err := m.NetworkParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceGroupParams) UnmarshalBinary(b []byte) error {
	var res InterfaceGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
