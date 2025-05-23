// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProgressTaskEvent Specifies the details about the various events which are created during the execution of Progress Task.
//
// swagger:model ProgressTaskEvent
type ProgressTaskEvent struct {

	// Specifies the log message describing the current event.
	Message *string `json:"message,omitempty"`

	// Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
	OccuredAtUsecs *int64 `json:"occuredAtUsecs,omitempty"`
}

// Validate validates this progress task event
func (m *ProgressTaskEvent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this progress task event based on context it is used
func (m *ProgressTaskEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProgressTaskEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressTaskEvent) UnmarshalBinary(b []byte) error {
	var res ProgressTaskEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
