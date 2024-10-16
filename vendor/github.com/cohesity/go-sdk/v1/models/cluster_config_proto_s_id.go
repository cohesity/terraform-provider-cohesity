// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterConfigProtoSID Represents the security identifier that uniquely defines a security
// principal. SIDs are associated with users and groups.
// Reference: https://msdn.microsoft.com/en-us/library/aa379597.aspx
//
// swagger:model ClusterConfigProto_SID
type ClusterConfigProtoSID struct {

	// The authority under which the SID was created. This is always 6 bytes
	// long.
	IdentifierAuthority []uint8 `json:"identifierAuthority"`

	// The revision level of the SID.
	RevisionLevel *int32 `json:"revisionLevel,omitempty"`

	// List of ids relative to the identifier_authority that uniquely
	// identify a principal. The last entry in this list is the RID, which
	// uniquely identifies the principal within a domain.
	SubAuthority []uint32 `json:"subAuthority"`
}

// Validate validates this cluster config proto s ID
func (m *ClusterConfigProtoSID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster config proto s ID based on context it is used
func (m *ClusterConfigProtoSID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterConfigProtoSID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterConfigProtoSID) UnmarshalBinary(b []byte) error {
	var res ClusterConfigProtoSID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
