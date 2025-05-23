// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatastoreParams Datastore params.
//
// Specifies the datastore params.
//
// swagger:model DatastoreParams
type DatastoreParams struct {

	// Specifies the Id of the datastore.
	ID *int64 `json:"id,omitempty"`

	// Specifies the max concurrent stream per datastore.
	MaxConcurrentStreams *int32 `json:"maxConcurrentStreams,omitempty"`
}

// Validate validates this datastore params
func (m *DatastoreParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this datastore params based on context it is used
func (m *DatastoreParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatastoreParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatastoreParams) UnmarshalBinary(b []byte) error {
	var res DatastoreParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
