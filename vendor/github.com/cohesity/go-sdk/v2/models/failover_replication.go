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

// FailoverReplication Specifies the details of a failover replication.
//
// swagger:model FailoverReplication
type FailoverReplication struct {

	// Specifies the replication id.
	ID *string `json:"id,omitempty"`

	// Specifies the replication status.
	// Enum: ["Running","Succeeded","Failed"]
	Status *string `json:"status,omitempty"`

	// Specifies the error details if replication status is 'Failed'.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Specifies the replication start time in micro seconds.
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`

	// Specifies the replication complete time in micro seconds.
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`

	// Specifies the percentage completed in the replication.
	PercentageCompleted *int32 `json:"percentageCompleted,omitempty"`

	// Specifies the total amount of logical data to be transferred for this replication.
	LogicalSizeBytes *int64 `json:"logicalSizeBytes,omitempty"`

	// Specifies the number of logical bytes transferred for this replication so far. This value can never exceed the total logical size of the replicated view.
	LogicalBytesTransferred *int64 `json:"logicalBytesTransferred,omitempty"`

	// Specifies the number of bytes sent over the wire for this replication so far.
	PhysicalBytesTransferred *int64 `json:"physicalBytesTransferred,omitempty"`

	// Specifies the failover target cluster id.
	TargetClusterID *int64 `json:"targetClusterId,omitempty"`

	// Specifies the failover target cluster incarnation id.
	TargetClusterIncarnationID *int64 `json:"targetClusterIncarnationId,omitempty"`

	// Specifies the failover target cluster name.
	TargetClusterName *string `json:"targetClusterName,omitempty"`
}

// Validate validates this failover replication
func (m *FailoverReplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var failoverReplicationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Running","Succeeded","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		failoverReplicationTypeStatusPropEnum = append(failoverReplicationTypeStatusPropEnum, v)
	}
}

const (

	// FailoverReplicationStatusRunning captures enum value "Running"
	FailoverReplicationStatusRunning string = "Running"

	// FailoverReplicationStatusSucceeded captures enum value "Succeeded"
	FailoverReplicationStatusSucceeded string = "Succeeded"

	// FailoverReplicationStatusFailed captures enum value "Failed"
	FailoverReplicationStatusFailed string = "Failed"
)

// prop value enum
func (m *FailoverReplication) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, failoverReplicationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FailoverReplication) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this failover replication based on context it is used
func (m *FailoverReplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FailoverReplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FailoverReplication) UnmarshalBinary(b []byte) error {
	var res FailoverReplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
