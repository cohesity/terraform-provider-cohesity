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

// AwsTargetParams Defines the Aws target related information required for recovery.
//
// swagger:model AwsTargetParams
type AwsTargetParams struct {

	// Custom destination Server configuration parameters where the RDS
	// Postgres database will be recovered."
	CustomServerConfig *AwsTargetParamsNetworkConfig `json:"customServerConfig,omitempty"`

	// If set to true means we are recovering to a know source and
	// 'known_source_config' will be populated else 'custom_server_config' will
	// be populated.
	IsKnownSource *bool `json:"isKnownSource,omitempty"`

	// Populated in case of a known target.
	KnownSourceConfig *AwsTargetParamsNetworkConfig `json:"knownSourceConfig,omitempty"`
}

// Validate validates this aws target params
func (m *AwsTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomServerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKnownSourceConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsTargetParams) validateCustomServerConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomServerConfig) { // not required
		return nil
	}

	if m.CustomServerConfig != nil {
		if err := m.CustomServerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customServerConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customServerConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AwsTargetParams) validateKnownSourceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.KnownSourceConfig) { // not required
		return nil
	}

	if m.KnownSourceConfig != nil {
		if err := m.KnownSourceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("knownSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("knownSourceConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this aws target params based on the context it is used
func (m *AwsTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomServerConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKnownSourceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsTargetParams) contextValidateCustomServerConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomServerConfig != nil {

		if swag.IsZero(m.CustomServerConfig) { // not required
			return nil
		}

		if err := m.CustomServerConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customServerConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customServerConfig")
			}
			return err
		}
	}

	return nil
}

func (m *AwsTargetParams) contextValidateKnownSourceConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.KnownSourceConfig != nil {

		if swag.IsZero(m.KnownSourceConfig) { // not required
			return nil
		}

		if err := m.KnownSourceConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("knownSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("knownSourceConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsTargetParams) UnmarshalBinary(b []byte) error {
	var res AwsTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
