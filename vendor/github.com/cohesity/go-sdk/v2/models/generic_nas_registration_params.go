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

// GenericNasRegistrationParams Generic NAS Protection Source.
//
// Specifies parameters to register GenericNas MountPoint.
//
// swagger:model GenericNasRegistrationParams
type GenericNasRegistrationParams struct {

	// Specifies the MountPoint for Generic NAS Source.
	// Required: true
	MountPoint *string `json:"mountPoint"`

	// Specifies the mode of the source.
	// 'kNfs3' indicates NFS3 mode.
	// 'kNfs4_1' indicates NFS4.1 mode.
	// 'kCifs1' indicates SMB mode.
	// Required: true
	// Enum: ["kNfs4_1","kNfs3","kCifs1"]
	Mode *string `json:"mode"`

	// Specifies the Description for Generic NAS Source.
	Description *string `json:"description,omitempty"`

	// Specifies if validation has to be skipped while registering the mount point.
	SkipValidation *bool `json:"skipValidation,omitempty"`

	// Specifies a distinct value that's unique to a source.
	UID *UniversalID `json:"uid,omitempty"`

	// Specifies the credentials to mount a SMB view. Must be specified if protocol is set to SMB.
	SmbMountCredentials *SmbMountCredentials `json:"smbMountCredentials,omitempty"`

	// Specifies the source throttling parameters to be used during registration of the NAS source.
	ThrottlingConfig *NasThrottlingConfig `json:"throttlingConfig,omitempty"`
}

// Validate validates this generic nas registration params
func (m *GenericNasRegistrationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMountPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmbMountCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThrottlingConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericNasRegistrationParams) validateMountPoint(formats strfmt.Registry) error {

	if err := validate.Required("mountPoint", "body", m.MountPoint); err != nil {
		return err
	}

	return nil
}

var genericNasRegistrationParamsTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNfs4_1","kNfs3","kCifs1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		genericNasRegistrationParamsTypeModePropEnum = append(genericNasRegistrationParamsTypeModePropEnum, v)
	}
}

const (

	// GenericNasRegistrationParamsModeKNfs41 captures enum value "kNfs4_1"
	GenericNasRegistrationParamsModeKNfs41 string = "kNfs4_1"

	// GenericNasRegistrationParamsModeKNfs3 captures enum value "kNfs3"
	GenericNasRegistrationParamsModeKNfs3 string = "kNfs3"

	// GenericNasRegistrationParamsModeKCifs1 captures enum value "kCifs1"
	GenericNasRegistrationParamsModeKCifs1 string = "kCifs1"
)

// prop value enum
func (m *GenericNasRegistrationParams) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, genericNasRegistrationParamsTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GenericNasRegistrationParams) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *GenericNasRegistrationParams) validateUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UID) { // not required
		return nil
	}

	if m.UID != nil {
		if err := m.UID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uid")
			}
			return err
		}
	}

	return nil
}

func (m *GenericNasRegistrationParams) validateSmbMountCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.SmbMountCredentials) { // not required
		return nil
	}

	if m.SmbMountCredentials != nil {
		if err := m.SmbMountCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smbMountCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smbMountCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *GenericNasRegistrationParams) validateThrottlingConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ThrottlingConfig) { // not required
		return nil
	}

	if m.ThrottlingConfig != nil {
		if err := m.ThrottlingConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throttlingConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throttlingConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this generic nas registration params based on the context it is used
func (m *GenericNasRegistrationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSmbMountCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThrottlingConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericNasRegistrationParams) contextValidateUID(ctx context.Context, formats strfmt.Registry) error {

	if m.UID != nil {

		if swag.IsZero(m.UID) { // not required
			return nil
		}

		if err := m.UID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uid")
			}
			return err
		}
	}

	return nil
}

func (m *GenericNasRegistrationParams) contextValidateSmbMountCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.SmbMountCredentials != nil {

		if swag.IsZero(m.SmbMountCredentials) { // not required
			return nil
		}

		if err := m.SmbMountCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smbMountCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smbMountCredentials")
			}
			return err
		}
	}

	return nil
}

func (m *GenericNasRegistrationParams) contextValidateThrottlingConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ThrottlingConfig != nil {

		if swag.IsZero(m.ThrottlingConfig) { // not required
			return nil
		}

		if err := m.ThrottlingConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throttlingConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("throttlingConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericNasRegistrationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericNasRegistrationParams) UnmarshalBinary(b []byte) error {
	var res GenericNasRegistrationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
