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

// AntivirusServiceGroup Antivirus Group Config.
//
// Specifies the configuration settings for an Antivirus service group.
//
// swagger:model AntivirusServiceGroup
type AntivirusServiceGroup struct {

	// Specifies the Antivirus Services belonging to this antivirus group.
	AntivirusServices []*AntivirusServiceConfig `json:"antivirusServices"`

	// Specifies the description of the Antivirus service group.
	Description *string `json:"description,omitempty"`

	// Specifies the Id of the Antivirus service group.
	// Required: true
	ID *int64 `json:"id"`

	// Specifies whether the antivirus service group is enabled or not.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Specifies the name of the Antivirus service group.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this antivirus service group
func (m *AntivirusServiceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAntivirusServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntivirusServiceGroup) validateAntivirusServices(formats strfmt.Registry) error {
	if swag.IsZero(m.AntivirusServices) { // not required
		return nil
	}

	for i := 0; i < len(m.AntivirusServices); i++ {
		if swag.IsZero(m.AntivirusServices[i]) { // not required
			continue
		}

		if m.AntivirusServices[i] != nil {
			if err := m.AntivirusServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("antivirusServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("antivirusServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AntivirusServiceGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AntivirusServiceGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this antivirus service group based on the context it is used
func (m *AntivirusServiceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAntivirusServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntivirusServiceGroup) contextValidateAntivirusServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AntivirusServices); i++ {

		if m.AntivirusServices[i] != nil {

			if swag.IsZero(m.AntivirusServices[i]) { // not required
				return nil
			}

			if err := m.AntivirusServices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("antivirusServices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("antivirusServices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntivirusServiceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntivirusServiceGroup) UnmarshalBinary(b []byte) error {
	var res AntivirusServiceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
