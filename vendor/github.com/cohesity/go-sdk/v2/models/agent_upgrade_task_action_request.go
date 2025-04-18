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

// AgentUpgradeTaskActionRequest Upgrade task action params
//
// Specifies the params to perform an action on an agent upgrade task.
//
// swagger:model AgentUpgradeTaskActionRequest
type AgentUpgradeTaskActionRequest struct {

	// Specifies the action type.<br> 'Cancel' indicates that the upgrade task needs to be cancelled.
	// Enum: ["Retry","Cancel"]
	Action *string `json:"action,omitempty"`

	// Specifies the ID of the task.
	ID *int64 `json:"id,omitempty"`
}

// Validate validates this agent upgrade task action request
func (m *AgentUpgradeTaskActionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var agentUpgradeTaskActionRequestTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Retry","Cancel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		agentUpgradeTaskActionRequestTypeActionPropEnum = append(agentUpgradeTaskActionRequestTypeActionPropEnum, v)
	}
}

const (

	// AgentUpgradeTaskActionRequestActionRetry captures enum value "Retry"
	AgentUpgradeTaskActionRequestActionRetry string = "Retry"

	// AgentUpgradeTaskActionRequestActionCancel captures enum value "Cancel"
	AgentUpgradeTaskActionRequestActionCancel string = "Cancel"
)

// prop value enum
func (m *AgentUpgradeTaskActionRequest) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, agentUpgradeTaskActionRequestTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AgentUpgradeTaskActionRequest) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this agent upgrade task action request based on context it is used
func (m *AgentUpgradeTaskActionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AgentUpgradeTaskActionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentUpgradeTaskActionRequest) UnmarshalBinary(b []byte) error {
	var res AgentUpgradeTaskActionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
