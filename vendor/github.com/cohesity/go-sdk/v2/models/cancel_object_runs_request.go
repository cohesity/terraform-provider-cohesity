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

// CancelObjectRunsRequest Request to cancel object runs.
//
// swagger:model CancelObjectRunsRequest
type CancelObjectRunsRequest struct {

	// Specifies a list of object runs to cancel.
	ObjectRuns []*CancelObjectRunsParams `json:"objectRuns"`
}

// Validate validates this cancel object runs request
func (m *CancelObjectRunsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectRuns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelObjectRunsRequest) validateObjectRuns(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectRuns) { // not required
		return nil
	}

	for i := 0; i < len(m.ObjectRuns); i++ {
		if swag.IsZero(m.ObjectRuns[i]) { // not required
			continue
		}

		if m.ObjectRuns[i] != nil {
			if err := m.ObjectRuns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectRuns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectRuns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cancel object runs request based on the context it is used
func (m *CancelObjectRunsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectRuns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelObjectRunsRequest) contextValidateObjectRuns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ObjectRuns); i++ {

		if m.ObjectRuns[i] != nil {

			if swag.IsZero(m.ObjectRuns[i]) { // not required
				return nil
			}

			if err := m.ObjectRuns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectRuns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectRuns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CancelObjectRunsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelObjectRunsRequest) UnmarshalBinary(b []byte) error {
	var res CancelObjectRunsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
