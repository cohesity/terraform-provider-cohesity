// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ObjectActionRequest Peform an action on an Object.
//
// Specifies the request to peform an action on an Object.
//
// swagger:model ObjectActionRequest
type ObjectActionRequest struct {
	CommonObjectActionRequest

	// Specifies the actions for VMware Objects.
	VmwareParams *VmwareObjectActionParams `json:"vmwareParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ObjectActionRequest) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonObjectActionRequest
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonObjectActionRequest = aO0

	// AO1
	var dataAO1 struct {
		VmwareParams *VmwareObjectActionParams `json:"vmwareParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.VmwareParams = dataAO1.VmwareParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ObjectActionRequest) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonObjectActionRequest)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		VmwareParams *VmwareObjectActionParams `json:"vmwareParams,omitempty"`
	}

	dataAO1.VmwareParams = m.VmwareParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this object action request
func (m *ObjectActionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonObjectActionRequest
	if err := m.CommonObjectActionRequest.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmwareParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectActionRequest) validateVmwareParams(formats strfmt.Registry) error {

	if swag.IsZero(m.VmwareParams) { // not required
		return nil
	}

	if m.VmwareParams != nil {
		if err := m.VmwareParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this object action request based on the context it is used
func (m *ObjectActionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonObjectActionRequest
	if err := m.CommonObjectActionRequest.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmwareParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectActionRequest) contextValidateVmwareParams(ctx context.Context, formats strfmt.Registry) error {

	if m.VmwareParams != nil {

		if swag.IsZero(m.VmwareParams) { // not required
			return nil
		}

		if err := m.VmwareParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectActionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectActionRequest) UnmarshalBinary(b []byte) error {
	var res ObjectActionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
