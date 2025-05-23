// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpmiSelVerboseEntry Specifies each entry in the sel verbose response.
//
// swagger:model IpmiSelVerboseEntry
type IpmiSelVerboseEntry struct {

	// Specifies the ID corresponding to record in SEL(System Event Log) for given node.
	RecordID *string `json:"recordId,omitempty"`

	// Specifies the type of SEL record corresponding to the entry.
	RecordType *string `json:"recordType,omitempty"`

	// Specifies the time stamp at which the record is added to SEL.
	RecordTimestamp *string `json:"recordTimestamp,omitempty"`

	// Corresponds to source of component that generated the record.
	GeneratorID *string `json:"generatorId,omitempty"`

	// Specifies the version of the Event Message Format for the record added to SEL.
	EvmRevision *string `json:"evmRevision,omitempty"`

	// Specifies the type of the sensor corresponding to the current SEL record.
	SensorType *string `json:"sensorType,omitempty"`

	// Specifies the sensor number corresponding to the current SEL record.
	SensorNumber *string `json:"sensorNumber,omitempty"`

	// Specifies the type of event occurred.
	EventType *string `json:"eventType,omitempty"`

	// Specifies whether the event occurred is assertion or deassertion.
	EventDirection *string `json:"eventDirection,omitempty"`

	// Specifies additional info about the event.
	EventData *string `json:"eventData,omitempty"`

	// Specifies a short description corresponding to the sensor event for which record is added to SEL.
	RecordDescription *string `json:"recordDescription,omitempty"`
}

// Validate validates this ipmi sel verbose entry
func (m *IpmiSelVerboseEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ipmi sel verbose entry based on context it is used
func (m *IpmiSelVerboseEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpmiSelVerboseEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpmiSelVerboseEntry) UnmarshalBinary(b []byte) error {
	var res IpmiSelVerboseEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
