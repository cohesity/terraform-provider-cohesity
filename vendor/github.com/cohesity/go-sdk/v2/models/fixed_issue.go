// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FixedIssue Fixed Issue.
//
// Specifies the description of a fixed issue.
//
// swagger:model FixedIssue
type FixedIssue struct {

	// Specifies a unique number of the bug.
	ID int64 `json:"id,omitempty"`

	// Specifies the description of fix made for the issue.
	ReleaseNote string `json:"releaseNote,omitempty"`
}

// Validate validates this fixed issue
func (m *FixedIssue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fixed issue based on context it is used
func (m *FixedIssue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FixedIssue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FixedIssue) UnmarshalBinary(b []byte) error {
	var res FixedIssue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
