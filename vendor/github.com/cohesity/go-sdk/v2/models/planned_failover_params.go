// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlannedFailoverParams Specifies parameters of a planned failover.
//
// swagger:model PlannedFailoverParams
type PlannedFailoverParams struct {

	// Spcifies the planned failover type.<br> 'Prepare' indicates this is a preparation for failover.<br> 'Finalize' indicates this is finalization of failover. After this is done, the view can be used as source view.
	// Required: true
	// Enum: ["Prepare","Finalize"]
	Type *string `json:"type"`

	// Specifies parameters of preparation of a planned failover.
	PreparePlannedFailverParams *PreparePlannedFailverParams `json:"preparePlannedFailverParams,omitempty"`
}

// Validate validates this planned failover params
func (m *PlannedFailoverParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreparePlannedFailverParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var plannedFailoverParamsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Prepare","Finalize"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		plannedFailoverParamsTypeTypePropEnum = append(plannedFailoverParamsTypeTypePropEnum, v)
	}
}

const (

	// PlannedFailoverParamsTypePrepare captures enum value "Prepare"
	PlannedFailoverParamsTypePrepare string = "Prepare"

	// PlannedFailoverParamsTypeFinalize captures enum value "Finalize"
	PlannedFailoverParamsTypeFinalize string = "Finalize"
)

// prop value enum
func (m *PlannedFailoverParams) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, plannedFailoverParamsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlannedFailoverParams) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PlannedFailoverParams) validatePreparePlannedFailverParams(formats strfmt.Registry) error {
	if swag.IsZero(m.PreparePlannedFailverParams) { // not required
		return nil
	}

	if m.PreparePlannedFailverParams != nil {
		if err := m.PreparePlannedFailverParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preparePlannedFailverParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preparePlannedFailverParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this planned failover params based on the context it is used
func (m *PlannedFailoverParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreparePlannedFailverParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlannedFailoverParams) contextValidatePreparePlannedFailverParams(ctx context.Context, formats strfmt.Registry) error {

	if m.PreparePlannedFailverParams != nil {

		if swag.IsZero(m.PreparePlannedFailverParams) { // not required
			return nil
		}

		if err := m.PreparePlannedFailverParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preparePlannedFailverParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preparePlannedFailverParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlannedFailoverParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlannedFailoverParams) UnmarshalBinary(b []byte) error {
	var res PlannedFailoverParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
