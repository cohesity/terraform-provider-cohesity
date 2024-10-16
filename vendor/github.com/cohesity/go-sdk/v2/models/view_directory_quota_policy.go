// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ViewDirectoryQuotaPolicy Specifies a quota for View directory.
//
// swagger:model ViewDirectoryQuotaPolicy
type ViewDirectoryQuotaPolicy struct {

	// Specifies if an alert should be triggered when the usage of this resource exceeds this quota limit. This limit is optional and is specified in bytes. If no value is specified, there is no limit.
	AlertLimitBytes *int64 `json:"alertLimitBytes,omitempty"`

	// Specifies an optional quota limit on the usage allowed for this resource. This limit is specified in bytes. If no value is specified, there is no limit.
	HardLimitBytes *int64 `json:"hardLimitBytes,omitempty"`
}

// Validate validates this view directory quota policy
func (m *ViewDirectoryQuotaPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this view directory quota policy based on context it is used
func (m *ViewDirectoryQuotaPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ViewDirectoryQuotaPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewDirectoryQuotaPolicy) UnmarshalBinary(b []byte) error {
	var res ViewDirectoryQuotaPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
