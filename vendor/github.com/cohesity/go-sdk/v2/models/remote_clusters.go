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

// RemoteClusters Specifies a list of Remote Cluster registered with the cluster.
//
// swagger:model RemoteClusters
type RemoteClusters struct {

	// Specifies the list of Remote Clusters.
	RemoteClusters []*RemoteCluster `json:"remoteClusters"`
}

// Validate validates this remote clusters
func (m *RemoteClusters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemoteClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteClusters) validateRemoteClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.RemoteClusters); i++ {
		if swag.IsZero(m.RemoteClusters[i]) { // not required
			continue
		}

		if m.RemoteClusters[i] != nil {
			if err := m.RemoteClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remoteClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remoteClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this remote clusters based on the context it is used
func (m *RemoteClusters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRemoteClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteClusters) contextValidateRemoteClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RemoteClusters); i++ {

		if m.RemoteClusters[i] != nil {

			if swag.IsZero(m.RemoteClusters[i]) { // not required
				return nil
			}

			if err := m.RemoteClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remoteClusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remoteClusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteClusters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteClusters) UnmarshalBinary(b []byte) error {
	var res RemoteClusters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
