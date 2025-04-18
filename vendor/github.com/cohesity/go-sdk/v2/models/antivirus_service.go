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

// AntivirusService Specifies an Antivirus Service.
//
// swagger:model AntivirusService
type AntivirusService struct {

	// Specifies the ICAP Uri for the Antivirus Service.
	// Required: true
	IcapURI *string `json:"icapUri"`

	// Specifies the description for the Antivirus Service.
	Description *string `json:"description,omitempty"`

	// Specifies the tag of the Antivirus Service.
	// Read Only: true
	Tag *string `json:"tag,omitempty"`

	// Specifies the tag id of the Antivirus Service.
	// Read Only: true
	TagID *int64 `json:"tagId,omitempty"`
}

// Validate validates this antivirus service
func (m *AntivirusService) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIcapURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntivirusService) validateIcapURI(formats strfmt.Registry) error {

	if err := validate.Required("icapUri", "body", m.IcapURI); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this antivirus service based on the context it is used
func (m *AntivirusService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTagID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AntivirusService) contextValidateTag(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

func (m *AntivirusService) contextValidateTagID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tagId", "body", m.TagID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntivirusService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntivirusService) UnmarshalBinary(b []byte) error {
	var res AntivirusService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
