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

// DataTieringTask Specifies the data tiering task details.
//
// swagger:model DataTieringTask
type DataTieringTask struct {
	CommonDataTieringTaskResponse

	// Specifies downtiering policy for data tiering task.
	DowntieringPolicy *DowntieringPolicy `json:"downtieringPolicy,omitempty"`

	// Specifies uptiering policy for data tiering task.
	UptieringPolicy *UptieringPolicy `json:"uptieringPolicy,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DataTieringTask) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonDataTieringTaskResponse
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonDataTieringTaskResponse = aO0

	// AO1
	var dataAO1 struct {
		DowntieringPolicy *DowntieringPolicy `json:"downtieringPolicy,omitempty"`

		UptieringPolicy *UptieringPolicy `json:"uptieringPolicy,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DowntieringPolicy = dataAO1.DowntieringPolicy

	m.UptieringPolicy = dataAO1.UptieringPolicy

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DataTieringTask) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonDataTieringTaskResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		DowntieringPolicy *DowntieringPolicy `json:"downtieringPolicy,omitempty"`

		UptieringPolicy *UptieringPolicy `json:"uptieringPolicy,omitempty"`
	}

	dataAO1.DowntieringPolicy = m.DowntieringPolicy

	dataAO1.UptieringPolicy = m.UptieringPolicy

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this data tiering task
func (m *DataTieringTask) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonDataTieringTaskResponse
	if err := m.CommonDataTieringTaskResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDowntieringPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUptieringPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTieringTask) validateDowntieringPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.DowntieringPolicy) { // not required
		return nil
	}

	if m.DowntieringPolicy != nil {
		if err := m.DowntieringPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downtieringPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downtieringPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *DataTieringTask) validateUptieringPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.UptieringPolicy) { // not required
		return nil
	}

	if m.UptieringPolicy != nil {
		if err := m.UptieringPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uptieringPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uptieringPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this data tiering task based on the context it is used
func (m *DataTieringTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonDataTieringTaskResponse
	if err := m.CommonDataTieringTaskResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDowntieringPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUptieringPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataTieringTask) contextValidateDowntieringPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.DowntieringPolicy != nil {

		if swag.IsZero(m.DowntieringPolicy) { // not required
			return nil
		}

		if err := m.DowntieringPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downtieringPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downtieringPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *DataTieringTask) contextValidateUptieringPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.UptieringPolicy != nil {

		if swag.IsZero(m.UptieringPolicy) { // not required
			return nil
		}

		if err := m.UptieringPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uptieringPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uptieringPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataTieringTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTieringTask) UnmarshalBinary(b []byte) error {
	var res DataTieringTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
