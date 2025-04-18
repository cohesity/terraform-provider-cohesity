// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataTieringAnalysisShareInfo Specifies the info for a particular share.
//
// swagger:model DataTieringAnalysisShareInfo
type DataTieringAnalysisShareInfo struct {

	// Specifies the id of the share.
	// Required: true
	ShareID *int64 `json:"shareId"`
}

// Validate validates this data tiering analysis share info
func (m *DataTieringAnalysisShareInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShareID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTieringAnalysisShareInfo) validateShareID(formats strfmt.Registry) error {

	if err := validate.Required("shareId", "body", m.ShareID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this data tiering analysis share info based on context it is used
func (m *DataTieringAnalysisShareInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataTieringAnalysisShareInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTieringAnalysisShareInfo) UnmarshalBinary(b []byte) error {
	var res DataTieringAnalysisShareInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
