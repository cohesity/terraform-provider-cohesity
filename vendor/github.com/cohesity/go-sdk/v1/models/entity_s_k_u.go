// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EntitySKU Sku of the entity.
//
// swagger:model Entity_SKU
type EntitySKU struct {

	// Capacity of the sku.
	// For azure sql dbs, this is the number of cores.
	Capacity *int32 `json:"capacity,omitempty"`

	// Can be one of Name_Type enum above.
	Name *string `json:"name,omitempty"`

	// Enum representation of name for UI selection purpose.
	NameType *int32 `json:"nameType,omitempty"`

	// Can be one of Tier_Type enum above.
	Tier *string `json:"tier,omitempty"`

	// Enum representation of tier for UI selection purpose.
	TierType *int32 `json:"tierType,omitempty"`
}

// Validate validates this entity s k u
func (m *EntitySKU) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this entity s k u based on context it is used
func (m *EntitySKU) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EntitySKU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntitySKU) UnmarshalBinary(b []byte) error {
	var res EntitySKU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
