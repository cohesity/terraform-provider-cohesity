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

// DeleteViewUsersQuotaParameters Delete View Users Quota Parameters.
//
// Specifies the user ids to remove the corresponding quota overrides in view.
//
// swagger:model DeleteViewUsersQuotaParameters
type DeleteViewUsersQuotaParameters struct {

	// Delete all existing user quota override policies.
	DeleteAll *bool `json:"deleteAll,omitempty"`

	// The user ids whose policy needs to be deleted.
	UserIds []*UserID `json:"userIds"`

	// View name of input view.
	ViewName *string `json:"viewName,omitempty"`
}

// Validate validates this delete view users quota parameters
func (m *DeleteViewUsersQuotaParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteViewUsersQuotaParameters) validateUserIds(formats strfmt.Registry) error {
	if swag.IsZero(m.UserIds) { // not required
		return nil
	}

	for i := 0; i < len(m.UserIds); i++ {
		if swag.IsZero(m.UserIds[i]) { // not required
			continue
		}

		if m.UserIds[i] != nil {
			if err := m.UserIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userIds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete view users quota parameters based on the context it is used
func (m *DeleteViewUsersQuotaParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteViewUsersQuotaParameters) contextValidateUserIds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserIds); i++ {

		if m.UserIds[i] != nil {

			if swag.IsZero(m.UserIds[i]) { // not required
				return nil
			}

			if err := m.UserIds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userIds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteViewUsersQuotaParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteViewUsersQuotaParameters) UnmarshalBinary(b []byte) error {
	var res DeleteViewUsersQuotaParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
