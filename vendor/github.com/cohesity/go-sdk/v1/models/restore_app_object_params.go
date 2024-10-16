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

// RestoreAppObjectParams Message for application object restore params.
//
// swagger:model RestoreAppObjectParams
type RestoreAppObjectParams struct {

	// The AD specific application object restore params. Only applicable if the
	// RestoreAppObject.app_entity is of type kAD.
	AdRestoreParams *RestoreADAppObjectParams `json:"adRestoreParams,omitempty"`

	// Id of finished clone task which has to be refreshed with different data.
	CloneTaskID *int64 `json:"cloneTaskId,omitempty"`

	// The Exchange specific application object restore params. Only applicable
	// if the RestoreAppObject.app_entity is of type kExchange.
	ExchangeRestoreParams *RestoreExchangeParams `json:"exchangeRestoreParams,omitempty"`

	// The Oracle specific application object restore params. Only applicable if
	// the RestoreAppObject.app_entity is of type kOracle.
	//
	// Note: Only one of sql_restore_params and oracle_restore_params can be set.
	OracleRestoreParams *RestoreOracleAppObjectParams `json:"oracleRestoreParams,omitempty"`

	// The SQL specific application object restore params. Only applicable
	// if the RestoreAppObject.app_entity is of type kSQL.
	SQLRestoreParams *RestoreSQLAppObjectParams `json:"sqlRestoreParams,omitempty"`

	// The target host if the application is to be restored to a different
	// host. If this is empty, then we are restoring to the original host, which
	// is the owner entity.
	TargetHost *EntityProto `json:"targetHost,omitempty"`

	// The registered source managing the target host. If this is empty, then the
	// target host has the same parent source as the owner entity.
	TargetHostParentSource *EntityProto `json:"targetHostParentSource,omitempty"`
}

// Validate validates this restore app object params
func (m *RestoreAppObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdRestoreParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExchangeRestoreParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOracleRestoreParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSQLRestoreParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetHostParentSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreAppObjectParams) validateAdRestoreParams(formats strfmt.Registry) error {
	if swag.IsZero(m.AdRestoreParams) { // not required
		return nil
	}

	if m.AdRestoreParams != nil {
		if err := m.AdRestoreParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adRestoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adRestoreParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) validateExchangeRestoreParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ExchangeRestoreParams) { // not required
		return nil
	}

	if m.ExchangeRestoreParams != nil {
		if err := m.ExchangeRestoreParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exchangeRestoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exchangeRestoreParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) validateOracleRestoreParams(formats strfmt.Registry) error {
	if swag.IsZero(m.OracleRestoreParams) { // not required
		return nil
	}

	if m.OracleRestoreParams != nil {
		if err := m.OracleRestoreParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracleRestoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oracleRestoreParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) validateSQLRestoreParams(formats strfmt.Registry) error {
	if swag.IsZero(m.SQLRestoreParams) { // not required
		return nil
	}

	if m.SQLRestoreParams != nil {
		if err := m.SQLRestoreParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sqlRestoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sqlRestoreParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) validateTargetHost(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetHost) { // not required
		return nil
	}

	if m.TargetHost != nil {
		if err := m.TargetHost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetHost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetHost")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) validateTargetHostParentSource(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetHostParentSource) { // not required
		return nil
	}

	if m.TargetHostParentSource != nil {
		if err := m.TargetHostParentSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetHostParentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetHostParentSource")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this restore app object params based on the context it is used
func (m *RestoreAppObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdRestoreParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExchangeRestoreParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOracleRestoreParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSQLRestoreParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetHostParentSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreAppObjectParams) contextValidateAdRestoreParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AdRestoreParams != nil {

		if swag.IsZero(m.AdRestoreParams) { // not required
			return nil
		}

		if err := m.AdRestoreParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adRestoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adRestoreParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) contextValidateExchangeRestoreParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ExchangeRestoreParams != nil {

		if swag.IsZero(m.ExchangeRestoreParams) { // not required
			return nil
		}

		if err := m.ExchangeRestoreParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exchangeRestoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exchangeRestoreParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) contextValidateOracleRestoreParams(ctx context.Context, formats strfmt.Registry) error {

	if m.OracleRestoreParams != nil {

		if swag.IsZero(m.OracleRestoreParams) { // not required
			return nil
		}

		if err := m.OracleRestoreParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oracleRestoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oracleRestoreParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) contextValidateSQLRestoreParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SQLRestoreParams != nil {

		if swag.IsZero(m.SQLRestoreParams) { // not required
			return nil
		}

		if err := m.SQLRestoreParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sqlRestoreParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sqlRestoreParams")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) contextValidateTargetHost(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetHost != nil {

		if swag.IsZero(m.TargetHost) { // not required
			return nil
		}

		if err := m.TargetHost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetHost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetHost")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreAppObjectParams) contextValidateTargetHostParentSource(ctx context.Context, formats strfmt.Registry) error {

	if m.TargetHostParentSource != nil {

		if swag.IsZero(m.TargetHostParentSource) { // not required
			return nil
		}

		if err := m.TargetHostParentSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("targetHostParentSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("targetHostParentSource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreAppObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreAppObjectParams) UnmarshalBinary(b []byte) error {
	var res RestoreAppObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
