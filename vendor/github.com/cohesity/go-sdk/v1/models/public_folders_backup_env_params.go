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

// PublicFoldersBackupEnvParams O365 PublicFolders Backup Environment Parameters.
//
// Message to capture any additional backup params for Public Folders within
// Office365 environment.
//
// swagger:model PublicFoldersBackupEnvParams
type PublicFoldersBackupEnvParams struct {

	// The filtering policy describes which paths within a Public Folder should
	// be excluded within the backup. If this is not specified, then the entire
	// Public Folders will be backed up.
	FilteringPolicy *FilteringPolicyProto `json:"filteringPolicy,omitempty"`
}

// Validate validates this public folders backup env params
func (m *PublicFoldersBackupEnvParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilteringPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicFoldersBackupEnvParams) validateFilteringPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.FilteringPolicy) { // not required
		return nil
	}

	if m.FilteringPolicy != nil {
		if err := m.FilteringPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filteringPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filteringPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this public folders backup env params based on the context it is used
func (m *PublicFoldersBackupEnvParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilteringPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicFoldersBackupEnvParams) contextValidateFilteringPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.FilteringPolicy != nil {

		if swag.IsZero(m.FilteringPolicy) { // not required
			return nil
		}

		if err := m.FilteringPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filteringPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filteringPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicFoldersBackupEnvParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicFoldersBackupEnvParams) UnmarshalBinary(b []byte) error {
	var res PublicFoldersBackupEnvParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
