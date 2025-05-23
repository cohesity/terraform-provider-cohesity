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

// PstParam Specifies the PST conversion specific parameters.
//
// swagger:model PstParam
type PstParam struct {

	// Specifies Password to be set for generated PSTs.
	// Required: true
	Password *string `json:"password"`

	// Specifies if create a PST or MSG for input items.
	CreatePst *bool `json:"createPst,omitempty"`

	// Specifies PST size threshold in bytes.
	SizeThresholdBytes *int64 `json:"sizeThresholdBytes,omitempty"`

	// If true, a separate download file will be made for each snapshot. If false, a single common download file will be used for all snapshots. Default is false.
	SeparateDownloadFiles *bool `json:"separateDownloadFiles,omitempty"`
}

// Validate validates this pst param
func (m *PstParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PstParam) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pst param based on context it is used
func (m *PstParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PstParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PstParam) UnmarshalBinary(b []byte) error {
	var res PstParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
