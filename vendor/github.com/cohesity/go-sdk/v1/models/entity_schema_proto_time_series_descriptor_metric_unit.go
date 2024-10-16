// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EntitySchemaProtoTimeSeriesDescriptorMetricUnit Metric Unit.
//
// Specifies the unit of measure for the metric.
// O specifies a unit of space used such as free disk space.
// 1 specifies a Unix epoch Timestamp (in microseconds).
// 2 specifies a Unix epoch Timestamp (in milliseconds).
// 3 specifies a Unix epoch Timestamp (in seconds).
// 4 specifies a Unix epoch Timestamp (in minutes).
// 5 specifies a counter such as the read IO metric.
// 6 specifies the temperature in Centigrade.
// 7 specifies the temperature in Fahrenheit.
// 8 specifies revolutions per minute such as a CPU fan speed.
// 9 specifies a percentage such as CPU or memory usage.
//
// swagger:model EntitySchemaProto_TimeSeriesDescriptor_MetricUnit
type EntitySchemaProtoTimeSeriesDescriptorMetricUnit struct {

	// type
	Type *int32 `json:"type,omitempty"`
}

// Validate validates this entity schema proto time series descriptor metric unit
func (m *EntitySchemaProtoTimeSeriesDescriptorMetricUnit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this entity schema proto time series descriptor metric unit based on context it is used
func (m *EntitySchemaProtoTimeSeriesDescriptorMetricUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EntitySchemaProtoTimeSeriesDescriptorMetricUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntitySchemaProtoTimeSeriesDescriptorMetricUnit) UnmarshalBinary(b []byte) error {
	var res EntitySchemaProtoTimeSeriesDescriptorMetricUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
