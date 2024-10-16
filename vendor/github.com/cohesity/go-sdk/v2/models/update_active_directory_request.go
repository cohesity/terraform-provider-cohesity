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

// UpdateActiveDirectoryRequest Specifies the request to create an Active Directory.
//
// swagger:model UpdateActiveDirectoryRequest
type UpdateActiveDirectoryRequest struct {
	CommonActiveDirectoryParams

	// Specifies if specified machine accounts should overwrite existing machine accounts.
	OverwriteMachineAccounts *bool `json:"overwriteMachineAccounts,omitempty"`

	// Specifies level of transitive Active Directory trust domains to be used.
	TransitiveAdTrustLevelLimit *int32 `json:"transitiveAdTrustLevelLimit,omitempty"`

	// Specifies the params of a user with administrative privilege of this Active Directory. This field is mandatory if machine accounts are updated.
	ActiveDirectoryAdminParams *ActiveDirectoryAdminParams `json:"activeDirectoryAdminParams,omitempty"`

	// Specifies the params of the user id mapping info of an Active Directory.
	IDMappingParams *IDMappingParams `json:"idMappingParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UpdateActiveDirectoryRequest) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonActiveDirectoryParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonActiveDirectoryParams = aO0

	// AO1
	var dataAO1 struct {
		OverwriteMachineAccounts *bool `json:"overwriteMachineAccounts,omitempty"`

		TransitiveAdTrustLevelLimit *int32 `json:"transitiveAdTrustLevelLimit,omitempty"`

		ActiveDirectoryAdminParams *ActiveDirectoryAdminParams `json:"activeDirectoryAdminParams,omitempty"`

		IDMappingParams *IDMappingParams `json:"idMappingParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.OverwriteMachineAccounts = dataAO1.OverwriteMachineAccounts

	m.TransitiveAdTrustLevelLimit = dataAO1.TransitiveAdTrustLevelLimit

	m.ActiveDirectoryAdminParams = dataAO1.ActiveDirectoryAdminParams

	m.IDMappingParams = dataAO1.IDMappingParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UpdateActiveDirectoryRequest) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonActiveDirectoryParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		OverwriteMachineAccounts *bool `json:"overwriteMachineAccounts,omitempty"`

		TransitiveAdTrustLevelLimit *int32 `json:"transitiveAdTrustLevelLimit,omitempty"`

		ActiveDirectoryAdminParams *ActiveDirectoryAdminParams `json:"activeDirectoryAdminParams,omitempty"`

		IDMappingParams *IDMappingParams `json:"idMappingParams,omitempty"`
	}

	dataAO1.OverwriteMachineAccounts = m.OverwriteMachineAccounts

	dataAO1.TransitiveAdTrustLevelLimit = m.TransitiveAdTrustLevelLimit

	dataAO1.ActiveDirectoryAdminParams = m.ActiveDirectoryAdminParams

	dataAO1.IDMappingParams = m.IDMappingParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this update active directory request
func (m *UpdateActiveDirectoryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonActiveDirectoryParams
	if err := m.CommonActiveDirectoryParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActiveDirectoryAdminParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDMappingParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateActiveDirectoryRequest) validateActiveDirectoryAdminParams(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveDirectoryAdminParams) { // not required
		return nil
	}

	if m.ActiveDirectoryAdminParams != nil {
		if err := m.ActiveDirectoryAdminParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activeDirectoryAdminParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activeDirectoryAdminParams")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateActiveDirectoryRequest) validateIDMappingParams(formats strfmt.Registry) error {

	if swag.IsZero(m.IDMappingParams) { // not required
		return nil
	}

	if m.IDMappingParams != nil {
		if err := m.IDMappingParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idMappingParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("idMappingParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update active directory request based on the context it is used
func (m *UpdateActiveDirectoryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonActiveDirectoryParams
	if err := m.CommonActiveDirectoryParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActiveDirectoryAdminParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIDMappingParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateActiveDirectoryRequest) contextValidateActiveDirectoryAdminParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ActiveDirectoryAdminParams != nil {

		if swag.IsZero(m.ActiveDirectoryAdminParams) { // not required
			return nil
		}

		if err := m.ActiveDirectoryAdminParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activeDirectoryAdminParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activeDirectoryAdminParams")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateActiveDirectoryRequest) contextValidateIDMappingParams(ctx context.Context, formats strfmt.Registry) error {

	if m.IDMappingParams != nil {

		if swag.IsZero(m.IDMappingParams) { // not required
			return nil
		}

		if err := m.IDMappingParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idMappingParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("idMappingParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateActiveDirectoryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateActiveDirectoryRequest) UnmarshalBinary(b []byte) error {
	var res UpdateActiveDirectoryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
