// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ViewFailover Specifies the failover status of a view.
//
// swagger:model ViewFailover
type ViewFailover struct {

	// Specifies if the view is ready for failover.
	IsFailoverReady *bool `json:"isFailoverReady,omitempty"`

	// Specifies the remote view id.
	RemoteViewID *int64 `json:"remoteViewId,omitempty"`

	// Specifies the remote cluster id.
	RemoteClusterID *int64 `json:"remoteClusterId,omitempty"`

	// Specifies the remote cluster incarnation id.
	RemoteClusterIncarnationID *int64 `json:"remoteClusterIncarnationId,omitempty"`

	// View uid.
	ViewUID *string `json:"viewUid,omitempty"`
}

// Validate validates this view failover
func (m *ViewFailover) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this view failover based on context it is used
func (m *ViewFailover) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ViewFailover) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewFailover) UnmarshalBinary(b []byte) error {
	var res ViewFailover
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
