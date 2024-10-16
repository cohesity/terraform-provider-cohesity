// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PrivateTaskEvent Events that clients can attach to a task.
//
// If a task has many small steps (some of them conditional), then a client can
// create a single task and attach each small step as an event.
// The client can collect a few of them, and report progress on a batch.
// The events act like a log for the task, and Pulse does not interpret them.
//
// swagger:model PrivateTaskEvent
type PrivateTaskEvent struct {

	// Message associated with the event.
	EventMsg *string `json:"eventMsg,omitempty"`

	// How much the owning task completed when this event occurred.
	OwnerPercentFinished *float32 `json:"ownerPercentFinished,omitempty"`

	// How much work was remaining for the owning task when this event occurred.
	OwnerRemainingWorkCount *int64 `json:"ownerRemainingWorkCount,omitempty"`

	// The timestamp at which the event occurred.
	TimestampSecs *int64 `json:"timestampSecs,omitempty"`
}

// Validate validates this private task event
func (m *PrivateTaskEvent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this private task event based on context it is used
func (m *PrivateTaskEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrivateTaskEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrivateTaskEvent) UnmarshalBinary(b []byte) error {
	var res PrivateTaskEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
