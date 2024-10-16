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

// ProtectedObject Protected Object.
//
// Provides details about a Protected Object.
//
// swagger:model ProtectedObject
type ProtectedObject struct {

	// Specifies the id of the jobs protecting this Protection Source.
	JobID *UniversalID `json:"jobId,omitempty"`

	// If protection fails then specifies why the protection failed on this
	// object.
	ProtectionFauilureReason *string `json:"protectionFauilureReason,omitempty"`

	// Specifies the id of the Protection Source.
	ProtectionSourceID *int64 `json:"protectionSourceId,omitempty"`
}

// Validate validates this protected object
func (m *ProtectedObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectedObject) validateJobID(formats strfmt.Registry) error {
	if swag.IsZero(m.JobID) { // not required
		return nil
	}

	if m.JobID != nil {
		if err := m.JobID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jobId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this protected object based on the context it is used
func (m *ProtectedObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJobID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectedObject) contextValidateJobID(ctx context.Context, formats strfmt.Registry) error {

	if m.JobID != nil {

		if swag.IsZero(m.JobID) { // not required
			return nil
		}

		if err := m.JobID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jobId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectedObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectedObject) UnmarshalBinary(b []byte) error {
	var res ProtectedObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
