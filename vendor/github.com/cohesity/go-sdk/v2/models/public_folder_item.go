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

// PublicFolderItem PublicFolderItem
//
// Specifies an Public folder indexed item.
//
// swagger:model PublicFolderItem
type PublicFolderItem struct {
	CommonIndexedObjectParams

	// Specifies the Public folder item type.
	Type *string `json:"type,omitempty"`

	// Specifies the id of the indexed item.
	ID *string `json:"id,omitempty"`

	// Specifies the subject of the indexed item.
	Subject *string `json:"subject,omitempty"`

	// Specifies whether the item has any attachments
	HasAttachments *bool `json:"hasAttachments,omitempty"`

	// Specifies the item class of the indexed item.
	ItemClass *string `json:"itemClass,omitempty"`

	// Specifies the Unix timestamp epoch in seconds at which this item is received.
	ReceivedTimeSecs *int64 `json:"receivedTimeSecs,omitempty"`

	// Specifies the size in bytes for the indexed item.
	ItemSize *int64 `json:"itemSize,omitempty"`

	// Specifies the id of parent folder the indexed item.
	ParentFolderID *string `json:"parentFolderId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PublicFolderItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonIndexedObjectParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonIndexedObjectParams = aO0

	// AO1
	var dataAO1 struct {
		Type *string `json:"type,omitempty"`

		ID *string `json:"id,omitempty"`

		Subject *string `json:"subject,omitempty"`

		HasAttachments *bool `json:"hasAttachments,omitempty"`

		ItemClass *string `json:"itemClass,omitempty"`

		ReceivedTimeSecs *int64 `json:"receivedTimeSecs,omitempty"`

		ItemSize *int64 `json:"itemSize,omitempty"`

		ParentFolderID *string `json:"parentFolderId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Type = dataAO1.Type

	m.ID = dataAO1.ID

	m.Subject = dataAO1.Subject

	m.HasAttachments = dataAO1.HasAttachments

	m.ItemClass = dataAO1.ItemClass

	m.ReceivedTimeSecs = dataAO1.ReceivedTimeSecs

	m.ItemSize = dataAO1.ItemSize

	m.ParentFolderID = dataAO1.ParentFolderID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PublicFolderItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonIndexedObjectParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Type *string `json:"type,omitempty"`

		ID *string `json:"id,omitempty"`

		Subject *string `json:"subject,omitempty"`

		HasAttachments *bool `json:"hasAttachments,omitempty"`

		ItemClass *string `json:"itemClass,omitempty"`

		ReceivedTimeSecs *int64 `json:"receivedTimeSecs,omitempty"`

		ItemSize *int64 `json:"itemSize,omitempty"`

		ParentFolderID *string `json:"parentFolderId,omitempty"`
	}

	dataAO1.Type = m.Type

	dataAO1.ID = m.ID

	dataAO1.Subject = m.Subject

	dataAO1.HasAttachments = m.HasAttachments

	dataAO1.ItemClass = m.ItemClass

	dataAO1.ReceivedTimeSecs = m.ReceivedTimeSecs

	dataAO1.ItemSize = m.ItemSize

	dataAO1.ParentFolderID = m.ParentFolderID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this public folder item
func (m *PublicFolderItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonIndexedObjectParams
	if err := m.CommonIndexedObjectParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this public folder item based on the context it is used
func (m *PublicFolderItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonIndexedObjectParams
	if err := m.CommonIndexedObjectParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PublicFolderItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicFolderItem) UnmarshalBinary(b []byte) error {
	var res PublicFolderItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
