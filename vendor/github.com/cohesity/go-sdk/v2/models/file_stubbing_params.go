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

// FileStubbingParams File stubbing params
//
// swagger:model FileStubbingParams
type FileStubbingParams struct {

	// Specifies whether to create a symlink for the migrated data from source to target.
	SkipBackSymlink *bool `json:"skipBackSymlink,omitempty"`

	// Specifies whether to remove the orphan data from the target if the symlink is removed from the source.
	AutoOrphanDataCleanup *bool `json:"autoOrphanDataCleanup,omitempty"`

	// downtiering file age
	DowntieringFileAge *DowntieringFileAgePolicy `json:"downtieringFileAge,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FileStubbingParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		SkipBackSymlink *bool `json:"skipBackSymlink,omitempty"`

		AutoOrphanDataCleanup *bool `json:"autoOrphanDataCleanup,omitempty"`

		DowntieringFileAge *DowntieringFileAgePolicy `json:"downtieringFileAge,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.SkipBackSymlink = dataAO0.SkipBackSymlink

	m.AutoOrphanDataCleanup = dataAO0.AutoOrphanDataCleanup

	m.DowntieringFileAge = dataAO0.DowntieringFileAge

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FileStubbingParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		SkipBackSymlink *bool `json:"skipBackSymlink,omitempty"`

		AutoOrphanDataCleanup *bool `json:"autoOrphanDataCleanup,omitempty"`

		DowntieringFileAge *DowntieringFileAgePolicy `json:"downtieringFileAge,omitempty"`
	}

	dataAO0.SkipBackSymlink = m.SkipBackSymlink

	dataAO0.AutoOrphanDataCleanup = m.AutoOrphanDataCleanup

	dataAO0.DowntieringFileAge = m.DowntieringFileAge

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this file stubbing params
func (m *FileStubbingParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDowntieringFileAge(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileStubbingParams) validateDowntieringFileAge(formats strfmt.Registry) error {

	if swag.IsZero(m.DowntieringFileAge) { // not required
		return nil
	}

	if m.DowntieringFileAge != nil {
		if err := m.DowntieringFileAge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downtieringFileAge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downtieringFileAge")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this file stubbing params based on the context it is used
func (m *FileStubbingParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDowntieringFileAge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FileStubbingParams) contextValidateDowntieringFileAge(ctx context.Context, formats strfmt.Registry) error {

	if m.DowntieringFileAge != nil {

		if swag.IsZero(m.DowntieringFileAge) { // not required
			return nil
		}

		if err := m.DowntieringFileAge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downtieringFileAge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downtieringFileAge")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FileStubbingParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileStubbingParams) UnmarshalBinary(b []byte) error {
	var res FileStubbingParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
