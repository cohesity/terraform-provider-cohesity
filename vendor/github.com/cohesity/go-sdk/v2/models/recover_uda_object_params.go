// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecoverUdaObjectParams Specifies details of objects to be recovered.
//
// swagger:model RecoverUdaObjectParams
type RecoverUdaObjectParams struct {

	// Specifies the ID of the object.
	ObjectID *int64 `json:"objectId,omitempty"`

	// Specifies the fully qualified name of the object to be restored.
	ObjectName *string `json:"objectName,omitempty"`

	// Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object.
	Overwrite *bool `json:"overwrite,omitempty"`

	// Specifies the new name to which the object should be renamed to after the recovery.
	RenameTo *string `json:"renameTo,omitempty"`
}

// Validate validates this recover uda object params
func (m *RecoverUdaObjectParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this recover uda object params based on context it is used
func (m *RecoverUdaObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecoverUdaObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverUdaObjectParams) UnmarshalBinary(b []byte) error {
	var res RecoverUdaObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
