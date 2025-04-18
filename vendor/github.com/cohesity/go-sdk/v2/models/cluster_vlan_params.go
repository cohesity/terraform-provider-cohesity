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

// ClusterVlanParams Cluster vlan parameters.
//
// swagger:model ClusterVlanParams
type ClusterVlanParams struct {
	CreateClusterVlanParams

	// Set to true when vlan app IP addresses are being used by apps. When this is set to true, the vlan interface can't be deleted.
	AppIpsInUse *bool `json:"appIpsInUse,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ClusterVlanParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CreateClusterVlanParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CreateClusterVlanParams = aO0

	// AO1
	var dataAO1 struct {
		AppIpsInUse *bool `json:"appIpsInUse,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AppIpsInUse = dataAO1.AppIpsInUse

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ClusterVlanParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CreateClusterVlanParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AppIpsInUse *bool `json:"appIpsInUse,omitempty"`
	}

	dataAO1.AppIpsInUse = m.AppIpsInUse

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this cluster vlan params
func (m *ClusterVlanParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CreateClusterVlanParams
	if err := m.CreateClusterVlanParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this cluster vlan params based on the context it is used
func (m *ClusterVlanParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CreateClusterVlanParams
	if err := m.CreateClusterVlanParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterVlanParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterVlanParams) UnmarshalBinary(b []byte) error {
	var res ClusterVlanParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
