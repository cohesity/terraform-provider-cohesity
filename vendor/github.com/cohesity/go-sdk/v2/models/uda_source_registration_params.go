// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UdaSourceRegistrationParams Register Universal Data Adapter source registration request parameters.
//
// Specifies parameters to register a Universal Data Adapter source.
//
// swagger:model UdaSourceRegistrationParams
type UdaSourceRegistrationParams struct {

	// Specifies the source type for Universal Data Adapter source.
	// Required: true
	SourceType *string `json:"sourceType"`

	// Specifies the OS type for Universal Data Adapter source.
	// Enum: ["kLinux","kWindows","kAix"]
	OsType *string `json:"osType,omitempty"`

	// Specifies the IPs/hostnames for the nodes forming the Universal Data Adapter source cluster.
	// Required: true
	// Min Items: 1
	// Unique: true
	Hosts []string `json:"hosts"`

	// Specifies the absolute path of scripts used to interact with the Universal Data Adapter source.
	// Required: true
	ScriptDir *string `json:"scriptDir"`

	// Specifies if SMB/NFS view mounting should be enabled on source. Default value is false.
	MountView *bool `json:"mountView,omitempty"`

	// Specifies custom arguments to be supplied to the source registration scripts. This field is deprecated. Use sourceRegistrationArguments instead.
	SourceRegistrationArgs *string `json:"sourceRegistrationArgs,omitempty"`

	// Specifies the map of custom arguments to be supplied to the source registration scripts.
	SourceRegistrationArguments []*KeyValuePair `json:"sourceRegistrationArguments"`

	// credentials
	Credentials *UdaSourceRegistrationParamsCredentials `json:"credentials,omitempty"`

	// view params
	ViewParams *UdaSourceRegistrationParamsViewParams `json:"viewParams,omitempty"`
}

// Validate validates this uda source registration params
func (m *UdaSourceRegistrationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScriptDir(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceRegistrationArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaSourceRegistrationParams) validateSourceType(formats strfmt.Registry) error {

	if err := validate.Required("sourceType", "body", m.SourceType); err != nil {
		return err
	}

	return nil
}

var udaSourceRegistrationParamsTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kLinux","kWindows","kAix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		udaSourceRegistrationParamsTypeOsTypePropEnum = append(udaSourceRegistrationParamsTypeOsTypePropEnum, v)
	}
}

const (

	// UdaSourceRegistrationParamsOsTypeKLinux captures enum value "kLinux"
	UdaSourceRegistrationParamsOsTypeKLinux string = "kLinux"

	// UdaSourceRegistrationParamsOsTypeKWindows captures enum value "kWindows"
	UdaSourceRegistrationParamsOsTypeKWindows string = "kWindows"

	// UdaSourceRegistrationParamsOsTypeKAix captures enum value "kAix"
	UdaSourceRegistrationParamsOsTypeKAix string = "kAix"
)

// prop value enum
func (m *UdaSourceRegistrationParams) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, udaSourceRegistrationParamsTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UdaSourceRegistrationParams) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("osType", "body", *m.OsType); err != nil {
		return err
	}

	return nil
}

func (m *UdaSourceRegistrationParams) validateHosts(formats strfmt.Registry) error {

	if err := validate.Required("hosts", "body", m.Hosts); err != nil {
		return err
	}

	iHostsSize := int64(len(m.Hosts))

	if err := validate.MinItems("hosts", "body", iHostsSize, 1); err != nil {
		return err
	}

	if err := validate.UniqueItems("hosts", "body", m.Hosts); err != nil {
		return err
	}

	return nil
}

func (m *UdaSourceRegistrationParams) validateScriptDir(formats strfmt.Registry) error {

	if err := validate.Required("scriptDir", "body", m.ScriptDir); err != nil {
		return err
	}

	return nil
}

func (m *UdaSourceRegistrationParams) validateSourceRegistrationArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceRegistrationArguments) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceRegistrationArguments); i++ {
		if swag.IsZero(m.SourceRegistrationArguments[i]) { // not required
			continue
		}

		if m.SourceRegistrationArguments[i] != nil {
			if err := m.SourceRegistrationArguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceRegistrationArguments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sourceRegistrationArguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UdaSourceRegistrationParams) validateCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *UdaSourceRegistrationParams) validateViewParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ViewParams) { // not required
		return nil
	}

	if m.ViewParams != nil {
		if err := m.ViewParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viewParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("viewParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this uda source registration params based on the context it is used
func (m *UdaSourceRegistrationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceRegistrationArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UdaSourceRegistrationParams) contextValidateSourceRegistrationArguments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SourceRegistrationArguments); i++ {

		if m.SourceRegistrationArguments[i] != nil {

			if swag.IsZero(m.SourceRegistrationArguments[i]) { // not required
				return nil
			}

			if err := m.SourceRegistrationArguments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceRegistrationArguments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sourceRegistrationArguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UdaSourceRegistrationParams) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {

		if swag.IsZero(m.Credentials) { // not required
			return nil
		}

		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *UdaSourceRegistrationParams) contextValidateViewParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ViewParams != nil {

		if swag.IsZero(m.ViewParams) { // not required
			return nil
		}

		if err := m.ViewParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viewParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("viewParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UdaSourceRegistrationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaSourceRegistrationParams) UnmarshalBinary(b []byte) error {
	var res UdaSourceRegistrationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UdaSourceRegistrationParamsCredentials Specifies credentials to access the Universal Data Adapter source. For e.g.: To perform backup and recovery tasks with Oracle Recovery Manager (RMAN), specify credentials for a user having 'SYSDBA' or 'SYSBACKUP' administrative privilege.
//
// swagger:model UdaSourceRegistrationParamsCredentials
type UdaSourceRegistrationParamsCredentials struct {

	// Specifies the password to access target entity.
	Password *string `json:"password,omitempty"`

	// Specifies the username to access target entity.
	Username *string `json:"username,omitempty"`
}

// Validate validates this uda source registration params credentials
func (m *UdaSourceRegistrationParamsCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this uda source registration params credentials based on context it is used
func (m *UdaSourceRegistrationParamsCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UdaSourceRegistrationParamsCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaSourceRegistrationParamsCredentials) UnmarshalBinary(b []byte) error {
	var res UdaSourceRegistrationParamsCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UdaSourceRegistrationParamsViewParams Specifies optional configuration parameters for the mounted view.
//
// swagger:model UdaSourceRegistrationParamsViewParams
type UdaSourceRegistrationParamsViewParams struct {

	// This field is deprecated and its value will be ignored. It was used to specify the absolute path on the host where the Cohesity view would be mounted. This is now controlled by the agent configuration and is common for all the Universal Data Adapter sources.
	MountDir *string `json:"mountDir,omitempty"`
}

// Validate validates this uda source registration params view params
func (m *UdaSourceRegistrationParamsViewParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this uda source registration params view params based on context it is used
func (m *UdaSourceRegistrationParamsViewParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UdaSourceRegistrationParamsViewParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UdaSourceRegistrationParamsViewParams) UnmarshalBinary(b []byte) error {
	var res UdaSourceRegistrationParamsViewParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
