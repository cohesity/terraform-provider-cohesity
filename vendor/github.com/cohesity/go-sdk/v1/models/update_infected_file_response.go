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

// UpdateInfectedFileResponse Specifies update infected files response.
//
// swagger:model UpdateInfectedFileResponse
type UpdateInfectedFileResponse struct {

	// Specifies the failed update infected files.
	UpdateFailedInfectedFiles []*InfectedFileID `json:"updateFailedInfectedFiles"`

	// Specifies the successfully updated infected files.
	UpdateSucceededInfectedFiles []*InfectedFileID `json:"updateSucceededInfectedFiles"`
}

// Validate validates this update infected file response
func (m *UpdateInfectedFileResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdateFailedInfectedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateSucceededInfectedFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateInfectedFileResponse) validateUpdateFailedInfectedFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateFailedInfectedFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.UpdateFailedInfectedFiles); i++ {
		if swag.IsZero(m.UpdateFailedInfectedFiles[i]) { // not required
			continue
		}

		if m.UpdateFailedInfectedFiles[i] != nil {
			if err := m.UpdateFailedInfectedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateFailedInfectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateFailedInfectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateInfectedFileResponse) validateUpdateSucceededInfectedFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateSucceededInfectedFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.UpdateSucceededInfectedFiles); i++ {
		if swag.IsZero(m.UpdateSucceededInfectedFiles[i]) { // not required
			continue
		}

		if m.UpdateSucceededInfectedFiles[i] != nil {
			if err := m.UpdateSucceededInfectedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateSucceededInfectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateSucceededInfectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update infected file response based on the context it is used
func (m *UpdateInfectedFileResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUpdateFailedInfectedFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateSucceededInfectedFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateInfectedFileResponse) contextValidateUpdateFailedInfectedFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UpdateFailedInfectedFiles); i++ {

		if m.UpdateFailedInfectedFiles[i] != nil {

			if swag.IsZero(m.UpdateFailedInfectedFiles[i]) { // not required
				return nil
			}

			if err := m.UpdateFailedInfectedFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateFailedInfectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateFailedInfectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateInfectedFileResponse) contextValidateUpdateSucceededInfectedFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UpdateSucceededInfectedFiles); i++ {

		if m.UpdateSucceededInfectedFiles[i] != nil {

			if swag.IsZero(m.UpdateSucceededInfectedFiles[i]) { // not required
				return nil
			}

			if err := m.UpdateSucceededInfectedFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateSucceededInfectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateSucceededInfectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateInfectedFileResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateInfectedFileResponse) UnmarshalBinary(b []byte) error {
	var res UpdateInfectedFileResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
