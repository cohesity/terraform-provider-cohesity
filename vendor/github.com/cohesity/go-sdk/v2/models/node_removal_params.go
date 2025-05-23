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

// NodeRemovalParams Node Removal Parameters.
//
// Specifies parameters to initiate/cancel node removal.
//
// swagger:model NodeRemovalParams
type NodeRemovalParams struct {

	// If true, cancels node removal that is already in progress.
	// Required: true
	Cancel *bool `json:"cancel"`

	// Specifies whether node being removed is offline.
	IsOffline *bool `json:"isOffline,omitempty"`

	// Specifies whether request is for pre-check validations only
	IsValidateOnly *bool `json:"isValidateOnly,omitempty"`

	// Specifies whether request is for clearing pre-check result only
	IsClearPreCheckResultOnly *bool `json:"isClearPreCheckResultOnly,omitempty"`
}

// Validate validates this node removal params
func (m *NodeRemovalParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCancel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeRemovalParams) validateCancel(formats strfmt.Registry) error {

	if err := validate.Required("cancel", "body", m.Cancel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this node removal params based on context it is used
func (m *NodeRemovalParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeRemovalParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeRemovalParams) UnmarshalBinary(b []byte) error {
	var res NodeRemovalParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
