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

// O365ConnectParams Office 365 Connect Params.
//
// Specifies an Object containing information about a registered Office 365
// source.
//
// swagger:model O365ConnectParams
type O365ConnectParams struct {

	// Specifies the parameters used for selectively discovering the office 365
	// objects during source registration or refresh.
	ObjectsDiscoveryParams *ObjectsDiscoveryParams `json:"ObjectsDiscoveryParams,omitempty"`

	// Specifies the settings for backups through M365 Backup Storage APIs.
	CsmParams *M365CsmParams `json:"csmParams,omitempty"`
}

// Validate validates this o365 connect params
func (m *O365ConnectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectsDiscoveryParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsmParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *O365ConnectParams) validateObjectsDiscoveryParams(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectsDiscoveryParams) { // not required
		return nil
	}

	if m.ObjectsDiscoveryParams != nil {
		if err := m.ObjectsDiscoveryParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ObjectsDiscoveryParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ObjectsDiscoveryParams")
			}
			return err
		}
	}

	return nil
}

func (m *O365ConnectParams) validateCsmParams(formats strfmt.Registry) error {
	if swag.IsZero(m.CsmParams) { // not required
		return nil
	}

	if m.CsmParams != nil {
		if err := m.CsmParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csmParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csmParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o365 connect params based on the context it is used
func (m *O365ConnectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectsDiscoveryParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCsmParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *O365ConnectParams) contextValidateObjectsDiscoveryParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectsDiscoveryParams != nil {

		if swag.IsZero(m.ObjectsDiscoveryParams) { // not required
			return nil
		}

		if err := m.ObjectsDiscoveryParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ObjectsDiscoveryParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ObjectsDiscoveryParams")
			}
			return err
		}
	}

	return nil
}

func (m *O365ConnectParams) contextValidateCsmParams(ctx context.Context, formats strfmt.Registry) error {

	if m.CsmParams != nil {

		if swag.IsZero(m.CsmParams) { // not required
			return nil
		}

		if err := m.CsmParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csmParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csmParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *O365ConnectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *O365ConnectParams) UnmarshalBinary(b []byte) error {
	var res O365ConnectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
