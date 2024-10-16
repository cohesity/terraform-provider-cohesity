// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SourcesForSid Sources for Sid.
//
// Protection Sources and Views With Access Permissions.
// Specifies the Protection Sources objects and Views that the specified
// principal has permissions to access. The principal is specified using
// a security identifier (SID).
//
// swagger:model SourcesForSid
type SourcesForSid struct {

	// Array of Protection Sources.
	//
	// Specifies the Protection Source objects that the specified principal
	// has permissions to access.
	ProtectionSources []*ProtectionSource `json:"protectionSources"`

	// Specifies the security identifier (SID) of the principal.
	Sid *string `json:"sid,omitempty"`

	// Array of View Names.
	//
	// Specifies the names of the Views that the specified principal has
	// permissions to access.
	Views []*View `json:"views"`
}

// Validate validates this sources for sid
func (m *SourcesForSid) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtectionSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViews(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourcesForSid) validateProtectionSources(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionSources) { // not required
		return nil
	}

	for i := 0; i < len(m.ProtectionSources); i++ {
		if swag.IsZero(m.ProtectionSources[i]) { // not required
			continue
		}

		if m.ProtectionSources[i] != nil {
			if err := m.ProtectionSources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionSources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionSources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SourcesForSid) validateViews(formats strfmt.Registry) error {
	if swag.IsZero(m.Views) { // not required
		return nil
	}

	for i := 0; i < len(m.Views); i++ {
		if swag.IsZero(m.Views[i]) { // not required
			continue
		}

		if m.Views[i] != nil {
			if err := m.Views[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("views" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("views" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sources for sid based on the context it is used
func (m *SourcesForSid) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtectionSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViews(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourcesForSid) contextValidateProtectionSources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProtectionSources); i++ {

		if m.ProtectionSources[i] != nil {

			if swag.IsZero(m.ProtectionSources[i]) { // not required
				return nil
			}

			if err := m.ProtectionSources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protectionSources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protectionSources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SourcesForSid) contextValidateViews(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Views); i++ {

		if m.Views[i] != nil {

			if swag.IsZero(m.Views[i]) { // not required
				return nil
			}

			if err := m.Views[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("views" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("views" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourcesForSid) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourcesForSid) UnmarshalBinary(b []byte) error {
	var res SourcesForSid
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
