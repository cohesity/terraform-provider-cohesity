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

// RecoverPureSanGroupNewSourceConfig Recover Pure group New Source Config.
//
// Specifies the new destination Source configuration where the Pure group will be recovered.
//
// swagger:model RecoverPureSanGroupNewSourceConfig
type RecoverPureSanGroupNewSourceConfig struct {

	// Specifies params to rename the recovered SAN group. If not specified, the original names of the group are preserved.
	RenameRecoveredGroupParams *RecoveredOrClonedVmsRenameConfig `json:"renameRecoveredGroupParams,omitempty"`

	// Specifies the id of the resource pool to recover the Pure SAN Volume to. This field must be specified if recoverToNewSource is true.
	ResourcePool *RecoveryObjectIdentifier `json:"resourcePool,omitempty"`

	// Specifies the id of the new target parent source to recover the Pure SAN group to. This field must be specified if recoverToNewSource is true.
	// Required: true
	Source *RecoveryObjectIdentifier `json:"source"`
}

// Validate validates this recover pure san group new source config
func (m *RecoverPureSanGroupNewSourceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRenameRecoveredGroupParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcePool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverPureSanGroupNewSourceConfig) validateRenameRecoveredGroupParams(formats strfmt.Registry) error {
	if swag.IsZero(m.RenameRecoveredGroupParams) { // not required
		return nil
	}

	if m.RenameRecoveredGroupParams != nil {
		if err := m.RenameRecoveredGroupParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("renameRecoveredGroupParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("renameRecoveredGroupParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverPureSanGroupNewSourceConfig) validateResourcePool(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourcePool) { // not required
		return nil
	}

	if m.ResourcePool != nil {
		if err := m.ResourcePool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourcePool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourcePool")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverPureSanGroupNewSourceConfig) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recover pure san group new source config based on the context it is used
func (m *RecoverPureSanGroupNewSourceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRenameRecoveredGroupParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourcePool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverPureSanGroupNewSourceConfig) contextValidateRenameRecoveredGroupParams(ctx context.Context, formats strfmt.Registry) error {

	if m.RenameRecoveredGroupParams != nil {

		if swag.IsZero(m.RenameRecoveredGroupParams) { // not required
			return nil
		}

		if err := m.RenameRecoveredGroupParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("renameRecoveredGroupParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("renameRecoveredGroupParams")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverPureSanGroupNewSourceConfig) contextValidateResourcePool(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourcePool != nil {

		if swag.IsZero(m.ResourcePool) { // not required
			return nil
		}

		if err := m.ResourcePool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourcePool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourcePool")
			}
			return err
		}
	}

	return nil
}

func (m *RecoverPureSanGroupNewSourceConfig) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoverPureSanGroupNewSourceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverPureSanGroupNewSourceConfig) UnmarshalBinary(b []byte) error {
	var res RecoverPureSanGroupNewSourceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
