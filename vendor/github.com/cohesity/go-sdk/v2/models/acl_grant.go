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

// ACLGrant Specifies an ACL grant.
//
// swagger:model AclGrant
type ACLGrant struct {

	// Specifies a list of permissions granted to the grantees.
	// Required: true
	// Min Items: 1
	Permissions []string `json:"permissions"`

	// Specifies the grantee.
	// Required: true
	Grantee *ACLGrantee `json:"grantee"`
}

// Validate validates this Acl grant
func (m *ACLGrant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var aclGrantPermissionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Read","Write","ReadACP","WriteACP","FullControl"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aclGrantPermissionsItemsEnum = append(aclGrantPermissionsItemsEnum, v)
	}
}

func (m *ACLGrant) validatePermissionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, aclGrantPermissionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ACLGrant) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	iPermissionsSize := int64(len(m.Permissions))

	if err := validate.MinItems("permissions", "body", iPermissionsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Permissions); i++ {

		// value enum
		if err := m.validatePermissionsItemsEnum("permissions"+"."+strconv.Itoa(i), "body", m.Permissions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ACLGrant) validateGrantee(formats strfmt.Registry) error {

	if err := validate.Required("grantee", "body", m.Grantee); err != nil {
		return err
	}

	if m.Grantee != nil {
		if err := m.Grantee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grantee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grantee")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Acl grant based on the context it is used
func (m *ACLGrant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGrantee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ACLGrant) contextValidateGrantee(ctx context.Context, formats strfmt.Registry) error {

	if m.Grantee != nil {

		if err := m.Grantee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grantee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grantee")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ACLGrant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ACLGrant) UnmarshalBinary(b []byte) error {
	var res ACLGrant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
