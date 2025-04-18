// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FilesAndFoldersObject Specifies a file or folder to download.
//
// swagger:model FilesAndFoldersObject
type FilesAndFoldersObject struct {

	// Specifies the absolute path of the file or folder.
	// Required: true
	AbsolutePath *string `json:"absolutePath"`

	// Specifies whether the file or folder object is a directory.
	IsDirectory *bool `json:"isDirectory,omitempty"`
}

// Validate validates this files and folders object
func (m *FilesAndFoldersObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbsolutePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilesAndFoldersObject) validateAbsolutePath(formats strfmt.Registry) error {

	if err := validate.Required("absolutePath", "body", m.AbsolutePath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this files and folders object based on context it is used
func (m *FilesAndFoldersObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FilesAndFoldersObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilesAndFoldersObject) UnmarshalBinary(b []byte) error {
	var res FilesAndFoldersObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
