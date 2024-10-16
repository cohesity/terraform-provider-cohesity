// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateInfectedFilesParameters Specifies the parameters of infected entities to be updated.
//
// swagger:model UpdateInfectedFilesParameters
type UpdateInfectedFilesParameters struct {

	// Specifies a list of infected entities to be updated.
	// Required: true
	// Min Items: 1
	InfectedFiles []*InfectedFile `json:"infectedFiles"`

	// Specifies the state[Quarantined, Unquarantined] of the infected entity.
	// Enum: ["Quarantined","Unquarantined"]
	State *string `json:"state,omitempty"`
}

// Validate validates this update infected files parameters
func (m *UpdateInfectedFilesParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfectedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateInfectedFilesParameters) validateInfectedFiles(formats strfmt.Registry) error {

	if err := validate.Required("infectedFiles", "body", m.InfectedFiles); err != nil {
		return err
	}

	iInfectedFilesSize := int64(len(m.InfectedFiles))

	if err := validate.MinItems("infectedFiles", "body", iInfectedFilesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.InfectedFiles); i++ {
		if swag.IsZero(m.InfectedFiles[i]) { // not required
			continue
		}

		if m.InfectedFiles[i] != nil {
			if err := m.InfectedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("infectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("infectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var updateInfectedFilesParametersTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Quarantined","Unquarantined"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateInfectedFilesParametersTypeStatePropEnum = append(updateInfectedFilesParametersTypeStatePropEnum, v)
	}
}

const (

	// UpdateInfectedFilesParametersStateQuarantined captures enum value "Quarantined"
	UpdateInfectedFilesParametersStateQuarantined string = "Quarantined"

	// UpdateInfectedFilesParametersStateUnquarantined captures enum value "Unquarantined"
	UpdateInfectedFilesParametersStateUnquarantined string = "Unquarantined"
)

// prop value enum
func (m *UpdateInfectedFilesParameters) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateInfectedFilesParametersTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateInfectedFilesParameters) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update infected files parameters based on the context it is used
func (m *UpdateInfectedFilesParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfectedFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateInfectedFilesParameters) contextValidateInfectedFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InfectedFiles); i++ {

		if m.InfectedFiles[i] != nil {

			if swag.IsZero(m.InfectedFiles[i]) { // not required
				return nil
			}

			if err := m.InfectedFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("infectedFiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("infectedFiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateInfectedFilesParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateInfectedFilesParameters) UnmarshalBinary(b []byte) error {
	var res UpdateInfectedFilesParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
