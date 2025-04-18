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

// CapacityTrendDatapoint A data point for the capacity trend analysis chart.
//
// swagger:model CapacityTrendDatapoint
type CapacityTrendDatapoint struct {

	// Specifies an array of tag and its corresponding statistic.
	DataPointStats []*CapacityTrendDatapointStats `json:"dataPointStats"`

	// Specifies the timestamp of this data point.
	DataPointTimestamp int64 `json:"dataPointTimestamp,omitempty"`

	// Specifies error messages, if any error is encountered while fetching the datapoint stats.
	ErrorMessages []string `json:"errorMessages"`
}

// Validate validates this capacity trend datapoint
func (m *CapacityTrendDatapoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataPointStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityTrendDatapoint) validateDataPointStats(formats strfmt.Registry) error {
	if swag.IsZero(m.DataPointStats) { // not required
		return nil
	}

	for i := 0; i < len(m.DataPointStats); i++ {
		if swag.IsZero(m.DataPointStats[i]) { // not required
			continue
		}

		if m.DataPointStats[i] != nil {
			if err := m.DataPointStats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataPointStats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataPointStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this capacity trend datapoint based on the context it is used
func (m *CapacityTrendDatapoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataPointStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityTrendDatapoint) contextValidateDataPointStats(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DataPointStats); i++ {

		if m.DataPointStats[i] != nil {

			if swag.IsZero(m.DataPointStats[i]) { // not required
				return nil
			}

			if err := m.DataPointStats[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataPointStats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataPointStats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapacityTrendDatapoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapacityTrendDatapoint) UnmarshalBinary(b []byte) error {
	var res CapacityTrendDatapoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
