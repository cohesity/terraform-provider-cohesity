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

// OutlookBackupParams Specifies Outlook job parameters applicable for all Office365 Environment type Protection Sources in a Protection Job.
//
// swagger:model OutlookBackupParams
type OutlookBackupParams struct {

	// should backup mailbox
	ShouldBackupMailbox *bool `json:"shouldBackupMailbox,omitempty"`

	// Specifies filters on the backup objects like files and directories. Specifying filters decide which objects within a source should be backed up. If this field is not specified, then all of the objects within the source will be backed up.
	FilePathFilter *FileFilteringPolicy `json:"filePathFilter,omitempty"`
}

// Validate validates this outlook backup params
func (m *OutlookBackupParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilePathFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutlookBackupParams) validateFilePathFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.FilePathFilter) { // not required
		return nil
	}

	if m.FilePathFilter != nil {
		if err := m.FilePathFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filePathFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filePathFilter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this outlook backup params based on the context it is used
func (m *OutlookBackupParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilePathFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutlookBackupParams) contextValidateFilePathFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.FilePathFilter != nil {

		if swag.IsZero(m.FilePathFilter) { // not required
			return nil
		}

		if err := m.FilePathFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filePathFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filePathFilter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutlookBackupParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutlookBackupParams) UnmarshalBinary(b []byte) error {
	var res OutlookBackupParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
