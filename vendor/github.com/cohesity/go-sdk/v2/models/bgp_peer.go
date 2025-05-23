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

// BgpPeer BgpPeer
//
// BGP peer information.
//
// swagger:model BgpPeer
type BgpPeer struct {

	// IP Address.
	AddressOrTag *string `json:"addressOrTag,omitempty"`

	// Peer's description.
	Description *string `json:"description,omitempty"`

	// Remote AS.
	RemoteAS *uint32 `json:"remoteAS,omitempty"`

	// Bgp timers.
	Timers *BgpTimers `json:"timers,omitempty"`
}

// Validate validates this bgp peer
func (m *BgpPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BgpPeer) validateTimers(formats strfmt.Registry) error {
	if swag.IsZero(m.Timers) { // not required
		return nil
	}

	if m.Timers != nil {
		if err := m.Timers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timers")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this bgp peer based on the context it is used
func (m *BgpPeer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BgpPeer) contextValidateTimers(ctx context.Context, formats strfmt.Registry) error {

	if m.Timers != nil {

		if swag.IsZero(m.Timers) { // not required
			return nil
		}

		if err := m.Timers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timers")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BgpPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BgpPeer) UnmarshalBinary(b []byte) error {
	var res BgpPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
