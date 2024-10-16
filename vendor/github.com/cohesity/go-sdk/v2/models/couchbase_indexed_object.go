// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CouchbaseIndexedObject CouchbaseIndexedObject
//
// Specifies a Couchbase indexed object.
//
// swagger:model CouchbaseIndexedObject
type CouchbaseIndexedObject struct {
	CommonIndexedObjectParams

	// Specifies the Couchbase Object Type. For Couchbase this is alywas set to Bucket.
	// Enum: ["CouchbaseBuckets"]
	Type *string `json:"type,omitempty"`

	// Specifies the id of the indexed object.
	ID *string `json:"id,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CouchbaseIndexedObject) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonIndexedObjectParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonIndexedObjectParams = aO0

	// AO1
	var dataAO1 struct {
		Type *string `json:"type,omitempty"`

		ID *string `json:"id,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Type = dataAO1.Type

	m.ID = dataAO1.ID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CouchbaseIndexedObject) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonIndexedObjectParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Type *string `json:"type,omitempty"`

		ID *string `json:"id,omitempty"`
	}

	dataAO1.Type = m.Type

	dataAO1.ID = m.ID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this couchbase indexed object
func (m *CouchbaseIndexedObject) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonIndexedObjectParams
	if err := m.CommonIndexedObjectParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var couchbaseIndexedObjectTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CouchbaseBuckets"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		couchbaseIndexedObjectTypeTypePropEnum = append(couchbaseIndexedObjectTypeTypePropEnum, v)
	}
}

// property enum
func (m *CouchbaseIndexedObject) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, couchbaseIndexedObjectTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CouchbaseIndexedObject) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this couchbase indexed object based on the context it is used
func (m *CouchbaseIndexedObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
func (m *CouchbaseIndexedObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CouchbaseIndexedObject) UnmarshalBinary(b []byte) error {
	var res CouchbaseIndexedObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
