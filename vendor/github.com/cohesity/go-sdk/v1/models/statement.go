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
)

// Statement Protobuf that describes the statement for a bucket policy.
//
// swagger:model Statement
type Statement struct {

	// This field includes keyword which map to permission for each S3
	// operation. We support wildcard('*' and '?') for this field.
	// Some of the valid formats are : "*",
	// "s3:*",
	// "s3:*Object"
	ActionVec []string `json:"actionVec"`

	// The field specified the conditions for when a policy is in effect.
	ConditionVec []*Condition `json:"conditionVec"`

	// This field specifies whether the statement results in an allow or an
	// explicit deny.
	IsAllow *bool `json:"isAllow,omitempty"`

	// If set, actions except the ones specified in action_vec would be
	// considered valid for evaluating the statement. This is set if JSON has
	// "NotAction" element.
	NegateActionVec *bool `json:"negateActionVec,omitempty"`

	// If set, users except the specified principal would be considered valid for
	// evaluating the statement. This is set if JSON has "NotPrincipal" element.
	NegatePrincipal *bool `json:"negatePrincipal,omitempty"`

	// If set, resources except the ones specified in resource_vec would be
	// considered valid for evaluating the statement. This is set if JSON has
	// "NotResource" element.
	NegateResourceVec *bool `json:"negateResourceVec,omitempty"`

	// This field indicates the user entities allowed to access the resource(s).
	Principal *Principal `json:"principal,omitempty"`

	// This field indicates the resource for which the statement is applied.
	// The format we will be using is "urn:csf:s3:::bucket_name/key_name".
	// 'csf' stands for Cohesity SmartFiles. We support wildcard('*' and '?') in
	// the key name.
	// Some of the valid formats are : "urn:csf:s3:::bucket_name",
	// "urn:csf:s3:::bucket_name/*",
	// "urn:csf:s3:::bucket_name/*/ab?"
	// We remove the common prefix 'urn:csf:s3:::bucket_name' from the string
	// and then store it in proto.
	ResourceVec []string `json:"resourceVec"`

	// Statement identifier.
	Sid *string `json:"sid,omitempty"`
}

// Validate validates this statement
func (m *Statement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrincipal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Statement) validateConditionVec(formats strfmt.Registry) error {
	if swag.IsZero(m.ConditionVec) { // not required
		return nil
	}

	for i := 0; i < len(m.ConditionVec); i++ {
		if swag.IsZero(m.ConditionVec[i]) { // not required
			continue
		}

		if m.ConditionVec[i] != nil {
			if err := m.ConditionVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditionVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditionVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Statement) validatePrincipal(formats strfmt.Registry) error {
	if swag.IsZero(m.Principal) { // not required
		return nil
	}

	if m.Principal != nil {
		if err := m.Principal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("principal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("principal")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this statement based on the context it is used
func (m *Statement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditionVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrincipal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Statement) contextValidateConditionVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConditionVec); i++ {

		if m.ConditionVec[i] != nil {

			if swag.IsZero(m.ConditionVec[i]) { // not required
				return nil
			}

			if err := m.ConditionVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditionVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditionVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Statement) contextValidatePrincipal(ctx context.Context, formats strfmt.Registry) error {

	if m.Principal != nil {

		if swag.IsZero(m.Principal) { // not required
			return nil
		}

		if err := m.Principal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("principal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("principal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Statement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Statement) UnmarshalBinary(b []byte) error {
	var res Statement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
