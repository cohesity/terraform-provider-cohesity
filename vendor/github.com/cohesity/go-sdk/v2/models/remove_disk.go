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

// RemoveDisk Disk Removal Response.
//
// Specifies details of disk removal response.
//
// swagger:model RemoveDisk
type RemoveDisk struct {

	// If true, Disk is marked for removal.
	MarkedForRemoval *bool `json:"markedForRemoval,omitempty"`

	// Specifies id of the disk.
	ID *int64 `json:"id,omitempty"`

	// Specifies the pre-check validations results.
	ValidationChecks []*PreCheckValidation `json:"validationChecks"`

	// Specifies the last run time of the pre-checks execution in Unix epoch timestamp (in seconds).
	TimestampSecs *int64 `json:"timestampSecs,omitempty"`
}

// Validate validates this remove disk
func (m *RemoveDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidationChecks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveDisk) validateValidationChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationChecks) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidationChecks); i++ {
		if swag.IsZero(m.ValidationChecks[i]) { // not required
			continue
		}

		if m.ValidationChecks[i] != nil {
			if err := m.ValidationChecks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this remove disk based on the context it is used
func (m *RemoveDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValidationChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveDisk) contextValidateValidationChecks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ValidationChecks); i++ {

		if m.ValidationChecks[i] != nil {

			if swag.IsZero(m.ValidationChecks[i]) { // not required
				return nil
			}

			if err := m.ValidationChecks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validationChecks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoveDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveDisk) UnmarshalBinary(b []byte) error {
	var res RemoveDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
