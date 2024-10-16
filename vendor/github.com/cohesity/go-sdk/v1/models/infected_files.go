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

// InfectedFiles Infected files.
//
// Specifies the Result parameters for all infected files.
//
// swagger:model InfectedFiles
type InfectedFiles struct {

	// Specifies the infected files.
	InfectedFiles []*InfectedFile `json:"infectedFiles"`

	// This cookie can be used in the succeeding call to list infected files to
	// get the next set of infected files. If set to nil, it means that there's
	// no more results that the server could provide.
	PaginationCookie *string `json:"paginationCookie,omitempty"`
}

// Validate validates this infected files
func (m *InfectedFiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfectedFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfectedFiles) validateInfectedFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.InfectedFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.InfectedFiles); i++ {
		if swag.IsZero(m.InfectedFiles[i]) { // not required
			continue
		}

		if m.InfectedFiles[i] != nil {
			if err := m.InfectedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("infectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("infectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this infected files based on the context it is used
func (m *InfectedFiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfectedFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InfectedFiles) contextValidateInfectedFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InfectedFiles); i++ {

		if m.InfectedFiles[i] != nil {

			if swag.IsZero(m.InfectedFiles[i]) { // not required
				return nil
			}

			if err := m.InfectedFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("infectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("infectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfectedFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfectedFiles) UnmarshalBinary(b []byte) error {
	var res InfectedFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
