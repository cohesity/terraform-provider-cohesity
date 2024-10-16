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
	"github.com/go-openapi/validate"
)

// UpdateProtectionObjectParameters Update Protection Object Parameters
//
// Specifies the parameters to update a Protection Object.
//
// swagger:model UpdateProtectionObjectParameters
type UpdateProtectionObjectParameters struct {

	// Specifies if the protection for the Protection Object is to be paused.
	PauseBackup *bool `json:"pauseBackup,omitempty"`

	// Specifies the unique id of the Protected Source to be updated.
	// Required: true
	ProtectedSourceUID *UniversalID `json:"protectedSourceUid"`

	// Specifies the unique id of the new RPO policy to assign to the object.
	RpoPolicyID *string `json:"rpoPolicyId,omitempty"`

	// Specifies the additional special settings for a Protected Source.
	SourceParameters []*SourceSpecialParameter `json:"sourceParameters"`
}

// Validate validates this update protection object parameters
func (m *UpdateProtectionObjectParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtectedSourceUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateProtectionObjectParameters) validateProtectedSourceUID(formats strfmt.Registry) error {

	if err := validate.Required("protectedSourceUid", "body", m.ProtectedSourceUID); err != nil {
		return err
	}

	if m.ProtectedSourceUID != nil {
		if err := m.ProtectedSourceUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protectedSourceUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protectedSourceUid")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateProtectionObjectParameters) validateSourceParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceParameters); i++ {
		if swag.IsZero(m.SourceParameters[i]) { // not required
			continue
		}

		if m.SourceParameters[i] != nil {
			if err := m.SourceParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sourceParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update protection object parameters based on the context it is used
func (m *UpdateProtectionObjectParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtectedSourceUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateProtectionObjectParameters) contextValidateProtectedSourceUID(ctx context.Context, formats strfmt.Registry) error {

	if m.ProtectedSourceUID != nil {

		if err := m.ProtectedSourceUID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protectedSourceUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protectedSourceUid")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateProtectionObjectParameters) contextValidateSourceParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SourceParameters); i++ {

		if m.SourceParameters[i] != nil {

			if swag.IsZero(m.SourceParameters[i]) { // not required
				return nil
			}

			if err := m.SourceParameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sourceParameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sourceParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProtectionObjectParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProtectionObjectParameters) UnmarshalBinary(b []byte) error {
	var res UpdateProtectionObjectParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
