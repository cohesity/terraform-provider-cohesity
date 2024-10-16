// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EntityProto Entity.
//
// Specifies the attributes and the latest statistics about an entity.
//
// swagger:model EntityProto
type EntityProto struct {

	// Array of Attributes.
	//
	// List of attributes of an entity.
	AttributeVec []*KeyValuePair `json:"attributeVec"`

	// Specifies a globally unique identifier of an entity.
	EntityID *EntityIdentifier `json:"entityId,omitempty"`

	// Array of Metric Statistics.
	//
	// List of the latest statistics for all metrics defined in the schema
	// that this entity belongs to.
	// If statistics for a metric is not available, then that data point is not
	// returned.
	LatestMetricVec []*MetricValue `json:"latestMetricVec"`
}

// Validate validates this entity proto
func (m *EntityProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestMetricVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityProto) validateAttributeVec(formats strfmt.Registry) error {
	if swag.IsZero(m.AttributeVec) { // not required
		return nil
	}

	for i := 0; i < len(m.AttributeVec); i++ {
		if swag.IsZero(m.AttributeVec[i]) { // not required
			continue
		}

		if m.AttributeVec[i] != nil {
			if err := m.AttributeVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributeVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntityProto) validateEntityID(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityID) { // not required
		return nil
	}

	if m.EntityID != nil {
		if err := m.EntityID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityId")
			}
			return err
		}
	}

	return nil
}

func (m *EntityProto) validateLatestMetricVec(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestMetricVec) { // not required
		return nil
	}

	for i := 0; i < len(m.LatestMetricVec); i++ {
		if swag.IsZero(m.LatestMetricVec[i]) { // not required
			continue
		}

		if m.LatestMetricVec[i] != nil {
			if err := m.LatestMetricVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("latestMetricVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("latestMetricVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this entity proto based on the context it is used
func (m *EntityProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributeVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestMetricVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityProto) contextValidateAttributeVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AttributeVec); i++ {

		if m.AttributeVec[i] != nil {

			if swag.IsZero(m.AttributeVec[i]) { // not required
				return nil
			}

			if err := m.AttributeVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributeVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributeVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EntityProto) contextValidateEntityID(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityID != nil {

		if swag.IsZero(m.EntityID) { // not required
			return nil
		}

		if err := m.EntityID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityId")
			}
			return err
		}
	}

	return nil
}

func (m *EntityProto) contextValidateLatestMetricVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LatestMetricVec); i++ {

		if m.LatestMetricVec[i] != nil {

			if swag.IsZero(m.LatestMetricVec[i]) { // not required
				return nil
			}

			if err := m.LatestMetricVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("latestMetricVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("latestMetricVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityProto) UnmarshalBinary(b []byte) error {
	var res EntityProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
