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

// BucketPolicy Protobuf that describes the policy set on a bucket.
//
// swagger:model BucketPolicy
type BucketPolicy struct {

	// The identifier for the bucket policy.
	ID *string `json:"id,omitempty"`

	// Raw JSON string of the stored policy.
	RawPolicyStr *string `json:"rawPolicyStr,omitempty"`

	// This field defines the statement to execute for each request.
	StatementVec []*Statement `json:"statementVec"`

	// This field specifies the language syntax rules that are to be used to
	// process the policy.
	Version *string `json:"version,omitempty"`
}

// Validate validates this bucket policy
func (m *BucketPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatementVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BucketPolicy) validateStatementVec(formats strfmt.Registry) error {
	if swag.IsZero(m.StatementVec) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementVec); i++ {
		if swag.IsZero(m.StatementVec[i]) { // not required
			continue
		}

		if m.StatementVec[i] != nil {
			if err := m.StatementVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statementVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statementVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bucket policy based on the context it is used
func (m *BucketPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatementVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BucketPolicy) contextValidateStatementVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatementVec); i++ {

		if m.StatementVec[i] != nil {

			if swag.IsZero(m.StatementVec[i]) { // not required
				return nil
			}

			if err := m.StatementVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statementVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statementVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BucketPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BucketPolicy) UnmarshalBinary(b []byte) error {
	var res BucketPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
