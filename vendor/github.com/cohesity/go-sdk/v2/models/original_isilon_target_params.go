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

// OriginalIsilonTargetParams Recover To Original Isilon Volume Target Params.
//
// Specifies the params of the original Isilon recovery target.
//
// swagger:model OriginalIsilonTargetParams
type OriginalIsilonTargetParams struct {

	// Specifies whether to overwrite existing file/folder during recovery.
	OverwriteExistingFile *bool `json:"overwriteExistingFile,omitempty"`

	// Specifies whether to preserve file/folder attributes during recovery.
	PreserveFileAttributes *bool `json:"preserveFileAttributes,omitempty"`

	// Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// Specifies whether encryption should be enabled during recovery.
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty"`

	// Specifies the list of IP addresses that are allowed or denied during recovery. Allowed IPs and Denied IPs cannot be used together.
	FilterIPConfig *FilterIPConfig `json:"filterIpConfig,omitempty"`

	// Specifies VLAN settings associated with the restore. If this is not specified, then the VLAN params will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for restores.
	VlanConfig *RecoveryVlanConfig `json:"vlanConfig,omitempty"`
}

// Validate validates this original isilon target params
func (m *OriginalIsilonTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterIPConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OriginalIsilonTargetParams) validateFilterIPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterIPConfig) { // not required
		return nil
	}

	if m.FilterIPConfig != nil {
		if err := m.FilterIPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filterIpConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filterIpConfig")
			}
			return err
		}
	}

	return nil
}

func (m *OriginalIsilonTargetParams) validateVlanConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanConfig) { // not required
		return nil
	}

	if m.VlanConfig != nil {
		if err := m.VlanConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this original isilon target params based on the context it is used
func (m *OriginalIsilonTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilterIPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlanConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OriginalIsilonTargetParams) contextValidateFilterIPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.FilterIPConfig != nil {

		if swag.IsZero(m.FilterIPConfig) { // not required
			return nil
		}

		if err := m.FilterIPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filterIpConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filterIpConfig")
			}
			return err
		}
	}

	return nil
}

func (m *OriginalIsilonTargetParams) contextValidateVlanConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.VlanConfig != nil {

		if swag.IsZero(m.VlanConfig) { // not required
			return nil
		}

		if err := m.VlanConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlanConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OriginalIsilonTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OriginalIsilonTargetParams) UnmarshalBinary(b []byte) error {
	var res OriginalIsilonTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
