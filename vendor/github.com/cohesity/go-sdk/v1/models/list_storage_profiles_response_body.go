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

// ListStorageProfilesResponseBody List Storage Profiles Response Body.
//
// Specifies the response to request to list the storage profiles associated
// with a VDC.
//
// swagger:model ListStorageProfilesResponseBody
type ListStorageProfilesResponseBody struct {

	// Specifies a list of storage profiles.
	StorageProfiles []*VcdStorageProfile `json:"storageProfiles"`
}

// Validate validates this list storage profiles response body
func (m *ListStorageProfilesResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorageProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListStorageProfilesResponseBody) validateStorageProfiles(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageProfiles); i++ {
		if swag.IsZero(m.StorageProfiles[i]) { // not required
			continue
		}

		if m.StorageProfiles[i] != nil {
			if err := m.StorageProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageProfiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list storage profiles response body based on the context it is used
func (m *ListStorageProfilesResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageProfiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListStorageProfilesResponseBody) contextValidateStorageProfiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StorageProfiles); i++ {

		if m.StorageProfiles[i] != nil {

			if swag.IsZero(m.StorageProfiles[i]) { // not required
				return nil
			}

			if err := m.StorageProfiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageProfiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListStorageProfilesResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListStorageProfilesResponseBody) UnmarshalBinary(b []byte) error {
	var res ListStorageProfilesResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
