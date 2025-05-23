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

// DeleteInfectedFiles Specifies a list of infected files.
//
// swagger:model DeleteInfectedFiles
type DeleteInfectedFiles struct {

	// Specifies the list of infected files that are successfully deleted.
	DeleteSucceededInfectedFiles []*InfectedFile `json:"deleteSucceededInfectedFiles"`

	// Specifies the list of infected files that failed deletion.
	DeleteFailedInfectedFiles []*InfectedFile `json:"deleteFailedInfectedFiles"`
}

// Validate validates this delete infected files
func (m *DeleteInfectedFiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteSucceededInfectedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteFailedInfectedFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteInfectedFiles) validateDeleteSucceededInfectedFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteSucceededInfectedFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.DeleteSucceededInfectedFiles); i++ {
		if swag.IsZero(m.DeleteSucceededInfectedFiles[i]) { // not required
			continue
		}

		if m.DeleteSucceededInfectedFiles[i] != nil {
			if err := m.DeleteSucceededInfectedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteSucceededInfectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deleteSucceededInfectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeleteInfectedFiles) validateDeleteFailedInfectedFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteFailedInfectedFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.DeleteFailedInfectedFiles); i++ {
		if swag.IsZero(m.DeleteFailedInfectedFiles[i]) { // not required
			continue
		}

		if m.DeleteFailedInfectedFiles[i] != nil {
			if err := m.DeleteFailedInfectedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteFailedInfectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deleteFailedInfectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete infected files based on the context it is used
func (m *DeleteInfectedFiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeleteSucceededInfectedFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteFailedInfectedFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteInfectedFiles) contextValidateDeleteSucceededInfectedFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeleteSucceededInfectedFiles); i++ {

		if m.DeleteSucceededInfectedFiles[i] != nil {

			if swag.IsZero(m.DeleteSucceededInfectedFiles[i]) { // not required
				return nil
			}

			if err := m.DeleteSucceededInfectedFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteSucceededInfectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deleteSucceededInfectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeleteInfectedFiles) contextValidateDeleteFailedInfectedFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeleteFailedInfectedFiles); i++ {

		if m.DeleteFailedInfectedFiles[i] != nil {

			if swag.IsZero(m.DeleteFailedInfectedFiles[i]) { // not required
				return nil
			}

			if err := m.DeleteFailedInfectedFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteFailedInfectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deleteFailedInfectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteInfectedFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteInfectedFiles) UnmarshalBinary(b []byte) error {
	var res DeleteInfectedFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
