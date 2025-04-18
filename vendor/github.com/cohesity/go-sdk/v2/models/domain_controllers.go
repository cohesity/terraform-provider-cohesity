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

// DomainControllers Specifies the domain controllers of a domain.
//
// swagger:model DomainControllers
type DomainControllers struct {

	// Specifies the domain name.
	DomainName *string `json:"domainName,omitempty"`

	// Specifies a list of domain controllers of the domain.
	Controllers []*DomainController `json:"controllers"`
}

// Validate validates this domain controllers
func (m *DomainControllers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControllers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainControllers) validateControllers(formats strfmt.Registry) error {
	if swag.IsZero(m.Controllers) { // not required
		return nil
	}

	for i := 0; i < len(m.Controllers); i++ {
		if swag.IsZero(m.Controllers[i]) { // not required
			continue
		}

		if m.Controllers[i] != nil {
			if err := m.Controllers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this domain controllers based on the context it is used
func (m *DomainControllers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControllers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainControllers) contextValidateControllers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Controllers); i++ {

		if m.Controllers[i] != nil {

			if swag.IsZero(m.Controllers[i]) { // not required
				return nil
			}

			if err := m.Controllers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("controllers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("controllers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainControllers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainControllers) UnmarshalBinary(b []byte) error {
	var res DomainControllers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
