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

// GetTenantStatsResult GetTenantStatsResult
//
// GetTenantStatsResult is the result of get tenantStats api.
//
// swagger:model GetTenantStatsResult
type GetTenantStatsResult struct {

	// Specifies an opaque string to pass to get the next set of active opens.
	// If null is returned, this response is the last set of active opens.
	Cookie *string `json:"cookie,omitempty"`

	// Specifies a list of tenant stats.
	StatsList []*TenantStats `json:"statsList"`
}

// Validate validates this get tenant stats result
func (m *GetTenantStatsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatsList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTenantStatsResult) validateStatsList(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsList) { // not required
		return nil
	}

	for i := 0; i < len(m.StatsList); i++ {
		if swag.IsZero(m.StatsList[i]) { // not required
			continue
		}

		if m.StatsList[i] != nil {
			if err := m.StatsList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get tenant stats result based on the context it is used
func (m *GetTenantStatsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatsList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTenantStatsResult) contextValidateStatsList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatsList); i++ {

		if m.StatsList[i] != nil {

			if swag.IsZero(m.StatsList[i]) { // not required
				return nil
			}

			if err := m.StatsList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTenantStatsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTenantStatsResult) UnmarshalBinary(b []byte) error {
	var res GetTenantStatsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
