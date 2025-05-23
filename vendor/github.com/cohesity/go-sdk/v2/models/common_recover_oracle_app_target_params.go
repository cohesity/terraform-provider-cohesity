// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CommonRecoverOracleAppTargetParams Recover Oracle App Snapshot Params.
//
// Specifies the snapshot parameters to recover Oracle databases.
//
// swagger:model CommonRecoverOracleAppTargetParams
type CommonRecoverOracleAppTargetParams struct {

	// Specifies the parameter whether the recovery should be performed to a new source or an original Source Target.
	// Required: true
	RecoverToNewSource *bool `json:"recoverToNewSource"`

	// Specifies the destination Source configuration parameters where the databases will be recovered. This is mandatory if recoverToNewSource is set to true.
	NewSourceConfig *RecoverOracleAppNewSourceConfig `json:"newSourceConfig,omitempty"`

	// Specifies the Source configuration if databases are being recovered to Original Source. If not specified, all the configuration parameters will be retained.
	OriginalSourceConfig *RecoverOracleAppOriginalSourceConfig `json:"originalSourceConfig,omitempty"`
}

// Validate validates this common recover oracle app target params
func (m *CommonRecoverOracleAppTargetParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoverToNewSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewSourceConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalSourceConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonRecoverOracleAppTargetParams) validateRecoverToNewSource(formats strfmt.Registry) error {

	if err := validate.Required("recoverToNewSource", "body", m.RecoverToNewSource); err != nil {
		return err
	}

	return nil
}

func (m *CommonRecoverOracleAppTargetParams) validateNewSourceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NewSourceConfig) { // not required
		return nil
	}

	if m.NewSourceConfig != nil {
		if err := m.NewSourceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newSourceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CommonRecoverOracleAppTargetParams) validateOriginalSourceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalSourceConfig) { // not required
		return nil
	}

	if m.OriginalSourceConfig != nil {
		if err := m.OriginalSourceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originalSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originalSourceConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this common recover oracle app target params based on the context it is used
func (m *CommonRecoverOracleAppTargetParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNewSourceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginalSourceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonRecoverOracleAppTargetParams) contextValidateNewSourceConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NewSourceConfig != nil {

		if swag.IsZero(m.NewSourceConfig) { // not required
			return nil
		}

		if err := m.NewSourceConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newSourceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CommonRecoverOracleAppTargetParams) contextValidateOriginalSourceConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginalSourceConfig != nil {

		if swag.IsZero(m.OriginalSourceConfig) { // not required
			return nil
		}

		if err := m.OriginalSourceConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originalSourceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originalSourceConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonRecoverOracleAppTargetParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonRecoverOracleAppTargetParams) UnmarshalBinary(b []byte) error {
	var res CommonRecoverOracleAppTargetParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
