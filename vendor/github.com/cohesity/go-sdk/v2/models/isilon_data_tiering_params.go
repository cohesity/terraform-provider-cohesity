// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IsilonDataTieringParams Specifies the parameters which are specific to Isilon related Protection Groups.
//
// swagger:model IsilonDataTieringParams
type IsilonDataTieringParams struct {

	// Specifies the objects to be included in the Protection Group.
	// Required: true
	// Min Items: 1
	// Unique: true
	Objects []*ProtectionObjectInput `json:"objects"`

	// Specifies the id of the root of data tiering source.
	SourceID *int64 `json:"sourceId,omitempty"`
}

// Validate validates this isilon data tiering params
func (m *IsilonDataTieringParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsilonDataTieringParams) validateObjects(formats strfmt.Registry) error {

	if err := validate.Required("objects", "body", m.Objects); err != nil {
		return err
	}

	iObjectsSize := int64(len(m.Objects))

	if err := validate.MinItems("objects", "body", iObjectsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("objects", "body", m.Objects); err != nil {
		return err
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this isilon data tiering params based on the context it is used
func (m *IsilonDataTieringParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsilonDataTieringParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IsilonDataTieringParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IsilonDataTieringParams) UnmarshalBinary(b []byte) error {
	var res IsilonDataTieringParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
