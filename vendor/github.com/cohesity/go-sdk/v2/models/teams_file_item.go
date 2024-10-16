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

// TeamsFileItem ChannelItem
//
// Specifies a M365 Teams channel file item.
//
// swagger:model TeamsFileItem
type TeamsFileItem struct {

	// Specifies the file type.
	// Enum: ["File","Directory","Symlink"]
	FileType *string `json:"fileType,omitempty"`

	// Specifies the size in bytes for the indexed item.
	ItemSize *int64 `json:"itemSize,omitempty"`

	// Specifies the Unix timestamp epoch in seconds at which this item is created.
	CreationTimeSecs *int64 `json:"creationTimeSecs,omitempty"`

	// Specifies the name of the drive location for this file.
	DriveName *string `json:"driveName,omitempty"`
}

// Validate validates this teams file item
func (m *TeamsFileItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var teamsFileItemTypeFileTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["File","Directory","Symlink"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		teamsFileItemTypeFileTypePropEnum = append(teamsFileItemTypeFileTypePropEnum, v)
	}
}

const (

	// TeamsFileItemFileTypeFile captures enum value "File"
	TeamsFileItemFileTypeFile string = "File"

	// TeamsFileItemFileTypeDirectory captures enum value "Directory"
	TeamsFileItemFileTypeDirectory string = "Directory"

	// TeamsFileItemFileTypeSymlink captures enum value "Symlink"
	TeamsFileItemFileTypeSymlink string = "Symlink"
)

// prop value enum
func (m *TeamsFileItem) validateFileTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, teamsFileItemTypeFileTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TeamsFileItem) validateFileType(formats strfmt.Registry) error {
	if swag.IsZero(m.FileType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFileTypeEnum("fileType", "body", *m.FileType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this teams file item based on context it is used
func (m *TeamsFileItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TeamsFileItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamsFileItem) UnmarshalBinary(b []byte) error {
	var res TeamsFileItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
