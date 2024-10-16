// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterConfigProtoQoSMappingQoSContext QoSContext captures the properties that are relevant for QoS in a
// request. If a new field is added to QoSContext, cluster_config.h should
// be changed to enhance the hasher (QoSContextHash) and comparator
// (QoSContextEqual) for QoSContext.
//
// swagger:model ClusterConfigProto_QoSMapping_QoSContext
type ClusterConfigProtoQoSMappingQoSContext struct {

	// Component from which request is coming.
	Component *int32 `json:"component,omitempty"`

	// Priority of a request.
	Priority *int32 `json:"priority,omitempty"`

	// type
	Type *int32 `json:"type,omitempty"`

	// View box id of a request.
	ViewBoxID *int64 `json:"viewBoxId,omitempty"`

	// View id of a request.
	ViewID *int64 `json:"viewId,omitempty"`
}

// Validate validates this cluster config proto qo s mapping qo s context
func (m *ClusterConfigProtoQoSMappingQoSContext) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster config proto qo s mapping qo s context based on context it is used
func (m *ClusterConfigProtoQoSMappingQoSContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterConfigProtoQoSMappingQoSContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterConfigProtoQoSMappingQoSContext) UnmarshalBinary(b []byte) error {
	var res ClusterConfigProtoQoSMappingQoSContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
