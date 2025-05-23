// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpmiSdrInfoEntry Specifies a single entry in the sdr info for the ipmi.
//
// swagger:model IpmiSdrInfoEntry
type IpmiSdrInfoEntry struct {

	// Specifies the type of sensor corresponding to the entry.
	SensorType *string `json:"sensorType,omitempty"`

	// Specifies the id of the sensor.
	ID *string `json:"id,omitempty"`

	// Specifies the status of the sensor.
	Status *string `json:"status,omitempty"`

	// Specifies the reading of the sensor.
	Reading *string `json:"reading,omitempty"`

	// Specifies the description of the event corresponding to the sensor entry.
	EventDescription *string `json:"eventDescription,omitempty"`
}

// Validate validates this ipmi sdr info entry
func (m *IpmiSdrInfoEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ipmi sdr info entry based on context it is used
func (m *IpmiSdrInfoEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpmiSdrInfoEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpmiSdrInfoEntry) UnmarshalBinary(b []byte) error {
	var res IpmiSdrInfoEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
