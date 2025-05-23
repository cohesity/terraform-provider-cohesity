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

// UdaIndexedObject UdaIndexedObject
//
// Specifies a Universal Data Adapter indexed object.
//
// swagger:model UdaIndexedObject
type UdaIndexedObject struct {
	CommonIndexedObjectParams

	// Specifies the id of the indexed object.
	ID *string `json:"id,omitempty"`

	// Specifies the full name of the indexed object.
	FullName *string `json:"fullName,omitempty"`

	// Specifies the type of the indexed object.
	ObjectType *string `json:"objectType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UdaIndexedObject) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonIndexedObjectParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonIndexedObjectParams = aO0

	// AO1
	var dataAO1 struct {
		ID *string `json:"id,omitempty"`

		FullName *string `json:"fullName,omitempty"`

		ObjectType *string `json:"objectType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	m.FullName = dataAO1.FullName

	m.ObjectType = dataAO1.ObjectType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UdaIndexedObject) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonIndexedObjectParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ID *string `json:"id,omitempty"`

		FullName *string `json:"fullName,omitempty"`

		ObjectType *string `json:"objectType,omitempty"`
	}

	dataAO1.ID = m.ID

	dataAO1.FullName = m.FullName

	dataAO1.ObjectType = m.ObjectType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this uda indexed object
func (m *UdaIndexedObject) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonIndexedObjectParams
	if err := m.CommonIndexedObjectParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this uda indexed object based on the context it is used
func (m *UdaIndexedObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonIndexedObjectParams
	if err := m.CommonIndexedObjectParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UdaIndexedObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaIndexedObject) UnmarshalBinary(b []byte) error {
	var res UdaIndexedObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
