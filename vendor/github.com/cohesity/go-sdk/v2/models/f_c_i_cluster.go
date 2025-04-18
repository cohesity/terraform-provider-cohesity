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

// FCICluster Specifies the details of a Failover Cluster Instance.
//
// swagger:model FCICluster
type FCICluster struct {

	// Specifies the unique identifier of the FCI.
	ID *string `json:"id,omitempty"`

	// Specifies the name of the FCI.
	Name *string `json:"name,omitempty"`

	// Specifies the error information if any associated with this FCI cluster.
	Error *Error `json:"error,omitempty"`

	// Specifies the resource information about the FCI.
	ResourceInfo *AppResource `json:"resourceInfo,omitempty"`

	// Specifies the list of SQL servers which belongs to the given FCI.
	Servers []*SQLServer `json:"servers"`

	// Indicates to the UI whether this FCI cluster should be selected by default
	IsSelectedByDefault *bool `json:"isSelectedByDefault,omitempty"`
}

// Validate validates this f c i cluster
func (m *FCICluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FCICluster) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *FCICluster) validateResourceInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceInfo) { // not required
		return nil
	}

	if m.ResourceInfo != nil {
		if err := m.ResourceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceInfo")
			}
			return err
		}
	}

	return nil
}

func (m *FCICluster) validateServers(formats strfmt.Registry) error {
	if swag.IsZero(m.Servers) { // not required
		return nil
	}

	for i := 0; i < len(m.Servers); i++ {
		if swag.IsZero(m.Servers[i]) { // not required
			continue
		}

		if m.Servers[i] != nil {
			if err := m.Servers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this f c i cluster based on the context it is used
func (m *FCICluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FCICluster) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *FCICluster) contextValidateResourceInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceInfo != nil {

		if swag.IsZero(m.ResourceInfo) { // not required
			return nil
		}

		if err := m.ResourceInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceInfo")
			}
			return err
		}
	}

	return nil
}

func (m *FCICluster) contextValidateServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Servers); i++ {

		if m.Servers[i] != nil {

			if swag.IsZero(m.Servers[i]) { // not required
				return nil
			}

			if err := m.Servers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FCICluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FCICluster) UnmarshalBinary(b []byte) error {
	var res FCICluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
