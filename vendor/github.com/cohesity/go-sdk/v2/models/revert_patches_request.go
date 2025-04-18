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

// RevertPatchesRequest Revert patch.
//
// Specifies the request to revert a patch. An optional patch level can be specified. When a patch level is specified, system keeps reverting patches until the given patch level is reverted. If no patch level is specified, just the last applied patch is reverted.
//
// swagger:model RevertPatchesRequest
type RevertPatchesRequest struct {

	// Specifies the name of the service whose patch should be reverted.
	// Required: true
	Service *string `json:"service"`
}

// Validate validates this revert patches request
func (m *RevertPatchesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RevertPatchesRequest) validateService(formats strfmt.Registry) error {

	if err := validate.Required("service", "body", m.Service); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this revert patches request based on context it is used
func (m *RevertPatchesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RevertPatchesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RevertPatchesRequest) UnmarshalBinary(b []byte) error {
	var res RevertPatchesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
