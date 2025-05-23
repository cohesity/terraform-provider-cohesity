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

// UptieringPolicy Specifies the data uptiering policy.
//
// swagger:model UptieringPolicy
type UptieringPolicy struct {

	// file age
	FileAge *UptieringFileAgePolicy `json:"fileAge,omitempty"`

	// If set, all files in the view will be uptiered regardless of
	// file_select_policy, num_file_access, hot_file_window, file_size
	// constraints.
	IncludeAllFiles *bool `json:"includeAllFiles,omitempty"`

	// target
	Target *UptieringTarget `json:"target,omitempty"`

	CommonTieringPolicy
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UptieringPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		FileAge *UptieringFileAgePolicy `json:"fileAge,omitempty"`

		IncludeAllFiles *bool `json:"includeAllFiles,omitempty"`

		Target *UptieringTarget `json:"target,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.FileAge = dataAO0.FileAge

	m.IncludeAllFiles = dataAO0.IncludeAllFiles

	m.Target = dataAO0.Target

	// AO1
	var aO1 CommonTieringPolicy
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CommonTieringPolicy = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UptieringPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		FileAge *UptieringFileAgePolicy `json:"fileAge,omitempty"`

		IncludeAllFiles *bool `json:"includeAllFiles,omitempty"`

		Target *UptieringTarget `json:"target,omitempty"`
	}

	dataAO0.FileAge = m.FileAge

	dataAO0.IncludeAllFiles = m.IncludeAllFiles

	dataAO0.Target = m.Target

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.CommonTieringPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this uptiering policy
func (m *UptieringPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with CommonTieringPolicy
	if err := m.CommonTieringPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UptieringPolicy) validateFileAge(formats strfmt.Registry) error {

	if swag.IsZero(m.FileAge) { // not required
		return nil
	}

	if m.FileAge != nil {
		if err := m.FileAge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileAge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileAge")
			}
			return err
		}
	}

	return nil
}

func (m *UptieringPolicy) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this uptiering policy based on the context it is used
func (m *UptieringPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileAge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with CommonTieringPolicy
	if err := m.CommonTieringPolicy.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UptieringPolicy) contextValidateFileAge(ctx context.Context, formats strfmt.Registry) error {

	if m.FileAge != nil {

		if swag.IsZero(m.FileAge) { // not required
			return nil
		}

		if err := m.FileAge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileAge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileAge")
			}
			return err
		}
	}

	return nil
}

func (m *UptieringPolicy) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UptieringPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UptieringPolicy) UnmarshalBinary(b []byte) error {
	var res UptieringPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
