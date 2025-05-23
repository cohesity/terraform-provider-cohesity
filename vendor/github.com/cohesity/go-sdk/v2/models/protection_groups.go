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

// ProtectionGroups Protection Group  response.
//
// swagger:model ProtectionGroups
type ProtectionGroups struct {

	// Specifies the list of Protection Groups which were returned by the request.
	ProtectionGroups []*ProtectionGroup `json:"protectionGroups"`

	// Specifies the information needed in order to support pagination. This will not be included for the last page of results.
	PaginationCookie *string `json:"paginationCookie,omitempty"`
}

// Validate validates this protection groups
func (m *ProtectionGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtectionGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionGroups) validateProtectionGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ProtectionGroups); i++ {
		if swag.IsZero(m.ProtectionGroups[i]) { // not required
			continue
		}

		if m.ProtectionGroups[i] != nil {
			if err := m.ProtectionGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this protection groups based on the context it is used
func (m *ProtectionGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtectionGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionGroups) contextValidateProtectionGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProtectionGroups); i++ {

		if m.ProtectionGroups[i] != nil {

			if swag.IsZero(m.ProtectionGroups[i]) { // not required
				return nil
			}

			if err := m.ProtectionGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionGroups) UnmarshalBinary(b []byte) error {
	var res ProtectionGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
