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

// TimeSeriesSchemaResponse Time series schema response.
//
// Specifies the Apollo schema to list the data point.
//
// swagger:model TimeSeriesSchemaResponse
type TimeSeriesSchemaResponse struct {

	// Specifies the list of the schema info for an entity.
	SchemaInfoList []*SchemaInfo `json:"schemaInfoList"`
}

// Validate validates this time series schema response
func (m *TimeSeriesSchemaResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchemaInfoList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeSeriesSchemaResponse) validateSchemaInfoList(formats strfmt.Registry) error {
	if swag.IsZero(m.SchemaInfoList) { // not required
		return nil
	}

	for i := 0; i < len(m.SchemaInfoList); i++ {
		if swag.IsZero(m.SchemaInfoList[i]) { // not required
			continue
		}

		if m.SchemaInfoList[i] != nil {
			if err := m.SchemaInfoList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schemaInfoList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schemaInfoList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this time series schema response based on the context it is used
func (m *TimeSeriesSchemaResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchemaInfoList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeSeriesSchemaResponse) contextValidateSchemaInfoList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SchemaInfoList); i++ {

		if m.SchemaInfoList[i] != nil {

			if swag.IsZero(m.SchemaInfoList[i]) { // not required
				return nil
			}

			if err := m.SchemaInfoList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schemaInfoList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("schemaInfoList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeSeriesSchemaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeSeriesSchemaResponse) UnmarshalBinary(b []byte) error {
	var res TimeSeriesSchemaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
