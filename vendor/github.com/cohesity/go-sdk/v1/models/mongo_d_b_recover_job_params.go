// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MongoDBRecoverJobParams Contains any additional mongodb environment specific params for the
// recover job.
//
// swagger:model MongoDBRecoverJobParams
type MongoDBRecoverJobParams struct {

	// Should the agent Recover users and roles in this job.
	RecoverUserRole *bool `json:"recoverUserRole,omitempty"`

	// Should the agent Recover zones/shard tags in this job.
	RecoverZonesTags *bool `json:"recoverZonesTags,omitempty"`

	// A suffix that is to be applied to all recovered entities
	Suffix *string `json:"suffix,omitempty"`
}

// Validate validates this mongo d b recover job params
func (m *MongoDBRecoverJobParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mongo d b recover job params based on context it is used
func (m *MongoDBRecoverJobParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongoDBRecoverJobParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDBRecoverJobParams) UnmarshalBinary(b []byte) error {
	var res MongoDBRecoverJobParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
