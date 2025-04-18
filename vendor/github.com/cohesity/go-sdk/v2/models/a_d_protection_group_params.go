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

// ADProtectionGroupParams Active Directory(AD) Protection Group Parameters.
//
// Specifies the parameters which are specific to Active directory related Protection Groups.
//
// swagger:model ADProtectionGroupParams
type ADProtectionGroupParams struct {

	// Specifies the list of object ids to be protected.
	// Required: true
	// Min Items: 1
	// Unique: true
	Objects []*ActiveDirectoryProtectionGroupObjectParams `json:"objects"`

	// Specifies the fields required to enable indexing of the protected objects such as files and directories.
	IndexingPolicy *IndexingPolicy `json:"indexingPolicy,omitempty"`
}

// Validate validates this a d protection group params
func (m *ADProtectionGroupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ADProtectionGroupParams) validateObjects(formats strfmt.Registry) error {

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

func (m *ADProtectionGroupParams) validateIndexingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexingPolicy) { // not required
		return nil
	}

	if m.IndexingPolicy != nil {
		if err := m.IndexingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("indexingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("indexingPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a d protection group params based on the context it is used
func (m *ADProtectionGroupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndexingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ADProtectionGroupParams) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ADProtectionGroupParams) contextValidateIndexingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IndexingPolicy != nil {

		if swag.IsZero(m.IndexingPolicy) { // not required
			return nil
		}

		if err := m.IndexingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("indexingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("indexingPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ADProtectionGroupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ADProtectionGroupParams) UnmarshalBinary(b []byte) error {
	var res ADProtectionGroupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
