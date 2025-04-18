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

// UpdateProtectionGroupRunResponseBody Update Protection Group Run Response Body.
//
// Specifies the response of update Protection Group Runs request.
//
// swagger:model UpdateProtectionGroupRunResponseBody
type UpdateProtectionGroupRunResponseBody struct {

	// Specifies a list of ids of Protection Group Runs that are successfully updated.
	SuccessfulRunIds []string `json:"successfulRunIds"`

	// Specifies a list of ids of Protection Group Runs that failed to update along with error details.
	FailedRuns []*FailedRunDetails `json:"failedRuns"`
}

// Validate validates this update protection group run response body
func (m *UpdateProtectionGroupRunResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailedRuns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateProtectionGroupRunResponseBody) validateFailedRuns(formats strfmt.Registry) error {
	if swag.IsZero(m.FailedRuns) { // not required
		return nil
	}

	for i := 0; i < len(m.FailedRuns); i++ {
		if swag.IsZero(m.FailedRuns[i]) { // not required
			continue
		}

		if m.FailedRuns[i] != nil {
			if err := m.FailedRuns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failedRuns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failedRuns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update protection group run response body based on the context it is used
func (m *UpdateProtectionGroupRunResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFailedRuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateProtectionGroupRunResponseBody) contextValidateFailedRuns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FailedRuns); i++ {

		if m.FailedRuns[i] != nil {

			if swag.IsZero(m.FailedRuns[i]) { // not required
				return nil
			}

			if err := m.FailedRuns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failedRuns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failedRuns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProtectionGroupRunResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProtectionGroupRunResponseBody) UnmarshalBinary(b []byte) error {
	var res UpdateProtectionGroupRunResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
