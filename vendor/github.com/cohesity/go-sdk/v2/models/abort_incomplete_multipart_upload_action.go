// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AbortIncompleteMultipartUploadAction Specifies the Lifecycle Abort Incomplete Multipart Upload Action.
//
// swagger:model AbortIncompleteMultipartUploadAction
type AbortIncompleteMultipartUploadAction struct {

	// Specifies the number of days after which to abort an incomplete multipart upload.
	Days *int64 `json:"days,omitempty"`
}

// Validate validates this abort incomplete multipart upload action
func (m *AbortIncompleteMultipartUploadAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this abort incomplete multipart upload action based on context it is used
func (m *AbortIncompleteMultipartUploadAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AbortIncompleteMultipartUploadAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AbortIncompleteMultipartUploadAction) UnmarshalBinary(b []byte) error {
	var res AbortIncompleteMultipartUploadAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
