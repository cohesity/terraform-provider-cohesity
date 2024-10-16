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

// ExchangeProtectionSource Exchange Protection Source.
//
// Specifies an object representing an Exchange entity.
// DAG - Database availability group
//
// swagger:model ExchangeProtectionSource
type ExchangeProtectionSource struct {

	// Specifies the Exchange DAG information if ExchangeProtectionSourceType is
	// 'kExchangeDAG'.
	DagInfo *DagInfo `json:"DagInfo,omitempty"`

	// Specifies the Exchange Application server information if
	// ExchangeProtectionSourceType is 'kExchangeNode'.
	ApplicationServerInfo *ApplicationServerInfo `json:"applicationServerInfo,omitempty"`

	// Specifies the Exchange DAG Database copy information if
	// ExchangeProtectionSourceType is 'kExchangeDAGDatabaseCopy'.
	DagDatabaseCopyInfo *ExchangeDatabaseCopyInfo `json:"dagDatabaseCopyInfo,omitempty"`

	// Specifies the Exchange DAG Database information if
	// ExchangeProtectionSourceType is 'kExchangeDAGDatabase'
	DagDatabaseInfo *ExchangeDAGDatabase `json:"dagDatabaseInfo,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// Specifies the entity id of the owner of the Exchange Protection Source.
	OwnerID *int64 `json:"ownerId,omitempty"`

	// Specifies the Exchange Standalone Database information if
	// ExchangeProtectionSourceType is 'kExchangeStandaloneDatabase'.
	StandaloneDatabaseCopyInfo *ExchangeDatabaseInfo `json:"standaloneDatabaseCopyInfo,omitempty"`

	// Specifies the type of the Exchange Protection Source.
	Type *int32 `json:"type,omitempty"`

	// Specifies the UUID for the Exchange entity.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this exchange protection source
func (m *ExchangeProtectionSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDagInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationServerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDagDatabaseCopyInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDagDatabaseInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandaloneDatabaseCopyInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeProtectionSource) validateDagInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DagInfo) { // not required
		return nil
	}

	if m.DagInfo != nil {
		if err := m.DagInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DagInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DagInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeProtectionSource) validateApplicationServerInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationServerInfo) { // not required
		return nil
	}

	if m.ApplicationServerInfo != nil {
		if err := m.ApplicationServerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationServerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationServerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeProtectionSource) validateDagDatabaseCopyInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DagDatabaseCopyInfo) { // not required
		return nil
	}

	if m.DagDatabaseCopyInfo != nil {
		if err := m.DagDatabaseCopyInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dagDatabaseCopyInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dagDatabaseCopyInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeProtectionSource) validateDagDatabaseInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DagDatabaseInfo) { // not required
		return nil
	}

	if m.DagDatabaseInfo != nil {
		if err := m.DagDatabaseInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dagDatabaseInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dagDatabaseInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeProtectionSource) validateStandaloneDatabaseCopyInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.StandaloneDatabaseCopyInfo) { // not required
		return nil
	}

	if m.StandaloneDatabaseCopyInfo != nil {
		if err := m.StandaloneDatabaseCopyInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standaloneDatabaseCopyInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standaloneDatabaseCopyInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this exchange protection source based on the context it is used
func (m *ExchangeProtectionSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDagInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplicationServerInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDagDatabaseCopyInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDagDatabaseInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandaloneDatabaseCopyInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeProtectionSource) contextValidateDagInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DagInfo != nil {

		if swag.IsZero(m.DagInfo) { // not required
			return nil
		}

		if err := m.DagInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DagInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DagInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeProtectionSource) contextValidateApplicationServerInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ApplicationServerInfo != nil {

		if swag.IsZero(m.ApplicationServerInfo) { // not required
			return nil
		}

		if err := m.ApplicationServerInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("applicationServerInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("applicationServerInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeProtectionSource) contextValidateDagDatabaseCopyInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DagDatabaseCopyInfo != nil {

		if swag.IsZero(m.DagDatabaseCopyInfo) { // not required
			return nil
		}

		if err := m.DagDatabaseCopyInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dagDatabaseCopyInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dagDatabaseCopyInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeProtectionSource) contextValidateDagDatabaseInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DagDatabaseInfo != nil {

		if swag.IsZero(m.DagDatabaseInfo) { // not required
			return nil
		}

		if err := m.DagDatabaseInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dagDatabaseInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dagDatabaseInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ExchangeProtectionSource) contextValidateStandaloneDatabaseCopyInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.StandaloneDatabaseCopyInfo != nil {

		if swag.IsZero(m.StandaloneDatabaseCopyInfo) { // not required
			return nil
		}

		if err := m.StandaloneDatabaseCopyInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standaloneDatabaseCopyInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standaloneDatabaseCopyInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeProtectionSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeProtectionSource) UnmarshalBinary(b []byte) error {
	var res ExchangeProtectionSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
