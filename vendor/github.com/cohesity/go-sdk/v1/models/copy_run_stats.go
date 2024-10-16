// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CopyRunStats Stats for one copy task or aggregated stats of a Copy Run in a
// Protection Job Run.
//
// swagger:model CopyRunStats
type CopyRunStats struct {

	// Specifies the time when this replication ended. If not set, then the
	// replication has not ended yet.
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`

	// Specifies whether this archival is incremental for archival targets.
	IsIncremental *bool `json:"isIncremental,omitempty"`

	// Specifies the number of logical bytes transferred for this replication
	// so far. This value can never exceed the total logical size of the
	// replicated view.
	LogicalBytesTransferred *int64 `json:"logicalBytesTransferred,omitempty"`

	// Specifies the total amount of logical data to be transferred for this
	// replication.
	LogicalSizeBytes *int64 `json:"logicalSizeBytes,omitempty"`

	// Specifies average logical bytes transfer rate in bytes per second for
	// archchival targets.
	LogicalTransferRateBps *int64 `json:"logicalTransferRateBps,omitempty"`

	// Specifies the number of physical bytes sent over the wire for
	// replication targets.
	PhysicalBytesTransferred *int64 `json:"physicalBytesTransferred,omitempty"`

	// Specifies the time when this replication was started. If not set, then
	// replication has not been started yet.
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`
}

// Validate validates this copy run stats
func (m *CopyRunStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this copy run stats based on context it is used
func (m *CopyRunStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CopyRunStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopyRunStats) UnmarshalBinary(b []byte) error {
	var res CopyRunStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
