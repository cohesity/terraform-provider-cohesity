// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChassisInfo Chassis Information.
//
// ChassisInfo is the struct for the Chassis.
//
// swagger:model ChassisInfo
type ChassisInfo struct {

	// ChassisId is a unique id assigned to the chassis.
	ChassisID *int64 `json:"chassisId,omitempty"`

	// ChassisName is the name of the chassis. This could be the chassis serial
	// number by default.
	ChassisName *string `json:"chassisName,omitempty"`

	// Chassis serial.
	ChassisSerial *string `json:"chassisSerial,omitempty"`

	// Location is the location of the chassis within the rack.
	Location *string `json:"location,omitempty"`

	// Rack is the rack within which this chassis lives.
	RackID *int64 `json:"rackId,omitempty"`
}

// Validate validates this chassis info
func (m *ChassisInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chassis info based on context it is used
func (m *ChassisInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChassisInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChassisInfo) UnmarshalBinary(b []byte) error {
	var res ChassisInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
