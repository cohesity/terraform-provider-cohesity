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

// FileNlmLocks Specifies per-file NLM locks
//
// swagger:model FileNlmLocks
type FileNlmLocks struct {

	// File identitfiers
	FileID *FileID `json:"fileId,omitempty"`

	// Specifies the list of NLM locks in a view.
	NlmLocks []*NlmLock `json:"nlmLocks"`
}

// Validate validates this file nlm locks
func (m *FileNlmLocks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNlmLocks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileNlmLocks) validateFileID(formats strfmt.Registry) error {
	if swag.IsZero(m.FileID) { // not required
		return nil
	}

	if m.FileID != nil {
		if err := m.FileID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileId")
			}
			return err
		}
	}

	return nil
}

func (m *FileNlmLocks) validateNlmLocks(formats strfmt.Registry) error {
	if swag.IsZero(m.NlmLocks) { // not required
		return nil
	}

	for i := 0; i < len(m.NlmLocks); i++ {
		if swag.IsZero(m.NlmLocks[i]) { // not required
			continue
		}

		if m.NlmLocks[i] != nil {
			if err := m.NlmLocks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nlmLocks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nlmLocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this file nlm locks based on the context it is used
func (m *FileNlmLocks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNlmLocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileNlmLocks) contextValidateFileID(ctx context.Context, formats strfmt.Registry) error {

	if m.FileID != nil {

		if swag.IsZero(m.FileID) { // not required
			return nil
		}

		if err := m.FileID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileId")
			}
			return err
		}
	}

	return nil
}

func (m *FileNlmLocks) contextValidateNlmLocks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NlmLocks); i++ {

		if m.NlmLocks[i] != nil {

			if swag.IsZero(m.NlmLocks[i]) { // not required
				return nil
			}

			if err := m.NlmLocks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nlmLocks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nlmLocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileNlmLocks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileNlmLocks) UnmarshalBinary(b []byte) error {
	var res FileNlmLocks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
