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

// RemoteDisks Remote disks
//
// Specifies a list of remote disks.
//
// swagger:model RemoteDisks
type RemoteDisks struct {

	// Specifies a list of remote disks.
	RemoteDisks []*RemoteDisk `json:"remoteDisks"`
}

// Validate validates this remote disks
func (m *RemoteDisks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemoteDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteDisks) validateRemoteDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.RemoteDisks); i++ {
		if swag.IsZero(m.RemoteDisks[i]) { // not required
			continue
		}

		if m.RemoteDisks[i] != nil {
			if err := m.RemoteDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remoteDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remoteDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this remote disks based on the context it is used
func (m *RemoteDisks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRemoteDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteDisks) contextValidateRemoteDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RemoteDisks); i++ {

		if m.RemoteDisks[i] != nil {

			if swag.IsZero(m.RemoteDisks[i]) { // not required
				return nil
			}

			if err := m.RemoteDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remoteDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("remoteDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteDisks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteDisks) UnmarshalBinary(b []byte) error {
	var res RemoteDisks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
