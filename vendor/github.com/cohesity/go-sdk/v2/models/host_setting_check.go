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

// HostSettingCheck Specifies the host checking details.
//
// swagger:model HostSettingCheck
type HostSettingCheck struct {

	// Specifies the type of host checking that was performed.
	// Enum: ["IsAgentPortAccessible","IsAgentRunning","IsSQLWriterRunning","AreSQLInstancesRunning","CheckServiceLoginsConfig","CheckSQLFCIVIP","CheckSQLDiskSpace"]
	Type *string `json:"type,omitempty"`

	// Specifies the result of host checking performed by agent.
	// Enum: ["Pass","Fail","Warning"]
	Result *string `json:"result,omitempty"`
}

// Validate validates this host setting check
func (m *HostSettingCheck) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hostSettingCheckTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IsAgentPortAccessible","IsAgentRunning","IsSQLWriterRunning","AreSQLInstancesRunning","CheckServiceLoginsConfig","CheckSQLFCIVIP","CheckSQLDiskSpace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostSettingCheckTypeTypePropEnum = append(hostSettingCheckTypeTypePropEnum, v)
	}
}

const (

	// HostSettingCheckTypeIsAgentPortAccessible captures enum value "IsAgentPortAccessible"
	HostSettingCheckTypeIsAgentPortAccessible string = "IsAgentPortAccessible"

	// HostSettingCheckTypeIsAgentRunning captures enum value "IsAgentRunning"
	HostSettingCheckTypeIsAgentRunning string = "IsAgentRunning"

	// HostSettingCheckTypeIsSQLWriterRunning captures enum value "IsSQLWriterRunning"
	HostSettingCheckTypeIsSQLWriterRunning string = "IsSQLWriterRunning"

	// HostSettingCheckTypeAreSQLInstancesRunning captures enum value "AreSQLInstancesRunning"
	HostSettingCheckTypeAreSQLInstancesRunning string = "AreSQLInstancesRunning"

	// HostSettingCheckTypeCheckServiceLoginsConfig captures enum value "CheckServiceLoginsConfig"
	HostSettingCheckTypeCheckServiceLoginsConfig string = "CheckServiceLoginsConfig"

	// HostSettingCheckTypeCheckSQLFCIVIP captures enum value "CheckSQLFCIVIP"
	HostSettingCheckTypeCheckSQLFCIVIP string = "CheckSQLFCIVIP"

	// HostSettingCheckTypeCheckSQLDiskSpace captures enum value "CheckSQLDiskSpace"
	HostSettingCheckTypeCheckSQLDiskSpace string = "CheckSQLDiskSpace"
)

// prop value enum
func (m *HostSettingCheck) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hostSettingCheckTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HostSettingCheck) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

var hostSettingCheckTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pass","Fail","Warning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostSettingCheckTypeResultPropEnum = append(hostSettingCheckTypeResultPropEnum, v)
	}
}

const (

	// HostSettingCheckResultPass captures enum value "Pass"
	HostSettingCheckResultPass string = "Pass"

	// HostSettingCheckResultFail captures enum value "Fail"
	HostSettingCheckResultFail string = "Fail"

	// HostSettingCheckResultWarning captures enum value "Warning"
	HostSettingCheckResultWarning string = "Warning"
)

// prop value enum
func (m *HostSettingCheck) validateResultEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hostSettingCheckTypeResultPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HostSettingCheck) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	// value enum
	if err := m.validateResultEnum("result", "body", *m.Result); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this host setting check based on context it is used
func (m *HostSettingCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostSettingCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostSettingCheck) UnmarshalBinary(b []byte) error {
	var res HostSettingCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
