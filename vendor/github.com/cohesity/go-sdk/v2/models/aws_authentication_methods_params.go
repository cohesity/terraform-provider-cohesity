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

// AwsAuthenticationMethodsParams AWS External Targets Authentication Methods Params.
//
// Specifies the Authentication Methods parameters which are specific to AWS related External Targets.
//
// swagger:model AwsAuthenticationMethodsParams
type AwsAuthenticationMethodsParams struct {
	CommonAuthenticationMethodParams

	// i am user params
	IAmUserParams *AwsIAmUserParams `json:"iAmUserParams,omitempty"`

	// i am role params
	IAmRoleParams *AwsIAmRoleParams `json:"iAmRoleParams,omitempty"`

	// i am roles anywhere params
	IAmRolesAnywhereParams *AwsIAmRolesAnywhereParams `json:"iAmRolesAnywhereParams,omitempty"`

	// use s t s params
	UseSTSParams *AwsUseSTSParams `json:"useSTSParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AwsAuthenticationMethodsParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonAuthenticationMethodParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonAuthenticationMethodParams = aO0

	// AO1
	var dataAO1 struct {
		IAmUserParams *AwsIAmUserParams `json:"iAmUserParams,omitempty"`

		IAmRoleParams *AwsIAmRoleParams `json:"iAmRoleParams,omitempty"`

		IAmRolesAnywhereParams *AwsIAmRolesAnywhereParams `json:"iAmRolesAnywhereParams,omitempty"`

		UseSTSParams *AwsUseSTSParams `json:"useSTSParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.IAmUserParams = dataAO1.IAmUserParams

	m.IAmRoleParams = dataAO1.IAmRoleParams

	m.IAmRolesAnywhereParams = dataAO1.IAmRolesAnywhereParams

	m.UseSTSParams = dataAO1.UseSTSParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AwsAuthenticationMethodsParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonAuthenticationMethodParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		IAmUserParams *AwsIAmUserParams `json:"iAmUserParams,omitempty"`

		IAmRoleParams *AwsIAmRoleParams `json:"iAmRoleParams,omitempty"`

		IAmRolesAnywhereParams *AwsIAmRolesAnywhereParams `json:"iAmRolesAnywhereParams,omitempty"`

		UseSTSParams *AwsUseSTSParams `json:"useSTSParams,omitempty"`
	}

	dataAO1.IAmUserParams = m.IAmUserParams

	dataAO1.IAmRoleParams = m.IAmRoleParams

	dataAO1.IAmRolesAnywhereParams = m.IAmRolesAnywhereParams

	dataAO1.UseSTSParams = m.UseSTSParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this aws authentication methods params
func (m *AwsAuthenticationMethodsParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonAuthenticationMethodParams
	if err := m.CommonAuthenticationMethodParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIAmUserParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIAmRoleParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIAmRolesAnywhereParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseSTSParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsAuthenticationMethodsParams) validateIAmUserParams(formats strfmt.Registry) error {

	if swag.IsZero(m.IAmUserParams) { // not required
		return nil
	}

	if m.IAmUserParams != nil {
		if err := m.IAmUserParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iAmUserParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iAmUserParams")
			}
			return err
		}
	}

	return nil
}

func (m *AwsAuthenticationMethodsParams) validateIAmRoleParams(formats strfmt.Registry) error {

	if swag.IsZero(m.IAmRoleParams) { // not required
		return nil
	}

	if m.IAmRoleParams != nil {
		if err := m.IAmRoleParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iAmRoleParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iAmRoleParams")
			}
			return err
		}
	}

	return nil
}

func (m *AwsAuthenticationMethodsParams) validateIAmRolesAnywhereParams(formats strfmt.Registry) error {

	if swag.IsZero(m.IAmRolesAnywhereParams) { // not required
		return nil
	}

	if m.IAmRolesAnywhereParams != nil {
		if err := m.IAmRolesAnywhereParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iAmRolesAnywhereParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iAmRolesAnywhereParams")
			}
			return err
		}
	}

	return nil
}

func (m *AwsAuthenticationMethodsParams) validateUseSTSParams(formats strfmt.Registry) error {

	if swag.IsZero(m.UseSTSParams) { // not required
		return nil
	}

	if m.UseSTSParams != nil {
		if err := m.UseSTSParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("useSTSParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("useSTSParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this aws authentication methods params based on the context it is used
func (m *AwsAuthenticationMethodsParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonAuthenticationMethodParams
	if err := m.CommonAuthenticationMethodParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIAmUserParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIAmRoleParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIAmRolesAnywhereParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUseSTSParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsAuthenticationMethodsParams) contextValidateIAmUserParams(ctx context.Context, formats strfmt.Registry) error {

	if m.IAmUserParams != nil {

		if swag.IsZero(m.IAmUserParams) { // not required
			return nil
		}

		if err := m.IAmUserParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iAmUserParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iAmUserParams")
			}
			return err
		}
	}

	return nil
}

func (m *AwsAuthenticationMethodsParams) contextValidateIAmRoleParams(ctx context.Context, formats strfmt.Registry) error {

	if m.IAmRoleParams != nil {

		if swag.IsZero(m.IAmRoleParams) { // not required
			return nil
		}

		if err := m.IAmRoleParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iAmRoleParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iAmRoleParams")
			}
			return err
		}
	}

	return nil
}

func (m *AwsAuthenticationMethodsParams) contextValidateIAmRolesAnywhereParams(ctx context.Context, formats strfmt.Registry) error {

	if m.IAmRolesAnywhereParams != nil {

		if swag.IsZero(m.IAmRolesAnywhereParams) { // not required
			return nil
		}

		if err := m.IAmRolesAnywhereParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iAmRolesAnywhereParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iAmRolesAnywhereParams")
			}
			return err
		}
	}

	return nil
}

func (m *AwsAuthenticationMethodsParams) contextValidateUseSTSParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UseSTSParams != nil {

		if swag.IsZero(m.UseSTSParams) { // not required
			return nil
		}

		if err := m.UseSTSParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("useSTSParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("useSTSParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsAuthenticationMethodsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsAuthenticationMethodsParams) UnmarshalBinary(b []byte) error {
	var res AwsAuthenticationMethodsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
