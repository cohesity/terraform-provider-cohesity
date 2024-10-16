// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CouchbaseCluster Couchbase Cluster Info.
//
// Specifies an Object containing information about a couchbase cluster.
//
// swagger:model CouchbaseCluster
type CouchbaseCluster struct {

	// Seeds of this Couchbase Cluster.
	Seeds []string `json:"seeds"`
}

// Validate validates this couchbase cluster
func (m *CouchbaseCluster) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this couchbase cluster based on context it is used
func (m *CouchbaseCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CouchbaseCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CouchbaseCluster) UnmarshalBinary(b []byte) error {
	var res CouchbaseCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
