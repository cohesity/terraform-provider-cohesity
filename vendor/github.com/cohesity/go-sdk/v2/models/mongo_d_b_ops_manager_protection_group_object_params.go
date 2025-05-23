// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MongoDBOpsManagerProtectionGroupObjectParams Specifies the object identifier to create MongoDB Protection Group.
//
// swagger:model MongoDBOpsManagerProtectionGroupObjectParams
type MongoDBOpsManagerProtectionGroupObjectParams struct {

	// Specifies the ID of the object.
	ID *int64 `json:"id,omitempty"`

	// Specifies the fully qualified name of the object.
	Name *string `json:"name,omitempty"`
}

// Validate validates this mongo d b ops manager protection group object params
func (m *MongoDBOpsManagerProtectionGroupObjectParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mongo d b ops manager protection group object params based on context it is used
func (m *MongoDBOpsManagerProtectionGroupObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongoDBOpsManagerProtectionGroupObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDBOpsManagerProtectionGroupObjectParams) UnmarshalBinary(b []byte) error {
	var res MongoDBOpsManagerProtectionGroupObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
