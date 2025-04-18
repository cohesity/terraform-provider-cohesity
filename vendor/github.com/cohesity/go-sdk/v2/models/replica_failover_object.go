// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicaFailoverObject Specifies the object paring of replicated object and failover object created after restore.
//
// swagger:model ReplicaFailoverObject
type ReplicaFailoverObject struct {

	// Specifies the object Id existing on the replciation cluster.
	ReplicaObjectID *int64 `json:"replicaObjectId,omitempty"`

	// Specifies the corrosponding object id of the failover object.
	FailoverObjectID *int64 `json:"failoverObjectId,omitempty"`
}

// Validate validates this replica failover object
func (m *ReplicaFailoverObject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this replica failover object based on context it is used
func (m *ReplicaFailoverObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicaFailoverObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicaFailoverObject) UnmarshalBinary(b []byte) error {
	var res ReplicaFailoverObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
