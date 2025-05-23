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

// ViewClientsSummary Specifies the View Client summary.
//
// swagger:model ViewClientsSummary
type ViewClientsSummary struct {

	// Specifies a list of summary.
	Summary []*ViewClientsSummaryInfo `json:"summary"`
}

// Validate validates this view clients summary
func (m *ViewClientsSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewClientsSummary) validateSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	for i := 0; i < len(m.Summary); i++ {
		if swag.IsZero(m.Summary[i]) { // not required
			continue
		}

		if m.Summary[i] != nil {
			if err := m.Summary[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summary" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("summary" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this view clients summary based on the context it is used
func (m *ViewClientsSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewClientsSummary) contextValidateSummary(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Summary); i++ {

		if m.Summary[i] != nil {

			if swag.IsZero(m.Summary[i]) { // not required
				return nil
			}

			if err := m.Summary[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summary" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("summary" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ViewClientsSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewClientsSummary) UnmarshalBinary(b []byte) error {
	var res ViewClientsSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
