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

// CloseSmbFileOpenParams Specifies params to close active SMB file open.
//
// swagger:model CloseSmbFileOpenParams
type CloseSmbFileOpenParams struct {

	// Specifies the filepath in the Cohesity View relative to the root filesystem. If this field is specified, viewName field must also be specified.
	// Required: true
	FilePath *string `json:"filePath"`

	// Specifies the name of the Cohesity View in which to search. If a view name is not specified, all the views in the Cluster are searched. This field is mandatory if filePath field is specified.
	// Required: true
	ViewName *string `json:"viewName"`

	// Specifies the id of the active open.
	// Required: true
	OpenID *int64 `json:"openId"`
}

// Validate validates this close smb file open params
func (m *CloseSmbFileOpenParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloseSmbFileOpenParams) validateFilePath(formats strfmt.Registry) error {

	if err := validate.Required("filePath", "body", m.FilePath); err != nil {
		return err
	}

	return nil
}

func (m *CloseSmbFileOpenParams) validateViewName(formats strfmt.Registry) error {

	if err := validate.Required("viewName", "body", m.ViewName); err != nil {
		return err
	}

	return nil
}

func (m *CloseSmbFileOpenParams) validateOpenID(formats strfmt.Registry) error {

	if err := validate.Required("openId", "body", m.OpenID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this close smb file open params based on context it is used
func (m *CloseSmbFileOpenParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloseSmbFileOpenParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloseSmbFileOpenParams) UnmarshalBinary(b []byte) error {
	var res CloseSmbFileOpenParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
