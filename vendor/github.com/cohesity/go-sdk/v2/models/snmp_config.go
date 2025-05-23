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

// SnmpConfig SNMP configuration for this cluster.
//
// swagger:model SnmpConfig
type SnmpConfig struct {

	// SnmpVersion is the SNMP version to talk with SNMP agent. It is SNMP V2 or SNMP V3.
	// Enum: ["kSnmpV2","kSnmpV3"]
	Version *string `json:"version,omitempty"`

	// Server is the IP address of Network Management System.
	Server *string `json:"server,omitempty"`

	// AgentPort is the TCP port SNMP agent is using.
	AgentPort *int32 `json:"agentPort,omitempty"`

	// TrapPort is the TCP port SNMP agent is using.
	TrapPort *int32 `json:"trapPort,omitempty"`

	// ReadUser is SNMP read user for SNMP Agent.
	ReadUser *SnmpUser `json:"readUser,omitempty"`

	// WriteUser is SNMP write user for SNMP Agent.
	WriteUser *SnmpUser `json:"writeUser,omitempty"`

	// TrapUser is SNMP trap user for SNMP trap daemon.
	TrapUser *SnmpUser `json:"trapUser,omitempty"`

	// Operation is the operation of configuring SNMP services.
	// Enum: ["kOperationEnable","kOperationDisable"]
	Operation *string `json:"operation,omitempty"`

	// Vip is the IP address SNMP agent and SNMP Trap Daemon will use. It should be one of the VIPs assigned to the cluster.
	Vip *string `json:"vip,omitempty"`

	// SystemInfo contains the information relating to the SNMP system.
	SystemInfo *SnmpSysInfo `json:"systemInfo,omitempty"`
}

// Validate validates this snmp config
func (m *SnmpConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWriteUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrapUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var snmpConfigTypeVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kSnmpV2","kSnmpV3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snmpConfigTypeVersionPropEnum = append(snmpConfigTypeVersionPropEnum, v)
	}
}

const (

	// SnmpConfigVersionKSnmpV2 captures enum value "kSnmpV2"
	SnmpConfigVersionKSnmpV2 string = "kSnmpV2"

	// SnmpConfigVersionKSnmpV3 captures enum value "kSnmpV3"
	SnmpConfigVersionKSnmpV3 string = "kSnmpV3"
)

// prop value enum
func (m *SnmpConfig) validateVersionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snmpConfigTypeVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnmpConfig) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	// value enum
	if err := m.validateVersionEnum("version", "body", *m.Version); err != nil {
		return err
	}

	return nil
}

func (m *SnmpConfig) validateReadUser(formats strfmt.Registry) error {
	if swag.IsZero(m.ReadUser) { // not required
		return nil
	}

	if m.ReadUser != nil {
		if err := m.ReadUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("readUser")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpConfig) validateWriteUser(formats strfmt.Registry) error {
	if swag.IsZero(m.WriteUser) { // not required
		return nil
	}

	if m.WriteUser != nil {
		if err := m.WriteUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writeUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("writeUser")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpConfig) validateTrapUser(formats strfmt.Registry) error {
	if swag.IsZero(m.TrapUser) { // not required
		return nil
	}

	if m.TrapUser != nil {
		if err := m.TrapUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trapUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trapUser")
			}
			return err
		}
	}

	return nil
}

var snmpConfigTypeOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kOperationEnable","kOperationDisable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snmpConfigTypeOperationPropEnum = append(snmpConfigTypeOperationPropEnum, v)
	}
}

const (

	// SnmpConfigOperationKOperationEnable captures enum value "kOperationEnable"
	SnmpConfigOperationKOperationEnable string = "kOperationEnable"

	// SnmpConfigOperationKOperationDisable captures enum value "kOperationDisable"
	SnmpConfigOperationKOperationDisable string = "kOperationDisable"
)

// prop value enum
func (m *SnmpConfig) validateOperationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snmpConfigTypeOperationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnmpConfig) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperationEnum("operation", "body", *m.Operation); err != nil {
		return err
	}

	return nil
}

func (m *SnmpConfig) validateSystemInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemInfo) { // not required
		return nil
	}

	if m.SystemInfo != nil {
		if err := m.SystemInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("systemInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snmp config based on the context it is used
func (m *SnmpConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReadUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWriteUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrapUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpConfig) contextValidateReadUser(ctx context.Context, formats strfmt.Registry) error {

	if m.ReadUser != nil {

		if swag.IsZero(m.ReadUser) { // not required
			return nil
		}

		if err := m.ReadUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("readUser")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpConfig) contextValidateWriteUser(ctx context.Context, formats strfmt.Registry) error {

	if m.WriteUser != nil {

		if swag.IsZero(m.WriteUser) { // not required
			return nil
		}

		if err := m.WriteUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("writeUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("writeUser")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpConfig) contextValidateTrapUser(ctx context.Context, formats strfmt.Registry) error {

	if m.TrapUser != nil {

		if swag.IsZero(m.TrapUser) { // not required
			return nil
		}

		if err := m.TrapUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trapUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trapUser")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpConfig) contextValidateSystemInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemInfo != nil {

		if swag.IsZero(m.SystemInfo) { // not required
			return nil
		}

		if err := m.SystemInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("systemInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpConfig) UnmarshalBinary(b []byte) error {
	var res SnmpConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
