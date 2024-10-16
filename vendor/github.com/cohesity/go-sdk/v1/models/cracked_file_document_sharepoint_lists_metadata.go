// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CrackedFileDocumentSharepointListsMetadata Contains metadata specific to sharepoint lists.
//
// swagger:model CrackedFileDocument_SharepointListsMetadata
type CrackedFileDocumentSharepointListsMetadata struct {

	// No. of items in the sharepoint list.
	ItemCount *int32 `json:"itemCount,omitempty"`

	// Name of the list.
	ListName *string `json:"listName,omitempty"`

	// type
	Type *int32 `json:"type,omitempty"`
}

// Validate validates this cracked file document sharepoint lists metadata
func (m *CrackedFileDocumentSharepointListsMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cracked file document sharepoint lists metadata based on context it is used
func (m *CrackedFileDocumentSharepointListsMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CrackedFileDocumentSharepointListsMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrackedFileDocumentSharepointListsMetadata) UnmarshalBinary(b []byte) error {
	var res CrackedFileDocumentSharepointListsMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
