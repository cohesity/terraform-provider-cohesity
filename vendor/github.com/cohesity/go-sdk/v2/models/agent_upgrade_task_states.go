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
)

// AgentUpgradeTaskStates Agent upgrade task list
//
// List of agent upgrade tasks.
//
// swagger:model AgentUpgradeTaskStates
type AgentUpgradeTaskStates struct {

	// Specifies the list of agent upgrade tasks.
	Tasks []*AgentUpgradeTaskState `json:"tasks"`
}

// Validate validates this agent upgrade task states
func (m *AgentUpgradeTaskStates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentUpgradeTaskStates) validateTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {
		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this agent upgrade task states based on the context it is used
func (m *AgentUpgradeTaskStates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentUpgradeTaskStates) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tasks); i++ {

		if m.Tasks[i] != nil {

			if swag.IsZero(m.Tasks[i]) { // not required
				return nil
			}

			if err := m.Tasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentUpgradeTaskStates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentUpgradeTaskStates) UnmarshalBinary(b []byte) error {
	var res AgentUpgradeTaskStates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
