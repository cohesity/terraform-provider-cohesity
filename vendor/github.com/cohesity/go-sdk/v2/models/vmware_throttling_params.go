// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareThrottlingParams Throttling Params.
//
// Specifies throttling params.
//
// swagger:model VmwareThrottlingParams
type VmwareThrottlingParams struct {

	// If the latency of a datastore is above this value, then a new backup task that uses the datastore won't be started.
	NewTaskLatencyThresholdMsecs *int64 `json:"newTaskLatencyThresholdMsecs,omitempty"`

	// If the latency of a datastore is above this value, then an existing backup task that uses the datastore will start getting throttled.
	ActiveTaskLatencyThresholdMsecs *int64 `json:"activeTaskLatencyThresholdMsecs,omitempty"`

	// If this value is > 0 and the number of streams concurrently active on a datastore is equal to it, then any further requests to access the datastore would be denied until the number of active streams reduces. This applies for all the datastores in the specified host.
	MaxConcurrentStreams *int32 `json:"maxConcurrentStreams,omitempty"`

	// Specifies the number of VMs of a vCenter that can be backed up concurrently.
	MaxConcurrentBackups *int32 `json:"maxConcurrentBackups,omitempty"`
}

// Validate validates this vmware throttling params
func (m *VmwareThrottlingParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vmware throttling params based on context it is used
func (m *VmwareThrottlingParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VmwareThrottlingParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareThrottlingParams) UnmarshalBinary(b []byte) error {
	var res VmwareThrottlingParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
