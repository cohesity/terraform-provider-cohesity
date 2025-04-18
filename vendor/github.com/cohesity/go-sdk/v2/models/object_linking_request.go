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

// ObjectLinkingRequest Request for linking replicated objects to failover objects on replication cluster.
//
// swagger:model ObjectLinkingRequest
type ObjectLinkingRequest struct {

	// Specifies the objectMap that will be used to create linking between given replicated source entity and newly restored entity on erplication cluster.
	ObjectMap []*ReplicaFailoverObject `json:"objectMap"`
}

// Validate validates this object linking request
func (m *ObjectLinkingRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectLinkingRequest) validateObjectMap(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectMap) { // not required
		return nil
	}

	for i := 0; i < len(m.ObjectMap); i++ {
		if swag.IsZero(m.ObjectMap[i]) { // not required
			continue
		}

		if m.ObjectMap[i] != nil {
			if err := m.ObjectMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this object linking request based on the context it is used
func (m *ObjectLinkingRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectLinkingRequest) contextValidateObjectMap(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ObjectMap); i++ {

		if m.ObjectMap[i] != nil {

			if swag.IsZero(m.ObjectMap[i]) { // not required
				return nil
			}

			if err := m.ObjectMap[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objectMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectLinkingRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectLinkingRequest) UnmarshalBinary(b []byte) error {
	var res ObjectLinkingRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
