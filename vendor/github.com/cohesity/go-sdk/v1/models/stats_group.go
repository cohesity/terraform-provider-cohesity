// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatsGroup StatsGroup.
//
// StatsGroup describes the details of a stats group. A stats group is a basic
// group of usage stats, it is the usage of a tenant within a storage domain
// may also for a specific consumer type.
//
// swagger:model StatsGroup
type StatsGroup struct {

	// Specifies the consumer information of this group.
	Consumer *Consumer `json:"consumer,omitempty"`

	// Specifies the entity id of the group.
	EntityID *string `json:"entityId,omitempty"`

	// Specifies the id of the group.
	ID *int64 `json:"id,omitempty"`

	// Specifies the name of the group.
	Name *string `json:"name,omitempty"`

	// Specifies the id of the organization (tenant) with respect to this group.
	TenantID *string `json:"tenantId,omitempty"`

	// Specifies the name of the organization (tenant) with respect to this
	// group.
	TenantName *string `json:"tenantName,omitempty"`

	// Specifies the id of the view box (storage domain) with respect to this
	// group.
	ViewBoxID *int64 `json:"viewBoxId,omitempty"`

	// Specifies the name of the view box (storage domain) with respect to this
	// group.
	ViewBoxName *string `json:"viewBoxName,omitempty"`
}

// Validate validates this stats group
func (m *StatsGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsumer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatsGroup) validateConsumer(formats strfmt.Registry) error {
	if swag.IsZero(m.Consumer) { // not required
		return nil
	}

	if m.Consumer != nil {
		if err := m.Consumer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consumer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consumer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this stats group based on the context it is used
func (m *StatsGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsumer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatsGroup) contextValidateConsumer(ctx context.Context, formats strfmt.Registry) error {

	if m.Consumer != nil {

		if swag.IsZero(m.Consumer) { // not required
			return nil
		}

		if err := m.Consumer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consumer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consumer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatsGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatsGroup) UnmarshalBinary(b []byte) error {
	var res StatsGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
