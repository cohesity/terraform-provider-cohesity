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

// AcropolisProtectionGroupObjectParams Specifies an object protected by a Acropolis Protection Group.
//
// swagger:model AcropolisProtectionGroupObjectParams
type AcropolisProtectionGroupObjectParams struct {

	// Specifies the ID of the object.
	// Required: true
	ID *int64 `json:"id"`

	// Specifies the name of the virtual machine.
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Specifies a list of disks to exclude from being protected. This is only applicable to VM objects.
	ExcludeDisks []*AcropolisDiskInfo `json:"excludeDisks"`

	// Specifies a list of disks to include in the protection. This is only applicable to VM objects.
	IncludeDisks []*AcropolisDiskInfo `json:"includeDisks"`
}

// Validate validates this acropolis protection group object params
func (m *AcropolisProtectionGroupObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludeDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncludeDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcropolisProtectionGroupObjectParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AcropolisProtectionGroupObjectParams) validateExcludeDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.ExcludeDisks); i++ {
		if swag.IsZero(m.ExcludeDisks[i]) { // not required
			continue
		}

		if m.ExcludeDisks[i] != nil {
			if err := m.ExcludeDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excludeDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excludeDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AcropolisProtectionGroupObjectParams) validateIncludeDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.IncludeDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.IncludeDisks); i++ {
		if swag.IsZero(m.IncludeDisks[i]) { // not required
			continue
		}

		if m.IncludeDisks[i] != nil {
			if err := m.IncludeDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("includeDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("includeDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this acropolis protection group object params based on the context it is used
func (m *AcropolisProtectionGroupObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExcludeDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncludeDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcropolisProtectionGroupObjectParams) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AcropolisProtectionGroupObjectParams) contextValidateExcludeDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExcludeDisks); i++ {

		if m.ExcludeDisks[i] != nil {

			if swag.IsZero(m.ExcludeDisks[i]) { // not required
				return nil
			}

			if err := m.ExcludeDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("excludeDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("excludeDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AcropolisProtectionGroupObjectParams) contextValidateIncludeDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncludeDisks); i++ {

		if m.IncludeDisks[i] != nil {

			if swag.IsZero(m.IncludeDisks[i]) { // not required
				return nil
			}

			if err := m.IncludeDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("includeDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("includeDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AcropolisProtectionGroupObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcropolisProtectionGroupObjectParams) UnmarshalBinary(b []byte) error {
	var res AcropolisProtectionGroupObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
