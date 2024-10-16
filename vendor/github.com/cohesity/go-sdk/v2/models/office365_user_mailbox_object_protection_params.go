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

// Office365UserMailboxObjectProtectionParams Specifies the params to create a User Mailbox Object Protection.
//
// swagger:model Office365UserMailboxObjectProtectionParams
type Office365UserMailboxObjectProtectionParams struct {
	Office365ObjectProtectionCommonParams

	// Array of prefixes used to exclude folders which are by default included. Two kinds of filters are supported. a) prefix which always starts with '/'. b) posix which always starts with empty quotes(''). Regular expressions are not supported. If not specified, all folders which are included by default will be included. These prefixes have no effect on folders that are excluded by default. The only folders excluded by default are documented with includeFolders.
	ExcludeFolders []string `json:"excludeFolders"`

	// Array of prefixes used to include folders which are by default excluded. Two kinds of filters are supported. a) prefix which always starts with '/'. b) posix which always starts with empty quotes(''). Regular expressions are not supported. If not specified, all folders which are excluded by default will be excluded. These prefixes have no effect on folders that are included by default. All folders are included by default except for the Recoverable Items folder.
	IncludeFolders []string `json:"includeFolders"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Office365UserMailboxObjectProtectionParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Office365ObjectProtectionCommonParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Office365ObjectProtectionCommonParams = aO0

	// AO1
	var dataAO1 struct {
		ExcludeFolders []string `json:"excludeFolders"`

		IncludeFolders []string `json:"includeFolders"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ExcludeFolders = dataAO1.ExcludeFolders

	m.IncludeFolders = dataAO1.IncludeFolders

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Office365UserMailboxObjectProtectionParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Office365ObjectProtectionCommonParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ExcludeFolders []string `json:"excludeFolders"`

		IncludeFolders []string `json:"includeFolders"`
	}

	dataAO1.ExcludeFolders = m.ExcludeFolders

	dataAO1.IncludeFolders = m.IncludeFolders

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this office365 user mailbox object protection params
func (m *Office365UserMailboxObjectProtectionParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Office365ObjectProtectionCommonParams
	if err := m.Office365ObjectProtectionCommonParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this office365 user mailbox object protection params based on the context it is used
func (m *Office365UserMailboxObjectProtectionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Office365ObjectProtectionCommonParams
	if err := m.Office365ObjectProtectionCommonParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Office365UserMailboxObjectProtectionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Office365UserMailboxObjectProtectionParams) UnmarshalBinary(b []byte) error {
	var res Office365UserMailboxObjectProtectionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
