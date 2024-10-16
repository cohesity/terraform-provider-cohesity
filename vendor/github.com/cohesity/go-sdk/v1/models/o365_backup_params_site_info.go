// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// O365BackupParamsSiteInfo o365 backup params site info
//
// swagger:model O365BackupParams_SiteInfo
type O365BackupParamsSiteInfo struct {

	// Id of the site.
	ID *string `json:"id,omitempty"`

	// Name of the site.
	Name *string `json:"name,omitempty"`
}

// Validate validates this o365 backup params site info
func (m *O365BackupParamsSiteInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this o365 backup params site info based on context it is used
func (m *O365BackupParamsSiteInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *O365BackupParamsSiteInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *O365BackupParamsSiteInfo) UnmarshalBinary(b []byte) error {
	var res O365BackupParamsSiteInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
