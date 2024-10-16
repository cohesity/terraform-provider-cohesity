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

// GetViewBoxStatsResult GetViewBoxStatsResult
//
// GetViewBoxStatsResult is the result of get viewBoxStats api.
//
// swagger:model GetViewBoxStatsResult
type GetViewBoxStatsResult struct {

	// Specifies a list of view box stats.
	StatsList []*StorageDomainStats `json:"statsList"`
}

// Validate validates this get view box stats result
func (m *GetViewBoxStatsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatsList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetViewBoxStatsResult) validateStatsList(formats strfmt.Registry) error {
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

// ContextValidate validate this get view box stats result based on the context it is used
func (m *GetViewBoxStatsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatsList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetViewBoxStatsResult) contextValidateStatsList(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetViewBoxStatsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetViewBoxStatsResult) UnmarshalBinary(b []byte) error {
	var res GetViewBoxStatsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
