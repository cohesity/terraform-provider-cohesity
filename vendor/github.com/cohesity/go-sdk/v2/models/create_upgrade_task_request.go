// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateUpgradeTaskRequest Create agent upgrade task params
//
// Specifies the params to create a schedule based agent upgrade task.
//
// swagger:model CreateUpgradeTaskRequest
type CreateUpgradeTaskRequest struct {

	// Specifies the name of the task.
	Name *string `json:"name,omitempty"`

	// Specifies the description of the task.
	Description *string `json:"description,omitempty"`

	// Specifies agent IDs to be upgraded in the task.
	AgentIDs []int64 `json:"agentIDs"`

	// Specifies the start time of the task specified by the user as a Unix epoch Timestamp (in microseconds). If no schedule is specified, the agents will be upgraded immediately.
	ScheduleTimeUsecs *int64 `json:"scheduleTimeUsecs,omitempty"`

	// Specifies the time before which the upgrade task should start execution as a Unix epoch Timestamp (in microseconds). If this is not specified the task will start anytime after scheduleTimeUsecs.
	ScheduleEndTimeUsecs *int64 `json:"scheduleEndTimeUsecs,omitempty"`

	// Specifies the task that needs to be retried.
	RetryTaskID *int64 `json:"retryTaskId,omitempty"`
}

// Validate validates this create upgrade task request
func (m *CreateUpgradeTaskRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create upgrade task request based on context it is used
func (m *CreateUpgradeTaskRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateUpgradeTaskRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUpgradeTaskRequest) UnmarshalBinary(b []byte) error {
	var res CreateUpgradeTaskRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
