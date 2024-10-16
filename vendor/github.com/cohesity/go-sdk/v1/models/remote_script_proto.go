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

// RemoteScriptProto Message to encapsulate the information of script that can be executed either
// before or after the backup is taken.
//
// swagger:model RemoteScriptProto
type RemoteScriptProto struct {

	// Connector params for the remote host where script is located and is
	// executed.
	RemoteHostParams *RemoteHostConnectorParams `json:"remoteHostParams,omitempty"`

	// Contains script path and its optional params. For AWS, Azure and GCP this
	// will have script details for machines using bash.
	Script *ScriptPathAndParams `json:"script,omitempty"`

	// Execution status of the script.
	Status *ScriptExecutionStatus `json:"status,omitempty"`

	// Contains script path and its optional params. For AWS, Azure and GCP this
	// will have script details for machines using powershell.
	WindowsScript *ScriptPathAndParams `json:"windowsScript,omitempty"`
}

// Validate validates this remote script proto
func (m *RemoteScriptProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemoteHostParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScript(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *RemoteScriptProto) validateRemoteHostParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteHostParams) { // not required
		return nil
	}

	if m.RemoteHostParams != nil {
		if err := m.RemoteHostParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remoteHostParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remoteHostParams")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteScriptProto) validateScript(formats strfmt.Registry) error {
	if swag.IsZero(m.Script) { // not required
		return nil
	}

	if m.Script != nil {
		if err := m.Script.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteScriptProto) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteScriptProto) validateWindowsScript(formats strfmt.Registry) error {
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

// ContextValidate validate this remote script proto based on the context it is used
func (m *RemoteScriptProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRemoteHostParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScript(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
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

func (m *RemoteScriptProto) contextValidateRemoteHostParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RemoteHostParams != nil {

		if swag.IsZero(m.RemoteHostParams) { // not required
			return nil
		}

		if err := m.RemoteHostParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remoteHostParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remoteHostParams")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteScriptProto) contextValidateScript(ctx context.Context, formats strfmt.Registry) error {

	if m.Script != nil {

		if swag.IsZero(m.Script) { // not required
			return nil
		}

		if err := m.Script.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("script")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("script")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteScriptProto) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteScriptProto) contextValidateWindowsScript(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RemoteScriptProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteScriptProto) UnmarshalBinary(b []byte) error {
	var res RemoteScriptProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
