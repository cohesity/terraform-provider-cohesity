// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ArchivalMediaInfo Specifies the media information in QStar Archival.
//
// swagger:model ArchivalMediaInfo
type ArchivalMediaInfo struct {

	// Specifies a unique identifier for the media.
	// Read Only: true
	Barcode *string `json:"barcode,omitempty"`

	// Specifies the location of the offline media as recorded by the backup administrator using media management software.
	// Read Only: true
	Location *string `json:"location,omitempty"`

	// Specifies a flag that indicates if the media is online or offline. Offline media must be manually loaded into the media library before a recovery can occur.
	IsOnline *bool `json:"isOnline,omitempty"`
}

// Validate validates this archival media info
func (m *ArchivalMediaInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this archival media info based on the context it is used
func (m *ArchivalMediaInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBarcode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArchivalMediaInfo) contextValidateBarcode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "barcode", "body", m.Barcode); err != nil {
		return err
	}

	return nil
}

func (m *ArchivalMediaInfo) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "location", "body", m.Location); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArchivalMediaInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchivalMediaInfo) UnmarshalBinary(b []byte) error {
	var res ArchivalMediaInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
