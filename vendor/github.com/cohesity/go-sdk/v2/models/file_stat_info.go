// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FileStatInfo filestatInfo is the stat information for the file.
//
// swagger:model FileStatInfo
type FileStatInfo struct {

	// If this is a file, the size of the file as returned by stat. This is Bytes
	Size *int64 `json:"size,omitempty"`

	// If this is a file, the last modified time as returned by stat.
	ModifiedTimeUsecs *int64 `json:"modifiedTimeUsecs,omitempty"`

	// Specifies the file type.
	// Enum: ["File","Directory","Symlink"]
	Type *string `json:"type,omitempty"`

	// Source inode id metadata for certain adapters e.g. Netapp.
	BackupSourceInodeID *int64 `json:"backupSourceInodeId,omitempty"`

	// Metadata for the sharepoint item in browse. This will be set only if we are quering for sharepoint site. This is used to further different between item types, for example a kDirectory item in sharepoint site could be a Document library (kSiteDoclib), list (kSiteList) or folder (kDirectory).
	SharePointItemMetadata *SharepointItemMetadata `json:"sharePointItemMetadata,omitempty"`
}

// Validate validates this file stat info
func (m *FileStatInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharePointItemMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fileStatInfoTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["File","Directory","Symlink"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileStatInfoTypeTypePropEnum = append(fileStatInfoTypeTypePropEnum, v)
	}
}

const (

	// FileStatInfoTypeFile captures enum value "File"
	FileStatInfoTypeFile string = "File"

	// FileStatInfoTypeDirectory captures enum value "Directory"
	FileStatInfoTypeDirectory string = "Directory"

	// FileStatInfoTypeSymlink captures enum value "Symlink"
	FileStatInfoTypeSymlink string = "Symlink"
)

// prop value enum
func (m *FileStatInfo) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileStatInfoTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileStatInfo) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *FileStatInfo) validateSharePointItemMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.SharePointItemMetadata) { // not required
		return nil
	}

	if m.SharePointItemMetadata != nil {
		if err := m.SharePointItemMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharePointItemMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharePointItemMetadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file stat info based on the context it is used
func (m *FileStatInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSharePointItemMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileStatInfo) contextValidateSharePointItemMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.SharePointItemMetadata != nil {

		if swag.IsZero(m.SharePointItemMetadata) { // not required
			return nil
		}

		if err := m.SharePointItemMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharePointItemMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharePointItemMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileStatInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileStatInfo) UnmarshalBinary(b []byte) error {
	var res FileStatInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
