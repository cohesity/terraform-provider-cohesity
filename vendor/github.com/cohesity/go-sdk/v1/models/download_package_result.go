// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DownloadPackageResult Download Package Result.
//
// Specifies the result of a request to download a package to a Cluster.
//
// swagger:model DownloadPackageResult
type DownloadPackageResult struct {

	// Specifies a message describing the result of the request to download
	// a package to a Cluster.
	Message *string `json:"message,omitempty"`
}

// Validate validates this download package result
func (m *DownloadPackageResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this download package result based on context it is used
func (m *DownloadPackageResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DownloadPackageResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DownloadPackageResult) UnmarshalBinary(b []byte) error {
	var res DownloadPackageResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
