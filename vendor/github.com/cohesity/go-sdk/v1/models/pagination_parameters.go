// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaginationParameters Specifies the cursor based pagination parameters for Protection Source and
// its children. Pagination is supported at a given level within the Protection
// Source Hierarchy with the help of before or after cursors.
// A Cursor will always refer to a specific source within the source dataset
// but will be invalidated if the item is removed.
//
// swagger:model PaginationParameters
type PaginationParameters struct {

	// Specifies the entity id starting from which the items are to be returned.
	AfterCursorEntityID *int64 `json:"afterCursorEntityId,omitempty"`

	// Specifies the entity id upto which the items are to be returned.
	BeforeCursorEntityID *int64 `json:"beforeCursorEntityId,omitempty"`

	// Specifies the entity id for the Node at any level within the Source entity
	// hierarchy whose children are to be paginated.
	NodeID *int64 `json:"nodeId,omitempty"`

	// Specifies the maximum number of entities to be returned within the page.
	PageSize *int64 `json:"pageSize,omitempty"`
}

// Validate validates this pagination parameters
func (m *PaginationParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pagination parameters based on context it is used
func (m *PaginationParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaginationParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationParameters) UnmarshalBinary(b []byte) error {
	var res PaginationParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
