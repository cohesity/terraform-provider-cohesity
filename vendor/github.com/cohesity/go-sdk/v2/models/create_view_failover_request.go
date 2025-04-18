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

// CreateViewFailoverRequest Specifies the request parameters to create a view failover task.
//
// swagger:model CreateViewFailoverRequest
type CreateViewFailoverRequest struct {

	// Specifies the failover type.<br> 'Planned' indicates this is a planned failover.<br> 'Unplanned' indicates this is an unplanned failover, which is used when source cluster is down.
	// Required: true
	// Enum: ["Planned","Unplanned"]
	Type *string `json:"type"`

	// Specifies parameters to create a planned failover.
	PlannedFailoverParams *PlannedFailoverParams `json:"plannedFailoverParams,omitempty"`

	// Specifies parameters to create an unplanned failover.
	UnplannedFailoverParams *UnplannedFailoverParams `json:"unplannedFailoverParams,omitempty"`
}

// Validate validates this create view failover request
func (m *CreateViewFailoverRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlannedFailoverParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnplannedFailoverParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createViewFailoverRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Planned","Unplanned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createViewFailoverRequestTypeTypePropEnum = append(createViewFailoverRequestTypeTypePropEnum, v)
	}
}

const (

	// CreateViewFailoverRequestTypePlanned captures enum value "Planned"
	CreateViewFailoverRequestTypePlanned string = "Planned"

	// CreateViewFailoverRequestTypeUnplanned captures enum value "Unplanned"
	CreateViewFailoverRequestTypeUnplanned string = "Unplanned"
)

// prop value enum
func (m *CreateViewFailoverRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createViewFailoverRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateViewFailoverRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CreateViewFailoverRequest) validatePlannedFailoverParams(formats strfmt.Registry) error {
	if swag.IsZero(m.PlannedFailoverParams) { // not required
		return nil
	}

	if m.PlannedFailoverParams != nil {
		if err := m.PlannedFailoverParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plannedFailoverParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plannedFailoverParams")
			}
			return err
		}
	}

	return nil
}

func (m *CreateViewFailoverRequest) validateUnplannedFailoverParams(formats strfmt.Registry) error {
	if swag.IsZero(m.UnplannedFailoverParams) { // not required
		return nil
	}

	if m.UnplannedFailoverParams != nil {
		if err := m.UnplannedFailoverParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unplannedFailoverParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unplannedFailoverParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create view failover request based on the context it is used
func (m *CreateViewFailoverRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlannedFailoverParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnplannedFailoverParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateViewFailoverRequest) contextValidatePlannedFailoverParams(ctx context.Context, formats strfmt.Registry) error {

	if m.PlannedFailoverParams != nil {

		if swag.IsZero(m.PlannedFailoverParams) { // not required
			return nil
		}

		if err := m.PlannedFailoverParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plannedFailoverParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plannedFailoverParams")
			}
			return err
		}
	}

	return nil
}

func (m *CreateViewFailoverRequest) contextValidateUnplannedFailoverParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UnplannedFailoverParams != nil {

		if swag.IsZero(m.UnplannedFailoverParams) { // not required
			return nil
		}

		if err := m.UnplannedFailoverParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unplannedFailoverParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unplannedFailoverParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateViewFailoverRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateViewFailoverRequest) UnmarshalBinary(b []byte) error {
	var res CreateViewFailoverRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
