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

// CommonPrePostCloudScriptParams Common Pre Post Cloud Script Params
//
// Specifies the common params for PrePost backup scripts specific for cloud-adapters. They have two different scripts for the two different shell types - windows and linux
//
// swagger:model CommonPrePostCloudScriptParams
type CommonPrePostCloudScriptParams struct {

	// Specifies the script details that will be specific to linux machines and executed on bash.
	LinuxScript *CommonPreBackupScriptParams `json:"linuxScript,omitempty"`

	// Specifies the script details that will be specific to windows machines and executed on powershell.
	WindowsScript *CommonPreBackupScriptParams `json:"windowsScript,omitempty"`
}

// Validate validates this common pre post cloud script params
func (m *CommonPrePostCloudScriptParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinuxScript(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWindowsScript(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonPrePostCloudScriptParams) validateLinuxScript(formats strfmt.Registry) error {
	if swag.IsZero(m.LinuxScript) { // not required
		return nil
	}

	if m.LinuxScript != nil {
		if err := m.LinuxScript.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linuxScript")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("linuxScript")
			}
			return err
		}
	}

	return nil
}

func (m *CommonPrePostCloudScriptParams) validateWindowsScript(formats strfmt.Registry) error {
	if swag.IsZero(m.WindowsScript) { // not required
		return nil
	}

	if m.WindowsScript != nil {
		if err := m.WindowsScript.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("windowsScript")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("windowsScript")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this common pre post cloud script params based on the context it is used
func (m *CommonPrePostCloudScriptParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinuxScript(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWindowsScript(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonPrePostCloudScriptParams) contextValidateLinuxScript(ctx context.Context, formats strfmt.Registry) error {

	if m.LinuxScript != nil {

		if swag.IsZero(m.LinuxScript) { // not required
			return nil
		}

		if err := m.LinuxScript.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("linuxScript")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("linuxScript")
			}
			return err
		}
	}

	return nil
}

func (m *CommonPrePostCloudScriptParams) contextValidateWindowsScript(ctx context.Context, formats strfmt.Registry) error {

	if m.WindowsScript != nil {

		if swag.IsZero(m.WindowsScript) { // not required
			return nil
		}

		if err := m.WindowsScript.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("windowsScript")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("windowsScript")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonPrePostCloudScriptParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonPrePostCloudScriptParams) UnmarshalBinary(b []byte) error {
	var res CommonPrePostCloudScriptParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
