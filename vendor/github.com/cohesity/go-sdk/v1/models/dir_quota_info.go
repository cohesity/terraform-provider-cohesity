// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DirQuotaInfo Directory Quota Info.
//
// Specifies the configuration and policy details for directory quota in a
// view.
//
// swagger:model DirQuotaInfo
type DirQuotaInfo struct {

	// Specifies the directory quota configuration.
	Config *DirQuotaConfig `json:"config,omitempty"`

	// This cookie can be used in the succeeding call to list user quotas and
	// usages to get the next set of user quota overrides. If set to nil, it
	// means that there's no more results that the server could provide.
	Cookie *int64 `json:"cookie,omitempty"`

	// Specifies the list of directory quota policies applied on the view.
	Quotas []*DirQuotaPolicy `json:"quotas"`
}

// Validate validates this dir quota info
func (m *DirQuotaInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuotas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirQuotaInfo) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *DirQuotaInfo) validateQuotas(formats strfmt.Registry) error {
	if swag.IsZero(m.Quotas) { // not required
		return nil
	}

	for i := 0; i < len(m.Quotas); i++ {
		if swag.IsZero(m.Quotas[i]) { // not required
			continue
		}

		if m.Quotas[i] != nil {
			if err := m.Quotas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quotas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("quotas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dir quota info based on the context it is used
func (m *DirQuotaInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuotas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirQuotaInfo) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if swag.IsZero(m.Config) { // not required
			return nil
		}

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *DirQuotaInfo) contextValidateQuotas(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Quotas); i++ {

		if m.Quotas[i] != nil {

			if swag.IsZero(m.Quotas[i]) { // not required
				return nil
			}

			if err := m.Quotas[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quotas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("quotas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirQuotaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirQuotaInfo) UnmarshalBinary(b []byte) error {
	var res DirQuotaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
