// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ViewBoxPairInfo Storage Domain (View Box) Pairing.
//
// Specifies a pairing between a Storage Domain (View Box) on local Cluster
// with a Storage Domain (View Box) on a remote Cluster.
// When replication is configured between a local Cluster and a
// remote Cluster, the Snapshots are replicated from the specified
// Storage Domain (View Box) on the local Cluster to the Storage Domain
// (View Box) on the remote Cluster. See the online help for details about the
// supported Storage Domains (View Box) pairings between Clusters.
//
// swagger:model ViewBoxPairInfo
type ViewBoxPairInfo struct {

	// Specifies the id of the Storage Domain (View Box) on the local Cluster.
	LocalViewBoxID *int64 `json:"localViewBoxId,omitempty"`

	// Specifies the name of the Storage Domain (View Box) on the local Cluster.
	LocalViewBoxName *string `json:"localViewBoxName,omitempty"`

	// Specifies the id of the Storage Domain (View Box) on the remote Cluster.
	RemoteViewBoxID *int64 `json:"remoteViewBoxId,omitempty"`

	// Specifies the name of the Storage Domain (View Box) on the remote Cluster.
	RemoteViewBoxName *string `json:"remoteViewBoxName,omitempty"`
}

// Validate validates this view box pair info
func (m *ViewBoxPairInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this view box pair info based on context it is used
func (m *ViewBoxPairInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ViewBoxPairInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewBoxPairInfo) UnmarshalBinary(b []byte) error {
	var res ViewBoxPairInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
