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

// BondSlaveInfo bond slave info
//
// swagger:model BondSlaveInfo
type BondSlaveInfo struct {

	// Bond seocondary link state.
	LinkState *string `json:"linkState,omitempty"`

	// Mac address of the bond secondary interface.
	MacAddr *string `json:"macAddr,omitempty"`

	// Bond secondary name.
	Name *string `json:"name,omitempty"`

	// Bond seocondarys slot info.
	Slot *string `json:"slot,omitempty"`

	// Bond seocondary Speed.
	Speed *string `json:"speed,omitempty"`

	// Interface Stats.
	Stats *InterfaceStats `json:"stats,omitempty"`

	// Bond secondary uplink switch info.
	UplinkSwitchInfo *UplinkSwitchInfo `json:"uplinkSwitchInfo,omitempty"`
}

// Validate validates this bond slave info
func (m *BondSlaveInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUplinkSwitchInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BondSlaveInfo) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *BondSlaveInfo) validateUplinkSwitchInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.UplinkSwitchInfo) { // not required
		return nil
	}

	if m.UplinkSwitchInfo != nil {
		if err := m.UplinkSwitchInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uplinkSwitchInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uplinkSwitchInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this bond slave info based on the context it is used
func (m *BondSlaveInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUplinkSwitchInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BondSlaveInfo) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {

		if swag.IsZero(m.Stats) { // not required
			return nil
		}

		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *BondSlaveInfo) contextValidateUplinkSwitchInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.UplinkSwitchInfo != nil {

		if swag.IsZero(m.UplinkSwitchInfo) { // not required
			return nil
		}

		if err := m.UplinkSwitchInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uplinkSwitchInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uplinkSwitchInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BondSlaveInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BondSlaveInfo) UnmarshalBinary(b []byte) error {
	var res BondSlaveInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
