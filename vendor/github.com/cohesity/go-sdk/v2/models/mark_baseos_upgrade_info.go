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

// MarkBaseosUpgradeInfo Specifies info about BaseOS upgrade operation.
//
// swagger:model MarkBaseosUpgradeInfo
type MarkBaseosUpgradeInfo struct {

	// Specifies whether the operation is set or not.
	// Required: true
	SetOperation *bool `json:"setOperation"`

	// Specifies optional message related to operation status.
	Message *string `json:"message,omitempty"`
}

// Validate validates this mark baseos upgrade info
func (m *MarkBaseosUpgradeInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSetOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarkBaseosUpgradeInfo) validateSetOperation(formats strfmt.Registry) error {

	if err := validate.Required("setOperation", "body", m.SetOperation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mark baseos upgrade info based on context it is used
func (m *MarkBaseosUpgradeInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MarkBaseosUpgradeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarkBaseosUpgradeInfo) UnmarshalBinary(b []byte) error {
	var res MarkBaseosUpgradeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
